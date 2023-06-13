// Copyright (c) 2022 The illium developers
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/project-illium/ilxd/blockchain"
	"github.com/project-illium/ilxd/mempool"
	"github.com/project-illium/ilxd/types"
	"github.com/project-illium/ilxd/types/blocks"
	"github.com/project-illium/ilxd/types/transactions"
	"sort"
	"sync"
	"time"
)

const (
	BlockGenerationInterval = time.Second
	BlockVersion            = 1
)

type BlockGenerator struct {
	privKey        crypto.PrivKey
	ownPeerID      peer.ID
	ownPeerIDBytes []byte
	mpool          *mempool.Mempool
	tickInterval   time.Duration
	chain          *blockchain.Blockchain
	broadcast      func(blk *blocks.XThinnerBlock) error
	active         bool
	activeMtx      sync.RWMutex
	quit           chan struct{}
}

func NewBlockGenerator(opts ...Option) (*BlockGenerator, error) {
	var cfg config
	for _, opt := range opts {
		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	ownPeerID, err := peer.IDFromPrivateKey(cfg.privKey)
	if err != nil {
		return nil, err
	}
	ownPeerIDBytes, err := ownPeerID.Marshal()
	if err != nil {
		return nil, err
	}

	if cfg.tickInterval == time.Duration(0) {
		cfg.tickInterval = BlockGenerationInterval
	}

	g := &BlockGenerator{
		ownPeerID:      ownPeerID,
		ownPeerIDBytes: ownPeerIDBytes,
		privKey:        cfg.privKey,
		mpool:          cfg.mpool,
		tickInterval:   cfg.tickInterval,
		chain:          cfg.chain,
		broadcast:      cfg.broadcastFunc,
		activeMtx:      sync.RWMutex{},
		active:         false,
	}

	return g, nil
}
func (g *BlockGenerator) Start() {
	g.activeMtx.Lock()
	defer g.activeMtx.Unlock()
	if g.active {
		return
	}
	g.active = true

	g.quit = make(chan struct{})
	go g.eventLoop()
	log.Info("Block generator active")
}

func (g *BlockGenerator) Close() {
	g.activeMtx.Lock()
	defer g.activeMtx.Unlock()

	if g.active {
		g.active = false
		close(g.quit)
	}
}

func (g *BlockGenerator) Active() bool {
	g.activeMtx.RLock()
	defer g.activeMtx.RUnlock()

	return g.active
}

func (g *BlockGenerator) eventLoop() {
	ticker := time.NewTicker(g.tickInterval)
	for {
		select {
		case <-ticker.C:
			val := g.chain.WeightedRandomValidator()
			if val == g.ownPeerID {
				if err := g.generateBlock(); err != nil {
					log.Warnf("Error in block generator: %s", err.Error())
				}
			}

		case <-g.quit:
			return
		}
	}
}

func (g *BlockGenerator) generateBlock() error {
	ok, err := g.chain.IsProducerUnderLimit(g.ownPeerID)
	if err != nil {
		return err
	}
	if !ok {
		fmt.Println("overlimit")
		return nil
	}

	bestID, height, timestamp := g.chain.BestBlock()

	now := time.Now()
	blockTime := now
	if !now.After(timestamp) {
		blockTime = timestamp.Add(time.Second)
	}
	// Don't generate a block if the timestamp would be too far into the future.
	if blockTime.After(now.Add(blockchain.MaxBlockFutureTime)) {
		return nil
	}

	blk := &blocks.Block{
		Header: &blocks.BlockHeader{
			Version:     BlockVersion,
			Height:      height + 1,
			Parent:      bestID[:],
			Timestamp:   blockTime.Unix(),
			Producer_ID: g.ownPeerIDBytes,
		},
	}

	// The consensus rules prevent a stake tx and a spend of a staked
	// nullifier from being in the same block. We'll loop through
	// and remove any spends of stake if they were in the mempool.
	txs := g.mpool.GetTransactions()
	if len(txs) == 0 {
		return nil
	}
	checkNullifiers := make(map[types.Nullifier]bool)
	for _, tx := range txs {
		if stake := tx.GetStakeTransaction(); stake != nil {
			checkNullifiers[types.NewNullifier(stake.Nullifier)] = true
		}
	}
	for txid, tx := range txs {
		switch t := tx.Tx.(type) {
		case *transactions.Transaction_StandardTransaction:
			for _, n := range t.StandardTransaction.Nullifiers {
				if checkNullifiers[types.NewNullifier(n)] {
					delete(txs, txid)
				}
			}
		case *transactions.Transaction_MintTransaction:
			for _, n := range t.MintTransaction.Nullifiers {
				if checkNullifiers[types.NewNullifier(n)] {
					delete(txs, txid)
				}
			}
		}
	}
	blk.Transactions = make([]*transactions.Transaction, 0, len(txs))
	for _, tx := range txs {
		blk.Transactions = append(blk.Transactions, tx)
	}

	sort.Sort(mempool.TxSorter(blk.Transactions))

	merkleRoot := blockchain.TransactionsMerkleRoot(blk.Transactions)
	blk.Header.TxRoot = merkleRoot[:]

	sigHash, err := blk.Header.SigHash()
	if err != nil {
		return err
	}
	sig, err := g.privKey.Sign(sigHash)
	if err != nil {
		return err
	}
	blk.Header.Signature = sig

	if err := g.chain.CheckConnectBlock(blk); err != nil {
		return err
	}

	xthinnerBlock, err := g.mpool.EncodeXthinner(blk.Txids())
	if err != nil {
		return err
	}
	xthinnerBlock.Header = blk.Header

	return g.broadcast(xthinnerBlock)
}
