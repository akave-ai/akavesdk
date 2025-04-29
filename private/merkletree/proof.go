// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package merkletree provides memtree and merkle proof functionality.
package merkletree

import (
	"bytes"
	"crypto/sha256"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/storage/sealer/fr32"
	pool "github.com/libp2p/go-buffer-pool"
	"golang.org/x/xerrors"
)

// MaxMemtreeSize maximum memtree size.
const MaxMemtreeSize = 256 << 20

// BuildSha254Memtree builds a sha256 memtree from the input data
// Returned slice should be released to the pool after use.
func BuildSha254Memtree(rawIn io.Reader, size abi.UnpaddedPieceSize) ([]byte, error) {
	if size.Padded() > MaxMemtreeSize {
		return nil, xerrors.Errorf("piece too large for memtree: %d", size)
	}

	unpadBuf := pool.Get(int(size))
	// read into unpadBuf
	_, err := io.ReadFull(rawIn, unpadBuf)
	if err != nil {
		pool.Put(unpadBuf)
		return nil, xerrors.Errorf("failed to read into unpadBuf: %w", err)
	}

	nLeaves := int64(size.Padded()) / NodeSize
	totalNodes, levelSizes := computeTotalNodes(nLeaves, 2)
	memtreeBuf := pool.Get(int(totalNodes * NodeSize))

	fr32.Pad(unpadBuf, memtreeBuf[:size.Padded()])
	pool.Put(unpadBuf)

	d := sha256.New()

	levelStarts := make([]int64, len(levelSizes))
	levelStarts[0] = 0
	for i := 1; i < len(levelSizes); i++ {
		levelStarts[i] = levelStarts[i-1] + levelSizes[i-1]*NodeSize
	}

	for level := 1; level < len(levelSizes); level++ {
		levelNodes := levelSizes[level]
		prevLevelStart := levelStarts[level-1]
		currLevelStart := levelStarts[level]

		for i := int64(0); i < levelNodes; i++ {
			leftOffset := prevLevelStart + (2*i)*NodeSize
			d.Reset()
			d.Write(memtreeBuf[leftOffset : leftOffset+(NodeSize*2)])

			outOffset := currLevelStart + i*NodeSize

			// sum calls append, so we give it a zero len slice at the correct offset
			d.Sum(memtreeBuf[outOffset:outOffset])

			// set top bits to 00
			memtreeBuf[outOffset+NodeSize-1] &= 0x3F
		}
	}

	return memtreeBuf, nil
}

// NodeSize node size.
const NodeSize = 32

// Node byte representation of node.
type Node [NodeSize]byte

// RawMerkleProof contains leaf, proof, root.
type RawMerkleProof struct {
	Leaf  [32]byte
	Proof [][32]byte
	Root  [32]byte
}

// MemtreeProof generates a Merkle proof for the given leaf index from the memtree.
// The memtree is a byte slice containing all the nodes of the Merkle tree, including leaves and internal nodes.
func MemtreeProof(memtree []byte, leafIndex int64) (*RawMerkleProof, error) {
	// Currently, the implementation supports only binary trees (arity == 2)
	const arity = 2

	// Calculate the total number of nodes in the memtree
	totalNodes := int64(len(memtree)) / NodeSize

	// Reconstruct level sizes from the total number of nodes
	// Starting from the number of leaves, compute the number of nodes at each level
	nLeaves := (totalNodes + 1) / 2

	currLevelCount := nLeaves
	levelSizes := []int64{}
	totalNodesCheck := int64(0)

	for {
		levelSizes = append(levelSizes, currLevelCount)
		totalNodesCheck += currLevelCount

		if currLevelCount == 1 {
			break
		}
		// Compute the number of nodes in the next level
		currLevelCount = (currLevelCount + int64(arity) - 1) / int64(arity)
	}

	// Verify that the reconstructed total nodes match the actual total nodes
	if totalNodesCheck != totalNodes {
		return nil, xerrors.New("invalid memtree size; reconstructed total nodes do not match")
	}

	// Compute the starting byte offset for each level in memtree
	levelStarts := make([]int64, len(levelSizes))
	var offset int64
	for i, size := range levelSizes {
		levelStarts[i] = offset
		offset += size * NodeSize
	}

	// Validate the leaf index
	if leafIndex < 0 || leafIndex >= levelSizes[0] {
		return nil, xerrors.Errorf("invalid leaf index %d for %d leaves", leafIndex, levelSizes[0])
	}

	// Initialize the proof structure
	proof := &RawMerkleProof{
		Proof: make([][NodeSize]byte, 0, len(levelSizes)-1),
	}

	// Extract the leaf hash from the memtree
	leafOffset := levelStarts[0] + leafIndex*NodeSize
	copy(proof.Leaf[:], memtree[leafOffset:leafOffset+NodeSize])

	// Build the proof by collecting sibling hashes at each level
	index := leafIndex
	for level := 0; level < len(levelSizes)-1; level++ {
		siblingIndex := index ^ 1 // Toggle the last bit to get the sibling index

		siblingOffset := levelStarts[level] + siblingIndex*NodeSize
		var siblingHash [NodeSize]byte
		copy(siblingHash[:], memtree[siblingOffset:siblingOffset+NodeSize])
		proof.Proof = append(proof.Proof, siblingHash)

		// Move up to the parent index
		index /= int64(arity)
	}

	// Extract the root hash from the memtree
	rootOffset := levelStarts[len(levelSizes)-1]
	copy(proof.Root[:], memtree[rootOffset:rootOffset+NodeSize])
	return proof, nil
}

func computeTotalNodes(nLeaves, arity int64) (int64, []int64) {
	totalNodes := int64(0)
	levelCounts := []int64{}
	currLevelCount := nLeaves
	for currLevelCount > 0 {
		levelCounts = append(levelCounts, currLevelCount)
		totalNodes += currLevelCount
		if currLevelCount == 1 {
			break
		}
		currLevelCount = (currLevelCount + arity - 1) / arity
	}
	return totalNodes, levelCounts
}

// hash function to compute SHA-256 hash.
func hash(data []byte) []byte {
	sum := sha256.Sum256(data)
	sum[len(sum)-1] &= 0x3F
	return sum[:]
}

// VerifyMerkleProof verifies the provided Merkle proof.
// It checks if the reconstructed root matches the proof's root.
func VerifyMerkleProof(proof *RawMerkleProof, leafIndex int64) bool {
	currentHash := proof.Leaf[:] // Start with the leaf hash

	index := leafIndex
	for _, sibling := range proof.Proof {
		if index%2 == 0 {
			// Current hash is a left child
			currentHash = hash(append(currentHash, sibling[:]...))
		} else {
			// Current hash is a right child
			currentHash = hash(append(sibling[:], currentHash...))
		}
		index /= 2 // Move up to the parent index
	}

	// Compare the reconstructed root with the provided root
	return bytes.Equal(currentHash, proof.Root[:])
}
