package core

// ##CROSS: transfer log

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params/predeploy"
	"github.com/holiman/uint256"
)

var (
	transferLogSig common.Hash
)

func init() {
	ab, _ := predeploy.CrossExMetaData.GetAbi()
	transferLogSig = ab.Events["CrossTransfer"].ID
}

// AddTransferLog adds transfer log into state
func AddTransferLog(
	state vm.StateDB,
	sender, recipient common.Address,
	amount, senderBalance, recipientBalance *uint256.Int,
) {
	// ignore if amount is 0
	if amount.Sign() <= 0 {
		return
	}

	dataInputs := []*uint256.Int{
		amount,
		senderBalance,
		recipientBalance,
	}

	var data []byte
	for _, v := range dataInputs {
		data = append(data, common.LeftPadBytes(v.Bytes(), 32)...)
	}

	// add transfer log
	state.AddLog(&types.Log{
		Address: predeploy.CrossExAddr,
		Topics: []common.Hash{
			transferLogSig,
			common.BytesToHash(sender.Bytes()),
			common.BytesToHash(recipient.Bytes()),
		},
		Data: data,
	})
}
