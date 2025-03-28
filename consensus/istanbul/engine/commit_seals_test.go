package engine

import (
	"bytes"
	"crypto/ecdsa"
	crand "crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestVerifyCommittedSeals(t *testing.T) {
	// 1 proposer 2 validator. the length of all validators is 3.

	signer, _ := ecdsa.GenerateKey(crypto.S256(), crand.Reader)
	validator1, _ := ecdsa.GenerateKey(crypto.S256(), crand.Reader)
	validator2, _ := ecdsa.GenerateKey(crypto.S256(), crand.Reader)

	signerAddress := crypto.PubkeyToAddress(signer.PublicKey)
	validators := make([]common.Address, 0)
	validators = append(validators, signerAddress)
	validators = append(validators, crypto.PubkeyToAddress(validator1.PublicKey))
	validators = append(validators, crypto.PubkeyToAddress(validator2.PublicKey))

	valSet := validator.NewSet(validators, istanbul.NewRoundRobinProposerPolicy())

	engine := &Engine{}
	expectedIstExtra := &types.IstanbulExtra{
		VanityData:   []byte{},
		Validators:   validators,
		Round:        0,
		Vote:         nil,
		RandomReveal: []byte{},
	}

	var buf bytes.Buffer
	_ = expectedIstExtra.EncodeRLP(&buf)

	encodedBytes := buf.Bytes()

	header := &types.Header{
		Number: big.NewInt(12),
		Extra:  encodedBytes}

	data := PrepareCommittedSeal(header, 0)

	seal, _ := crypto.Sign(data, signer)

	commitSeals := make([][]byte, 1)
	commitSeals[0] = seal

	_ = engine.CommitHeader(header, commitSeals, big.NewInt(0))

	parent := make([]*types.Header, 0)
	err := engine.verifyCommittedSeals(new(core.BlockChain), header, parent, valSet)
	assert.NoError(t, err, "In the 3 validators case, only 1 commit can pass this verification")
}
