package ethapi

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/xsleonard/go-merkle"
	"golang.org/x/crypto/sha3"
)

// ##CROSS: cross api

type CrossAPI struct {
	b Backend
}

// NewCrossAPI creates a new CrossAPI.
func NewCrossAPI(b Backend) *CrossAPI {
	return &CrossAPI{b: b}
}

// ##CROSS: checkpoint
// GetCheckpointProof returns a merkle proof for a given block number.
func (api *CrossAPI) GetCheckpointProof(ctx context.Context, blockNr, startNr, endNr rpc.BlockNumber) ([]hexutil.Bytes, error) {
	// does not accept aliases like "latest"
	if blockNr < 0 || startNr < 0 || endNr < 0 {
		return nil, errors.New("invalid block number")
	}
	if startNr >= endNr || !(startNr <= blockNr && blockNr <= endNr) {
		return nil, errors.New("invalid range")
	}
	if endNr-startNr > 86400 {
		return nil, errors.New("range too large")
	}

	// collect block hashes
	hashes := make([][]byte, 0, endNr-startNr+1)
	for n := startNr; n <= endNr; n++ {
		h, err := api.b.HeaderByNumber(ctx, n)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, h.Hash().Bytes())
	}

	// generate merkle proof
	tree, err := generateMerkleTree(hashes)
	if err != nil {
		return nil, err
	}

	return generateMerkleProof(tree, uint64(blockNr-startNr))
}

func generateMerkleTree(hashes [][]byte) (*merkle.Tree, error) {
	length := uint64(len(hashes))
	if length == 0 {
		return nil, errors.New("empty range")
	}

	// hashes will be leaf nodes of the merkle tree
	leaves := make([][]byte, nextPowerOfTwo(length)) // number of leaf nodes should be a power of 2
	copy(leaves, hashes)

	// build a merkle tree
	tree := merkle.NewTreeWithOpts(merkle.TreeOptions{EnableHashSorting: false, DisableHashLeaves: true})
	if err := tree.Generate(leaves, sha3.NewLegacyKeccak256()); err != nil {
		return nil, err
	}

	return &tree, nil
}

func generateMerkleProof(tree *merkle.Tree, index uint64) ([]hexutil.Bytes, error) {
	proof := make([]hexutil.Bytes, 0)

	if tree.Height() == 0 {
		return nil, errors.New("tree has no levels")
	}
	if index >= uint64(len(tree.Leaves())) {
		return nil, errors.New("index out of range")
	}

	// Start from the leaf node
	currentIndex := index

	// Traverse up the tree, collecting sibling nodes for the proof
	for level := len(tree.Levels) - 1; level > 0; level-- {
		nodes := tree.Levels[level]

		// Determine if current node is left or right child
		isLeftChild := currentIndex%2 == 0

		// Get sibling index
		siblingIndex := currentIndex
		if isLeftChild {
			siblingIndex = currentIndex + 1
		} else {
			siblingIndex = currentIndex - 1
		}

		// Add sibling to proof if it exists
		if siblingIndex < uint64(len(nodes)) {
			proof = append(proof, nodes[siblingIndex].Hash)
		}

		// Move to parent node for next iteration
		currentIndex = currentIndex / 2
	}
	return proof, nil
}

func nextPowerOfTwo(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	// http://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++

	return n
}

// ##
