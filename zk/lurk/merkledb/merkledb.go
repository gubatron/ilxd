// Copyright (c) 2022 The illium developers
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package merkledb

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/go-datastore"
	"github.com/project-illium/ilxd/params/hash"
	"github.com/project-illium/ilxd/repo"
	"github.com/project-illium/ilxd/types"
	"sync"
)

// MerkleProof represents a merkle inclusion or exclusion proof
// that links the data to a root hash.
type MerkleProof []types.ID

// Node represents a branch in the merkle tree. The left and right
// values can either be:
// - A zero byte array (represents a nil value)
// - A hash of data
// - A hash of a child node
type Node struct {
	left  types.ID
	right types.ID
}

// Left returns the left branch of the node.
func (n *Node) Left() types.ID {
	var id types.ID
	copy(id[:], n.left[:])
	return id
}

// Right returns the right branch of the node.
func (n *Node) Right() types.ID {
	var id types.ID
	copy(id[:], n.right[:])
	return id
}

// Hash returns the hash of the concatenation of the left and
// right branches.
func (n *Node) Hash() types.ID {
	return types.NewID(hash.HashMerkleBranches(n.left.Bytes(), n.right.Bytes()))
}

func (n *Node) copy() *Node {
	n2 := &Node{}
	copy(n2.left[:], n.left[:])
	copy(n2.right[:], n.right[:])
	return n2
}

// MerkleDB is a key value database which, in addition to storing key/value
// pairs, maintains a binary merkle tree of the data stored in the database.
// This allows us to create a root hash for the database as well as create
// compact inclusion and exclusion proofs for that data.
type MerkleDB struct {
	ds  repo.Datastore
	mtx sync.RWMutex
}

// NewMerkleDB returns a new database
func NewMerkleDB(ds repo.Datastore) (*MerkleDB, error) {
	dbtx, err := ds.NewTransaction(context.Background(), false)
	if err != nil {
		return nil, err
	}
	err = putRoot(dbtx, &Node{
		left:  types.NewID(nil),
		right: types.NewID(nil),
	})
	if err != nil {
		return nil, err
	}

	return &MerkleDB{ds: ds, mtx: sync.RWMutex{}}, dbtx.Commit(context.Background())
}

// Put a new key/value pair into the database. This operation
// will update the database's merkle root. It will also override
// any value that is currently stored in the database at this
// key.
func (mdb *MerkleDB) Put(key types.ID, value []byte) error {
	mdb.mtx.Lock()
	defer mdb.mtx.Unlock()

	var (
		nodeMap  = make(map[int]*nodeTracker)
		toDelete = make(map[types.ID]struct{})
	)
	dbtx, err := mdb.ds.NewTransaction(context.Background(), false)
	if err != nil {
		return err
	}

	if err := putValue(dbtx, key, value); err != nil {
		return err
	}
	if err := putKey(dbtx, types.NewIDFromData(value), key); err != nil {
		return err
	}

	rootNode, err := fetchRoot(dbtx)
	if err != nil {
		return err
	}

	if err := recursivePut(dbtx, key, value, rootNode, 0, nodeMap, toDelete); err != nil {
		return err
	}

	if err := putRoot(dbtx, rootNode); err != nil {
		return err
	}

	return dbtx.Commit(context.Background())
}

