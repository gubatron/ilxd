// Copyright (c) 2022 The illium developers
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package sync

import (
	"context"
	"github.com/go-test/deep"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/project-illium/ilxd/blockchain/harness"
	"github.com/project-illium/ilxd/net"
	"github.com/project-illium/ilxd/params"
	"github.com/project-illium/ilxd/types/blocks"
	"github.com/project-illium/ilxd/types/transactions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChainService(t *testing.T) {
	mn := mocknet.New()
	host1, err := mn.GenPeer()
	assert.NoError(t, err)
	network1, err := net.NewNetwork(context.Background(), []net.Option{
		net.WithHost(host1),
		net.Params(&params.RegestParams),
		net.BlockValidator(func(*blocks.XThinnerBlock, peer.ID) error {
			return nil
		}),
		net.MempoolValidator(func(transaction *transactions.Transaction) error {
			return nil
		}),
	}...)

	testHarness1, err := harness.NewTestHarness(harness.DefaultOptions())
	assert.NoError(t, err)

	err = testHarness1.GenerateBlocks(10)
	assert.NoError(t, err)

	ctx1, _ := context.WithCancel(context.Background())
	service1 := NewChainService(ctx1, testHarness1.Blockchain(), network1, testHarness1.Blockchain().Params())

	host2, err := mn.GenPeer()
	assert.NoError(t, err)
	network2, err := net.NewNetwork(context.Background(), []net.Option{
		net.WithHost(host2),
		net.Params(&params.RegestParams),
		net.BlockValidator(func(*blocks.XThinnerBlock, peer.ID) error {
			return nil
		}),
		net.MempoolValidator(func(transaction *transactions.Transaction) error {
			return nil
		}),
	}...)

	testHarness2, err := harness.NewTestHarness(harness.DefaultOptions())
	assert.NoError(t, err)

	err = testHarness2.GenerateBlocks(10)
	assert.NoError(t, err)

	ctx2, _ := context.WithCancel(context.Background())
	service2 := NewChainService(ctx2, testHarness2.Blockchain(), network2, testHarness2.Blockchain().Params())

	assert.NoError(t, mn.LinkAll())
	assert.NoError(t, mn.ConnectAllButSelf())

	b5, err := testHarness1.Blockchain().GetBlockByHeight(5)
	assert.NoError(t, err)

	ret, err := service2.GetBlockTxids(host1.ID(), b5.ID())
	assert.NoError(t, err)
	assert.Equal(t, b5.Txids(), ret)

	b4, err := testHarness2.Blockchain().GetBlockByHeight(4)
	assert.NoError(t, err)

	ret, err = service1.GetBlockTxids(host2.ID(), b4.ID())
	assert.NoError(t, err)
	assert.Equal(t, b4.Txids(), ret)

	ret2, err := service2.GetBlockTxs(host1.ID(), b5.ID(), []uint32{0})
	assert.NoError(t, err)
	assert.Empty(t, deep.Equal(b5.GetTransactions(), ret2))

	ret2, err = service1.GetBlockTxs(host2.ID(), b4.ID(), []uint32{0})
	assert.NoError(t, err)
	assert.Empty(t, deep.Equal(b4.GetTransactions(), ret2))

}