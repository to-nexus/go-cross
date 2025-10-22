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

// StakeHubMetaData contains all meta data concerning the StakeHub contract.
var StakeHubMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"function\",\"name\":\"addToBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyValidator\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getKeeper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorAddress\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeAmount\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorAddress\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"leaveValidator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minStakeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFromBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeeper\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStakeAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnbondingPeriod\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setvalidatorThreshold\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbondingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperChanged\",\"inputs\":[{\"name\":\"previousKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperUpdated\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinStakeAmountUpdated\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeClaimed\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnBlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnbondingPeriodUpdated\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorApplied\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorLeft\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"leavingTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorThresholdUpdated\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyLeaving\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateOperatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInLeavingState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyKeeper\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StillInUnbondingPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "StakeHub",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f5f3560e01c806302c7553214610d1c5780630c2eb40314610cc8578063114eaf5514610c34578063391b6f4e14610bff5780633f4ba83a14610b7157806340550a1c14610b1a578063417c73a714610a9a5780634a49ac4c14610a0a5780634fd101d7146109ec5780635c975abb146109bd578063603eff8b146108ba5780636cf6d6751461089c5780637071688a1461087e578063748747e6146107b85780638456cb59146107185780638f9801e414610681578063bff02e20146105d7578063c4d66de8146103ff578063d1379332146103b7578063eb321173146101e4578063eb4af04514610180578063f158628d146101365763f188768414610117575f80fd5b3461013357806003193601126101335760209054604051908152f35b80fd5b5034610133576020366003190112610133576004356001600160a01b0381169081900361017c5790602091815260068252604060018060a01b0391205416604051908152f35b5080fd5b50346101335760203660031901126101335760043561100833036101cf576020817f8448c02797b448f4946bc25b3bf925e5556d1df822c944da701c54bab8a3162f928455604051908152a180f35b630f22c43960e41b8252611008600452602482fd5b50346101335780600319360112610133576101fd6110f5565b60ff61020833610eca565b54166103a457610223335f52600460205260405f2054151590565b15610395577f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c6103865760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d33815260056020526040812060048101548015610377574210610368578054600182015491839055906001600160a01b03168280808085335af13d15610363573d67ffffffffffffffff811161034f57604051906102da601f8201601f191660200183610f02565b81528460203d92013e5b1561033b576040519182527f8766f9fdb3d7ff5c4fe7ccf145668fbc991a789c4d655f9535b0e08c4964a04d60203393a3807f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d80f35b63c39ba1a960e01b83526004829052602483fd5b634e487b7160e01b85526041600452602485fd5b6102e4565b630e8c992760e01b8252600482fd5b63281f6b9160e01b8352600483fd5b633ee5aeb560e01b8152600490fd5b632ec5b44960e01b8152600490fd5b63033788ff60e01b815233600452602490fd5b5034610133576020366003190112610133576004356001600160a01b0381169081900361017c579060209181526005825260016040818060a01b039220015416604051908152f35b5034610133576020366003190112610133576004356001600160a01b0381169081900361017c575f51602061122a5f395f51905f52549060ff8260401c16159167ffffffffffffffff8116801590816105cf575b60011490816105c5575b1590816105bc575b506105ad5767ffffffffffffffff1981166001175f51602061122a5f395f51905f525582610581575b50413303610572573a6105635769021e19e0c9bab24000008355600a60015562093a806002556104bc6111be565b6104c46111be565b6bffffffffffffffffffffffff60a01b5f5160206111ea5f395f51905f525416175f5160206111ea5f395f51905f52556104fc6111be565b6105046111be565b61050b5780f35b68ff0000000000000000195f51602061122a5f395f51905f5254165f51602061122a5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6383f1b1d360e01b8352600483fd5b63022d8c9560e31b8352600483fd5b68ffffffffffffffffff191668010000000000000001175f51602061122a5f395f51905f52555f61048e565b63f92ee8a960e01b8452600484fd5b9050155f610465565b303b15915061045d565b849150610453565b5034610133576040366003190112610133576105f7602435600435610fcb565b604051606080825284519082018190526080820195919460200190825b81811061066257505050838503602085015260208084519687815201930190945b80861061064a57505082935060408301520390f35b90926020806001928651815201940195019490610635565b82516001600160a01b0316885260209788019790920191600101610614565b50346101335760203660031901126101335760043561100833036101cf5780156106d7576020817f324143af489ab0ba77b0d3580e2486ee612b673b87db2179c589f61ac532fd5992600155604051908152a180f35b60849060405190632c648cf160e01b82526040600483015260126044830152711d985b1a59185d1bdc951a1c995cda1bdb1960721b60648301526024820152fd5b50346101335780600319360112610133575f5160206111ea5f395f51905f52546001600160a01b031633036107a95761074f6110f5565b6107576110f5565b600160ff195f51602061120a5f395f51905f525416175f51602061120a5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b63c60eb33560e01b8152600490fd5b5034610133576020366003190112610133576004356001600160a01b0381169081900361017c5761100833036101cf578015610849575f5160206111ea5f395f51905f5254816001600160a01b0382167f068b48a2fe7f498b57ff6da64f075ae658fde8d77124b092e62b3dc58d91ce358580a36001600160a01b031916175f5160206111ea5f395f51905f525580f35b60848260405190632c648cf160e01b825260406004830152600660448301526535b2b2b832b960d11b60648301526024820152fd5b50346101335780600319360112610133576020600354604051908152f35b50346101335780600319360112610133576020600254604051908152f35b50346101335780600319360112610133576108d36110f5565b60ff6108de33610eca565b54166103a4576108f9335f52600460205260405f2054151590565b1561039557338152600560205260408120600481019081546109ae5761092160025442610f38565b8092556001808060a01b03910154166040519182527fe73d82fe262870ff059c299a893937ea90edc32a71143234a83299297ff8966660203393a36110063b1561013357604051632961046560e21b815233600482015281908181602481836110065af180156109a3576109925750f35b8161099c91610f02565b6101335780f35b6040513d84823e3d90fd5b6333dd977160e01b8352600483fd5b5034610133578060031936011261013357602060ff5f51602061120a5f395f51905f5254166040519015158152f35b50346101335780600319360112610133576020600154604051908152f35b5034610133576020366003190112610133576004356001600160a01b03811690818103610a96575f5160206111ea5f395f51905f52546001600160a01b03163303610a8757610a5890610eca565b805460ff191690557fe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de40118280a280f35b63c60eb33560e01b8352600483fd5b8280fd5b5034610133576020366003190112610133576004356001600160a01b03811690818103610a96575f5160206111ea5f395f51905f52546001600160a01b03163303610a8757610ae890610eca565b805460ff191660011790557f7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec61941868280a280f35b5034610133576020366003190112610133576004356001600160a01b0381169081900361017c57816040916020935260058352206004810154159081610b66575b506040519015158152f35b90505415155f610b5b565b50346101335780600319360112610133575f5160206111ea5f395f51905f52546001600160a01b031633036107a957610ba861111c565b610bb061111c565b60ff195f51602061120a5f395f51905f5254165f51602061120a5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a180f35b50346101335780600319360112610133575f5160206111ea5f395f51905f52546040516001600160a01b039091168152602090f35b50346101335760203660031901126101335760043561100833036101cf578015610c8a576020817f38a1644f59872db6fd17fdced395495fcaa3ca7d2825a0704db5a3acbd1dd06492600255604051908152a180f35b60849060405190632c648cf160e01b825260406004830152600f60448301526e1d5b989bdb991a5b99d4195c9a5bd9608a1b60648301526024820152fd5b5034610133576020366003190112610133576004356001600160a01b0381169081900361017c579060209181526005825260408120906004820154155f14610d155750545b604051908152f35b9050610d0d565b506020366003190112610e77576004356001600160a01b03811690819003610e7757610d466110f5565b60ff610d5133610eca565b5416610eb7578015610ea8575f543410610e9957610d7a335f52600460205260405f2054151590565b610e8a575f818152600660205260409020546001600160a01b0316610e7b57610da233611144565b505f8181526006602090815260408083208054336001600160a01b0319918216811790925581855260058452828520600181018054831688179055600281018054909216831790915534808255426003830155600490910194909455905192835292917f0d83b503e698d0879dc7ad5505a19535ed048371a78d4b5908e85ca253a31a3c9190a36110063b15610e7757604051632961046560e21b81523360048201525f81602481836110065af18015610e6c57610e5e575080f35b610e6a91505f90610f02565b005b6040513d5f823e3d90fd5b5f80fd5b63813f326760e01b5f5260045ffd5b639cc20ced60e01b5f5260045ffd5b63e008b5f960e01b5f5260045ffd5b63713ce51160e01b5f5260045ffd5b63033788ff60e01b5f523360045260245ffd5b6001600160a01b03165f9081527f46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e5016020526040902090565b90601f8019910116810190811067ffffffffffffffff821117610f2457604052565b634e487b7160e01b5f52604160045260245ffd5b91908201809211610f4557565b634e487b7160e01b5f52601160045260245ffd5b67ffffffffffffffff8111610f245760051b60200190565b90610f7b82610f59565b610f886040519182610f02565b8281528092610f99601f1991610f59565b0190602036910137565b8051821015610fb75760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b9060035490818310156110c057806110b75750610fe9815b83610f38565b928184116110ae575b828403938411610f4557611007849394610f71565b9361101184610f71565b935f91600354925b8281106110265750505050565b6110308183610f38565b9084821015610fb7577fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b909101546001600160a01b039081165f9081526005602052604090206001818101549093921661108a838c610fa3565b5260048101546110a757545b6110a0828a610fa3565b5201611019565b505f611096565b92508092610ff2565b610fe990610fe3565b5090506020916040516110d38482610f02565b5f81525f368137604051936110e88186610f02565b5f8552505f368137929190565b60ff5f51602061120a5f395f51905f52541661110d57565b63d93c066560e01b5f5260045ffd5b60ff5f51602061120a5f395f51905f5254161561113557565b638dfc202b60e01b5f5260045ffd5b805f52600460205260405f2054155f146111b95760035468010000000000000000811015610f245760018101600355600354811015610fb7577fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018190556003545f9182526004602052604090912055600190565b505f90565b60ff5f51602061122a5f395f51905f525460401c16156111da57565b631afcd79f60e31b5f5260045ffdfe46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e500cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212200577972b1b716c6114e5f070f1ca38fd84e1e906349d69857c4b1d0aeafaefd464736f6c634300081c0033",
}