// Get returns a value from the database for a given key along with
// a merkle proof linking the value to the database's root hash. This
// method with not return an exclusion proof if the value does not exist
// just an error.
func (mdb *MerkleDB) Get(key types.ID) ([]byte, MerkleProof, error) {
	mdb.mtx.RLock()
	defer mdb.mtx.RUnlock()

	dbtx, err := mdb.ds.NewTransaction(context.Background(), true)
	if err != nil {
		return nil, nil, err
	}
	value, err := getValue(dbtx, key)
	if err != nil {
		return nil, nil, err
	}
	valHash := types.NewIDFromData(value)

	var nodes []*Node

	rootNode, err := fetchRoot(dbtx)
	if err != nil {
		return nil, nil, err
	}

	nodes = append(nodes, rootNode)

	var (
		keyBytes = key.Bytes()
		node     = rootNode.copy()
	)
	for i := 0; ; i++ {
		bit, err := getBit(keyBytes, i)
		if err != nil {
			return nil, nil, err
		}
		if bit == 0 {
			if node.left.Compare(valHash) == 0 {
				break
			} else if isNil(node.left) {
				return nil, nil, errors.New("node not found in tree")
			}
			node, err = fetchNode(dbtx, node.left)
			if err != nil {
				return nil, nil, err
			}
		} else {
			if node.right.Compare(valHash) == 0 {
				break
			} else if isNil(node.right) {
				return nil, nil, errors.New("node not found in tree")
			}
			node, err = fetchNode(dbtx, node.right)
			if err != nil {
				return nil, nil, err
			}
		}
		nodes = append(nodes, node)
	}
	proof := make(MerkleProof, len(nodes))
	for i, n := range nodes {
		bit, err := getBit(keyBytes, i)
		if err != nil {
			return nil, nil, err
		}
		if bit == 0 {
			proof[i] = n.right
		} else {
			proof[i] = n.left
		}
	}

	return value, proof, nil
}

// Exists returns whether the key exists in the database along with
// a merkle inclusion or exclusion proof which links either the value,
// or nil, to the database's merkle root.
func (mdb *MerkleDB) Exists(key types.ID) (bool, MerkleProof, error) {
	mdb.mtx.RLock()
	defer mdb.mtx.RUnlock()

	dbtx, err := mdb.ds.NewTransaction(context.Background(), true)
	if err != nil {
		return false, nil, err
	}

	var (
		nodes  []*Node
		exists bool
	)

	rootNode, err := fetchRoot(dbtx)
	if err != nil {
		return false, nil, err
	}

	nodes = append(nodes, rootNode)

	var (
		keyBytes = key.Bytes()
		node     = rootNode.copy()
	)
	for i := 0; ; i++ {
		bit, err := getBit(keyBytes, i)
		if err != nil {
			return false, nil, err
		}
		if bit == 0 {
			if isNil(node.left) {
				exists = false
				break
			}
			node, err = fetchNode(dbtx, node.left)
			if errors.Is(err, datastore.ErrNotFound) {
				exists = true
				break
			} else if err != nil {
				return false, nil, err
			}
		} else {
			if isNil(node.right) {
				exists = false
				break
			}
			node, err = fetchNode(dbtx, node.right)
			if errors.Is(err, datastore.ErrNotFound) {
				exists = true
				break
			} else if err != nil {
				return false, nil, err
			}
		}
		nodes = append(nodes, node)
	}

	proof := make(MerkleProof, len(nodes))
	for i, n := range nodes {
		bit, err := getBit(keyBytes, i)
		if err != nil {
			return false, nil, err
		}
		if bit == 0 {
			proof[i] = n.right
		} else {
			proof[i] = n.left
		}
	}
	return exists, proof, nil
}

// Delete removes a key/value pair from the databae. In the tree
// structure the left will be set to the nil hash.
func (mdb *MerkleDB) Delete(key types.ID) error {
	mdb.mtx.Lock()
	defer mdb.mtx.Unlock()

	var (
		nodeMap  = make(map[int]*nodeTracker)
		toDelete = make(map[types.ID]struct{})
	)
	dbtx, err := mdb.ds.NewTransaction(context.Background(), false)
	if err != nil {
		return err
	}

	val, err := getValue(dbtx, key)
	if err != nil {
		return err
	}

	rootNode, err := fetchRoot(dbtx)
	if err != nil {
		return err
	}

	if err := recursiveDelete(dbtx, key, val, rootNode, 0, nodeMap, toDelete); err != nil {
		return err
	}

	if err := putRoot(dbtx, rootNode); err != nil {
		return err
	}

	if err := deleteValue(dbtx, key); err != nil {
		return err
	}
	if err := deleteKey(dbtx, types.NewIDFromData(val)); err != nil {
		return err
	}

	return dbtx.Commit(context.Background())
}

