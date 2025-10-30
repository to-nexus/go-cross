// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package breakpoint

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// CrossGovernorMetaData contains all meta data concerning the CrossGovernor contract.
var CrossGovernorMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BALLOT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"COUNTING_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVote\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteBySig\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteWithReason\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteWithReasonAndParams\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteWithReasonAndParamsBySig\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getKeeper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposalId\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotesWithParams\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashProposal\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalDeadline\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalDetails\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalDetailsAt\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalEta\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalNeedsQueuing\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalProposer\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalSnapshot\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalVotes\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"againstVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"forVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"abstainVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"propose\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queue\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queue\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"quorum\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumDenominator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumNumerator\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumNumerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relay\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"removeFromBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeeper\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProposalThreshold\",\"inputs\":[{\"name\":\"newProposalThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVotingDelay\",\"inputs\":[{\"name\":\"newVotingDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVotingPeriod\",\"inputs\":[{\"name\":\"newVotingPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIGovernor.ProposalState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timelock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC5805\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateQuorumNumerator\",\"inputs\":[{\"name\":\"newQuorumNumerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTimelock\",\"inputs\":[{\"name\":\"newTimelock\",\"type\":\"address\",\"internalType\":\"contractTimelockControllerUpgradeable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperChanged\",\"inputs\":[{\"name\":\"previousKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalCanceled\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalCreated\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"targets\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"signatures\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"},{\"name\":\"voteStart\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"voteEnd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalExecuted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalQueued\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"etaSeconds\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalThresholdSet\",\"inputs\":[{\"name\":\"oldProposalThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newProposalThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumNumeratorUpdated\",\"inputs\":[{\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newQuorumNumerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockChange\",\"inputs\":[{\"name\":\"oldTimelock\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTimelock\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnBlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCast\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCastWithParams\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VotingDelaySet\",\"inputs\":[{\"name\":\"oldVotingDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotingDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VotingPeriodSet\",\"inputs\":[{\"name\":\"oldVotingPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotingPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorAlreadyCastVote\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorAlreadyQueuedProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorDisabledDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorInsufficientProposerVotes\",\"inputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorInvalidProposalLength\",\"inputs\":[{\"name\":\"targets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"calldatas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"values\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorInvalidQuorumFraction\",\"inputs\":[{\"name\":\"quorumNumerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumDenominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorInvalidSignature\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorInvalidVoteParams\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorInvalidVoteType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorInvalidVotingPeriod\",\"inputs\":[{\"name\":\"votingPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorNonexistentProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorNotQueuedProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GovernorOnlyExecutor\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorQueueNotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GovernorRestrictedProposer\",\"inputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorUnableToCancel\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GovernorUnexpectedProposalState\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"uint8\",\"internalType\":\"enumIGovernor.ProposalState\"},{\"name\":\"expectedStates\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyKeeper\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "CrossGovernor",
	BinRuntime: "0x60806040526004361015610022575b3615610018575f80fd5b610020612eba565b005b5f3560e01c806301ffc9a71461040c57806302a251a31461040757806306f3f9e61461040257806306fdde03146103fd578063143489d0146103f8578063150b7a02146103f3578063160cbed7146103ee57806316e9eaec146103e95780632656227d146103e45780632d63f693146103df5780632e82db94146103da5780632fe3e261146103d5578063391b6f4e146103d05780633932abb1146103cb5780633e4f49e6146103c65780633f4ba83a146103c157806340e58ee5146103bc578063417c73a7146103b757806343859632146103b2578063452115d6146103ad5780634a49ac4c146103a85780634bf5d7e9146103a3578063544ffc9c1461039e57806354fd4d501461039957806356781388146103945780635b8d0e0d1461038f5780635c975abb1461038a5780635f398a141461038557806360c4247f14610380578063748747e61461037b57806379051887146103765780637b3c71d3146103715780637d5e81e21461036c5780637ecebe00146103675780638456cb591461036257806384b0196e1461035d5780638ff262e31461035857806391ddadf41461035357806397c3d3341461034e5780639a802a6d14610349578063a7713a7014610344578063a890c9101461033f578063a8f8a66814610317578063a9a952941461033a578063ab58fb8e14610335578063b58131b014610330578063bc197c811461032b578063c01f9e3714610326578063c28bc2fa14610321578063c4d66de81461031c578063c59057e414610317578063d33219b414610312578063da35c6641461030d578063dd4e2ba514610308578063ddf0b00914610303578063deaaa7cc146102fe578063e540d01d146102f9578063eb9019d4146102f4578063ece40cc1146102ef578063f23a6e61146102ea578063f8ce560a146102e5578063fc0c546a146102e05763fe0d94c10361000e57612038565b612004565b611fe6565b611f74565b611f50565b611ebe565b611e8c565b611e52565b611dfd565b611d9e565b611d75565b611d41565b6119fc565b611bd2565b611b67565b611b49565b611a9c565b611a73565b611a34565b611a18565b61196c565b611941565b611869565b61184e565b611824565b6116ec565b61160c565b61151e565b6114c4565b611416565b6113c1565b611394565b6112ac565b61128e565b61120e565b6111e0565b61113a565b6110d4565b6110a9565b61105c565b61102d565b610f8f565b610f73565b610f12565b610e71565b610e15565b610d7a565b610d4b565b610ce9565b610cb5565b610c7b565b610bc2565b610ba4565b610b8d565b610b5e565b610a3c565b6107f3565b6106d0565b6105de565b6104d8565b6104a0565b346104925760203660031901126104925760043563ffffffff60e01b8116809103610492576366defe7760e11b8114908115610481575b8115610470575b811561045f575b506040519015158152602090f35b6301ffc9a760e01b1490505f610451565b630271189760e51b8114915061044a565b6332a2ad4360e11b81149150610443565b5f80fd5b5f91031261049257565b34610492575f3660031901126104925760206104d063ffffffff5f5160206152985f395f51905f525460301c1690565b604051908152f35b34610492576020366003190112610492576004356104f4612edc565b6064811161058f576001600160d01b0361050c613780565b1690610516612bc0565b916001600160d01b038211610577577f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b46339979261055b906001600160d01b03841690614b95565b505060408051918252602082019290925290819081015b0390a1005b506306dfcc6560e41b5f5260d060045260245260445ffd5b63243e544560e01b5f52600452606460245260445ffd5b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9060206105db9281815201906105a6565b90565b34610492575f366003190112610492576040515f5f5160206154185f395f51905f525461060a81612088565b80845290600181169081156106ac5750600114610642575b61063e8361063281850382610741565b604051918291826105ca565b0390f35b5f5160206154185f395f51905f525f9081527fda13dda7583a39a3cd73e8830529c760837228fa4683752c823b17e10548aad5939250905b80821061069257509091508101602001610632610622565b91926001816020925483858801015201910190929161067a565b60ff191660208086019190915291151560051b840190910191506106329050610622565b34610492576020366003190112610492576004355f9081525f5160206152785f395f51905f5260209081526040909120546001600160a01b03166040516001600160a01b039091168152f35b6001600160a01b0381160361049257565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b0382111761076257604052565b61072d565b60405190610776608083610741565b565b60405190610776604083610741565b6001600160401b03811161076257601f01601f191660200190565b9291926107ae82610787565b916107bc6040519384610741565b829481845281830111610492578281602093845f960137010152565b9080601f83011215610492578160206105db933591016107a2565b346104925760803660031901126104925761080f60043561071c565b61081a60243561071c565b6064356001600160401b038111610492576108399036906004016107d8565b50610842612e0f565b306001600160a01b039091160361086557604051630a85bd0160e11b8152602090f35b637485328f60e11b5f5260045ffd5b6001600160401b0381116107625760051b60200190565b9080601f830112156104925781356108a281610874565b926108b06040519485610741565b81845260208085019260051b82010192831161049257602001905b8282106108d85750505090565b6020809183356108e78161071c565b8152019101906108cb565b9080601f8301121561049257813561090981610874565b926109176040519485610741565b81845260208085019260051b82010192831161049257602001905b82821061093f5750505090565b8135815260209182019101610932565b9080601f8301121561049257813561096681610874565b926109746040519485610741565b81845260208085019260051b820101918383116104925760208201905b8382106109a057505050505090565b81356001600160401b038111610492576020916109c2878480948801016107d8565b815201910190610991565b6080600319820112610492576004356001600160401b03811161049257816109f79160040161088b565b916024356001600160401b0381116104925782610a16916004016108f2565b91604435906001600160401b03821161049257610a359160040161094f565b9060643590565b346104925760206104d0610a4f366109cd565b92919091612107565b90602080835192838152019201905f5b818110610a755750505090565b82516001600160a01b0316845260209384019390920191600101610a68565b90602080835192838152019201905f5b818110610ab15750505090565b8251845260209384019390920191600101610aa4565b9080602083519182815201916020808360051b8301019401925f915b838310610af257505050505090565b9091929394602080610b10600193601f1986820301875289516105a6565b97019301930191939290610ae3565b949392610b4b606093610b3d610b599460808a5260808a0190610a58565b9088820360208a0152610a94565b908682036040880152610ac7565b930152565b346104925760203660031901126104925761063e610b7d600435612498565b9060409492945194859485610b1f565b60206104d0610b9b366109cd565b9291909161255c565b346104925760203660031901126104925760206104d06004356126bd565b34610492576020366003190112610492576004355f5160206152d85f395f51905f5254811015610c7657610c5e905f5160206152d85f395f51905f525f527f417ab5363cd721cbd7a5446b963bb4274a7db26ef2bb28819834095c165ca17a0154610c6c610c2f82612498565b9392959091610c50604051978897885260a0602089015260a0880190610a58565b908682036040880152610a94565b908482036060860152610ac7565b9060808301520390f35b612534565b34610492575f3660031901126104925760206040517f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118152f35b34610492575f366003190112610492575f5160206152585f395f51905f52546040516001600160a01b039091168152602090f35b34610492575f36600319011261049257602065ffffffffffff5f5160206152985f395f51905f525416604051908152f35b634e487b7160e01b5f52602160045260245ffd5b60081115610d3857565b610d1a565b6008811015610d3857602452565b3461049257602036600319011261049257610d676004356130f3565b6040516008821015610d38576020918152f35b34610492575f366003190112610492575f5160206152585f395f51905f52546001600160a01b03163303610e0657610db0613203565b610db8613203565b60ff195f5160206153d85f395f51905f5254165f5160206153d85f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b63c60eb33560e01b5f5260045ffd5b34610492576020366003190112610492576004355f525f5160206153785f395f51905f5260205261002060405f20600381015490610e5281612322565b90610e6b6002610e6460018401612373565b92016123bb565b916126e3565b3461049257602036600319011261049257600435610e8e8161071c565b5f5160206152585f395f51905f52546001600160a01b03163303610e06576001600160a01b03165f8181527f46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e50160205260408120805460ff191660011790557f7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec61941869080a2005b3461049257604036600319011261049257602060ff610f67602435600435610f398261071c565b5f525f5160206153985f395f51905f528452600360405f20019060018060a01b03165f5260205260405f2090565b54166040519015158152f35b346104925760206104d0610f86366109cd565b929190916126e3565b3461049257602036600319011261049257600435610fac8161071c565b5f5160206152585f395f51905f52546001600160a01b03163303610e06576001600160a01b03165f8181527f46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e50160205260408120805460ff191690557fe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de40119080a2005b34610492575f3660031901126104925761063e61104861289d565b6040519182916020835260208301906105a6565b34610492576020366003190112610492576004355f525f5160206153985f395f51905f52602052606060405f20805490600260018201549101549060405192835260208301526040820152f35b34610492575f3660031901126104925761063e61104861295f565b6024359060ff8216820361049257565b346104925760403660031901126104925760206104d06004356110f56110c4565b604051916111038584610741565b5f8352339061322b565b9181601f84011215610492578235916001600160401b038311610492576020838186019501011161049257565b346104925760c0366003190112610492576004356111566110c4565b90604435906111648261071c565b6064356001600160401b0381116104925761118390369060040161110d565b6084356001600160401b038111610492576111a29036906004016107d8565b9160a435946001600160401b0386116104925761063e966111ca6111d09736906004016107d8565b9561297e565b6040519081529081906020820190565b34610492575f36600319011261049257602060ff5f5160206153d85f395f51905f5254166040519015158152f35b346104925760803660031901126104925760043561122a6110c4565b906044356001600160401b0381116104925761124a90369060040161110d565b9190926064356001600160401b038111610492576112839461127361127b9236906004016107d8565b9436916107a2565b9133906132a8565b604051908152602090f35b346104925760203660031901126104925760206104d06004356133b9565b34610492576020366003190112610492576004356112c98161071c565b611008330361136f576001600160a01b0316801561133b575f5160206152585f395f51905f5254816001600160a01b0382167f068b48a2fe7f498b57ff6da64f075ae658fde8d77124b092e62b3dc58d91ce355f80a36001600160a01b031916175f5160206152585f395f51905f5255005b6084604051632c648cf160e01b815260406004820152600660448201526535b2b2b832b960d11b60648201525f6024820152fd5b630f22c43960e41b5f5261100860045260245ffd5b65ffffffffffff81160361049257565b34610492576020366003190112610492576100206004356113b481611384565b6113bc612edc565b6134ec565b34610492576060366003190112610492576004356113dd6110c4565b90604435906001600160401b0382116104925760209261140e6114076104d094369060040161110d565b36916107a2565b91339061322b565b34610492576080366003190112610492576004356001600160401b0381116104925761144690369060040161088b565b6024356001600160401b038111610492576114659036906004016108f2565b906044356001600160401b0381116104925761148590369060040161094f565b90606435916001600160401b03831161049257366023840112156104925761063e936114be6111d09436906024816004013591016107a2565b92612aa1565b34610492576020366003190112610492576004356114e18161071c565b60018060a01b03165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052602060405f2054604051908152f35b34610492575f366003190112610492575f5160206152585f395f51905f52546001600160a01b03163303610e06576115546135c8565b61155c6135c8565b600160ff195f5160206153d85f395f51905f525416175f5160206153d85f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b916115e2906115d46105db97959693600f60f81b865260e0602087015260e08601906105a6565b9084820360408601526105a6565b60608301949094526001600160a01b031660808201525f60a082015280830360c090910152610a94565b34610492575f366003190112610492577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005415806116c3575b15611686576116526135ef565b61165a6136ba565b9061063e60405161166c602082610741565b5f80825236602083013760405193849330914691866115ad565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015415611645565b34610492576080366003190112610492576004356117086110c4565b90604435916117168361071c565b6064356001600160401b038111610492576117e861173b6117ec9236906004016107d8565b6001600160a01b0386165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052604090208054600181019091556117e29060405160208101917ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d7835288604083015260ff8816606083015260018060a01b038a16608083015260a082015260a081526117da60c082610741565b519020613f9a565b86614026565b1590565b61180757906111d09161063e9361180161284e565b9261322b565b6394ab6c0760e01b5f526001600160a01b03831660045260245b5ffd5b34610492575f36600319011261049257602061183e612bc0565b65ffffffffffff60405191168152f35b34610492575f36600319011261049257602060405160648152f35b34610492576060366003190112610492576004356118868161071c565b6044356024356001600160401b038211610492576118aa60209236906004016107d8565b505f5160206153385f395f51905f5254604051630748d63560e31b81526001600160a01b0394851660048201526024810192909252909283916044918391165afa801561193c5761063e915f9161190d575b506040519081529081906020820190565b61192f915060203d602011611935575b6119278183610741565b810190612e2a565b5f6118fc565b503d61191d565b612843565b34610492575f3660031901126104925760206001600160d01b03611963613780565b16604051908152f35b34610492576020366003190112610492576004356119898161071c565b611991612edc565b5f5160206153f85f395f51905f5254604080516001600160a01b03808416825290931660208401819052927f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b2264019190a16001600160a01b031916175f5160206153f85f395f51905f5255005b346104925760206104d0611a0f366109cd565b92919091612de1565b3461049257602036600319011261049257602060405160018152f35b346104925760203660031901126104925760206104d06004355f525f5160206152785f395f51905f5260205265ffffffffffff600160405f2001541690565b34610492575f3660031901126104925760205f5160206153585f395f51905f5254604051908152f35b346104925760a036600319011261049257611ab860043561071c565b611ac360243561071c565b6044356001600160401b03811161049257611ae29036906004016108f2565b506064356001600160401b03811161049257611b029036906004016108f2565b506084356001600160401b03811161049257611b229036906004016107d8565b5061063e611b2e612c46565b6040516001600160e01b031990911681529081906020820190565b346104925760203660031901126104925760206104d0600435612c71565b606036600319011261049257600435611b7f8161071c565b6024356044356001600160401b03811161049257610020925f92611ba88493369060040161110d565b9190611bb2612edc565b826040519384928337810185815203925af1611bcc612cd0565b906137dc565b3461049257602036600319011261049257600435611bef8161071c565b5f5160206154385f395f51905f5254906001600160401b03611c2060ff604085901c1615936001600160401b031690565b1680159081611d39575b6001149081611d2f575b159081611d26575b50611d1757611c7f9082611c7660016001600160401b03195f5160206154385f395f51905f525416175f5160206154385f395f51905f5255565b611ce257612cff565b611c8557005b611caf60ff60401b195f5160206154385f395f51905f5254165f5160206154385f395f51905f5255565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2908060208101610572565b611d12600160401b60ff60401b195f5160206154385f395f51905f525416175f5160206154385f395f51905f5255565b612cff565b63f92ee8a960e01b5f5260045ffd5b9050155f611c3c565b303b159150611c34565b839150611c2a565b34610492575f366003190112610492575f5160206153f85f395f51905f52546040516001600160a01b039091168152602090f35b34610492575f3660031901126104925760205f5160206152d85f395f51905f5254604051908152f35b34610492575f3660031901126104925761063e604051611dbf604082610741565b602081527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e60208201526040519182916020835260208301906105a6565b34610492576020366003190112610492576004355f525f5160206153785f395f51905f5260205261002060405f20600381015490611e3a81612322565b90611e4c6002610e6460018401612373565b91612107565b34610492575f3660031901126104925760206040517ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d78152f35b346104925760203660031901126104925760043563ffffffff811681036104925761002090611eb9612edc565b613b1d565b3461049257604036600319011261049257600435611edb8161071c565b60206024355f604051611eee8482610741565b525f5160206153385f395f51905f5254604051630748d63560e31b81526001600160a01b0394851660048201526024810192909252909283916044918391165afa801561193c5761063e915f9161190d57506040519081529081906020820190565b3461049257602036600319011261049257610020600435611f6f612edc565b613bb6565b346104925760a036600319011261049257611f9060043561071c565b611f9b60243561071c565b6084356001600160401b03811161049257611fba9036906004016107d8565b50611fc3612e0f565b306001600160a01b03909116036108655760405163f23a6e6160e01b8152602090f35b346104925760203660031901126104925760206104d0600435612e39565b34610492575f366003190112610492575f5160206153385f395f51905f52546040516001600160a01b039091168152602090f35b6020366003190112610492576004355f525f5160206153785f395f51905f5260205261002060405f2060038101549061207081612322565b906120826002610e6460018401612373565b9161255c565b90600182811c921680156120b6575b60208310146120a257565b634e487b7160e01b5f52602260045260245ffd5b91607f1691612097565b5f527f0d5829787b8befdbc6044ef7457d8a95c2a04bc99235349f1a212c063e59d40160205260405f2090565b5f525f5160206152785f395f51905f5260205260405f2090565b9290919261211782858584612de1565b9361212185612f8e565b505f5160206153f85f395f51905f525461214b906001600160a01b03165b6001600160a01b031690565b936040519363793d064960e11b8552602085600481895afa94851561193c575f95612301575b503060601b6bffffffffffffffffffffffff191618946020604051809263b1c5f42760e01b825281806121aa8b89898c60048601613d58565b03915afa90811561193c575f916122e2575b506121c6876120c0565b555f5160206153f85f395f51905f52546121e8906001600160a01b031661213f565b90813b15610492575f809461221487604051998a97889687956308f2a0bb60e41b875260048701613d81565b03925af190811561193c5761223892612233926122c8575b5042613ddc565b6141f2565b65ffffffffffff8116156122b9576122b3816122947f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda289293600161227a876120ed565b019065ffffffffffff1665ffffffffffff19825416179055565b6040805185815265ffffffffffff909216602083015290918291820190565b0390a190565b634844252360e11b5f5260045ffd5b806122d65f6122dc93610741565b80610496565b5f61222c565b6122fb915060203d602011611935576119278183610741565b5f6121bc565b61231b91955060203d602011611935576119278183610741565b935f612171565b90604051918281549182825260208201905f5260205f20925f5b81811061235157505061077692500383610741565b84546001600160a01b031683526001948501948794506020909301920161233c565b90604051918281549182825260208201905f5260205f20925f5b8181106123a257505061077692500383610741565b845483526001948501948794506020909301920161238d565b908154916123c883610874565b926123d66040519485610741565b80845260208401915f5260205f20915f905b8282106123f55750505050565b6040515f855461240481612088565b8084529060018116908115612475575060011461243e575b506001928261243085946020940382610741565b8152019401910190926123e8565b5f878152602081209092505b81831061245f5750508101602001600161241c565b600181602092548386880101520192019161244a565b60ff191660208581019190915291151560051b840190910191506001905061241c565b805f525f5160206153785f395f51905f5260205260405f209160405190608082018281106001600160401b03821117610762576040526124d784612322565b82526124e560018501612373565b916020810192835260036124fb600287016123bb565b95604083019687520154936060820194808652156125225750519151935192519193929190565b636ad0607560e01b5f5260045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b8051821015610c765760209160051b010190565b9392909161256c82828588612de1565b9260209561257e60206010178661300e565b5061259e61258b866120ed565b805460ff60f01b1916600160f01b179055565b6125a6612e0f565b306001600160a01b0390911603612652575b846125c595969750613df9565b306125d161213f612e0f565b141580612625575b612610575b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f9080602081016122b3565b5f5f5160206153b85f395f51905f52556125de565b5061264d6117e85f5160206153b85f395f51905f52546001600160801b0381169060801c1490565b6125d9565b925f95929491955b84518110156126ae576001903061268461213f612677848a612548565b516001600160a01b031690565b14612690575b0161265a565b6126a961269d8289612548565b518a815191012061304c565b61268a565b509195508594909391926125b8565b5f525f5160206152785f395f51905f5260205265ffffffffffff60405f205460a01c1690565b939291906126f383838388612de1565b946126fd866130f3565b6008811015610d38571580612818575b156128015761271d949550612de1565b612728603b8261300e565b5061274a612735826120ed565b80546001600160f81b0316600160f81b179055565b6040518181527f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c90602090a161277f816120c0565b5480612789575090565b5f5160206153f85f395f51905f52546127aa906001600160a01b031661213f565b803b156104925760405163c4d252f560e01b815260048101929092525f908290602490829084905af1801561193c576127ed575b505f6127e9826120c0565b5590565b806122d65f6127fb93610741565b5f6127de565b85638fe5d8a960e01b5f526004523360245260445ffd5b505f8681525f5160206152785f395f51905f5260205260409020546001600160a01b0316331461270d565b6040513d5f823e3d90fd5b6040519061285d602083610741565b5f8252565b60405190612871604083610741565b601d82527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c740000006020830152565b5f5160206153385f395f51905f5254604051634bf5d7e960e01b8152905f90829060049082906001600160a01b03165afa5f91816128e4575b506105db57506105db612862565b9091503d805f833e6128f68183610741565b810190602081830312610492578051906001600160401b038211610492570181601f820112156104925780519061292c82610787565b9261293a6040519485610741565b8284526020838301011161049257815f9260208093018386015e83010152905f6128d6565b6040519061296e604083610741565b60018252603160f81b6020830152565b9390929196956117e8612a5691612a508a61299a3687896107a2565b6001600160a01b0382165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590602081519101208b5160208d0120906040519260208401947f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81186528d604086015260ff8d16606086015260018060a01b0316608085015260a084015260c083015260e082015260e081526117da61010082610741565b8a614026565b612a71576105db959691612a6b9136916107a2565b926132a8565b6394ab6c0760e01b5f526001600160a01b03871660045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b9193929093612ab08233613560565b15612bad575f5160206153585f395f51905f52549485612ad8575b6105db94955033936146e4565b5f1965ffffffffffff612ae9612bc0565b160165ffffffffffff8111612ba85765ffffffffffff16955f604051612b10602082610741565b525f5160206153385f395f51905f5254604051630748d63560e31b81523360048201526024810198909852602090889060449082906001600160a01b03165afa96871561193c575f97612b87575b50808710612b6c5750612acb565b636121770b60e11b5f5233600452602487905260445260645ffd5b612ba191975060203d602011611935576119278183610741565b955f612b5e565b612a8d565b63d9b3955760e01b5f523360045260245ffd5b5f5160206153385f395f51905f52546040516324776b7d60e21b815290602090829060049082906001600160a01b03165afa5f9181612c09575b506105db57506105db436141f2565b9091506020813d602011612c3e575b81612c2560209383610741565b810103126104925751612c3781611384565b905f612bfa565b3d9150612c18565b5f5160206153f85f395f51905f5254306001600160a01b03909116036108655763bc197c8160e01b90565b805f525f5160206152785f395f51905f5260205265ffffffffffff60405f205460a01c16905f525f5160206152785f395f51905f5260205263ffffffff60405f205460d01c160165ffffffffffff8111612ba85765ffffffffffff1690565b3d15612cfa573d90612ce182610787565b91612cef6040519384610741565b82523d5f602084013e565b606090565b413303612dd2573a612dc357612d136148fd565b612d1b6148fd565b60018060a01b03166bffffffffffffffffffffffff60a01b5f5160206152585f395f51905f525416175f5160206152585f395f51905f5255612d5b6148fd565b612d636148fd565b612d93604051612d74604082610741565b600d81526c21b937b9b9a3b7bb32b93737b960991b6020820152613802565b612d9b6148fd565b612da3613996565b612dab6139c8565b612db3613a2f565b612dbb6148fd565b610776613aa4565b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b9290612dfb612e0992604051948593602085019788610b1f565b03601f198101835282610741565b51902090565b5f5160206153f85f395f51905f52546001600160a01b031690565b90816020910312610492575190565b5f5160206153385f395f51905f5254604051632394e7a360e21b8152600481018390529190602090839060249082906001600160a01b03165afa90811561193c576105db925f92612e95575b50612e8f906133b9565b90613c25565b612e8f919250612eb39060203d602011611935576119278183610741565b9190612e85565b5f5160206153f85f395f51905f5254306001600160a01b039091160361086557565b612ee4612e0f565b336001600160a01b0390911603612f4f57612efd612e0f565b306001600160a01b0390911603612f1057565b612f1936610787565b612f266040519182610741565b3681526020810190365f83375f602036830101525190205b80612f47613c91565b03612f3e5750565b6347096e4760e01b5f523360045260245ffd5b6008811015610d385760ff600191161b90565b600452606491906008811015610d38576024525f604452565b612f97816130f3565b906010612fa383612f62565b1615612fad575090565b6331b75e4d60e01b5f52600452612fc49150610d3d565b601060445260645ffd5b612fd7816130f3565b906002612fe383612f62565b1615612fed575090565b6331b75e4d60e01b5f526004526130049150610d3d565b600260445260645ffd5b90613018826130f3565b918161302384612f62565b161561302e57505090565b6331b75e4d60e01b5f5260045261304482610d3d565b60445260645ffd5b5f5160206153b85f395f51905f5254908160801c6001600160801b03806001830116931683146130c9575f527f7c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb0360205260405f20555f5160206153b85f395f51905f52906001600160801b0382549181199060801b169116179055565b634e487b715f5260416020526024601cfd5b90816020910312610492575180151581036104925790565b6130fc81613e8f565b9061310682610d2e565b600582036131ff5761311891506120c0565b545f5160206153f85f395f51905f525461313a906001600160a01b031661213f565b604051632c258a9f60e11b815260048101839052602081602481855afa90811561193c575f916131e0575b5015613172575050600590565b604051632ab0f52960e01b81526004810192909252602090829060249082905afa90811561193c575f916131b1575b50156131ac57600790565b600290565b6131d3915060203d6020116131d9575b6131cb8183610741565b8101906130db565b5f6131a1565b503d6131c1565b6131f9915060203d6020116131d9576131cb8183610741565b5f613165565b5090565b60ff5f5160206153d85f395f51905f5254161561321c57565b638dfc202b60e01b5f5260045ffd5b916105db939160405193613240602086610741565b5f85526132a8565b93909260ff613274936105db97958752166020860152604085015260a0606085015260a08401906105a6565b9160808184039101526105a6565b909260ff6080936105db96958452166020830152604082015281606082015201906105a6565b929190926132b581612fce565b506132bf816126bd565b5f5160206153385f395f51905f5254604051630748d63560e31b81526001600160a01b03808816600483018190526024830194909452929692909160209183916044918391165afa90811561193c576133229285915f93613398575b50846140ce565b948051155f14613365575061335f7fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4938660405194859485613282565b0390a290565b61335f907fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712948760405195869586613248565b6133b291935060203d602011611935576119278183610741565b915f61331b565b65ffffffffffff90806133d85f5160206152b85f395f51905f526141b7565b9491501611155f146133f257505b6001600160d01b031690565b6133fc91506141f2565b5f5160206152b85f395f51905f5254905f82916005841161348c575b61343093505f5160206152b85f395f51905f52614e66565b8061343c57505f6133e6565b61348061344b61348792613765565b5f5160206152b85f395f51905f525f527f293b0181c8ec34cd3252e741689bdc21b70ee7a0ec76216439035a5c3718909b0190565b5460301c90565b6133e6565b919261349781614d08565b8103908111612ba857613430935f5160206152b85f395f51905f525f5265ffffffffffff8260205f2001541665ffffffffffff8516105f146134da575091613418565b9291506134e690613dce565b90613418565b7fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93604065ffffffffffff805f5160206152985f395f51905f5254169382519485521692836020820152a165ffffffffffff195f5160206152985f395f51905f525416175f5160206152985f395f51905f5255565b908051603481106135c05760131981830101516001600160b01b03191669dc8f8d908f908c9a8dc360b01b016135c05761359f9160291982019061424d565b90159182156135ad57505090565b6001600160a01b03918216911614919050565b505050600190565b60ff5f5160206153d85f395f51905f5254166135e057565b63d93c066560e01b5f5260045ffd5b6040515f5160206152f85f395f51905f5254815f61360c83612088565b808352926001811690811561369b5750600114613630575b6105db92500382610741565b505f5160206152f85f395f51905f525f90815290917f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5b81831061367f5750509060206105db92820101613624565b6020919350806001915483858801015201910190918392613667565b602092506105db94915060ff191682840152151560051b820101613624565b6040515f5160206153185f395f51905f5254815f6136d783612088565b808352926001811690811561369b57506001146136fa576105db92500382610741565b505f5160206153185f395f51905f525f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b8183106137495750509060206105db92820101613624565b6020919350806001915483858801015201910190918392613731565b5f19810191908211612ba857565b91908203918211612ba857565b5f5160206152b85f395f51905f52548061379957505f90565b805f19810111612ba8575f5160206152b85f395f51905f525f527f293b0181c8ec34cd3252e741689bdc21b70ee7a0ec76216439035a5c3718909a015460301c90565b90919061077657508051156137f357602081519101fd5b63d6bda27560e01b5f5260045ffd5b9061380b6148fd565b61381361295f565b9161381c6148fd565b8051926001600160401b0384116107625761384d846138485f5160206152f85f395f51905f5254612088565b6143e6565b602093601f81116001146138f75790613883826138979361077696975f916138ec575b508160011b915f199060031b1c19161790565b5f5160206152f85f395f51905f5255614928565b6138bf5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10055565b6138e75f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10155565b614a3b565b90508501515f613870565b5f5160206152f85f395f51905f525f52601f1981167f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d905f5b81811061397e575091600191610776969782613897969510613966575b5050811b015f5160206152f85f395f51905f5255614928565b8601515f1960f88460031b161c191690555f8061394d565b85880151835560209788019760019093019201613930565b61399e6148fd565b6139a66148fd565b5f5160206153385f395f51905f5280546001600160a01b031916611006179055565b6139d06148fd565b6139d86148fd565b7f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b463399760406001600160d01b03613a0b613780565b16613a1e6004613a19612bc0565b614b95565b5050815190815260046020820152a1565b613a376148fd565b613a3f6148fd565b6110075f5160206153f85f395f51905f52547f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b2264016040805160018060a01b0384168152846020820152a16001600160a01b031916175f5160206153f85f395f51905f5255565b613aac6148fd565b613ab46148fd565b613ac0620151806134ec565b613acc62093a80613b1d565b5f5160206153585f395f51905f5254604080519182525f60208301527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc0546191a15f5f5160206153585f395f51905f5255565b63ffffffff8116908115613ba35769ffffffff000000000000907f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e882860405f5160206152985f395f51905f52549481519063ffffffff8760301c1682526020820152a160301b169069ffffffff0000000000001916175f5160206152985f395f51905f5255565b63f1cfbf0560e01b5f525f60045260245ffd5b5f5160206153585f395f51905f525460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc0546191a15f5160206153585f395f51905f5255565b8115613c11570490565b634e487b7160e01b5f52601260045260245ffd5b5f19828209908281029182808210910303908115613c86578160641115613c7f577f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c28f5c29936064910990828211900360fe1b910360021c170290565b6011613de9565b505060649004919050565b5f5160206153b85f395f51905f5254906001600160801b0382169160801c8214613d4657815f527f7c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb0360205260405f2054916001600160801b0381165f527f7c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb036020525f60408120556001600160801b038060015f5160206153b85f395f51905f52930116166001600160801b0319825416179055565b634e487b715f5260316020526024601cfd5b949392610b4b608093610b3d613d769460a08a5260a08a0190610a58565b935f60608201520152565b9192613db060a094613da2613dbe949998979960c0875260c0870190610a58565b908582036020870152610a94565b908382036040850152610ac7565b945f606083015260808201520152565b9060018201809211612ba857565b91908201809211612ba857565b634e487b715f526020526024601cfd5b5f5160206153f85f395f51905f525490949192916001600160a01b03909116906bffffffffffffffffffffffff193060601b16823b1561049257613e575f956040519788968795869563e38335e560e01b8752189260048601613d58565b039134905af1801561193c57613e76575b50613e735f916120c0565b55565b80613e825f8093610741565b800312610492575f613e68565b613e98816120ed565b5460f881901c9060f01c60ff16613f9357613f8d57613eb6816126bd565b8015613f7957613ed2613ec7612bc0565b65ffffffffffff1690565b80911015613f7357613ee382612c71565b10613eee5750600190565b613efa6117e882614c8d565b8015613f44575b15613f0c5750600390565b613f36905f525f5160206152785f395f51905f5260205265ffffffffffff600160405f2001541690565b613f3f57600490565b600590565b50613f6e6117e8825f525f5160206153985f395f51905f5260205260405f20600181015490541090565b613f01565b50505f90565b636ad0607560e01b5f52600482905260245ffd5b50600290565b5050600790565b604290613fa56151a8565b613fad615212565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a08152613ffe60c082610741565b519020906040519161190160f01b8352600283015260228201522090565b60041115610d3857565b9190823b61404e579061403891614cce565b506140428161401c565b1591826135ad57505090565b915f92612dfb61408485946040519283916020830195630b135d3f60e11b875260248401526040604484015260648301906105a6565b51915afa614090612cd0565b816140c0575b8161409f575090565b90506140bc630b135d3f60e11b9160208082518301019101612e2a565b1490565b905060208151101590614096565b6140ef909291925f525f5160206153985f395f51905f5260205260405f2090565b916003830161411861411183839060018060a01b03165f5260205260405f2090565b5460ff1690565b61419b5761413c60ff9392614149929060018060a01b03165f5260205260405f2090565b805460ff19166001179055565b1680614160575061415b828254613ddc565b905590565b60018103614177575060010161415b828254613ddc565b60020361418c5760020161415b828254613ddc565b6303599be160e11b5f5260045ffd5b6371c6af4960e01b5f526001600160a01b03821660045260245ffd5b805490816141c95750505f905f905f90565b815f19810111612ba8575f525f199060205f2001015460019165ffffffffffff82169160301c90565b65ffffffffffff811161420a5765ffffffffffff1690565b6306dfcc6560e41b5f52603060045260245260445ffd5b908160011b9180830460021490151715612ba857565b908160041b9180830460101490151715612ba857565b91908251821180156142d4575b6142ad5761426781613dce565b8211806142b6575b61427a901515614221565b60280180602811612ba85761428f8284613773565b036142ad5761429d92614eca565b90916001600160a01b0390911690565b5050505f905f90565b50602081840101516001600160f01b03191661060f60f31b1461426f565b5081811161425a565b5f5160206152d85f395f51905f5254600160401b81101561076257600181015f5160206152d85f395f51905f52555f5160206152d85f395f51905f5254811015610c76575f5160206152d85f395f51905f525f527f417ab5363cd721cbd7a5446b963bb4274a7db26ef2bb28819834095c165ca17a0155565b818110614361575050565b5f8155600101614356565b9181811061437957505050565b610776925f5260205f209182019101614356565b8151916001600160401b03831161076257600160401b8311610762576020906143bb8484548186558561436c565b01905f5260205f205f5b8381106143d25750505050565b6001906020845194019381840155016143c5565b90601f82116143f3575050565b610776915f5160206152f85f395f51905f525f5260205f20906020601f840160051c8301931061442b575b601f0160051c0190614356565b909150819061441e565b90601f8211614442575050565b610776915f5160206154185f395f51905f525f5260205f20906020601f840160051c8301931061442b57601f0160051c0190614356565b9190601f811161448857505050565b610776925f5260205f20906020601f840160051c8301931061442b57601f0160051c0190614356565b91909182516001600160401b038111610762576144d8816144d28454612088565b84614479565b6020601f82116001146145175781906145089394955f9261450c575b50508160011b915f199060031b1c19161790565b9055565b015190505f806144f4565b601f1982169061452a845f5260205f2090565b915f5b8181106145645750958360019596971061454c575b505050811b019055565b01515f1960f88460031b161c191690555f8080614542565b9192602060018192868b01518155019401920161452d565b815191600160401b83116107625781548383558084106145d6575b5060206145a99101915f5260205f2090565b5f915b8383106145b95750505050565b60016020826145ca839451866144b1565b019201920191906145ac565b825f528360205f2091820191015b8181106145f15750614597565b806145fe60019254612088565b8061460b575b50016145e4565b601f8111831461462057505f81555b5f614604565b6146429083601f614634855f5260205f2090565b920160051c82019101614356565b5f818152602081208183555561461a565b9080518051906001600160401b03821161076257600160401b8211610762576020906146848386548188558761436c565b01835f5260205f205f5b8381106146c757505050506060816146af602060039401516001860161438d565b6146c060408201516002860161457c565b0151910155565b82516001600160a01b03168183015560209092019160010161468e565b9194939483516146fc60208601918220838587612de1565b9484518451908181148015906148f2575b80156148ea575b6148cd57505065ffffffffffff61473c61472d886120ed565b5460a01c65ffffffffffff1690565b166148b0577f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e08697986105db9761486384896147a4614779612bc0565b65ffffffffffff61479d65ffffffffffff5f5160206152985f395f51905f52541690565b9116613ddc565b896147c363ffffffff5f5160206152985f395f51905f525460301c1690565b6148416147cf886120ed565b80546001600160a01b0319166001600160a01b038b161781556148186147f4866141f2565b825465ffffffffffff60a01b191660a09190911b65ffffffffffff60a01b16178255565b6148218361517d565b815463ffffffff60d01b191660d09190911b63ffffffff60d01b16179055565b8c61485661484f8651614f74565b9285613ddc565b94604051998a998a614fbd565b0390a161486f876142dd565b5190209161487b610767565b9384526020840152604083015260608201526148ab835f525f5160206153785f395f51905f5260205260405f2090565b614653565b611821866148bd816130f3565b6331b75e4d60e01b5f5290612f75565b8451630447b05d60e41b5f5260049190915260245260445260645ffd5b508015614714565b50845181141561470d565b60ff5f5160206154385f395f51905f525460401c161561491957565b631afcd79f60e31b5f5260045ffd5b9081516001600160401b03811161076257614967816149545f5160206153185f395f51905f5254612088565b5f5160206153185f395f51905f52614479565b602092601f82116001146149a757614996929382915f9261450c5750508160011b915f199060031b1c19161790565b5f5160206153185f395f51905f5255565b5f5160206153185f395f51905f525f52601f198216937f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75915f5b868110614a235750836001959610614a0b575b505050811b015f5160206153185f395f51905f5255565b01515f1960f88460031b161c191690555f80806149f4565b919260206001819286850151815501940192016149e1565b90614a446148fd565b81516001600160401b03811161076257614a7481614a6f5f5160206154185f395f51905f5254612088565b614435565b602092601f8211600114614ab457614aa3929382915f9261450c5750508160011b915f199060031b1c19161790565b5f5160206154185f395f51905f5255565b5f5160206154185f395f51905f525f52601f198216937fda13dda7583a39a3cd73e8830529c760837228fa4683752c823b17e10548aad5915f5b868110614b305750836001959610614b18575b505050811b015f5160206154185f395f51905f5255565b01515f1960f88460031b161c191690555f8080614b01565b91926020600181928685015181550194019201614aee565b805490600160401b8210156107625760018201808255821015610c76575f9081526020908190208351939091015160301b65ffffffffffff191665ffffffffffff90931692909217910155565b5f5160206152b85f395f51905f52549192918015614c645761344b614bb991613765565b908154614bd5614bce8265ffffffffffff1690565b9160301c90565b9265ffffffffffff808416921691808311614c5557869203614c1357614c0f92509065ffffffffffff82549181199060301b169116179055565b9190565b5050614c0f90614c32614c24610778565b65ffffffffffff9092168252565b6001600160d01b03851660208201525b5f5160206152b85f395f51905f52614b48565b632520601d60e01b5f5260045ffd5b50614c8890614c74614c24610778565b6001600160d01b0384166020820152614c42565b5f9190565b805f525f5160206153985f395f51905f52602052614cb6614cb160405f20926126bd565b612e39565b600260018301549201548201809211612ba857111590565b8151919060418303614cfe57614cf79250602082015190606060408401519301515f1a90615089565b9192909190565b50505f9160029190565b60018111156105db57806001600160801b821015614e29575b614dcf614dc5614dbb614db1614da7614d9d614d8c614dd69760048a600160401b614ddb9c1015614e1c575b640100000000811015614e0f575b62010000811015614e02575b610100811015614df5575b6010811015614de8575b1015614de0575b60030260011c90565b614d96818b613c07565b0160011c90565b614d96818a613c07565b614d968189613c07565b614d968188613c07565b614d968187613c07565b614d968186613c07565b8093613c07565b821190565b900390565b60011b614d83565b60041c9160021b91614d7c565b60081c9160041b91614d72565b60101c9160081b91614d67565b60201c9160101b91614d5b565b60401c9160201b91614d4d565b5050614ddb614dd6614dcf614dc5614dbb614db1614da7614d9d614d8c614e508a60801c90565b9850600160401b9750614d219650505050505050565b91905b838210614e765750505090565b9091928083169080841860011c8201809211612ba857845f5265ffffffffffff8260205f2001541665ffffffffffff8416105f14614eb85750925b9190614e69565b939250614ec490613dce565b91614eb1565b929092614ed684613dce565b831180614f56575b614ef0614ef791949293941515614221565b5f95613ddc565b915b818310614f095750505060019190565b9092919360ff614f2a614f256020888601015160ff60f81b1690565b61510b565b1690600f8211614f4b5790614f40600192614237565b019401919290614ef9565b505f94508493505050565b50602084820101516001600160f01b03191661060f60f31b14614ede565b90614f7e82610874565b614f8b6040519182610741565b8281528092614f9c601f1991610874565b01905f5b828110614fac57505050565b806060602080938501015201614fa0565b959998969794939192614ffe93614ff092885260018060a01b031660208801526101206040880152610120870190610a58565b908582036060870152610a94565b968388036080850152815180895260208901906020808260051b8c01019401915f905b82821061505d57505050506105db969750906150449184820360a0860152610ac7565b9360c083015260e08201526101008184039101526105a6565b9091929460208061507b6001938f601f1990820301865289516105a6565b970192019201909291615021565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411615100579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa1561193c575f516001600160a01b038116156150f657905f905f90565b505f906001905f90565b5050505f9160039190565b60f81c602f811180615173575b1561512757602f190160ff1690565b6060811180615169575b15615140576056190160ff1690565b604081118061515f575b15615159576036190160ff1690565b5060ff90565b506047811061514a565b5060678110615131565b50603a8110615118565b63ffffffff81116151915763ffffffff1690565b6306dfcc6560e41b5f52602060045260245260445ffd5b6151b06135ef565b80519081156151c0576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005480156151ed5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b61521a6136ba565b805190811561522a576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015480156151ed579056fe46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e5007c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb0100d7616c8fe29c6c2fbe1d0c5bc8f2faa4c35b43746e70b24b4d532752affd01e770710421fd2cad75ad828c61aa98f2d77d423a440b67872d0f65554148e0007fd223d3380145bd26132714391e777c488a0df7ac2dd4b66419d8549fb3a600a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1033ba4977254e415696610a40ebf2258dbfa0ec6a2ff64e84bfe715ff16977cc0000d7616c8fe29c6c2fbe1d0c5bc8f2faa4c35b43746e70b24b4d532752affd007fd223d3380145bd26132714391e777c488a0df7ac2dd4b66419d8549fb3a601a1cefa0f43667ef127a258e673c94202a79b656e62899531c4376d87a7f398007c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb02cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033000d5829787b8befdbc6044ef7457d8a95c2a04bc99235349f1a212c063e59d4007c712897014dbe49c045ef1299aa2d5f9e67e48eea4403efa21f1e0f3ac0cb00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220a1f6380e919ca0f61680bb3ea5e464fcf8de5cc52e20bc52874acd3a22c9f86464736f6c634300081c0033",
}

