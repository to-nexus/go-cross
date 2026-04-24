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

// ValidatorSlashMetaData contains all meta data concerning the ValidatorSlash contract.
var ValidatorSlashMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"function\",\"name\":\"getSlashedValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mitigate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOffline\",\"inputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	ID:         "ValidatorSlash",
	BinRuntime: "0x6080806040526004361015610012575f80fd5b5f3560e01c9081632ea22d0b14610533575080634ac9302e146104ce5780635b19c262146104b15780637ac6010d146104025780638129fc1c146102bc5780638d02d4be1461028e578063cec26e14146100935763e2b33d5b14610074575f80fd5b3461008f575f36600319011261008f57602060405160048152f35b5f80fd5b3461008f57602036600319011261008f5760043567ffffffffffffffff811161008f573660238201121561008f5780600401359067ffffffffffffffff821161008f573660248360051b8301011161008f5741331480610286575b15610277575f5b82811015610275576024600582901b830101356001600160a01b038116919082900361008f57604051631015428760e21b8152600481018390526020816024816110015afa90811561020a575f9161023b575b501561023257815f525f60205260405f2080549043821461022357600191610215576101738461080c565b5081808201555b438155018054837f17bddadfd7ec8898c3b9eadd0cf5ae77ba8d5df3a50e96ab86ec2dd711aa8fbb6020604051848152a260035411156101c1575b50600191505b016100f5565b5f90556110023b1561008f576040519163111a439760e21b835260048301525f82602481836110025af191821561020a57600192156101b5575f610204916105b1565b846101b5565b6040513d5f823e3d90fd5b81810182815401905561017a565b638aaaa79d60e01b5f5260045ffd5b600191506101bb565b90506020813d821161026d575b81610255602093836105b1565b8101031261008f5751801515810361008f5785610148565b3d9150610248565b005b633d45932f60e11b5f5260045ffd5b503a156100ee565b3461008f575f36600319011261008f57413314806102b4575b15610277576102756105d3565b503a156102a7565b3461008f575f36600319011261008f575f5160206108675f395f51905f525460ff8160401c16159067ffffffffffffffff8116801590816103fa575b60011490816103f0575b1590816103e7575b506103d85767ffffffffffffffff1981166001175f5160206108675f395f51905f5255816103ac575b50413314806103a4575b1561027757603260035561034d57005b68ff0000000000000000195f5160206108675f395f51905f5254165f5160206108675f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b503a1561033d565b68ffffffffffffffffff191668010000000000000001175f5160206108675f395f51905f525581610333565b63f92ee8a960e01b5f5260045ffd5b9050158361030a565b303b159150610302565b8391506102f8565b3461008f57602036600319011261008f57600435611013330361049c5760048110610458576020817f44cf37fc8d2185243e5927aa52de715bc8ff0d81f4785d5d391bc8178f1a2cdc92600355604051908152a1005b60849060405190632c648cf160e01b82526040600483015260156044830152741bd9999b1a5b9954db185cda151a1c995cda1bdb19605a1b60648301526024820152fd5b630f22c43960e41b5f5261101360045260245ffd5b3461008f575f36600319011261008f576020600354604051908152f35b3461008f575f36600319011261008f576104e66106ac565b6040518091602082016020835281518091526020604084019201905f5b818110610511575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610503565b3461008f57602036600319011261008f576004356001600160a01b0381169081900361008f575f525f60205260405f206040820182811067ffffffffffffffff82111761059d57604092602091845260018354938483520154918291015282519182526020820152f35b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761059d57604052565b600154156106aa5760035460021c6105e96106ac565b905f5b82518110156106a557600190818060a01b0360208260051b8601015116805f525f6020528260405f2001848154115f1461065957602081867f5bd07e8a6722365853b9a8e0af5f9e09a944cbdfc7e86bb21c93cc71486036bd935403809155604051908152a25b016105ec565b50805f525f6020525f836040822082815501556106758161071d565b507f5bd07e8a6722365853b9a8e0af5f9e09a944cbdfc7e86bb21c93cc71486036bd60206040515f8152a2610653565b505050565b565b60405190600154808352826020810160015f5260205f20925f5b8181106106db5750506106aa925003836105b1565b84548352600194850194879450602090930192016106c6565b8054821015610709575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b5f818152600260205260409020548015610806575f1981018181116107f2576001545f198101919082116107f2578181036107a4575b5050506001548015610790575f190161076d8160016106f4565b8154905f199060031b1b191690556001555f5260026020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b6107dc6107b56107c69360016106f4565b90549060031b1c92839260016106f4565b819391549060031b91821b915f19901b19161790565b90555f52600260205260405f20555f8080610753565b634e487b7160e01b5f52601160045260245ffd5b50505f90565b805f52600260205260405f2054155f14610861576001546801000000000000000081101561059d5761084a6107c682600185940160015560016106f4565b9055600154905f52600260205260405f2055600190565b505f9056fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212202f3c22a745d49a71a4d5a17cd24b2afe1712752da3d4e5a3b4e3517931c247bb64736f6c63430008210033",
}

// ValidatorSlash is an auto generated Go binding around an Ethereum contract.
type ValidatorSlash struct {
	abi abi.ABI
}

// NewValidatorSlash creates a new instance of ValidatorSlash.
func NewValidatorSlash() *ValidatorSlash {
	parsed, err := ValidatorSlashMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ValidatorSlash{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ValidatorSlash) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetSlashedValidators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4ac9302e.
//
// Solidity: function getSlashedValidators() view returns(address[])
func (validatorSlash *ValidatorSlash) PackGetSlashedValidators() []byte {
	enc, err := validatorSlash.abi.Pack("getSlashedValidators")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetSlashedValidators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4ac9302e.
//
// Solidity: function getSlashedValidators() view returns(address[])
func (validatorSlash *ValidatorSlash) UnpackGetSlashedValidators(data []byte) ([]common.Address, error) {
	out, err := validatorSlash.abi.Unpack("getSlashedValidators", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (validatorSlash *ValidatorSlash) PackInitialize() []byte {
	enc, err := validatorSlash.abi.Pack("initialize")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMitigate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8d02d4be.
//
// Solidity: function mitigate() returns()
func (validatorSlash *ValidatorSlash) PackMitigate() []byte {
	enc, err := validatorSlash.abi.Pack("mitigate")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSlashOffline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcec26e14.
//
// Solidity: function slashOffline(address[] validatorAddrs) returns()
func (validatorSlash *ValidatorSlash) PackSlashOffline(validatorAddrs []common.Address) []byte {
	enc, err := validatorSlash.abi.Pack("slashOffline", validatorAddrs)
	if err != nil {
		panic(err)
	}
	return enc
}