// Root returns the database's root has.
func (mdb *MerkleDB) Root() (types.ID, error) {
	mdb.mtx.RLock()
	defer mdb.mtx.RUnlock()

	dbtx, err := mdb.ds.NewTransaction(context.Background(), false)
	if err != nil {
		return types.ID{}, err
	}

	rootNode, err := fetchRoot(dbtx)
	if err != nil {
		return types.ID{}, err
	}
	return rootNode.Hash(), nil
}

func ValidateProof(key types.ID, value []byte, root types.ID, proof MerkleProof) (bool, error) {
	keyBytes := key.Bytes()

	if len(proof) < 1 {
		return false, nil
	}

	dataHash := types.NewIDFromData(value)
	if value == nil {
		dataHash = types.NewID(nil)
	}

	hashVal := dataHash.Clone()
	for i := len(proof) - 1; i >= 0; i-- {
		bit, err := getBit(keyBytes, i)
		if err != nil {
			return false, err
		}
		h := proof[i]
		var node Node
		if bit == 0 {
			node.left = hashVal
			node.right = h
		} else {
			node.left = h
			node.right = hashVal
		}

		hashVal = node.Hash()
	}
	return hashVal.Compare(root) == 0, nil
}

type nodeTracker struct {
	node *Node
	left bool
}

func recursivePut(dbtx datastore.Txn, key types.ID, value []byte, n *Node, level int, nodeMap map[int]*nodeTracker, toDelete map[types.ID]struct{}) error {
	// Get the bit at the level
	bit, err := getBit(key.Bytes(), level)
	if err != nil {
		return err
	}

	toDelete[n.Hash()] = struct{}{}

	if bit == 0 { // zero
		if isNil(n.left) {
			// If the left hash is zero this means there are no nodes
			// below this level, and we can set the data here.
			//
			// Nil data is a delete op so use the nil hash here rather
			// than a hash of a zero byte slice.
			if value == nil {
				n.left = types.NewID(nil)
			} else {
				n.left = types.NewIDFromData(value)
			}
		} else {
			// If it's not nil this means either:
			// - There are child nodes below this one.
			// - This node holds the hash of the data and there are no children.
			next, err := fetchNode(dbtx, n.left)
			// If the node is not found it must be that this node holds data
			// and there are no children. So we'll create a nil child.
			if errors.Is(err, datastore.ErrNotFound) {
				nextKey, err := getKey(dbtx, n.left)
				if err != nil {
					return err
				}
				nextBit, err := getBit(nextKey.Bytes(), level+1)
				if err != nil {
					return err
				}
				if nextBit == 0 {
					next = &Node{
						left:  n.left.Clone(),
						right: types.NewID(nil),
					}
				} else {
					next = &Node{
						left:  types.NewID(nil),
						right: n.left.Clone(),
					}
				}
				if err := putNode(dbtx, next); err != nil {
					return err
				}
			} else if err != nil {
				return err
			}
			nodeMap[level] = &nodeTracker{
				node: n,
				left: true,
			}
			return recursivePut(dbtx, key, value, next, level+1, nodeMap, toDelete)
		}
	} else { // one
		if isNil(n.right) {
			// If the right hash is zero this means there are no nodes
			// below this level, and we can set the data here.
			//
			// Nil data is a delete op so use the nil hash here rather
			// than a hash of a zero byte slice.
			if value == nil {
				n.right = types.NewID(nil)
			} else {
				n.right = types.NewIDFromData(value)
			}
		} else {
			// If it's not nil this means either:
			// - There are child nodes below this one.
			// - This node holds the hash of the data and there are no children.
			next, err := fetchNode(dbtx, n.right)
			// If the node is not found it must be that this node holds data
			// and there are no children. So we'll create a nil child.
			if errors.Is(err, datastore.ErrNotFound) {
				nextKey, err := getKey(dbtx, n.right)
				if err != nil {
					return err
				}
				nextBit, err := getBit(nextKey.Bytes(), level+1)
				if err != nil {
					return err
				}
				if nextBit == 0 {
					next = &Node{
						left:  n.right.Clone(),
						right: types.NewID(nil),
					}
				} else {
					next = &Node{
						left:  types.NewID(nil),
						right: n.right.Clone(),
					}
				}
				if err := putNode(dbtx, next); err != nil {
					return err
				}
			} else if err != nil {
				return err
			}
			nodeMap[level] = &nodeTracker{
				node: n,
				left: false,
			}
			return recursivePut(dbtx, key, value, next, level+1, nodeMap, toDelete)
		}
	}
	// Put the updated node to the database.
	if err := putNode(dbtx, n); err != nil {
		return err
	}
	nodeMap[level] = &nodeTracker{
		node: n,
	}

	// Loop through the map of nodes we've loaded, set the child hash
	// correctly, and update the datastore.
	for i := level - 1; i >= 0; i-- {
		prev := nodeMap[i]
		if err := deleteNode(dbtx, prev.node.Hash()); err != nil {
			return err
		}
		if prev.left {
			prev.node.left = nodeMap[i+1].node.Hash()
		} else {
			prev.node.right = nodeMap[i+1].node.Hash()
		}
		if err := putNode(dbtx, prev.node); err != nil {
			return err
		}
	}
	for id := range toDelete {
		if err := deleteNode(dbtx, id); err != nil {
			return err
		}
	}
	return nil
}