// CrossGovernor is an auto generated Go binding around an Ethereum contract.
type CrossGovernor struct {
	abi abi.ABI
}

// NewCrossGovernor creates a new instance of CrossGovernor.
func NewCrossGovernor() *CrossGovernor {
	parsed, err := CrossGovernorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &CrossGovernor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *CrossGovernor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (crossGovernor *CrossGovernor) PackBALLOTTYPEHASH() []byte {
	enc, err := crossGovernor.abi.Pack("BALLOT_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBALLOTTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (crossGovernor *CrossGovernor) UnpackBALLOTTYPEHASH(data []byte) ([32]byte, error) {
	out, err := crossGovernor.abi.Unpack("BALLOT_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackCLOCKMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (crossGovernor *CrossGovernor) PackCLOCKMODE() []byte {
	enc, err := crossGovernor.abi.Pack("CLOCK_MODE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCLOCKMODE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (crossGovernor *CrossGovernor) UnpackCLOCKMODE(data []byte) (string, error) {
	out, err := crossGovernor.abi.Unpack("CLOCK_MODE", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackCOUNTINGMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (crossGovernor *CrossGovernor) PackCOUNTINGMODE() []byte {
	enc, err := crossGovernor.abi.Pack("COUNTING_MODE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCOUNTINGMODE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (crossGovernor *CrossGovernor) UnpackCOUNTINGMODE(data []byte) (string, error) {
	out, err := crossGovernor.abi.Unpack("COUNTING_MODE", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackEXTENDEDBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (crossGovernor *CrossGovernor) PackEXTENDEDBALLOTTYPEHASH() []byte {
	enc, err := crossGovernor.abi.Pack("EXTENDED_BALLOT_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEXTENDEDBALLOTTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (crossGovernor *CrossGovernor) UnpackEXTENDEDBALLOTTYPEHASH(data []byte) ([32]byte, error) {
	out, err := crossGovernor.abi.Unpack("EXTENDED_BALLOT_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackAddToBlackList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (crossGovernor *CrossGovernor) PackAddToBlackList(account common.Address) []byte {
	enc, err := crossGovernor.abi.Pack("addToBlackList", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (crossGovernor *CrossGovernor) PackCancel(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("cancel", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancel0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (crossGovernor *CrossGovernor) PackCancel0(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := crossGovernor.abi.Pack("cancel0", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCancel0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (crossGovernor *CrossGovernor) UnpackCancel0(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("cancel0", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCastVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (crossGovernor *CrossGovernor) PackCastVote(proposalId *big.Int, support uint8) []byte {
	enc, err := crossGovernor.abi.Pack("castVote", proposalId, support)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCastVote is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (crossGovernor *CrossGovernor) UnpackCastVote(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("castVote", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCastVoteBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (crossGovernor *CrossGovernor) PackCastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) []byte {
	enc, err := crossGovernor.abi.Pack("castVoteBySig", proposalId, support, voter, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCastVoteBySig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (crossGovernor *CrossGovernor) UnpackCastVoteBySig(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("castVoteBySig", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCastVoteWithReason is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (crossGovernor *CrossGovernor) PackCastVoteWithReason(proposalId *big.Int, support uint8, reason string) []byte {
	enc, err := crossGovernor.abi.Pack("castVoteWithReason", proposalId, support, reason)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCastVoteWithReason is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (crossGovernor *CrossGovernor) UnpackCastVoteWithReason(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("castVoteWithReason", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCastVoteWithReasonAndParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (crossGovernor *CrossGovernor) PackCastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) []byte {
	enc, err := crossGovernor.abi.Pack("castVoteWithReasonAndParams", proposalId, support, reason, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCastVoteWithReasonAndParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (crossGovernor *CrossGovernor) UnpackCastVoteWithReasonAndParams(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("castVoteWithReasonAndParams", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCastVoteWithReasonAndParamsBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (crossGovernor *CrossGovernor) PackCastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) []byte {
	enc, err := crossGovernor.abi.Pack("castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCastVoteWithReasonAndParamsBySig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (crossGovernor *CrossGovernor) UnpackCastVoteWithReasonAndParamsBySig(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("castVoteWithReasonAndParamsBySig", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackClock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (crossGovernor *CrossGovernor) PackClock() []byte {
	enc, err := crossGovernor.abi.Pack("clock")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackClock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (crossGovernor *CrossGovernor) UnpackClock(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("clock", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (crossGovernor *CrossGovernor) PackEip712Domain() []byte {
	enc, err := crossGovernor.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (crossGovernor *CrossGovernor) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := crossGovernor.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, err

}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (crossGovernor *CrossGovernor) PackExecute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := crossGovernor.abi.Pack("execute", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExecute is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (crossGovernor *CrossGovernor) UnpackExecute(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("execute", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackExecute0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (crossGovernor *CrossGovernor) PackExecute0(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("execute0", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetKeeper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x391b6f4e.
//
// Solidity: function getKeeper() view returns(address)
func (crossGovernor *CrossGovernor) PackGetKeeper() []byte {
	enc, err := crossGovernor.abi.Pack("getKeeper")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetKeeper is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x391b6f4e.
//
// Solidity: function getKeeper() view returns(address)
func (crossGovernor *CrossGovernor) UnpackGetKeeper(data []byte) (common.Address, error) {
	out, err := crossGovernor.abi.Unpack("getKeeper", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetProposalId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (crossGovernor *CrossGovernor) PackGetProposalId(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := crossGovernor.abi.Pack("getProposalId", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetProposalId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackGetProposalId(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("getProposalId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (crossGovernor *CrossGovernor) PackGetVotes(account common.Address, timepoint *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("getVotes", account, timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetVotes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackGetVotes(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("getVotes", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetVotesWithParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (crossGovernor *CrossGovernor) PackGetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) []byte {
	enc, err := crossGovernor.abi.Pack("getVotesWithParams", account, timepoint, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetVotesWithParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackGetVotesWithParams(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("getVotesWithParams", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackHasVoted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (crossGovernor *CrossGovernor) PackHasVoted(proposalId *big.Int, account common.Address) []byte {
	enc, err := crossGovernor.abi.Pack("hasVoted", proposalId, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasVoted is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (crossGovernor *CrossGovernor) UnpackHasVoted(data []byte) (bool, error) {
	out, err := crossGovernor.abi.Unpack("hasVoted", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackHashProposal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (crossGovernor *CrossGovernor) PackHashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := crossGovernor.abi.Pack("hashProposal", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHashProposal is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (crossGovernor *CrossGovernor) UnpackHashProposal(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("hashProposal", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address keeper) returns()
func (crossGovernor *CrossGovernor) PackInitialize(keeper common.Address) []byte {
	enc, err := crossGovernor.abi.Pack("initialize", keeper)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (crossGovernor *CrossGovernor) PackName() []byte {
	enc, err := crossGovernor.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (crossGovernor *CrossGovernor) UnpackName(data []byte) (string, error) {
	out, err := crossGovernor.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (crossGovernor *CrossGovernor) PackNonces(owner common.Address) []byte {
	enc, err := crossGovernor.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (crossGovernor *CrossGovernor) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := crossGovernor.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (crossGovernor *CrossGovernor) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := crossGovernor.abi.Unpack("onERC1155BatchReceived", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (crossGovernor *CrossGovernor) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := crossGovernor.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (crossGovernor *CrossGovernor) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := crossGovernor.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (crossGovernor *CrossGovernor) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := crossGovernor.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (crossGovernor *CrossGovernor) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := crossGovernor.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.
//
// Solidity: function pause() returns()
func (crossGovernor *CrossGovernor) PackPause() []byte {
	enc, err := crossGovernor.abi.Pack("pause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (crossGovernor *CrossGovernor) PackPaused() []byte {
	enc, err := crossGovernor.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (crossGovernor *CrossGovernor) UnpackPaused(data []byte) (bool, error) {
	out, err := crossGovernor.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackProposalCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (crossGovernor *CrossGovernor) PackProposalCount() []byte {
	enc, err := crossGovernor.abi.Pack("proposalCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProposalCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackProposalCount(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("proposalCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackProposalDeadline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (crossGovernor *CrossGovernor) PackProposalDeadline(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("proposalDeadline", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProposalDeadline is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackProposalDeadline(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("proposalDeadline", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackProposalDetails is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16e9eaec.
//
// Solidity: function proposalDetails(uint256 proposalId) view returns(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (crossGovernor *CrossGovernor) PackProposalDetails(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("proposalDetails", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// ProposalDetailsOutput serves as a container for the return parameters of contract
// method ProposalDetails.
type ProposalDetailsOutput struct {
	Targets         []common.Address
	Values          []*big.Int
	Calldatas       [][]byte
	DescriptionHash [32]byte
}

// UnpackProposalDetails is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x16e9eaec.
//
// Solidity: function proposalDetails(uint256 proposalId) view returns(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (crossGovernor *CrossGovernor) UnpackProposalDetails(data []byte) (ProposalDetailsOutput, error) {
	out, err := crossGovernor.abi.Unpack("proposalDetails", data)
	outstruct := new(ProposalDetailsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Targets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Calldatas = *abi.ConvertType(out[2], new([][]byte)).(*[][]byte)
	outstruct.DescriptionHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	return *outstruct, err

}

// PackProposalDetailsAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e82db94.
//
// Solidity: function proposalDetailsAt(uint256 index) view returns(uint256 proposalId, address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (crossGovernor *CrossGovernor) PackProposalDetailsAt(index *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("proposalDetailsAt", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// ProposalDetailsAtOutput serves as a container for the return parameters of contract
// method ProposalDetailsAt.
type ProposalDetailsAtOutput struct {
	ProposalId      *big.Int
	Targets         []common.Address
	Values          []*big.Int
	Calldatas       [][]byte
	DescriptionHash [32]byte
}

// UnpackProposalDetailsAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2e82db94.
//
// Solidity: function proposalDetailsAt(uint256 index) view returns(uint256 proposalId, address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (crossGovernor *CrossGovernor) UnpackProposalDetailsAt(data []byte) (ProposalDetailsAtOutput, error) {
	out, err := crossGovernor.abi.Unpack("proposalDetailsAt", data)
	outstruct := new(ProposalDetailsAtOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ProposalId = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Targets = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Calldatas = *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)
	outstruct.DescriptionHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	return *outstruct, err

}

// PackProposalEta is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (crossGovernor *CrossGovernor) PackProposalEta(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("proposalEta", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProposalEta is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackProposalEta(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("proposalEta", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackProposalNeedsQueuing is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (crossGovernor *CrossGovernor) PackProposalNeedsQueuing(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("proposalNeedsQueuing", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProposalNeedsQueuing is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (crossGovernor *CrossGovernor) UnpackProposalNeedsQueuing(data []byte) (bool, error) {
	out, err := crossGovernor.abi.Unpack("proposalNeedsQueuing", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackProposalProposer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (crossGovernor *CrossGovernor) PackProposalProposer(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("proposalProposer", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProposalProposer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (crossGovernor *CrossGovernor) UnpackProposalProposer(data []byte) (common.Address, error) {
	out, err := crossGovernor.abi.Unpack("proposalProposer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackProposalSnapshot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (crossGovernor *CrossGovernor) PackProposalSnapshot(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("proposalSnapshot", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProposalSnapshot is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackProposalSnapshot(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("proposalSnapshot", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (crossGovernor *CrossGovernor) PackProposalThreshold() []byte {
	enc, err := crossGovernor.abi.Pack("proposalThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProposalThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackProposalThreshold(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("proposalThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackProposalVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (crossGovernor *CrossGovernor) PackProposalVotes(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("proposalVotes", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// ProposalVotesOutput serves as a container for the return parameters of contract
// method ProposalVotes.
type ProposalVotesOutput struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}

// UnpackProposalVotes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (crossGovernor *CrossGovernor) UnpackProposalVotes(data []byte) (ProposalVotesOutput, error) {
	out, err := crossGovernor.abi.Unpack("proposalVotes", data)
	outstruct := new(ProposalVotesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AgainstVotes = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.ForVotes = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.AbstainVotes = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackPropose is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (crossGovernor *CrossGovernor) PackPropose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) []byte {
	enc, err := crossGovernor.abi.Pack("propose", targets, values, calldatas, description)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPropose is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (crossGovernor *CrossGovernor) UnpackPropose(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("propose", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackQueue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (crossGovernor *CrossGovernor) PackQueue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := crossGovernor.abi.Pack("queue", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackQueue is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (crossGovernor *CrossGovernor) UnpackQueue(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("queue", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackQueue0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (crossGovernor *CrossGovernor) PackQueue0(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("queue0", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackQuorum is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (crossGovernor *CrossGovernor) PackQuorum(timepoint *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("quorum", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackQuorum is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackQuorum(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("quorum", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackQuorumDenominator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (crossGovernor *CrossGovernor) PackQuorumDenominator() []byte {
	enc, err := crossGovernor.abi.Pack("quorumDenominator")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackQuorumDenominator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackQuorumDenominator(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("quorumDenominator", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (crossGovernor *CrossGovernor) PackQuorumNumerator(timepoint *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("quorumNumerator", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackQuorumNumerator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackQuorumNumerator(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("quorumNumerator", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackQuorumNumerator0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (crossGovernor *CrossGovernor) PackQuorumNumerator0() []byte {
	enc, err := crossGovernor.abi.Pack("quorumNumerator0")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackQuorumNumerator0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackQuorumNumerator0(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("quorumNumerator0", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (crossGovernor *CrossGovernor) PackRelay(target common.Address, value *big.Int, data []byte) []byte {
	enc, err := crossGovernor.abi.Pack("relay", target, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRemoveFromBlackList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (crossGovernor *CrossGovernor) PackRemoveFromBlackList(account common.Address) []byte {
	enc, err := crossGovernor.abi.Pack("removeFromBlackList", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetKeeper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x748747e6.
//
// Solidity: function setKeeper(address keeper) returns()
func (crossGovernor *CrossGovernor) PackSetKeeper(keeper common.Address) []byte {
	enc, err := crossGovernor.abi.Pack("setKeeper", keeper)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (crossGovernor *CrossGovernor) PackSetProposalThreshold(newProposalThreshold *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("setProposalThreshold", newProposalThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (crossGovernor *CrossGovernor) PackSetVotingDelay(newVotingDelay *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("setVotingDelay", newVotingDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (crossGovernor *CrossGovernor) PackSetVotingPeriod(newVotingPeriod uint32) []byte {
	enc, err := crossGovernor.abi.Pack("setVotingPeriod", newVotingPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (crossGovernor *CrossGovernor) PackState(proposalId *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("state", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (crossGovernor *CrossGovernor) UnpackState(data []byte) (uint8, error) {
	out, err := crossGovernor.abi.Unpack("state", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (crossGovernor *CrossGovernor) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := crossGovernor.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (crossGovernor *CrossGovernor) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := crossGovernor.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (crossGovernor *CrossGovernor) PackTimelock() []byte {
	enc, err := crossGovernor.abi.Pack("timelock")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTimelock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (crossGovernor *CrossGovernor) UnpackTimelock(data []byte) (common.Address, error) {
	out, err := crossGovernor.abi.Unpack("timelock", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (crossGovernor *CrossGovernor) PackToken() []byte {
	enc, err := crossGovernor.abi.Pack("token")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (crossGovernor *CrossGovernor) UnpackToken(data []byte) (common.Address, error) {
	out, err := crossGovernor.abi.Unpack("token", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (crossGovernor *CrossGovernor) PackUnpause() []byte {
	enc, err := crossGovernor.abi.Pack("unpause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (crossGovernor *CrossGovernor) PackUpdateQuorumNumerator(newQuorumNumerator *big.Int) []byte {
	enc, err := crossGovernor.abi.Pack("updateQuorumNumerator", newQuorumNumerator)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (crossGovernor *CrossGovernor) PackUpdateTimelock(newTimelock common.Address) []byte {
	enc, err := crossGovernor.abi.Pack("updateTimelock", newTimelock)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (crossGovernor *CrossGovernor) PackVersion() []byte {
	enc, err := crossGovernor.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (crossGovernor *CrossGovernor) UnpackVersion(data []byte) (string, error) {
	out, err := crossGovernor.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (crossGovernor *CrossGovernor) PackVotingDelay() []byte {
	enc, err := crossGovernor.abi.Pack("votingDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackVotingDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackVotingDelay(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("votingDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (crossGovernor *CrossGovernor) PackVotingPeriod() []byte {
	enc, err := crossGovernor.abi.Pack("votingPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackVotingPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (crossGovernor *CrossGovernor) UnpackVotingPeriod(data []byte) (*big.Int, error) {
	out, err := crossGovernor.abi.Unpack("votingPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// CrossGovernorBlackListed represents a BlackListed event raised by the CrossGovernor contract.
type CrossGovernorBlackListed struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossGovernorBlackListedEventName = "BlackListed"

// ContractEventName returns the user-defined event name.
func (CrossGovernorBlackListed) ContractEventName() string {
	return CrossGovernorBlackListedEventName
}

// UnpackBlackListedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BlackListed(address indexed account)
func (crossGovernor *CrossGovernor) UnpackBlackListedEvent(log *types.Log) (*CrossGovernorBlackListed, error) {
	event := "BlackListed"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorBlackListed)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorEIP712DomainChanged represents a EIP712DomainChanged event raised by the CrossGovernor contract.
type CrossGovernorEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const CrossGovernorEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (CrossGovernorEIP712DomainChanged) ContractEventName() string {
	return CrossGovernorEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (crossGovernor *CrossGovernor) UnpackEIP712DomainChangedEvent(log *types.Log) (*CrossGovernorEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorInitialized represents a Initialized event raised by the CrossGovernor contract.
type CrossGovernorInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossGovernorInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (CrossGovernorInitialized) ContractEventName() string {
	return CrossGovernorInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (crossGovernor *CrossGovernor) UnpackInitializedEvent(log *types.Log) (*CrossGovernorInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorInitialized)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorKeeperChanged represents a KeeperChanged event raised by the CrossGovernor contract.
type CrossGovernorKeeperChanged struct {
	PreviousKeeper common.Address
	NewKeeper      common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const CrossGovernorKeeperChangedEventName = "KeeperChanged"

// ContractEventName returns the user-defined event name.
func (CrossGovernorKeeperChanged) ContractEventName() string {
	return CrossGovernorKeeperChangedEventName
}

// UnpackKeeperChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event KeeperChanged(address indexed previousKeeper, address indexed newKeeper)
func (crossGovernor *CrossGovernor) UnpackKeeperChangedEvent(log *types.Log) (*CrossGovernorKeeperChanged, error) {
	event := "KeeperChanged"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorKeeperChanged)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorPaused represents a Paused event raised by the CrossGovernor contract.
type CrossGovernorPaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossGovernorPausedEventName = "Paused"

// ContractEventName returns the user-defined event name.
func (CrossGovernorPaused) ContractEventName() string {
	return CrossGovernorPausedEventName
}

// UnpackPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Paused(address account)
func (crossGovernor *CrossGovernor) UnpackPausedEvent(log *types.Log) (*CrossGovernorPaused, error) {
	event := "Paused"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorPaused)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorProposalCanceled represents a ProposalCanceled event raised by the CrossGovernor contract.
type CrossGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const CrossGovernorProposalCanceledEventName = "ProposalCanceled"

// ContractEventName returns the user-defined event name.
func (CrossGovernorProposalCanceled) ContractEventName() string {
	return CrossGovernorProposalCanceledEventName
}

// UnpackProposalCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (crossGovernor *CrossGovernor) UnpackProposalCanceledEvent(log *types.Log) (*CrossGovernorProposalCanceled, error) {
	event := "ProposalCanceled"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorProposalCanceled)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorProposalCreated represents a ProposalCreated event raised by the CrossGovernor contract.
type CrossGovernorProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         *types.Log // Blockchain specific contextual infos
}

const CrossGovernorProposalCreatedEventName = "ProposalCreated"

// ContractEventName returns the user-defined event name.
func (CrossGovernorProposalCreated) ContractEventName() string {
	return CrossGovernorProposalCreatedEventName
}

// UnpackProposalCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (crossGovernor *CrossGovernor) UnpackProposalCreatedEvent(log *types.Log) (*CrossGovernorProposalCreated, error) {
	event := "ProposalCreated"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorProposalCreated)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorProposalExecuted represents a ProposalExecuted event raised by the CrossGovernor contract.
type CrossGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const CrossGovernorProposalExecutedEventName = "ProposalExecuted"

// ContractEventName returns the user-defined event name.
func (CrossGovernorProposalExecuted) ContractEventName() string {
	return CrossGovernorProposalExecutedEventName
}

// UnpackProposalExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (crossGovernor *CrossGovernor) UnpackProposalExecutedEvent(log *types.Log) (*CrossGovernorProposalExecuted, error) {
	event := "ProposalExecuted"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorProposalExecuted)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorProposalQueued represents a ProposalQueued event raised by the CrossGovernor contract.
type CrossGovernorProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const CrossGovernorProposalQueuedEventName = "ProposalQueued"

// ContractEventName returns the user-defined event name.
func (CrossGovernorProposalQueued) ContractEventName() string {
	return CrossGovernorProposalQueuedEventName
}

// UnpackProposalQueuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (crossGovernor *CrossGovernor) UnpackProposalQueuedEvent(log *types.Log) (*CrossGovernorProposalQueued, error) {
	event := "ProposalQueued"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorProposalQueued)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorProposalThresholdSet represents a ProposalThresholdSet event raised by the CrossGovernor contract.
type CrossGovernorProposalThresholdSet struct {
	OldProposalThreshold *big.Int
	NewProposalThreshold *big.Int
	Raw                  *types.Log // Blockchain specific contextual infos
}

const CrossGovernorProposalThresholdSetEventName = "ProposalThresholdSet"

// ContractEventName returns the user-defined event name.
func (CrossGovernorProposalThresholdSet) ContractEventName() string {
	return CrossGovernorProposalThresholdSetEventName
}

// UnpackProposalThresholdSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (crossGovernor *CrossGovernor) UnpackProposalThresholdSetEvent(log *types.Log) (*CrossGovernorProposalThresholdSet, error) {
	event := "ProposalThresholdSet"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorProposalThresholdSet)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the CrossGovernor contract.
type CrossGovernorQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                *types.Log // Blockchain specific contextual infos
}

const CrossGovernorQuorumNumeratorUpdatedEventName = "QuorumNumeratorUpdated"

// ContractEventName returns the user-defined event name.
func (CrossGovernorQuorumNumeratorUpdated) ContractEventName() string {
	return CrossGovernorQuorumNumeratorUpdatedEventName
}

// UnpackQuorumNumeratorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (crossGovernor *CrossGovernor) UnpackQuorumNumeratorUpdatedEvent(log *types.Log) (*CrossGovernorQuorumNumeratorUpdated, error) {
	event := "QuorumNumeratorUpdated"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorQuorumNumeratorUpdated)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorTimelockChange represents a TimelockChange event raised by the CrossGovernor contract.
type CrossGovernorTimelockChange struct {
	OldTimelock common.Address
	NewTimelock common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const CrossGovernorTimelockChangeEventName = "TimelockChange"

// ContractEventName returns the user-defined event name.
func (CrossGovernorTimelockChange) ContractEventName() string {
	return CrossGovernorTimelockChangeEventName
}

// UnpackTimelockChangeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (crossGovernor *CrossGovernor) UnpackTimelockChangeEvent(log *types.Log) (*CrossGovernorTimelockChange, error) {
	event := "TimelockChange"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorTimelockChange)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorUnBlackListed represents a UnBlackListed event raised by the CrossGovernor contract.
type CrossGovernorUnBlackListed struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossGovernorUnBlackListedEventName = "UnBlackListed"

// ContractEventName returns the user-defined event name.
func (CrossGovernorUnBlackListed) ContractEventName() string {
	return CrossGovernorUnBlackListedEventName
}

// UnpackUnBlackListedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UnBlackListed(address indexed account)
func (crossGovernor *CrossGovernor) UnpackUnBlackListedEvent(log *types.Log) (*CrossGovernorUnBlackListed, error) {
	event := "UnBlackListed"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorUnBlackListed)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorUnpaused represents a Unpaused event raised by the CrossGovernor contract.
type CrossGovernorUnpaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const CrossGovernorUnpausedEventName = "Unpaused"

// ContractEventName returns the user-defined event name.
func (CrossGovernorUnpaused) ContractEventName() string {
	return CrossGovernorUnpausedEventName
}

// UnpackUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpaused(address account)
func (crossGovernor *CrossGovernor) UnpackUnpausedEvent(log *types.Log) (*CrossGovernorUnpaused, error) {
	event := "Unpaused"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorUnpaused)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorVoteCast represents a VoteCast event raised by the CrossGovernor contract.
type CrossGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        *types.Log // Blockchain specific contextual infos
}

const CrossGovernorVoteCastEventName = "VoteCast"

// ContractEventName returns the user-defined event name.
func (CrossGovernorVoteCast) ContractEventName() string {
	return CrossGovernorVoteCastEventName
}

// UnpackVoteCastEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (crossGovernor *CrossGovernor) UnpackVoteCastEvent(log *types.Log) (*CrossGovernorVoteCast, error) {
	event := "VoteCast"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorVoteCast)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the CrossGovernor contract.
type CrossGovernorVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const CrossGovernorVoteCastWithParamsEventName = "VoteCastWithParams"

// ContractEventName returns the user-defined event name.
func (CrossGovernorVoteCastWithParams) ContractEventName() string {
	return CrossGovernorVoteCastWithParamsEventName
}

// UnpackVoteCastWithParamsEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (crossGovernor *CrossGovernor) UnpackVoteCastWithParamsEvent(log *types.Log) (*CrossGovernorVoteCastWithParams, error) {
	event := "VoteCastWithParams"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorVoteCastWithParams)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorVotingDelaySet represents a VotingDelaySet event raised by the CrossGovernor contract.
type CrossGovernorVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const CrossGovernorVotingDelaySetEventName = "VotingDelaySet"

// ContractEventName returns the user-defined event name.
func (CrossGovernorVotingDelaySet) ContractEventName() string {
	return CrossGovernorVotingDelaySetEventName
}

// UnpackVotingDelaySetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (crossGovernor *CrossGovernor) UnpackVotingDelaySetEvent(log *types.Log) (*CrossGovernorVotingDelaySet, error) {
	event := "VotingDelaySet"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorVotingDelaySet)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CrossGovernorVotingPeriodSet represents a VotingPeriodSet event raised by the CrossGovernor contract.
type CrossGovernorVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const CrossGovernorVotingPeriodSetEventName = "VotingPeriodSet"

// ContractEventName returns the user-defined event name.
func (CrossGovernorVotingPeriodSet) ContractEventName() string {
	return CrossGovernorVotingPeriodSetEventName
}

// UnpackVotingPeriodSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (crossGovernor *CrossGovernor) UnpackVotingPeriodSetEvent(log *types.Log) (*CrossGovernorVotingPeriodSet, error) {
	event := "VotingPeriodSet"
	if log.Topics[0] != crossGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CrossGovernorVotingPeriodSet)
	if len(log.Data) > 0 {
		if err := crossGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range crossGovernor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (crossGovernor *CrossGovernor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["CheckpointUnorderedInsertion"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackCheckpointUnorderedInsertionError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["EnforcedPause"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackEnforcedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["ExpectedPause"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackExpectedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorAlreadyCastVote"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorAlreadyCastVoteError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorAlreadyQueuedProposal"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorAlreadyQueuedProposalError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorDisabledDeposit"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorDisabledDepositError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorInsufficientProposerVotes"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorInsufficientProposerVotesError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorInvalidProposalLength"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorInvalidProposalLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorInvalidQuorumFraction"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorInvalidQuorumFractionError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorInvalidSignature"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorInvalidVoteParams"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorInvalidVoteParamsError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorInvalidVoteType"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorInvalidVoteTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorInvalidVotingPeriod"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorInvalidVotingPeriodError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorNonexistentProposal"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorNonexistentProposalError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorNotQueuedProposal"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorNotQueuedProposalError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorOnlyExecutor"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorQueueNotImplemented"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorQueueNotImplementedError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorRestrictedProposer"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorRestrictedProposerError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorUnableToCancel"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorUnableToCancelError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["GovernorUnexpectedProposalState"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackGovernorUnexpectedProposalStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["InBlackList"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackInBlackListError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["OnlyKeeper"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackOnlyKeeperError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], crossGovernor.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return crossGovernor.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// CrossGovernorCheckpointUnorderedInsertion represents a CheckpointUnorderedInsertion error raised by the CrossGovernor contract.
type CrossGovernorCheckpointUnorderedInsertion struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CheckpointUnorderedInsertion()
func CrossGovernorCheckpointUnorderedInsertionErrorID() common.Hash {
	return common.HexToHash("0x2520601d9d60b717c34a36ad270857824c5a1ebbfd08ff39aba6930089495cfa")
}

// UnpackCheckpointUnorderedInsertionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CheckpointUnorderedInsertion()
func (crossGovernor *CrossGovernor) UnpackCheckpointUnorderedInsertionError(raw []byte) (*CrossGovernorCheckpointUnorderedInsertion, error) {
	out := new(CrossGovernorCheckpointUnorderedInsertion)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "CheckpointUnorderedInsertion", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorEnforcedPause represents a EnforcedPause error raised by the CrossGovernor contract.
type CrossGovernorEnforcedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EnforcedPause()
func CrossGovernorEnforcedPauseErrorID() common.Hash {
	return common.HexToHash("0xd93c0665d6c96d04a8f174024fc4ddd66c250604aff22bbec808de86dd3637e3")
}

// UnpackEnforcedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EnforcedPause()
func (crossGovernor *CrossGovernor) UnpackEnforcedPauseError(raw []byte) (*CrossGovernorEnforcedPause, error) {
	out := new(CrossGovernorEnforcedPause)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "EnforcedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorExpectedPause represents a ExpectedPause error raised by the CrossGovernor contract.
type CrossGovernorExpectedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExpectedPause()
func CrossGovernorExpectedPauseErrorID() common.Hash {
	return common.HexToHash("0x8dfc202bcfe9a735b559bee70674422512bc5c30f687046ae8778315fb81da44")
}

// UnpackExpectedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExpectedPause()
func (crossGovernor *CrossGovernor) UnpackExpectedPauseError(raw []byte) (*CrossGovernorExpectedPause, error) {
	out := new(CrossGovernorExpectedPause)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "ExpectedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorFailedCall represents a FailedCall error raised by the CrossGovernor contract.
type CrossGovernorFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func CrossGovernorFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (crossGovernor *CrossGovernor) UnpackFailedCallError(raw []byte) (*CrossGovernorFailedCall, error) {
	out := new(CrossGovernorFailedCall)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorAlreadyCastVote represents a GovernorAlreadyCastVote error raised by the CrossGovernor contract.
type CrossGovernorGovernorAlreadyCastVote struct {
	Voter common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorAlreadyCastVote(address voter)
func CrossGovernorGovernorAlreadyCastVoteErrorID() common.Hash {
	return common.HexToHash("0x71c6af4932ed543cdb181fcbb800f4b9940a2ccceeaee5e6e141de8c50856ede")
}

// UnpackGovernorAlreadyCastVoteError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorAlreadyCastVote(address voter)
func (crossGovernor *CrossGovernor) UnpackGovernorAlreadyCastVoteError(raw []byte) (*CrossGovernorGovernorAlreadyCastVote, error) {
	out := new(CrossGovernorGovernorAlreadyCastVote)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorAlreadyCastVote", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorAlreadyQueuedProposal represents a GovernorAlreadyQueuedProposal error raised by the CrossGovernor contract.
type CrossGovernorGovernorAlreadyQueuedProposal struct {
	ProposalId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorAlreadyQueuedProposal(uint256 proposalId)
func CrossGovernorGovernorAlreadyQueuedProposalErrorID() common.Hash {
	return common.HexToHash("0xf20e7d374e58691196b2e081c7753efc94203ab3520c58efe153076e279fde0a")
}

// UnpackGovernorAlreadyQueuedProposalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorAlreadyQueuedProposal(uint256 proposalId)
func (crossGovernor *CrossGovernor) UnpackGovernorAlreadyQueuedProposalError(raw []byte) (*CrossGovernorGovernorAlreadyQueuedProposal, error) {
	out := new(CrossGovernorGovernorAlreadyQueuedProposal)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorAlreadyQueuedProposal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorDisabledDeposit represents a GovernorDisabledDeposit error raised by the CrossGovernor contract.
type CrossGovernorGovernorDisabledDeposit struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorDisabledDeposit()
func CrossGovernorGovernorDisabledDepositErrorID() common.Hash {
	return common.HexToHash("0xe90a651e5fdea7022846d10b5f36994c22e1f46db4b5021aa3e26ad83b24bfd8")
}

// UnpackGovernorDisabledDepositError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorDisabledDeposit()
func (crossGovernor *CrossGovernor) UnpackGovernorDisabledDepositError(raw []byte) (*CrossGovernorGovernorDisabledDeposit, error) {
	out := new(CrossGovernorGovernorDisabledDeposit)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorDisabledDeposit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorInsufficientProposerVotes represents a GovernorInsufficientProposerVotes error raised by the CrossGovernor contract.
type CrossGovernorGovernorInsufficientProposerVotes struct {
	Proposer  common.Address
	Votes     *big.Int
	Threshold *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInsufficientProposerVotes(address proposer, uint256 votes, uint256 threshold)
func CrossGovernorGovernorInsufficientProposerVotesErrorID() common.Hash {
	return common.HexToHash("0xc242ee16ab08d11dbce60e744efdbd91b4e07ac4c074d993992519795a6324d0")
}

// UnpackGovernorInsufficientProposerVotesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInsufficientProposerVotes(address proposer, uint256 votes, uint256 threshold)
func (crossGovernor *CrossGovernor) UnpackGovernorInsufficientProposerVotesError(raw []byte) (*CrossGovernorGovernorInsufficientProposerVotes, error) {
	out := new(CrossGovernorGovernorInsufficientProposerVotes)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorInsufficientProposerVotes", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorInvalidProposalLength represents a GovernorInvalidProposalLength error raised by the CrossGovernor contract.
type CrossGovernorGovernorInvalidProposalLength struct {
	Targets   *big.Int
	Calldatas *big.Int
	Values    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidProposalLength(uint256 targets, uint256 calldatas, uint256 values)
func CrossGovernorGovernorInvalidProposalLengthErrorID() common.Hash {
	return common.HexToHash("0x447b05d0c41e339e22932be48ca2091a47f3c39df25e2152ad11a8729753b2b4")
}

// UnpackGovernorInvalidProposalLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidProposalLength(uint256 targets, uint256 calldatas, uint256 values)
func (crossGovernor *CrossGovernor) UnpackGovernorInvalidProposalLengthError(raw []byte) (*CrossGovernorGovernorInvalidProposalLength, error) {
	out := new(CrossGovernorGovernorInvalidProposalLength)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidProposalLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorInvalidQuorumFraction represents a GovernorInvalidQuorumFraction error raised by the CrossGovernor contract.
type CrossGovernorGovernorInvalidQuorumFraction struct {
	QuorumNumerator   *big.Int
	QuorumDenominator *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidQuorumFraction(uint256 quorumNumerator, uint256 quorumDenominator)
func CrossGovernorGovernorInvalidQuorumFractionErrorID() common.Hash {
	return common.HexToHash("0x243e5445050913bb1c3de7a2f82eba0c3b0b8a55c2aacf660392fa35087a1919")
}

// UnpackGovernorInvalidQuorumFractionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidQuorumFraction(uint256 quorumNumerator, uint256 quorumDenominator)
func (crossGovernor *CrossGovernor) UnpackGovernorInvalidQuorumFractionError(raw []byte) (*CrossGovernorGovernorInvalidQuorumFraction, error) {
	out := new(CrossGovernorGovernorInvalidQuorumFraction)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidQuorumFraction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorInvalidSignature represents a GovernorInvalidSignature error raised by the CrossGovernor contract.
type CrossGovernorGovernorInvalidSignature struct {
	Voter common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidSignature(address voter)
func CrossGovernorGovernorInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0x94ab6c07905fb25046d2809e4563b09690f891c9495bfe19551d602e7eddbb1b")
}

// UnpackGovernorInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidSignature(address voter)
func (crossGovernor *CrossGovernor) UnpackGovernorInvalidSignatureError(raw []byte) (*CrossGovernorGovernorInvalidSignature, error) {
	out := new(CrossGovernorGovernorInvalidSignature)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorInvalidVoteParams represents a GovernorInvalidVoteParams error raised by the CrossGovernor contract.
type CrossGovernorGovernorInvalidVoteParams struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidVoteParams()
func CrossGovernorGovernorInvalidVoteParamsErrorID() common.Hash {
	return common.HexToHash("0x867db7717d0cc803be5726127d33cc17ae07776705d8ba6659af8aaa5b418dd8")
}

// UnpackGovernorInvalidVoteParamsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidVoteParams()
func (crossGovernor *CrossGovernor) UnpackGovernorInvalidVoteParamsError(raw []byte) (*CrossGovernorGovernorInvalidVoteParams, error) {
	out := new(CrossGovernorGovernorInvalidVoteParams)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidVoteParams", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorInvalidVoteType represents a GovernorInvalidVoteType error raised by the CrossGovernor contract.
type CrossGovernorGovernorInvalidVoteType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidVoteType()
func CrossGovernorGovernorInvalidVoteTypeErrorID() common.Hash {
	return common.HexToHash("0x06b337c26289d63178b4d4ed5fd513f38a1d8832d0edd309ef07bfc9ba5caf49")
}

// UnpackGovernorInvalidVoteTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidVoteType()
func (crossGovernor *CrossGovernor) UnpackGovernorInvalidVoteTypeError(raw []byte) (*CrossGovernorGovernorInvalidVoteType, error) {
	out := new(CrossGovernorGovernorInvalidVoteType)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidVoteType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorInvalidVotingPeriod represents a GovernorInvalidVotingPeriod error raised by the CrossGovernor contract.
type CrossGovernorGovernorInvalidVotingPeriod struct {
	VotingPeriod *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidVotingPeriod(uint256 votingPeriod)
func CrossGovernorGovernorInvalidVotingPeriodErrorID() common.Hash {
	return common.HexToHash("0xf1cfbf057db43f9730bc42a52728d66da9151a5c6929758ee824e299f82f5689")
}

// UnpackGovernorInvalidVotingPeriodError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidVotingPeriod(uint256 votingPeriod)
func (crossGovernor *CrossGovernor) UnpackGovernorInvalidVotingPeriodError(raw []byte) (*CrossGovernorGovernorInvalidVotingPeriod, error) {
	out := new(CrossGovernorGovernorInvalidVotingPeriod)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidVotingPeriod", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorNonexistentProposal represents a GovernorNonexistentProposal error raised by the CrossGovernor contract.
type CrossGovernorGovernorNonexistentProposal struct {
	ProposalId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorNonexistentProposal(uint256 proposalId)
func CrossGovernorGovernorNonexistentProposalErrorID() common.Hash {
	return common.HexToHash("0x6ad06075316ea071ccae80931b756598be5aad3433b2c47b38607a8eec344a70")
}

// UnpackGovernorNonexistentProposalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorNonexistentProposal(uint256 proposalId)
func (crossGovernor *CrossGovernor) UnpackGovernorNonexistentProposalError(raw []byte) (*CrossGovernorGovernorNonexistentProposal, error) {
	out := new(CrossGovernorGovernorNonexistentProposal)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorNonexistentProposal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorNotQueuedProposal represents a GovernorNotQueuedProposal error raised by the CrossGovernor contract.
type CrossGovernorGovernorNotQueuedProposal struct {
	ProposalId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorNotQueuedProposal(uint256 proposalId)
func CrossGovernorGovernorNotQueuedProposalErrorID() common.Hash {
	return common.HexToHash("0xd5ddb8255ec3d5fb4ee2dd5d919eb1db6a2f1e4420bb74c3741830500cfb6a4f")
}

// UnpackGovernorNotQueuedProposalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorNotQueuedProposal(uint256 proposalId)
func (crossGovernor *CrossGovernor) UnpackGovernorNotQueuedProposalError(raw []byte) (*CrossGovernorGovernorNotQueuedProposal, error) {
	out := new(CrossGovernorGovernorNotQueuedProposal)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorNotQueuedProposal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorOnlyExecutor represents a GovernorOnlyExecutor error raised by the CrossGovernor contract.
type CrossGovernorGovernorOnlyExecutor struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorOnlyExecutor(address account)
func CrossGovernorGovernorOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x47096e4749c231af946d5efc74a7fd817371e756031e98787f18bf70aaa7753c")
}

// UnpackGovernorOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorOnlyExecutor(address account)
func (crossGovernor *CrossGovernor) UnpackGovernorOnlyExecutorError(raw []byte) (*CrossGovernorGovernorOnlyExecutor, error) {
	out := new(CrossGovernorGovernorOnlyExecutor)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorOnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorQueueNotImplemented represents a GovernorQueueNotImplemented error raised by the CrossGovernor contract.
type CrossGovernorGovernorQueueNotImplemented struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorQueueNotImplemented()
func CrossGovernorGovernorQueueNotImplementedErrorID() common.Hash {
	return common.HexToHash("0x90884a46490684db0a73766419939e5ace793ae0f80195a286e159115c6628a0")
}

// UnpackGovernorQueueNotImplementedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorQueueNotImplemented()
func (crossGovernor *CrossGovernor) UnpackGovernorQueueNotImplementedError(raw []byte) (*CrossGovernorGovernorQueueNotImplemented, error) {
	out := new(CrossGovernorGovernorQueueNotImplemented)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorQueueNotImplemented", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorRestrictedProposer represents a GovernorRestrictedProposer error raised by the CrossGovernor contract.
type CrossGovernorGovernorRestrictedProposer struct {
	Proposer common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorRestrictedProposer(address proposer)
func CrossGovernorGovernorRestrictedProposerErrorID() common.Hash {
	return common.HexToHash("0xd9b395579c6f1566cc7608394c53136f366f7fa719d01a941bee075ef8c704f4")
}

// UnpackGovernorRestrictedProposerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorRestrictedProposer(address proposer)
func (crossGovernor *CrossGovernor) UnpackGovernorRestrictedProposerError(raw []byte) (*CrossGovernorGovernorRestrictedProposer, error) {
	out := new(CrossGovernorGovernorRestrictedProposer)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorRestrictedProposer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorUnableToCancel represents a GovernorUnableToCancel error raised by the CrossGovernor contract.
type CrossGovernorGovernorUnableToCancel struct {
	ProposalId *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorUnableToCancel(uint256 proposalId, address account)
func CrossGovernorGovernorUnableToCancelErrorID() common.Hash {
	return common.HexToHash("0x8fe5d8a9809b4b1a3569a8d98ce25e21fb956a89b6b9a2741d15bc699f46ef07")
}

// UnpackGovernorUnableToCancelError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorUnableToCancel(uint256 proposalId, address account)
func (crossGovernor *CrossGovernor) UnpackGovernorUnableToCancelError(raw []byte) (*CrossGovernorGovernorUnableToCancel, error) {
	out := new(CrossGovernorGovernorUnableToCancel)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorUnableToCancel", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorGovernorUnexpectedProposalState represents a GovernorUnexpectedProposalState error raised by the CrossGovernor contract.
type CrossGovernorGovernorUnexpectedProposalState struct {
	ProposalId     *big.Int
	Current        uint8
	ExpectedStates [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorUnexpectedProposalState(uint256 proposalId, uint8 current, bytes32 expectedStates)
func CrossGovernorGovernorUnexpectedProposalStateErrorID() common.Hash {
	return common.HexToHash("0x31b75e4d4f8317c390cf01cbc79dfe4f67ce2d27f65a099074fdc67f00f76908")
}

// UnpackGovernorUnexpectedProposalStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorUnexpectedProposalState(uint256 proposalId, uint8 current, bytes32 expectedStates)
func (crossGovernor *CrossGovernor) UnpackGovernorUnexpectedProposalStateError(raw []byte) (*CrossGovernorGovernorUnexpectedProposalState, error) {
	out := new(CrossGovernorGovernorUnexpectedProposalState)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "GovernorUnexpectedProposalState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorInBlackList represents a InBlackList error raised by the CrossGovernor contract.
type CrossGovernorInBlackList struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InBlackList(address account)
func CrossGovernorInBlackListErrorID() common.Hash {
	return common.HexToHash("0x033788ffabe13708428b49b6cedeba51975d173696096da94c2097fc14da662e")
}

// UnpackInBlackListError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InBlackList(address account)
func (crossGovernor *CrossGovernor) UnpackInBlackListError(raw []byte) (*CrossGovernorInBlackList, error) {
	out := new(CrossGovernorInBlackList)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "InBlackList", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorInvalidAccountNonce represents a InvalidAccountNonce error raised by the CrossGovernor contract.
type CrossGovernorInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func CrossGovernorInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (crossGovernor *CrossGovernor) UnpackInvalidAccountNonceError(raw []byte) (*CrossGovernorInvalidAccountNonce, error) {
	out := new(CrossGovernorInvalidAccountNonce)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorInvalidInitialization represents a InvalidInitialization error raised by the CrossGovernor contract.
type CrossGovernorInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func CrossGovernorInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (crossGovernor *CrossGovernor) UnpackInvalidInitializationError(raw []byte) (*CrossGovernorInvalidInitialization, error) {
	out := new(CrossGovernorInvalidInitialization)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorInvalidValue represents a InvalidValue error raised by the CrossGovernor contract.
type CrossGovernorInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func CrossGovernorInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (crossGovernor *CrossGovernor) UnpackInvalidValueError(raw []byte) (*CrossGovernorInvalidValue, error) {
	out := new(CrossGovernorInvalidValue)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorNotInitializing represents a NotInitializing error raised by the CrossGovernor contract.
type CrossGovernorNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func CrossGovernorNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (crossGovernor *CrossGovernor) UnpackNotInitializingError(raw []byte) (*CrossGovernorNotInitializing, error) {
	out := new(CrossGovernorNotInitializing)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorOnlyCoinbase represents a OnlyCoinbase error raised by the CrossGovernor contract.
type CrossGovernorOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func CrossGovernorOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (crossGovernor *CrossGovernor) UnpackOnlyCoinbaseError(raw []byte) (*CrossGovernorOnlyCoinbase, error) {
	out := new(CrossGovernorOnlyCoinbase)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorOnlyKeeper represents a OnlyKeeper error raised by the CrossGovernor contract.
type CrossGovernorOnlyKeeper struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyKeeper()
func CrossGovernorOnlyKeeperErrorID() common.Hash {
	return common.HexToHash("0xc60eb3352805112c61799a78f842543bbf71f1a5e9cd075fb2e23066f7db6914")
}

// UnpackOnlyKeeperError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyKeeper()
func (crossGovernor *CrossGovernor) UnpackOnlyKeeperError(raw []byte) (*CrossGovernorOnlyKeeper, error) {
	out := new(CrossGovernorOnlyKeeper)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "OnlyKeeper", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorOnlySystemContract represents a OnlySystemContract error raised by the CrossGovernor contract.
type CrossGovernorOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func CrossGovernorOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (crossGovernor *CrossGovernor) UnpackOnlySystemContractError(raw []byte) (*CrossGovernorOnlySystemContract, error) {
	out := new(CrossGovernorOnlySystemContract)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the CrossGovernor contract.
type CrossGovernorOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func CrossGovernorOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (crossGovernor *CrossGovernor) UnpackOnlyZeroGasPriceError(raw []byte) (*CrossGovernorOnlyZeroGasPrice, error) {
	out := new(CrossGovernorOnlyZeroGasPrice)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the CrossGovernor contract.
type CrossGovernorReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func CrossGovernorReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (crossGovernor *CrossGovernor) UnpackReentrancyGuardReentrantCallError(raw []byte) (*CrossGovernorReentrancyGuardReentrantCall, error) {
	out := new(CrossGovernorReentrancyGuardReentrantCall)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CrossGovernorSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the CrossGovernor contract.
type CrossGovernorSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func CrossGovernorSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (crossGovernor *CrossGovernor) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*CrossGovernorSafeCastOverflowedUintDowncast, error) {
	out := new(CrossGovernorSafeCastOverflowedUintDowncast)
	if err := crossGovernor.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}
