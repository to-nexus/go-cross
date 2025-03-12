package core

// ##CROSS: transfer log

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/predeploys"
	"github.com/holiman/uint256"
)

var (
	transferLogSig  common.Hash
	transferLogAddr = common.BytesToHash(CrossExAddr.Bytes())
)

func init() {
	ab, _ := predeploys.CrossExMetaData.GetAbi()
	transferLogSig = ab.Events["LogTransfer"].ID
}

// AddTransferLog adds transfer log into state
func AddTransferLog(
	state vm.StateDB,
	sender, recipient common.Address,
	amount,
	input1, input2,
	output1, output2 *uint256.Int,
) {
	// ignore if amount is 0
	if amount.Sign() <= 0 {
		return
	}

	dataInputs := []*uint256.Int{
		amount,
		input1,
		input2,
		output1,
		output2,
	}

	var data []byte
	for _, v := range dataInputs {
		data = append(data, common.LeftPadBytes(v.Bytes(), 32)...)
	}

	// add transfer log
	state.AddLog(&types.Log{
		Address: CrossExAddr,
		Topics: []common.Hash{
			transferLogSig,
			transferLogAddr,
			common.BytesToHash(sender.Bytes()),
			common.BytesToHash(recipient.Bytes()),
		},
		Data: data,
	})
}
