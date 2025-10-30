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

// IstanbulParamMetaData contains all meta data concerning the IstanbulParam contract.
var IstanbulParamMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"function\",\"name\":\"getParamIndex\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"idx\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParams\",\"inputs\":[{\"name\":\"timepoint_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParamsByIndex\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"beneficiary_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBaseFeeChangeDenominator\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBeneficiary\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setElasticityMultiplier\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmptyBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEpochLength\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasLimit\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProposerPolicy\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoPastUpdate\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "IstanbulParam",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c806308a4f07214610bc9578063131b23b514610b8b5780631c846d5914610ade5780634195c17714610a4b5780634708026a146109db5780635941158a146109365780637b765eb11461089a5780637eaf5a081461086b578063827792671461078c57806383f93bae146106e45780639d28f9b014610648578063aa26095714610584578063c4d66de814610333578063cf22854b14610264578063ea821efc1461020c578063f9d260d3146101df5763ff035345146100d4575f80fd5b346101db5760403660031901126101db576004356100f0610c03565b61100833036101c6576001600160401b0361010a43610ea7565b16808311156101b057506001600160401b038116801515806101a5575b6101685750600461013a61016693610d09565b01805467ffffffffffffffff60a01b191660a09290921b67ffffffffffffffff60a01b16919091179055565b005b60849060405190632c648cf160e01b825260406004830152600e60448301526d70726f706f736572506f6c69637960901b60648301526024820152fd5b506001811415610127565b8263497653a360e11b5f5260045260245260445ffd5b630f22c43960e41b5f5261100860045260245ffd5b5f80fd5b346101db575f3660031901126101db5760206101fb5f54610ea7565b6001600160401b0360405191168152f35b346101db5761021a36610bed565b9061100833036101c6576001600160401b0361023543610ea7565b168082111561024e575061024a600391610d09565b0155005b9063497653a360e11b5f5260045260245260445ffd5b346101db5760403660031901126101db57600435610280610c03565b61100833036101c6576001600160401b0361029a43610ea7565b16808311156101b057506001600160401b03811680156102e7575060016102c361016693610d09565b0180546001600160c01b031660c09290921b6001600160c01b031916919091179055565b60849060405190632c648cf160e01b825260406004830152601860448301527f626173654665654368616e676544656e6f6d696e61746f72000000000000000060648301526024820152fd5b346101db5760203660031901126101db576004356001600160a01b038116908190036101db575f5160206112fc5f395f51905f5254906001600160401b0360ff8360401c161592168015908161057c575b6001149081610572575b159081610569575b5061055a578160016001600160401b03195f5160206112fc5f395f51905f525416175f5160206112fc5f395f51905f525561052a575b41330361051b573a61050c57604051906103e582610cb8565b5f825260208201905f8252604083015f815260608401905f825260808501905f825260a08601905f825260c087015f815260e08801915f83526101008901935f85526101208a01955f87526101408b01975f89526101608c01935f85526101808d019b5f8d528d6101c0906040519161045e8184610ce8565b3683376101a001528d61047043610ea7565b6001600160401b03169052620151809052600190525f90526113889052603c90525f905260039052600890525f90525f905252630c84588090526104b390610f0f565b506104ba57005b60ff60401b195f5160206112fc5f395f51905f5254165f5160206112fc5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b600160401b60ff60401b195f5160206112fc5f395f51905f525416175f5160206112fc5f395f51905f52556103cc565b63f92ee8a960e01b5f5260045ffd5b90501583610396565b303b15915061038e565b839150610384565b346101db5760403660031901126101db576004356105a0610c03565b9061100833036101c6576001600160401b036105bb43610ea7565b169182821115610632576001600160401b03169081156105f5576105e0600191610d09565b01805467ffffffffffffffff19169091179055005b60848260405190632c648cf160e01b825260406004830152600e60448301526d1c995c5d595cdd151a5b595bdd5d60921b60648301526024820152fd5b5063497653a360e11b5f5260045260245260445ffd5b346101db5760403660031901126101db57600435610664610c03565b61100833036101c6576001600160401b0361067e43610ea7565b16808311156101b057506001600160401b03811680156106aa57506106a561016692610d09565b610c19565b60849060405190632c648cf160e01b825260406004830152600b60448301526a0cae0dec6d098cadccee8d60ab1b60648301526024820152fd5b346101db5760403660031901126101db57600435610700610c03565b61100833036101c6576001600160401b0361071a43610ea7565b16808311156101b057506001600160401b03811680156107495750600161074361016693610d09565b01610c41565b60849060405190632c648cf160e01b825260406004830152601460448301527332b630b9ba34b1b4ba3ca6bab63a34b83634b2b960611b60648301526024820152fd5b346101db5760203660031901126101db576004356001600160401b03811681036101db576107b990610c6c565b5080546001820154600283015460038401546004850154600590950154604080516001600160401b03808816825287831c81166020830152608088811c82168385015260c098891c6060840152818816818401529287901c811660a0808401919091529287901c8116888301529590961c60e08701526101008601939093526101208501919091526001600160a01b03851661014085015293901c81166101608301529091166101808201526101a090f35b346101db5760203660031901126101db5760206101fb6001600160401b03610894600435610e5f565b16610ea7565b346101db5760403660031901126101db576004356108b6610c03565b61100833036101c6576001600160401b036108d043610ea7565b16808311156101b057506001600160401b03811680156108fc57506108f761016692610d09565b610c41565b60849060405190632c648cf160e01b825260406004830152600b60448301526a189b1bd8dad4195c9a5bd960aa1b60648301526024820152fd5b346101db5760403660031901126101db57600435610952610c03565b61100833036101c6576001600160401b0361096c43610ea7565b16808311156101b057506001600160401b038116801561099b5750600161099561016693610d09565b01610c19565b60849060405190632c648cf160e01b82526040600483015260116044830152701b585e14995c5d595cdd151a5b595bdd5d607a1b60648301526024820152fd5b346101db5760403660031901126101db576004356109f7610c03565b9061100833036101c6576001600160401b03610a1243610ea7565b168082111561024e5761016683610a2884610d09565b80546001600160c01b031660c09290921b6001600160c01b031916919091179055565b346101db5760403660031901126101db57600435610a67610c03565b9061100833036101c6576001600160401b03610a8243610ea7565b169182821115610632576001600160401b0316908115610aa7576105e0600591610d09565b60848260405190632c648cf160e01b825260406004830152600860448301526719d85cd31a5b5a5d60c21b60648301526024820152fd5b346101db5760403660031901126101db576024356001600160a01b03811690600435908290036101db5761100833036101c6576001600160401b03610b2243610ea7565b168082111561024e57508115610b5157610b3d600491610d09565b0180546001600160a01b0319169091179055005b60848260405190632c648cf160e01b825260406004830152600b60448301526a62656e656669636961727960a81b60648301526024820152fd5b346101db57610b9936610bed565b9061100833036101c6576001600160401b03610bb443610ea7565b168082111561024e575061024a600291610d09565b346101db5760203660031901126101db576107b9610be8600435610e5f565b610c6c565b60409060031901126101db576004359060243590565b602435906001600160401b03821682036101db57565b9067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b5f54811015610ca4575f8080526014919091027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630191565b634e487b7160e01b5f52603260045260245ffd5b6101c081019081106001600160401b03821117610cd457604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610cd457604052565b610d15610be882610e5f565b50908154916001600160401b03831692828414610e585760405193610d3985610cb8565b84526001600160401b038160401c1660208501526001600160401b038160801c16604085015260c01c606084015260018101546001600160401b03811660808501526001600160401b038160401c1660a08501526001600160401b038160801c1660c085015260c01c60e0840152600281015461010084015260038101546101208401526001600160401b03600482015460018060a01b03811661014086015260a01c166101608401526001600160401b03600582015416610180840152604051906006829101905f905b600e8210610e4257505050610e37610e3f939282610e2c6101c06001600160401b0395610ce8565b6101a0850152610ea7565b168152610f0f565b90565b6001602081928554815201930191019091610e04565b5091505090565b5f5480915f905b838210610e7a575b5050610e3f9150610ea7565b9091816001600160401b03610e8e85610c6c565b50541611610ea157506001820190610e66565b91610e6e565b6001600160401b038111610ec1576001600160401b031690565b6306dfcc6560e41b5f52604060045260245260445ffd5b90808214610f0b57908154905f915b600e8310610ef55750505050565b6001809194019283549481840155019192610ee7565b5050565b5f54600160401b811015610cd45760018101805f55610f2d82610c6c565b50508111611100575b801515806112ca575b15611114575f1981019080821161110057610f63610f5c83610c6c565b5091610c6c565b6110ed57818103610f76575b5050610f36565b600680836001600160401b03806110e6965416166001600160401b0319855416178455610fb06001600160401b03825460401c1685610c19565b610fc76001600160401b03825460801c1685610c41565b805484546001600160c01b03166001600160c01b031990911617845561105560018501600183016001600160401b0380825416166001600160401b03198354161782556110216001600160401b03825460401c1683610c19565b6110386001600160401b03825460801c1683610c41565b5481546001600160c01b03166001600160c01b0319909116179055565b60028181015490850155600380820154908501556004808201805491860180546001600160a01b039093166001600160a01b0319841681178255915467ffffffffffffffff60a01b166001600160e01b03199093169091179190911790556001600160401b036005820154166001600160401b036005860191166001600160401b0319825416179055019101610ed8565b5f80610f6f565b634e487b7160e01b5f525f60045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b91909161112081610c6c565b9390936110ed57806001600160401b03806101a0935116166001600160401b031986541617855561115e6001600160401b0360208301511686610c19565b6111756001600160401b0360408301511686610c41565b606081015185546001600160c01b031660c09190911b6001600160c01b031916178555611211600186016001600160401b0380608085015116166001600160401b03198254161781556111d56001600160401b0360a08501511682610c19565b6111ec6001600160401b0360c08501511682610c41565b60e083015181546001600160c01b031660c09190911b6001600160c01b031916179055565b6101008101516002860155610120810151600386015561014081015160048601805461016084015167ffffffffffffffff60a01b60a09190911b166001600160a01b039093166001600160e01b0319909116179190911790556001600160401b03610180820151166001600160401b036005870191166001600160401b031982541617905501515f5b600e81106112b35750506112af919250610c6c565b5090565b60019060208351930192600682880101550161129a565b505f198101818111611100576112e76001600160401b0391610c6c565b5054166001600160401b0383511610610f3f56fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212201e8a0590a7f0100f848ef2317eb628c6ebc4e1f3167031ce01881f330a364c4764736f6c634300081c0033",
}