func recursiveDelete(dbtx datastore.Txn, key types.ID, value []byte, n *Node, level int, nodeMap map[int]*nodeTracker, toDelete map[types.ID]struct{}) error {
	// Get the bit at the level
	bit, err := getBit(key.Bytes(), level)
	if err != nil {
		return err
	}

	toDelete[n.Hash()] = struct{}{}

	if bit == 0 { // zero
		if isNil(n.left) {
			return nil
		}
		if n.left.Compare(types.NewIDFromData(value)) == 0 {
			n.left = types.NewID(nil)
		} else {
			next, err := fetchNode(dbtx, n.left)
			if err != nil {
				return nil
			}
			nodeMap[level] = &nodeTracker{
				node: n,
				left: true,
			}
			return recursiveDelete(dbtx, key, value, next, level+1, nodeMap, toDelete)
		}
	} else { // one
		if isNil(n.right) {
			return nil
		}
		if n.right.Compare(types.NewIDFromData(value)) == 0 {
			n.right = types.NewID(nil)
		} else {
			next, err := fetchNode(dbtx, n.right)
			if err != nil {
				return nil
			}
			nodeMap[level] = &nodeTracker{
				node: n,
				left: false,
			}
			return recursiveDelete(dbtx, key, value, next, level+1, nodeMap, toDelete)
		}
	}
	// Put the updated node to the database.
	if err := putNode(dbtx, n); err != nil {
		return err
	}
	nodeMap[level] = &nodeTracker{
		node: n,
	}

	// Loop through the map of nodes we've loaded, set the child hash
	// correctly, and update the datastore.
	for i := level - 1; i >= 0; i-- {
		prev := nodeMap[i]
		if err := deleteNode(dbtx, prev.node.Hash()); err != nil {
			return err
		}
		if prev.left {
			h := nodeMap[i+1].node.Hash()
			if isNil(h) {
				prev.node.left = types.NewID(nil)
			} else {
				prev.node.left = h
			}
		} else {
			h := nodeMap[i+1].node.Hash()
			if isNil(h) {
				prev.node.right = types.NewID(nil)
			} else {
				prev.node.right = h
			}
		}
		if err := putNode(dbtx, prev.node); err != nil {
			return err
		}
	}
	for id := range toDelete {
		if err := deleteNode(dbtx, id); err != nil {
			return err
		}
	}
	return nil
}

func isNil(id types.ID) bool {
	return id.Compare(types.NewID(nil)) == 0 ||
		id.Compare(types.NewID(hash.HashMerkleBranches(types.NewID(nil).Bytes(), types.NewID(nil).Bytes()))) == 0
}

func getBit(slice []byte, pos int) (int, error) {
	if pos < 0 || pos >= len(slice)*8 {
		return 0, fmt.Errorf("position out of range")
	}
	byteIndex := pos / 8
	bitIndex := pos % 8

	value := slice[byteIndex]
	bit := (value >> (7 - bitIndex)) & 1
	return int(bit), nil
}
