// Copyright (c) 2022 Project Illium
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package consensus

import (
	"context"
	"fmt"
	ctxio "github.com/jbenet/go-context/io"
	inet "github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-msgio"
	"github.com/project-illium/ilxd/net"
	"github.com/project-illium/ilxd/params"
	"github.com/project-illium/ilxd/types"
	"github.com/project-illium/ilxd/types/blocks"
	"github.com/project-illium/ilxd/types/wire"
	"google.golang.org/protobuf/proto"
	"io"
	"math/rand"
	"sync"
	"time"
)

const (
	// AvalancheRequestTimeout is the amount of time to wait for a response to a
	// query
	AvalancheRequestTimeout = 1 * time.Minute

	// AvalancheFinalizationScore is the confidence score we consider to be final
	AvalancheFinalizationScore = 160

	// AvalancheTimeStep is the amount of time to wait between event ticks
	AvalancheTimeStep = time.Millisecond

	// AvalancheMaxInflightPoll is the max outstanding requests that we can have
	// for any inventory item.
	AvalancheMaxInflightPoll = 10

	// AvalancheMaxElementPoll is the maximum number of invs to send in a single
	// query
	AvalancheMaxElementPoll = 4096

	// DeleteInventoryAfter is the maximum time we'll keep a block in memory
	// if it hasn't been finalized by avalanche.
	DeleteInventoryAfter = time.Hour * 6

	// ConsensusProtocol is the libp2p network protocol ID
	ConsensusProtocol = "/consensus/"

	// ConsensusProtocolVersion is the version of the ConsensusProtocol
	ConsensusProtocolVersion = "1.0.0"

	// MaxRejectedCache is the maximum size of the rejected cache
	MaxRejectedCache = 200

	// MinConnectedStakeThreshold is the minimum percentage of the weighted stake
	// set we must be connected to in order to finalize blocks.
	MinConnectedStakeThreshold = .5
)

type blockRecord struct {
	height         uint32
	activeBit      uint8
	preferredBit   uint8
	timestamp      time.Time
	finalizedBits  types.ID
	bitVotes       *VoteRecord
	blockInventory map[types.ID]*VoteRecord
}

func (br *blockRecord) HasFinalized() bool {
	for _, vr := range br.blockInventory {
		if vr.hasFinalized() {
			return true
		}
	}
	return false
}

// requestExpirationMsg signifies a request has expired and
// should be removed from the map.
type requestExpirationMsg struct {
	key string
	p   peer.ID
}

// queryMsg signifies a query from another peer.
type queryMsg struct {
	request    *wire.MsgAvaRequest
	respChan   chan *wire.MsgAvaResponse
	remotePeer peer.ID
}

// newBlockMessage represents new work for the engine.
type newBlockMessage struct {
	header       *blocks.BlockHeader
	isAcceptable bool
	callback     chan<- Status
}

// registerVotesMsg signifies a response to a query from another peer.
type registerVotesMsg struct {
	p    peer.ID
	resp *wire.MsgAvaResponse
}

// RequestBlockFunc is called when the engine receives a query from a peer about
// and unknown block. It should attempt to download the block from the remote peer,
// validate it, then pass it into the engine.
type RequestBlockFunc func(blockID types.ID, remotePeer peer.ID)

// HasBlockFunc checks the blockchain to see if we already have the block.
type HasBlockFunc func(blockID types.ID) bool

// ConsensusEngine implements a form of the avalanche consensus protocol.
// It primarily consists of an event loop that polls the weighted list of
// validators for any unfinalized blocks and records the responses. Blocks
// finalize when the confidence level exceeds the threshold.
type ConsensusEngine struct {
	ctx          context.Context
	network      *net.Network
	params       *params.NetworkParams
	chooser      *BackoffChooser
	ms           net.MessageSender
	valConn      ValidatorSetConnection
	self         peer.ID
	wg           sync.WaitGroup
	requestBlock RequestBlockFunc
	hasBlock     HasBlockFunc
	quit         chan struct{}
	msgChan      chan interface{}
	print        bool

	blockRecords   map[uint32]*blockRecord
	conflicts      map[uint32][]types.ID
	rejectedBlocks map[types.ID]struct{}
	queries        map[string]RequestRecord
	callbacks      map[types.ID]chan<- Status
}

