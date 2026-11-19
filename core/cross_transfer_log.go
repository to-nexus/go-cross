package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

// ##CROSS: transfer log

var (
	transferLogSig common.Hash = common.HexToHash("0x1b49dfd76419ac50d37de77c8afd5e57b6472c9ddd8399f88dc61343356a462b")
)

// AddTransferLog adds a log from the CrossEx address for a native value transfer.
// Its topics are (transferLogSig, sender, recipient) and its data is amount, senderBalance and recipientBalance,
// each a 32-byte big-endian word.
// The balances are a snapshot right after this transfer. They leave out later transfers in the same transaction
// and the refund of unused gas at its end, so the sender's balance is lower than its final balance by that refund,
// unless a fee payer bought the gas. Read final balances from the state.
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
		b := v.Bytes32()
		data = append(data, b[:]...)
	}

	// add transfer log
	state.AddLog(&types.Log{
		Address: contracts.CrossExAddr,
		Topics: []common.Hash{
			transferLogSig,
			common.BytesToHash(sender.Bytes()),
			common.BytesToHash(recipient.Bytes()),
		},
		Data: data,
	})
}
