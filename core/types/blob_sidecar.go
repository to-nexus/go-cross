// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

// ##CROSS: blob sidecars

type BlobSidecars []*BlobSidecar

// Len returns the length of s.
func (s BlobSidecars) Len() int { return len(s) }

// EncodeIndex encodes the i'th BlobTxSidecar to w. Note that this does not check for errors
// because we assume that BlobSidecars will only ever contain valid sidecars
func (s BlobSidecars) EncodeIndex(i int, w *bytes.Buffer) {
	rlp.Encode(w, s[i])
}

func (s BlobSidecars) BlobTxSidecarList() []*BlobTxSidecar {
	var inner []*BlobTxSidecar
	for _, sidecar := range s {
		inner = append(inner, &(sidecar.BlobTxSidecar))
	}
	return inner
}

type BlobSidecar struct {
	BlobTxSidecar
	BlockNumber *big.Int    `json:"blockNumber"`
	BlockHash   common.Hash `json:"blockHash"`
	TxIndex     uint64      `json:"transactionIndex"`
	TxHash      common.Hash `json:"transactionHash"`
}

func NewBlobSidecarFromTx(tx *Transaction) *BlobSidecar {
	if tx.BlobTxSidecar() == nil {
		return nil
	}
	return &BlobSidecar{
		BlobTxSidecar: *tx.BlobTxSidecar(),
		TxHash:        tx.Hash(),
	}
}

func (s *BlobSidecar) SanityCheck(blockNumber *big.Int, blockHash common.Hash) error {
	if s.BlockNumber.Cmp(blockNumber) != 0 {
		return errors.New("BlobSidecar with wrong block number")
	}
	if s.BlockHash != blockHash {
		return errors.New("BlobSidecar with wrong block hash")
	}
	if len(s.Blobs) != len(s.Commitments) {
		return errors.New("BlobSidecar has wrong commitment length")
	}
	if len(s.Blobs) != len(s.Proofs) {
		return errors.New("BlobSidecar has wrong proof length")
	}
	return nil
}

func (s *BlobSidecar) MarshalJSON() ([]byte, error) {
	fields := map[string]interface{}{
		"blockHash":   s.BlockHash,
		"blockNumber": hexutil.EncodeUint64(s.BlockNumber.Uint64()),
		"txHash":      s.TxHash,
		"txIndex":     hexutil.EncodeUint64(s.TxIndex),
	}
	fields["blobSidecar"] = s.BlobTxSidecar
	return json.Marshal(fields)
}

func (s *BlobSidecar) UnmarshalJSON(input []byte) error {
	type blobSidecar struct {
		BlobSidecar BlobTxSidecar `json:"blobSidecar"`
		BlockNumber *hexutil.Big  `json:"blockNumber"`
		BlockHash   common.Hash   `json:"blockHash"`
		TxIndex     *hexutil.Big  `json:"txIndex"`
		TxHash      common.Hash   `json:"txHash"`
	}
	var blob blobSidecar
	if err := json.Unmarshal(input, &blob); err != nil {
		return err
	}
	s.BlobTxSidecar = blob.BlobSidecar
	if blob.BlockNumber == nil {
		return errors.New("missing required field 'blockNumber' for BlobSidecar")
	}
	s.BlockNumber = blob.BlockNumber.ToInt()
	s.BlockHash = blob.BlockHash
	if blob.TxIndex == nil {
		return errors.New("missing required field 'txIndex' for BlobSidecar")
	}
	s.TxIndex = blob.TxIndex.ToInt().Uint64()
	s.TxHash = blob.TxHash
	return nil
}

// ##