// NewConsensusEngine returns a new ConsensusEngine
func NewConsensusEngine(ctx context.Context, opts ...Option) (*ConsensusEngine, error) {
	var cfg config
	for _, opt := range opts {
		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	eng := &ConsensusEngine{
		ctx:            ctx,
		network:        cfg.network,
		valConn:        cfg.valConn,
		chooser:        NewBackoffChooser(cfg.chooser),
		params:         cfg.params,
		self:           cfg.self,
		ms:             net.NewMessageSender(cfg.network.Host(), cfg.params.ProtocolPrefix+ConsensusProtocol+ConsensusProtocolVersion),
		wg:             sync.WaitGroup{},
		requestBlock:   cfg.requestBlock,
		hasBlock:       cfg.hasBlock,
		quit:           make(chan struct{}),
		msgChan:        make(chan interface{}),
		blockRecords:   make(map[uint32]*blockRecord),
		rejectedBlocks: make(map[types.ID]struct{}),
		conflicts:      make(map[uint32][]types.ID),
		queries:        make(map[string]RequestRecord),
		callbacks:      make(map[types.ID]chan<- Status),
	}
	eng.network.Host().SetStreamHandler(eng.params.ProtocolPrefix+ConsensusProtocol+ConsensusProtocolVersion, eng.HandleNewStream)
	eng.wg.Add(1)
	go eng.handler()
	return eng, nil
}

// Close gracefully shutsdown the consensus engine
func (eng *ConsensusEngine) Close() {
	close(eng.quit)
	eng.wg.Wait()
}

func (eng *ConsensusEngine) handler() {
	eventLoopTicker := time.NewTicker(AvalancheTimeStep)
out:
	for {
		select {
		case m := <-eng.msgChan:
			switch msg := m.(type) {
			case *requestExpirationMsg:
				eng.handleRequestExpiration(msg.key, msg.p)
			case *queryMsg:
				eng.handleQuery(msg.request, msg.remotePeer, msg.respChan)
			case *newBlockMessage:
				eng.handleNewBlock(msg.header, msg.isAcceptable, msg.callback)
			case *registerVotesMsg:
				eng.handleRegisterVotes(msg.p, msg.resp)
			}
		case <-eventLoopTicker.C:
			eng.pollLoop()
		case <-eng.quit:
			break out
		}
	}
	eventLoopTicker.Stop()
	eng.wg.Done()
}

// NewBlock is used to pass new work in the engine. The callback channel will return the final
// status (either Finalized or Rejected). Unfinalized but NotPreffered blocks will remain active
// in the engine until a conflicting block at the same height is finalized. At that point the block
// will be marked as Rejected.
func (eng *ConsensusEngine) NewBlock(header *blocks.BlockHeader, isAcceptable bool, callback chan<- Status) {
	eng.msgChan <- &newBlockMessage{
		header:       header,
		isAcceptable: isAcceptable,
		callback:     callback,
	}
}

func (eng *ConsensusEngine) handleNewBlock(header *blocks.BlockHeader, isAcceptable bool, callback chan<- Status) {
	blockID := header.ID().Clone()
	_, ok := eng.rejectedBlocks[blockID]
	if ok {
		return
	}

	record, ok := eng.blockRecords[header.Height]
	if !ok {
		record = &blockRecord{
			timestamp:      time.Now(),
			height:         header.Height,
			activeBit:      0,
			preferredBit:   0x80,
			finalizedBits:  types.ID{},
			bitVotes:       NewBitVoteRecord(header.Height),
			blockInventory: make(map[types.ID]*VoteRecord),
		}
		eng.blockRecords[header.Height] = record
	}

	initialPreference := isAcceptable
	// If we already have a preferred block at this height set the initial
	// preference to false.
	for invID, voteRecord := range record.blockInventory {
		if blockID == invID {
			return
		}
		if voteRecord.isPreferred() {
			initialPreference = false
		}
		log.Debugf("[CONSENSUS] Conflicting blocks at height %d: %s, %s", record.height, invID, header.ID())
	}

	if record.activeBit > 0 && !compareBits(record.finalizedBits, blockID, record.activeBit-1) {
		initialPreference = false
		log.Debugf("[CONSENSUS] received new block with rejected starting bits: %s", blockID)
	}

	if initialPreference {
		record.preferredBit = getBit(blockID, record.activeBit)
	}

	vr := NewBlockVoteRecord(blockID, header.Height, isAcceptable, initialPreference)
	record.blockInventory[blockID] = vr

	eng.callbacks[blockID] = callback
}

// HandleNewStream handles incoming streams from peers. We use one stream for
// incoming and a separate one for outgoing.
func (eng *ConsensusEngine) HandleNewStream(s inet.Stream) {
	go eng.handleNewMessage(s)
}

func (eng *ConsensusEngine) handleNewMessage(s inet.Stream) {
	defer s.Close()
	contextReader := ctxio.NewReader(eng.ctx, s)
	reader := msgio.NewVarintReaderSize(contextReader, inet.MessageSizeMax)
	remotePeer := s.Conn().RemotePeer()
	defer reader.Close()
	ticker := time.NewTicker(time.Minute)

	for {
		select {
		case <-eng.ctx.Done():
			return
		case <-ticker.C:
			return
		default:
		}

		req := new(wire.MsgAvaRequest)
		msgBytes, err := reader.ReadMsg()
		if err != nil {
			reader.ReleaseMsg(msgBytes)
			if err == io.EOF || err == inet.ErrReset {
				s.Close()
				return
			}
			log.Debugf("Error reading from avalanche stream: peer: %s, error: %s", remotePeer, err.Error())
			s.Reset()
			return
		}
		if err := proto.Unmarshal(msgBytes, req); err != nil {
			reader.ReleaseMsg(msgBytes)
			log.Debugf("Error unmarshalling avalanche message: peer: %s, error: %s", remotePeer, err.Error())
			s.Reset()
			return
		}
		reader.ReleaseMsg(msgBytes)

		respCh := make(chan *wire.MsgAvaResponse)
		eng.msgChan <- &queryMsg{
			request:    req,
			respChan:   respCh,
			remotePeer: remotePeer,
		}

		respMsg := <-respCh
		err = net.WriteMsg(s, respMsg)
		if err != nil {
			log.Errorf("Error writing avalanche stream to peer %d", remotePeer)
			s.Reset()
		}
		ticker.Reset(time.Minute)
	}
}

func (eng *ConsensusEngine) handleQuery(req *wire.MsgAvaRequest, remotePeer peer.ID, respChan chan *wire.MsgAvaResponse) {
	if len(req.Invs) == 0 {
		log.Debugf("Received empty avalanche request from peer %s", remotePeer)
		eng.network.IncreaseBanscore(remotePeer, 30, 0)
		return
	}
	resp := &wire.MsgAvaResponse{
		Request_ID: req.Request_ID,
		BitVote:    nil,
		BlockVotes: make([]byte, len(req.Invs)),
	}
	blkrec, blockRecExists := eng.blockRecords[req.Height]
	if !blockRecExists {
		resp.BitVote = []byte{0x80}
	} else {
		if req.Bit == uint32(blkrec.activeBit) {
			resp.BitVote = []byte{blkrec.preferredBit}
		} else if req.Bit > uint32(blkrec.activeBit) {
			resp.BitVote = []byte{0x80}
		} else if req.Bit < uint32(blkrec.activeBit) {
			bit := getBit(blkrec.finalizedBits, uint8(req.Bit))
			resp.BitVote = []byte{bit}
		}
		if eng.print {
			fmt.Printf("*** %d %d %08b %d\n", req.Bit, blkrec.activeBit, blkrec.finalizedBits[0], resp.BitVote[0])
		}
	}

	for i, invBytes := range req.Invs {
		inv := types.NewID(invBytes)

		if _, exists := eng.rejectedBlocks[inv]; exists {
			resp.BlockVotes[i] = 0x00 // No vote
			continue
		}
		var (
			record *VoteRecord
			ok     bool
		)
		if blockRecExists {
			record, ok = blkrec.blockInventory[inv]
		}
		if ok {
			// We're only going to vote for items we have a record for.
			resp.BlockVotes[i] = 0x00 // No vote
			if record.isPreferred() {
				resp.BlockVotes[i] = 0x01 // Yes vote
			}
		} else {
			if eng.hasBlock(inv) {
				resp.BlockVotes[i] = 0x01
			} else {
				resp.BlockVotes[i] = 0x80 // Neutral vote
				// Request to download the block from the remote peer
				go eng.requestBlock(inv, remotePeer)
			}
		}
	}

	respChan <- resp
}

func (eng *ConsensusEngine) handleRequestExpiration(key string, p peer.ID) {
	eng.chooser.RegisterDialFailure(p)
	r, ok := eng.queries[key]
	if !ok {
		return
	}
	delete(eng.queries, key)
	invs := r.GetInvs()
	for inv := range invs {
		br, ok := eng.blockRecords[r.height]
		if ok {
			vr, ok := br.blockInventory[inv]
			if ok {
				vr.inflightRequests--
			}
		}
	}
}

func (eng *ConsensusEngine) queueMessageToPeer(req *wire.MsgAvaRequest, peer peer.ID) {
	var (
		key  = queryKey(req.Request_ID, peer.String())
		resp = new(wire.MsgAvaResponse)
	)

	if peer != eng.self {
		err := eng.ms.SendRequest(eng.ctx, peer, req, resp)
		if err != nil {
			eng.msgChan <- &requestExpirationMsg{key, peer}
			return
		}
	} else {
		// Sleep here to not artificially advantage our own node.
		time.Sleep(time.Millisecond * 20)

		respCh := make(chan *wire.MsgAvaResponse)
		eng.msgChan <- &queryMsg{
			request:    req,
			remotePeer: peer,
			respChan:   respCh,
		}
		resp = <-respCh
	}

	eng.msgChan <- &registerVotesMsg{
		p:    peer,
		resp: resp,
	}
}

func (eng *ConsensusEngine) handleRegisterVotes(p peer.ID, resp *wire.MsgAvaResponse) {
	eng.chooser.RegisterDialSuccess(p)
	key := queryKey(resp.Request_ID, p.String())

	r, ok := eng.queries[key]
	if !ok {
		log.Debugf("Received avalanche response from peer %s with an unknown request ID", p)
		eng.network.IncreaseBanscore(p, 30, 0)
		return
	}

	// Always delete the key if it's present
	delete(eng.queries, key)

	if r.IsExpired() {
		log.Debugf("Received avalanche response from peer %s with an expired request", p)
		eng.network.IncreaseBanscore(p, 0, 20)
		return
	}

	invs := r.GetInvs()
	if len(resp.BlockVotes) != len(invs) {
		log.Debugf("Received avalanche response from peer %s with incorrect number of block votes", p)
		eng.network.IncreaseBanscore(p, 30, 0)
		return
	}

	if len(resp.BitVote) != 1 {
		log.Debugf("Received avalanche response from peer %s with incorrect number of bit votes", p)
		eng.network.IncreaseBanscore(p, 30, 0)
		return
	}

	br, ok := eng.blockRecords[r.height]
	if !ok {
		return
	}

	var (
		i          = -1
		bitFlipped = false
	)

	for inv := range invs {
		i++
		vr, ok := br.blockInventory[inv]
		if !ok {
			// We are not voting on this anymore
			continue
		}
		vr.inflightRequests--
		if vr.hasFinalized() {
			continue
		}

		if !vr.regsiterVote(resp.BlockVotes[i]) {
			// This vote did not provide any extra information
			continue
		}

		if vr.isPreferred() {
			bit := getBit(vr.blockID, br.activeBit)
			if bit != br.preferredBit {
				bitFlipped = true
				br.preferredBit = bit
				br.bitVotes.Reset(bit == 1)
			}

			// We need to keep track of conflicting blocks
			// when this one becomes accepted we need to set the
			// confidence of the conflicts back to zero.
			for conflict, rec := range br.blockInventory {
				if conflict != vr.blockID {
					rec.Reset(false)
				}
			}
		}

		if vr.status() == StatusFinalized {
			log.Debugf("[CONSENSUS] Block finalized in %d votes", vr.totalVotes)
			callback, ok := eng.callbacks[inv]
			if ok && callback != nil {
				go func(cb chan<- Status, stat Status) {
					cb <- vr.status()
				}(callback, vr.status())
			}
			for conflict, rec := range br.blockInventory {
				if conflict != vr.blockID {
					rec.Reject()
					eng.limitRejected()
					eng.rejectedBlocks[conflict] = struct{}{}
					callback, ok := eng.callbacks[conflict]
					if ok && callback != nil {
						go func(cb chan<- Status, stat Status) {
							callback <- stat
						}(callback, rec.status())
					}
				}
			}
		}
	}

	if !br.HasFinalized() && !bitFlipped && r.activeBit == br.activeBit {
		if br.bitVotes.regsiterVote(resp.BitVote[0]) {
			br.preferredBit = boolToUint8(br.bitVotes.isPreferred())
			if br.bitVotes.hasFinalized() {
				var (
					id              types.ID
					preferredExists = false
					newlyFlipped    = false
					nextBit         = uint8(0x80)
				)
				setBit(&br.finalizedBits, br.activeBit, br.preferredBit == 1)
				if eng.print {
					fmt.Printf("bit %d finalized as %d\n", br.activeBit, br.preferredBit)
					br.bitVotes.printState()
				}
				i := 0
				for _, vr := range br.blockInventory {
					if !compareBits(vr.blockID, br.finalizedBits, br.activeBit) {
						vr.Reset(false)
					} else if vr.isPreferred() {
						id = vr.blockID
						preferredExists = true
					} else if vr.acceptable {
						id = vr.blockID
						preferredExists = true
						newlyFlipped = true
					}
					i++
				}
				if preferredExists {
					if newlyFlipped {
						br.blockInventory[id].Reset(true)
					}
					nextBit = getBit(id, br.activeBit+1)
				}

				if eng.print {
					fmt.Println("next bit ", nextBit)
				}
				br.activeBit++
				br.bitVotes.Reset(nextBit == 1)
				br.preferredBit = nextBit
			}
		}
	}
}

func (eng *ConsensusEngine) pollLoop() {
	if eng.valConn.ConnectedStakePercentage() < MinConnectedStakeThreshold {
		return
	}
	for height, record := range eng.blockRecords {
		if time.Since(record.timestamp) > DeleteInventoryAfter {
			delete(eng.blockRecords, height)
			continue
		}
		if record.HasFinalized() {
			continue
		}

		invs := eng.getInvsForNextPoll(height)
		if len(invs) == 0 {
			return
		}

		p := eng.chooser.WeightedRandomValidator()
		if p == "" {
			for _, inv := range record.blockInventory {
				inv.inflightRequests--
			}
			continue
		}
		requestID := rand.Uint32()

		key := queryKey(requestID, p.String())
		eng.queries[key] = NewRequestRecord(time.Now().Unix(), height, record.activeBit, invs)

		invList := make([][]byte, 0, len(invs))
		for _, inv := range invs {
			b := inv.Clone()
			invList = append(invList, b[:])
		}

		req := &wire.MsgAvaRequest{
			Request_ID: requestID,
			Height:     height,
			Bit:        uint32(record.activeBit),
			Invs:       invList,
		}

		go eng.queueMessageToPeer(req, p)
	}
}

func (eng *ConsensusEngine) getInvsForNextPoll(height uint32) []types.ID {
	var invs []types.ID
	br, ok := eng.blockRecords[height]
	if !ok {
		return nil
	}
	for id, r := range br.blockInventory {
		// Delete very old inventory that hasn't finalized
		if time.Since(r.timestamp) > DeleteInventoryAfter {
			delete(br.blockInventory, id)
			continue
		}

		if r.hasFinalized() {
			// If this has finalized we can just skip.
			continue
		}

		confidence := r.getConfidence()
		var maxInflight uint8
		if confidence < AvalancheFinalizationScore {
			maxInflight = uint8(AvalancheFinalizationScore - r.getConfidence())
		}

		if maxInflight < AvalancheMaxInflightPoll {
			maxInflight = AvalancheMaxInflightPoll
		}

		if r.inflightRequests >= maxInflight {
			// If we are already at the max inflight then continue
			continue
		}
		r.inflightRequests++

		// We don't have a decision, we need more votes.
		invs = append(invs, id)
	}

	if len(invs) >= AvalancheMaxElementPoll {
		invs = invs[:AvalancheMaxElementPoll]
	}

	return invs
}

func (eng *ConsensusEngine) limitRejected() {
	if len(eng.rejectedBlocks) >= MaxRejectedCache {
		for blockID := range eng.rejectedBlocks {
			delete(eng.rejectedBlocks, blockID)
			break
		}
	}
}

func queryKey(requestID uint32, peerID string) string {
	return fmt.Sprintf("%d|%s", requestID, peerID)
}

func getBit(id types.ID, pos uint8) uint8 {
	byteIndex := pos / 8
	bitIndex := pos % 8

	value := id[byteIndex]
	bit := (value >> (7 - bitIndex)) & 1
	return bit
}

func compareBits(a, b types.ID, x uint8) bool {
	// Determine the number of full bytes to compare
	fullBytes := (x + 1) / 8
	remainingBits := (x + 1) % 8

	// Compare full bytes
	for i := uint8(0); i < fullBytes; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	// Check the remaining bits if there are any
	if remainingBits != 0 {
		mask := byte(255 << (8 - remainingBits)) // Create a mask for the remaining bits
		if (a[fullBytes] & mask) != (b[fullBytes] & mask) {
			return false
		}
	}

	return true
}

func setBit(arr *types.ID, bitIndex uint8, value bool) {
	byteIndex := bitIndex / 8     // Find the index of the byte
	bitPosition := 7 - bitIndex%8 // Find the position of the bit within the byte (MSB to LSB)

	if value {
		// Set the bit to 1
		arr[byteIndex] |= 1 << bitPosition
	} else {
		// Set the bit to 0
		arr[byteIndex] &^= 1 << bitPosition
	}
}
