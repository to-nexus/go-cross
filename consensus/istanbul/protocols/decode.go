package protocols

import (
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/rlp"
)

func Decode(code uint64, data []byte) (Message, error) {
	switch code {
	case PreprepareCode:
		var preprepare Preprepare
		if err := rlp.DecodeBytes(data, &preprepare); err != nil {
			return nil, istanbul.ErrFailedDecodePreprepare
		}
		preprepare.code = PreprepareCode
		return &preprepare, nil
	case PrepareCode:
		var prepare Prepare
		if err := rlp.DecodeBytes(data, &prepare); err != nil {
			return nil, istanbul.ErrFailedDecodeCommit
		}
		prepare.code = PrepareCode
		return &prepare, nil
	case CommitCode:
		var commit Commit
		if err := rlp.DecodeBytes(data, &commit); err != nil {
			return nil, istanbul.ErrFailedDecodeCommit
		}
		commit.code = CommitCode
		return &commit, nil
	case RoundChangeCode:
		var roundChange RoundChange
		if err := rlp.DecodeBytes(data, &roundChange); err != nil {
			return nil, istanbul.ErrFailedDecodeRoundChange
		}
		roundChange.code = RoundChangeCode
		return &roundChange, nil
	}

	return nil, istanbul.ErrInvalidMessage
}