// StakeHub is an auto generated Go binding around an Ethereum contract.
type StakeHub struct {
	abi abi.ABI
}

// NewStakeHub creates a new instance of StakeHub.
func NewStakeHub() *StakeHub {
	parsed, err := StakeHubMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &StakeHub{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *StakeHub) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddToBlackList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (stakeHub *StakeHub) PackAddToBlackList(account common.Address) []byte {
	enc, err := stakeHub.abi.Pack("addToBlackList", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackApplyValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02c75532.
//
// Solidity: function applyValidator(address validatorAddr) payable returns()
func (stakeHub *StakeHub) PackApplyValidator(validatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("applyValidator", validatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackClaimStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb321173.
//
// Solidity: function claimStake() returns()
func (stakeHub *StakeHub) PackClaimStake() []byte {
	enc, err := stakeHub.abi.Pack("claimStake")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetKeeper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x391b6f4e.
//
// Solidity: function getKeeper() view returns(address)
func (stakeHub *StakeHub) PackGetKeeper() []byte {
	enc, err := stakeHub.abi.Pack("getKeeper")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetKeeper is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x391b6f4e.
//
// Solidity: function getKeeper() view returns(address)
func (stakeHub *StakeHub) UnpackGetKeeper(data []byte) (common.Address, error) {
	out, err := stakeHub.abi.Unpack("getKeeper", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetOperatorAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf158628d.
//
// Solidity: function getOperatorAddress(address validatorAddr) view returns(address)
func (stakeHub *StakeHub) PackGetOperatorAddress(validatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getOperatorAddress", validatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetOperatorAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf158628d.
//
// Solidity: function getOperatorAddress(address validatorAddr) view returns(address)
func (stakeHub *StakeHub) UnpackGetOperatorAddress(data []byte) (common.Address, error) {
	out, err := stakeHub.abi.Unpack("getOperatorAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetStakeAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c2eb403.
//
// Solidity: function getStakeAmount(address operatorAddr) view returns(uint256)
func (stakeHub *StakeHub) PackGetStakeAmount(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getStakeAmount", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetStakeAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0c2eb403.
//
// Solidity: function getStakeAmount(address operatorAddr) view returns(uint256)
func (stakeHub *StakeHub) UnpackGetStakeAmount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("getStakeAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetValidatorAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1379332.
//
// Solidity: function getValidatorAddress(address operatorAddr) view returns(address)
func (stakeHub *StakeHub) PackGetValidatorAddress(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorAddress", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetValidatorAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd1379332.
//
// Solidity: function getValidatorAddress(address operatorAddr) view returns(address)
func (stakeHub *StakeHub) UnpackGetValidatorAddress(data []byte) (common.Address, error) {
	out, err := stakeHub.abi.Unpack("getValidatorAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (stakeHub *StakeHub) PackGetValidatorCount() []byte {
	enc, err := stakeHub.abi.Pack("getValidatorCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetValidatorCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (stakeHub *StakeHub) UnpackGetValidatorCount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("getValidatorCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetValidators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbff02e20.
//
// Solidity: function getValidators(uint256 offset, uint256 limit) view returns(address[] validatorAddrs, uint256[] amounts, uint256 totalLength)
func (stakeHub *StakeHub) PackGetValidators(offset *big.Int, limit *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("getValidators", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorsOutput serves as a container for the return parameters of contract
// method GetValidators.
type GetValidatorsOutput struct {
	ValidatorAddrs []common.Address
	Amounts        []*big.Int
	TotalLength    *big.Int
}

// UnpackGetValidators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbff02e20.
//
// Solidity: function getValidators(uint256 offset, uint256 limit) view returns(address[] validatorAddrs, uint256[] amounts, uint256 totalLength)
func (stakeHub *StakeHub) UnpackGetValidators(data []byte) (GetValidatorsOutput, error) {
	out, err := stakeHub.abi.Unpack("getValidators", data)
	outstruct := new(GetValidatorsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ValidatorAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.TotalLength = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address keeper) returns()
func (stakeHub *StakeHub) PackInitialize(keeper common.Address) []byte {
	enc, err := stakeHub.abi.Pack("initialize", keeper)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsActiveValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40550a1c.
//
// Solidity: function isActiveValidator(address operatorAddr) view returns(bool)
func (stakeHub *StakeHub) PackIsActiveValidator(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("isActiveValidator", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsActiveValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x40550a1c.
//
// Solidity: function isActiveValidator(address operatorAddr) view returns(bool)
func (stakeHub *StakeHub) UnpackIsActiveValidator(data []byte) (bool, error) {
	out, err := stakeHub.abi.Unpack("isActiveValidator", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackLeaveValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x603eff8b.
//
// Solidity: function leaveValidator() returns()
func (stakeHub *StakeHub) PackLeaveValidator() []byte {
	enc, err := stakeHub.abi.Pack("leaveValidator")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMinStakeAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (stakeHub *StakeHub) PackMinStakeAmount() []byte {
	enc, err := stakeHub.abi.Pack("minStakeAmount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMinStakeAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (stakeHub *StakeHub) UnpackMinStakeAmount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("minStakeAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.
//
// Solidity: function pause() returns()
func (stakeHub *StakeHub) PackPause() []byte {
	enc, err := stakeHub.abi.Pack("pause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (stakeHub *StakeHub) PackPaused() []byte {
	enc, err := stakeHub.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (stakeHub *StakeHub) UnpackPaused(data []byte) (bool, error) {
	out, err := stakeHub.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackRemoveFromBlackList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (stakeHub *StakeHub) PackRemoveFromBlackList(account common.Address) []byte {
	enc, err := stakeHub.abi.Pack("removeFromBlackList", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetKeeper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x748747e6.
//
// Solidity: function setKeeper(address keeper) returns()
func (stakeHub *StakeHub) PackSetKeeper(keeper common.Address) []byte {
	enc, err := stakeHub.abi.Pack("setKeeper", keeper)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinStakeAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 amount) returns()
func (stakeHub *StakeHub) PackSetMinStakeAmount(amount *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setMinStakeAmount", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetUnbondingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x114eaf55.
//
// Solidity: function setUnbondingPeriod(uint256 period) returns()
func (stakeHub *StakeHub) PackSetUnbondingPeriod(period *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setUnbondingPeriod", period)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetvalidatorThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f9801e4.
//
// Solidity: function setvalidatorThreshold(uint256 count) returns()
func (stakeHub *StakeHub) PackSetvalidatorThreshold(count *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setvalidatorThreshold", count)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUnbondingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (stakeHub *StakeHub) PackUnbondingPeriod() []byte {
	enc, err := stakeHub.abi.Pack("unbondingPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUnbondingPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (stakeHub *StakeHub) UnpackUnbondingPeriod(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("unbondingPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (stakeHub *StakeHub) PackUnpause() []byte {
	enc, err := stakeHub.abi.Pack("unpause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackValidatorThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (stakeHub *StakeHub) PackValidatorThreshold() []byte {
	enc, err := stakeHub.abi.Pack("validatorThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidatorThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (stakeHub *StakeHub) UnpackValidatorThreshold(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("validatorThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// StakeHubBlackListed represents a BlackListed event raised by the StakeHub contract.
type StakeHubBlackListed struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubBlackListedEventName = "BlackListed"

// ContractEventName returns the user-defined event name.
func (StakeHubBlackListed) ContractEventName() string {
	return StakeHubBlackListedEventName
}

// UnpackBlackListedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BlackListed(address indexed account)
func (stakeHub *StakeHub) UnpackBlackListedEvent(log *types.Log) (*StakeHubBlackListed, error) {
	event := "BlackListed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubBlackListed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubInitialized represents a Initialized event raised by the StakeHub contract.
type StakeHubInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (StakeHubInitialized) ContractEventName() string {
	return StakeHubInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (stakeHub *StakeHub) UnpackInitializedEvent(log *types.Log) (*StakeHubInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubInitialized)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubKeeperChanged represents a KeeperChanged event raised by the StakeHub contract.
type StakeHubKeeperChanged struct {
	PreviousKeeper common.Address
	NewKeeper      common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const StakeHubKeeperChangedEventName = "KeeperChanged"

// ContractEventName returns the user-defined event name.
func (StakeHubKeeperChanged) ContractEventName() string {
	return StakeHubKeeperChangedEventName
}

// UnpackKeeperChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event KeeperChanged(address indexed previousKeeper, address indexed newKeeper)
func (stakeHub *StakeHub) UnpackKeeperChangedEvent(log *types.Log) (*StakeHubKeeperChanged, error) {
	event := "KeeperChanged"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubKeeperChanged)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubKeeperUpdated represents a KeeperUpdated event raised by the StakeHub contract.
type StakeHubKeeperUpdated struct {
	Keeper common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const StakeHubKeeperUpdatedEventName = "KeeperUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubKeeperUpdated) ContractEventName() string {
	return StakeHubKeeperUpdatedEventName
}

// UnpackKeeperUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event KeeperUpdated(address keeper)
func (stakeHub *StakeHub) UnpackKeeperUpdatedEvent(log *types.Log) (*StakeHubKeeperUpdated, error) {
	event := "KeeperUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubKeeperUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubMinStakeAmountUpdated represents a MinStakeAmountUpdated event raised by the StakeHub contract.
type StakeHubMinStakeAmountUpdated struct {
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const StakeHubMinStakeAmountUpdatedEventName = "MinStakeAmountUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubMinStakeAmountUpdated) ContractEventName() string {
	return StakeHubMinStakeAmountUpdatedEventName
}

// UnpackMinStakeAmountUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinStakeAmountUpdated(uint256 amount)
func (stakeHub *StakeHub) UnpackMinStakeAmountUpdatedEvent(log *types.Log) (*StakeHubMinStakeAmountUpdated, error) {
	event := "MinStakeAmountUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubMinStakeAmountUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubPaused represents a Paused event raised by the StakeHub contract.
type StakeHubPaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubPausedEventName = "Paused"

// ContractEventName returns the user-defined event name.
func (StakeHubPaused) ContractEventName() string {
	return StakeHubPausedEventName
}

// UnpackPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Paused(address account)
func (stakeHub *StakeHub) UnpackPausedEvent(log *types.Log) (*StakeHubPaused, error) {
	event := "Paused"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubPaused)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubStakeClaimed represents a StakeClaimed event raised by the StakeHub contract.
type StakeHubStakeClaimed struct {
	ValidatorAddr common.Address
	OperatorAddr  common.Address
	Amount        *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const StakeHubStakeClaimedEventName = "StakeClaimed"

// ContractEventName returns the user-defined event name.
func (StakeHubStakeClaimed) ContractEventName() string {
	return StakeHubStakeClaimedEventName
}

// UnpackStakeClaimedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeClaimed(address indexed validatorAddr, address indexed operatorAddr, uint256 amount)
func (stakeHub *StakeHub) UnpackStakeClaimedEvent(log *types.Log) (*StakeHubStakeClaimed, error) {
	event := "StakeClaimed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubStakeClaimed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubUnBlackListed represents a UnBlackListed event raised by the StakeHub contract.
type StakeHubUnBlackListed struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubUnBlackListedEventName = "UnBlackListed"

// ContractEventName returns the user-defined event name.
func (StakeHubUnBlackListed) ContractEventName() string {
	return StakeHubUnBlackListedEventName
}

// UnpackUnBlackListedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UnBlackListed(address indexed account)
func (stakeHub *StakeHub) UnpackUnBlackListedEvent(log *types.Log) (*StakeHubUnBlackListed, error) {
	event := "UnBlackListed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubUnBlackListed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubUnbondingPeriodUpdated represents a UnbondingPeriodUpdated event raised by the StakeHub contract.
type StakeHubUnbondingPeriodUpdated struct {
	Period *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const StakeHubUnbondingPeriodUpdatedEventName = "UnbondingPeriodUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubUnbondingPeriodUpdated) ContractEventName() string {
	return StakeHubUnbondingPeriodUpdatedEventName
}

// UnpackUnbondingPeriodUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UnbondingPeriodUpdated(uint256 period)
func (stakeHub *StakeHub) UnpackUnbondingPeriodUpdatedEvent(log *types.Log) (*StakeHubUnbondingPeriodUpdated, error) {
	event := "UnbondingPeriodUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubUnbondingPeriodUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubUnpaused represents a Unpaused event raised by the StakeHub contract.
type StakeHubUnpaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubUnpausedEventName = "Unpaused"

// ContractEventName returns the user-defined event name.
func (StakeHubUnpaused) ContractEventName() string {
	return StakeHubUnpausedEventName
}

// UnpackUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpaused(address account)
func (stakeHub *StakeHub) UnpackUnpausedEvent(log *types.Log) (*StakeHubUnpaused, error) {
	event := "Unpaused"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubUnpaused)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubValidatorApplied represents a ValidatorApplied event raised by the StakeHub contract.
type StakeHubValidatorApplied struct {
	ValidatorAddr common.Address
	OperatorAddr  common.Address
	Amount        *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorAppliedEventName = "ValidatorApplied"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorApplied) ContractEventName() string {
	return StakeHubValidatorAppliedEventName
}

// UnpackValidatorAppliedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorApplied(address indexed validatorAddr, address indexed operatorAddr, uint256 amount)
func (stakeHub *StakeHub) UnpackValidatorAppliedEvent(log *types.Log) (*StakeHubValidatorApplied, error) {
	event := "ValidatorApplied"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorApplied)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubValidatorLeft represents a ValidatorLeft event raised by the StakeHub contract.
type StakeHubValidatorLeft struct {
	ValidatorAddr common.Address
	OperatorAddr  common.Address
	LeavingTime   *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorLeftEventName = "ValidatorLeft"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorLeft) ContractEventName() string {
	return StakeHubValidatorLeftEventName
}

// UnpackValidatorLeftEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorLeft(address indexed validatorAddr, address indexed operatorAddr, uint256 leavingTime)
func (stakeHub *StakeHub) UnpackValidatorLeftEvent(log *types.Log) (*StakeHubValidatorLeft, error) {
	event := "ValidatorLeft"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorLeft)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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

// StakeHubValidatorThresholdUpdated represents a ValidatorThresholdUpdated event raised by the StakeHub contract.
type StakeHubValidatorThresholdUpdated struct {
	Count *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorThresholdUpdatedEventName = "ValidatorThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorThresholdUpdated) ContractEventName() string {
	return StakeHubValidatorThresholdUpdatedEventName
}

// UnpackValidatorThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorThresholdUpdated(uint256 count)
func (stakeHub *StakeHub) UnpackValidatorThresholdUpdatedEvent(log *types.Log) (*StakeHubValidatorThresholdUpdated, error) {
	event := "ValidatorThresholdUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorThresholdUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
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
func (stakeHub *StakeHub) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["AlreadyLeaving"].ID.Bytes()[:4]) {
		return stakeHub.UnpackAlreadyLeavingError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["DuplicateOperatorAddress"].ID.Bytes()[:4]) {
		return stakeHub.UnpackDuplicateOperatorAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["DuplicateValidatorAddress"].ID.Bytes()[:4]) {
		return stakeHub.UnpackDuplicateValidatorAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["EnforcedPause"].ID.Bytes()[:4]) {
		return stakeHub.UnpackEnforcedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["ExpectedPause"].ID.Bytes()[:4]) {
		return stakeHub.UnpackExpectedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InBlackList"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInBlackListError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidValidatorAddress"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidValidatorAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotEnoughAmount"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotEnoughAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotInLeavingState"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotInLeavingStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotValidator"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotValidatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return stakeHub.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["OnlyKeeper"].ID.Bytes()[:4]) {
		return stakeHub.UnpackOnlyKeeperError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return stakeHub.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return stakeHub.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return stakeHub.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["StillInUnbondingPeriod"].ID.Bytes()[:4]) {
		return stakeHub.UnpackStillInUnbondingPeriodError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["TransferFailed"].ID.Bytes()[:4]) {
		return stakeHub.UnpackTransferFailedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// StakeHubAlreadyLeaving represents a AlreadyLeaving error raised by the StakeHub contract.
type StakeHubAlreadyLeaving struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyLeaving()
func StakeHubAlreadyLeavingErrorID() common.Hash {
	return common.HexToHash("0x33dd9771a04987f8b2f49f8bbcc151db435d1e17b8367051371cfd069b3fb717")
}

// UnpackAlreadyLeavingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyLeaving()
func (stakeHub *StakeHub) UnpackAlreadyLeavingError(raw []byte) (*StakeHubAlreadyLeaving, error) {
	out := new(StakeHubAlreadyLeaving)
	if err := stakeHub.abi.UnpackIntoInterface(out, "AlreadyLeaving", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubDuplicateOperatorAddress represents a DuplicateOperatorAddress error raised by the StakeHub contract.
type StakeHubDuplicateOperatorAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateOperatorAddress()
func StakeHubDuplicateOperatorAddressErrorID() common.Hash {
	return common.HexToHash("0x9cc20cedeb767979223b806ace03a4cd1b95ca667f7eb9f2be73281565920a41")
}

// UnpackDuplicateOperatorAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateOperatorAddress()
func (stakeHub *StakeHub) UnpackDuplicateOperatorAddressError(raw []byte) (*StakeHubDuplicateOperatorAddress, error) {
	out := new(StakeHubDuplicateOperatorAddress)
	if err := stakeHub.abi.UnpackIntoInterface(out, "DuplicateOperatorAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubDuplicateValidatorAddress represents a DuplicateValidatorAddress error raised by the StakeHub contract.
type StakeHubDuplicateValidatorAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateValidatorAddress()
func StakeHubDuplicateValidatorAddressErrorID() common.Hash {
	return common.HexToHash("0x813f326766a6a6c709fc0f61143d5c4586a07d19cd7652a1fcbe25d4f7562c0e")
}

// UnpackDuplicateValidatorAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateValidatorAddress()
func (stakeHub *StakeHub) UnpackDuplicateValidatorAddressError(raw []byte) (*StakeHubDuplicateValidatorAddress, error) {
	out := new(StakeHubDuplicateValidatorAddress)
	if err := stakeHub.abi.UnpackIntoInterface(out, "DuplicateValidatorAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubEnforcedPause represents a EnforcedPause error raised by the StakeHub contract.
type StakeHubEnforcedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EnforcedPause()
func StakeHubEnforcedPauseErrorID() common.Hash {
	return common.HexToHash("0xd93c0665d6c96d04a8f174024fc4ddd66c250604aff22bbec808de86dd3637e3")
}

// UnpackEnforcedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EnforcedPause()
func (stakeHub *StakeHub) UnpackEnforcedPauseError(raw []byte) (*StakeHubEnforcedPause, error) {
	out := new(StakeHubEnforcedPause)
	if err := stakeHub.abi.UnpackIntoInterface(out, "EnforcedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubExpectedPause represents a ExpectedPause error raised by the StakeHub contract.
type StakeHubExpectedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExpectedPause()
func StakeHubExpectedPauseErrorID() common.Hash {
	return common.HexToHash("0x8dfc202bcfe9a735b559bee70674422512bc5c30f687046ae8778315fb81da44")
}

// UnpackExpectedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExpectedPause()
func (stakeHub *StakeHub) UnpackExpectedPauseError(raw []byte) (*StakeHubExpectedPause, error) {
	out := new(StakeHubExpectedPause)
	if err := stakeHub.abi.UnpackIntoInterface(out, "ExpectedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInBlackList represents a InBlackList error raised by the StakeHub contract.
type StakeHubInBlackList struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InBlackList(address account)
func StakeHubInBlackListErrorID() common.Hash {
	return common.HexToHash("0x033788ffabe13708428b49b6cedeba51975d173696096da94c2097fc14da662e")
}

// UnpackInBlackListError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InBlackList(address account)
func (stakeHub *StakeHub) UnpackInBlackListError(raw []byte) (*StakeHubInBlackList, error) {
	out := new(StakeHubInBlackList)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InBlackList", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidInitialization represents a InvalidInitialization error raised by the StakeHub contract.
type StakeHubInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func StakeHubInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (stakeHub *StakeHub) UnpackInvalidInitializationError(raw []byte) (*StakeHubInvalidInitialization, error) {
	out := new(StakeHubInvalidInitialization)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidValidatorAddress represents a InvalidValidatorAddress error raised by the StakeHub contract.
type StakeHubInvalidValidatorAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValidatorAddress()
func StakeHubInvalidValidatorAddressErrorID() common.Hash {
	return common.HexToHash("0x713ce511aa50868d0c2d8b07e0b3982819c17809cacda19b9e7a5e3b4d4bb677")
}

// UnpackInvalidValidatorAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValidatorAddress()
func (stakeHub *StakeHub) UnpackInvalidValidatorAddressError(raw []byte) (*StakeHubInvalidValidatorAddress, error) {
	out := new(StakeHubInvalidValidatorAddress)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidValidatorAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidValue represents a InvalidValue error raised by the StakeHub contract.
type StakeHubInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func StakeHubInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (stakeHub *StakeHub) UnpackInvalidValueError(raw []byte) (*StakeHubInvalidValue, error) {
	out := new(StakeHubInvalidValue)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubNotEnoughAmount represents a NotEnoughAmount error raised by the StakeHub contract.
type StakeHubNotEnoughAmount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEnoughAmount()
func StakeHubNotEnoughAmountErrorID() common.Hash {
	return common.HexToHash("0xe008b5f9775a2bbe458e4c5378d46e4d64183e024a4785f91837700c840acfee")
}

// UnpackNotEnoughAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEnoughAmount()
func (stakeHub *StakeHub) UnpackNotEnoughAmountError(raw []byte) (*StakeHubNotEnoughAmount, error) {
	out := new(StakeHubNotEnoughAmount)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotEnoughAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubNotInLeavingState represents a NotInLeavingState error raised by the StakeHub contract.
type StakeHubNotInLeavingState struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInLeavingState()
func StakeHubNotInLeavingStateErrorID() common.Hash {
	return common.HexToHash("0x281f6b914c6cc246a19e3cbb0e0cb34cf54e3fd690ef2a5bd9c5ff78b9ddcd9a")
}

// UnpackNotInLeavingStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInLeavingState()
func (stakeHub *StakeHub) UnpackNotInLeavingStateError(raw []byte) (*StakeHubNotInLeavingState, error) {
	out := new(StakeHubNotInLeavingState)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotInLeavingState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubNotInitializing represents a NotInitializing error raised by the StakeHub contract.
type StakeHubNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func StakeHubNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (stakeHub *StakeHub) UnpackNotInitializingError(raw []byte) (*StakeHubNotInitializing, error) {
	out := new(StakeHubNotInitializing)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubNotValidator represents a NotValidator error raised by the StakeHub contract.
type StakeHubNotValidator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotValidator()
func StakeHubNotValidatorErrorID() common.Hash {
	return common.HexToHash("0x2ec5b449e1d63fa34c160c67ce2d5826d939017f27bc0c78ef142b8357c69f9e")
}

// UnpackNotValidatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotValidator()
func (stakeHub *StakeHub) UnpackNotValidatorError(raw []byte) (*StakeHubNotValidator, error) {
	out := new(StakeHubNotValidator)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotValidator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubOnlyCoinbase represents a OnlyCoinbase error raised by the StakeHub contract.
type StakeHubOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func StakeHubOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (stakeHub *StakeHub) UnpackOnlyCoinbaseError(raw []byte) (*StakeHubOnlyCoinbase, error) {
	out := new(StakeHubOnlyCoinbase)
	if err := stakeHub.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubOnlyKeeper represents a OnlyKeeper error raised by the StakeHub contract.
type StakeHubOnlyKeeper struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyKeeper()
func StakeHubOnlyKeeperErrorID() common.Hash {
	return common.HexToHash("0xc60eb3352805112c61799a78f842543bbf71f1a5e9cd075fb2e23066f7db6914")
}

// UnpackOnlyKeeperError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyKeeper()
func (stakeHub *StakeHub) UnpackOnlyKeeperError(raw []byte) (*StakeHubOnlyKeeper, error) {
	out := new(StakeHubOnlyKeeper)
	if err := stakeHub.abi.UnpackIntoInterface(out, "OnlyKeeper", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubOnlySystemContract represents a OnlySystemContract error raised by the StakeHub contract.
type StakeHubOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func StakeHubOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (stakeHub *StakeHub) UnpackOnlySystemContractError(raw []byte) (*StakeHubOnlySystemContract, error) {
	out := new(StakeHubOnlySystemContract)
	if err := stakeHub.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the StakeHub contract.
type StakeHubOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func StakeHubOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (stakeHub *StakeHub) UnpackOnlyZeroGasPriceError(raw []byte) (*StakeHubOnlyZeroGasPrice, error) {
	out := new(StakeHubOnlyZeroGasPrice)
	if err := stakeHub.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the StakeHub contract.
type StakeHubReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func StakeHubReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (stakeHub *StakeHub) UnpackReentrancyGuardReentrantCallError(raw []byte) (*StakeHubReentrancyGuardReentrantCall, error) {
	out := new(StakeHubReentrancyGuardReentrantCall)
	if err := stakeHub.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubStillInUnbondingPeriod represents a StillInUnbondingPeriod error raised by the StakeHub contract.
type StakeHubStillInUnbondingPeriod struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StillInUnbondingPeriod()
func StakeHubStillInUnbondingPeriodErrorID() common.Hash {
	return common.HexToHash("0x0e8c99272d5669d980928301edd9c9438012d2001950950ad47620b73840623a")
}

// UnpackStillInUnbondingPeriodError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StillInUnbondingPeriod()
func (stakeHub *StakeHub) UnpackStillInUnbondingPeriodError(raw []byte) (*StakeHubStillInUnbondingPeriod, error) {
	out := new(StakeHubStillInUnbondingPeriod)
	if err := stakeHub.abi.UnpackIntoInterface(out, "StillInUnbondingPeriod", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubTransferFailed represents a TransferFailed error raised by the StakeHub contract.
type StakeHubTransferFailed struct {
	Amount *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferFailed(uint256 amount)
func StakeHubTransferFailedErrorID() common.Hash {
	return common.HexToHash("0xc39ba1a9a812115efec13c9546b7b109574823d9c12ff09ebcc4ba5d47153ec6")
}

// UnpackTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferFailed(uint256 amount)
func (stakeHub *StakeHub) UnpackTransferFailedError(raw []byte) (*StakeHubTransferFailed, error) {
	out := new(StakeHubTransferFailed)
	if err := stakeHub.abi.UnpackIntoInterface(out, "TransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}