// IstanbulParam is an auto generated Go binding around an Ethereum contract.
type IstanbulParam struct {
	abi abi.ABI
}

// NewIstanbulParam creates a new instance of IstanbulParam.
func NewIstanbulParam() *IstanbulParam {
	parsed, err := IstanbulParamMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IstanbulParam{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IstanbulParam) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetParamIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7eaf5a08.
//
// Solidity: function getParamIndex(uint256 timepoint) view returns(uint64 idx)
func (istanbulParam *IstanbulParam) PackGetParamIndex(timepoint *big.Int) []byte {
	enc, err := istanbulParam.abi.Pack("getParamIndex", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetParamIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7eaf5a08.
//
// Solidity: function getParamIndex(uint256 timepoint) view returns(uint64 idx)
func (istanbulParam *IstanbulParam) UnpackGetParamIndex(data []byte) (uint64, error) {
	out, err := istanbulParam.abi.Unpack("getParamIndex", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackGetParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08a4f072.
//
// Solidity: function getParams(uint256 timepoint_) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, address beneficiary, uint64 proposerPolicy, uint64 gasLimit)
func (istanbulParam *IstanbulParam) PackGetParams(timepoint *big.Int) []byte {
	enc, err := istanbulParam.abi.Pack("getParams", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetParamsOutput serves as a container for the return parameters of contract
// method GetParams.
type GetParamsOutput struct {
	Timepoint                uint64
	EpochLength              uint64
	BlockPeriod              uint64
	EmptyBlockPeriod         uint64
	RequestTimeout           uint64
	MaxRequestTimeout        uint64
	ElasticityMultiplier     uint64
	BaseFeeChangeDenominator uint64
	MaxBaseFee               *big.Int
	MinBaseFee               *big.Int
	Beneficiary              common.Address
	ProposerPolicy           uint64
	GasLimit                 uint64
}

// UnpackGetParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x08a4f072.
//
// Solidity: function getParams(uint256 timepoint_) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, address beneficiary, uint64 proposerPolicy, uint64 gasLimit)
func (istanbulParam *IstanbulParam) UnpackGetParams(data []byte) (GetParamsOutput, error) {
	out, err := istanbulParam.abi.Unpack("getParams", data)
	outstruct := new(GetParamsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Timepoint = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.EpochLength = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BlockPeriod = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.EmptyBlockPeriod = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.RequestTimeout = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.MaxRequestTimeout = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.ElasticityMultiplier = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.BaseFeeChangeDenominator = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.MaxBaseFee = abi.ConvertType(out[8], new(big.Int)).(*big.Int)
	outstruct.MinBaseFee = abi.ConvertType(out[9], new(big.Int)).(*big.Int)
	outstruct.Beneficiary = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)
	outstruct.ProposerPolicy = *abi.ConvertType(out[11], new(uint64)).(*uint64)
	outstruct.GasLimit = *abi.ConvertType(out[12], new(uint64)).(*uint64)
	return *outstruct, err

}

// PackGetParamsByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82779267.
//
// Solidity: function getParamsByIndex(uint64 idx) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, address beneficiary, uint64 proposerPolicy, uint64 gasLimit)
func (istanbulParam *IstanbulParam) PackGetParamsByIndex(idx uint64) []byte {
	enc, err := istanbulParam.abi.Pack("getParamsByIndex", idx)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetParamsByIndexOutput serves as a container for the return parameters of contract
// method GetParamsByIndex.
type GetParamsByIndexOutput struct {
	Timepoint                uint64
	EpochLength              uint64
	BlockPeriod              uint64
	EmptyBlockPeriod         uint64
	RequestTimeout           uint64
	MaxRequestTimeout        uint64
	ElasticityMultiplier     uint64
	BaseFeeChangeDenominator uint64
	MaxBaseFee               *big.Int
	MinBaseFee               *big.Int
	Beneficiary              common.Address
	ProposerPolicy           uint64
	GasLimit                 uint64
}

// UnpackGetParamsByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82779267.
//
// Solidity: function getParamsByIndex(uint64 idx) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, address beneficiary, uint64 proposerPolicy, uint64 gasLimit)
func (istanbulParam *IstanbulParam) UnpackGetParamsByIndex(data []byte) (GetParamsByIndexOutput, error) {
	out, err := istanbulParam.abi.Unpack("getParamsByIndex", data)
	outstruct := new(GetParamsByIndexOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Timepoint = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.EpochLength = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BlockPeriod = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.EmptyBlockPeriod = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.RequestTimeout = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.MaxRequestTimeout = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.ElasticityMultiplier = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.BaseFeeChangeDenominator = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.MaxBaseFee = abi.ConvertType(out[8], new(big.Int)).(*big.Int)
	outstruct.MinBaseFee = abi.ConvertType(out[9], new(big.Int)).(*big.Int)
	outstruct.Beneficiary = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)
	outstruct.ProposerPolicy = *abi.ConvertType(out[11], new(uint64)).(*uint64)
	outstruct.GasLimit = *abi.ConvertType(out[12], new(uint64)).(*uint64)
	return *outstruct, err

}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address beneficiary_) returns()
func (istanbulParam *IstanbulParam) PackInitialize(beneficiary common.Address) []byte {
	enc, err := istanbulParam.abi.Pack("initialize", beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNumCheckpoints is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf9d260d3.
//
// Solidity: function numCheckpoints() view returns(uint64)
func (istanbulParam *IstanbulParam) PackNumCheckpoints() []byte {
	enc, err := istanbulParam.abi.Pack("numCheckpoints")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNumCheckpoints is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf9d260d3.
//
// Solidity: function numCheckpoints() view returns(uint64)
func (istanbulParam *IstanbulParam) UnpackNumCheckpoints(data []byte) (uint64, error) {
	out, err := istanbulParam.abi.Unpack("numCheckpoints", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackSetBaseFeeChangeDenominator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf22854b.
//
// Solidity: function setBaseFeeChangeDenominator(uint256 timepoint, uint64 baseFeeChangeDenominator) returns()
func (istanbulParam *IstanbulParam) PackSetBaseFeeChangeDenominator(timepoint *big.Int, baseFeeChangeDenominator uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setBaseFeeChangeDenominator", timepoint, baseFeeChangeDenominator)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetBeneficiary is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c846d59.
//
// Solidity: function setBeneficiary(uint256 timepoint, address beneficiary) returns()
func (istanbulParam *IstanbulParam) PackSetBeneficiary(timepoint *big.Int, beneficiary common.Address) []byte {
	enc, err := istanbulParam.abi.Pack("setBeneficiary", timepoint, beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetBlockPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b765eb1.
//
// Solidity: function setBlockPeriod(uint256 timepoint, uint64 blockPeriod) returns()
func (istanbulParam *IstanbulParam) PackSetBlockPeriod(timepoint *big.Int, blockPeriod uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setBlockPeriod", timepoint, blockPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetElasticityMultiplier is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83f93bae.
//
// Solidity: function setElasticityMultiplier(uint256 timepoint, uint64 elasticityMultiplier) returns()
func (istanbulParam *IstanbulParam) PackSetElasticityMultiplier(timepoint *big.Int, elasticityMultiplier uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setElasticityMultiplier", timepoint, elasticityMultiplier)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetEmptyBlockPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4708026a.
//
// Solidity: function setEmptyBlockPeriod(uint256 timepoint, uint64 emptyBlockPeriod) returns()
func (istanbulParam *IstanbulParam) PackSetEmptyBlockPeriod(timepoint *big.Int, emptyBlockPeriod uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setEmptyBlockPeriod", timepoint, emptyBlockPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetEpochLength is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9d28f9b0.
//
// Solidity: function setEpochLength(uint256 timepoint, uint64 epochLength) returns()
func (istanbulParam *IstanbulParam) PackSetEpochLength(timepoint *big.Int, epochLength uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setEpochLength", timepoint, epochLength)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetGasLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4195c177.
//
// Solidity: function setGasLimit(uint256 timepoint, uint64 gasLimit) returns()
func (istanbulParam *IstanbulParam) PackSetGasLimit(timepoint *big.Int, gasLimit uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setGasLimit", timepoint, gasLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMaxBaseFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x131b23b5.
//
// Solidity: function setMaxBaseFee(uint256 timepoint, uint256 maxBaseFee) returns()
func (istanbulParam *IstanbulParam) PackSetMaxBaseFee(timepoint *big.Int, maxBaseFee *big.Int) []byte {
	enc, err := istanbulParam.abi.Pack("setMaxBaseFee", timepoint, maxBaseFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMaxRequestTimeout is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5941158a.
//
// Solidity: function setMaxRequestTimeout(uint256 timepoint, uint64 maxRequestTimeout) returns()
func (istanbulParam *IstanbulParam) PackSetMaxRequestTimeout(timepoint *big.Int, maxRequestTimeout uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setMaxRequestTimeout", timepoint, maxRequestTimeout)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinBaseFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea821efc.
//
// Solidity: function setMinBaseFee(uint256 timepoint, uint256 minBaseFee) returns()
func (istanbulParam *IstanbulParam) PackSetMinBaseFee(timepoint *big.Int, minBaseFee *big.Int) []byte {
	enc, err := istanbulParam.abi.Pack("setMinBaseFee", timepoint, minBaseFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetProposerPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xff035345.
//
// Solidity: function setProposerPolicy(uint256 timepoint, uint64 proposerPolicy) returns()
func (istanbulParam *IstanbulParam) PackSetProposerPolicy(timepoint *big.Int, proposerPolicy uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setProposerPolicy", timepoint, proposerPolicy)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetRequestTimeout is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa260957.
//
// Solidity: function setRequestTimeout(uint256 timepoint, uint64 requestTimeout) returns()
func (istanbulParam *IstanbulParam) PackSetRequestTimeout(timepoint *big.Int, requestTimeout uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setRequestTimeout", timepoint, requestTimeout)
	if err != nil {
		panic(err)
	}
	return enc
}

// IstanbulParamInitialized represents a Initialized event raised by the IstanbulParam contract.
type IstanbulParamInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const IstanbulParamInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (IstanbulParamInitialized) ContractEventName() string {
	return IstanbulParamInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (istanbulParam *IstanbulParam) UnpackInitializedEvent(log *types.Log) (*IstanbulParamInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != istanbulParam.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IstanbulParamInitialized)
	if len(log.Data) > 0 {
		if err := istanbulParam.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range istanbulParam.abi.Events[event].Inputs {
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
func (istanbulParam *IstanbulParam) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["NoPastUpdate"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackNoPastUpdateError(raw[4:])
	}
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IstanbulParamInvalidInitialization represents a InvalidInitialization error raised by the IstanbulParam contract.
type IstanbulParamInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func IstanbulParamInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (istanbulParam *IstanbulParam) UnpackInvalidInitializationError(raw []byte) (*IstanbulParamInvalidInitialization, error) {
	out := new(IstanbulParamInvalidInitialization)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IstanbulParamInvalidValue represents a InvalidValue error raised by the IstanbulParam contract.
type IstanbulParamInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func IstanbulParamInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (istanbulParam *IstanbulParam) UnpackInvalidValueError(raw []byte) (*IstanbulParamInvalidValue, error) {
	out := new(IstanbulParamInvalidValue)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IstanbulParamNoPastUpdate represents a NoPastUpdate error raised by the IstanbulParam contract.
type IstanbulParamNoPastUpdate struct {
	Timepoint *big.Int
	Current   *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoPastUpdate(uint256 timepoint, uint256 current)
func IstanbulParamNoPastUpdateErrorID() common.Hash {
	return common.HexToHash("0x92eca746dec091e901b48412975f9b1a3003ef363e8373001bbb481486916aba")
}

// UnpackNoPastUpdateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoPastUpdate(uint256 timepoint, uint256 current)
func (istanbulParam *IstanbulParam) UnpackNoPastUpdateError(raw []byte) (*IstanbulParamNoPastUpdate, error) {
	out := new(IstanbulParamNoPastUpdate)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "NoPastUpdate", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IstanbulParamNotInitializing represents a NotInitializing error raised by the IstanbulParam contract.
type IstanbulParamNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func IstanbulParamNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (istanbulParam *IstanbulParam) UnpackNotInitializingError(raw []byte) (*IstanbulParamNotInitializing, error) {
	out := new(IstanbulParamNotInitializing)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IstanbulParamOnlyCoinbase represents a OnlyCoinbase error raised by the IstanbulParam contract.
type IstanbulParamOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func IstanbulParamOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (istanbulParam *IstanbulParam) UnpackOnlyCoinbaseError(raw []byte) (*IstanbulParamOnlyCoinbase, error) {
	out := new(IstanbulParamOnlyCoinbase)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IstanbulParamOnlySystemContract represents a OnlySystemContract error raised by the IstanbulParam contract.
type IstanbulParamOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func IstanbulParamOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (istanbulParam *IstanbulParam) UnpackOnlySystemContractError(raw []byte) (*IstanbulParamOnlySystemContract, error) {
	out := new(IstanbulParamOnlySystemContract)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IstanbulParamOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the IstanbulParam contract.
type IstanbulParamOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func IstanbulParamOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (istanbulParam *IstanbulParam) UnpackOnlyZeroGasPriceError(raw []byte) (*IstanbulParamOnlyZeroGasPrice, error) {
	out := new(IstanbulParamOnlyZeroGasPrice)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IstanbulParamSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the IstanbulParam contract.
type IstanbulParamSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func IstanbulParamSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (istanbulParam *IstanbulParam) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*IstanbulParamSafeCastOverflowedUintDowncast, error) {
	out := new(IstanbulParamSafeCastOverflowedUintDowncast)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}
