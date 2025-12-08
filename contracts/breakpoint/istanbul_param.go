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
	ABI:        "[{\"type\":\"function\",\"name\":\"getParamIndex\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"idx\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParams\",\"inputs\":[{\"name\":\"timepoint_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParamsByIndex\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBaseFeeChangeDenominator\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setElasticityMultiplier\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmptyBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEpochLength\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasLimit\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProposerPolicy\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"IndexOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoPastUpdate\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"current\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "IstanbulParam",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c80630cc94d1a14610b9b5780630e24e72414610af4578063114be79f14610a4e5780632267e331146109fa578063436acf0b1461092c5780634e099f57146108d45780635078b472146105b25780636eac0108146104fd57806382779267146104e157806391e96f70146104c35780639bc7540014610445578063a4ba5b8b14610399578063ace30dbc146102d8578063c848f595146101f9578063c8a37198146100fb5763f9d260d3146100c9575f80fd5b346100f7575f3660031901126100f75760206100e660015461117e565b6001600160401b0360405191168152f35b5f80fd5b346100f75760403660031901126100f757610114610c79565b61011c610c8f565b61101333036101e4576001600160401b036101364361117e565b166001600160401b038316818111156101cf5750506001600160401b0381168015610195575061016861019392610d1e565b805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b005b60849060405190632c648cf160e01b825260406004830152600b60448301526a189b1bd8dad4195c9a5bd960aa1b60648301526024820152fd5b636581309b60e11b5f5260045260245260445ffd5b630f22c43960e41b5f5261101360045260245ffd5b346100f75760403660031901126100f757610212610c79565b61021a610c8f565b61101333036101e4576001600160401b036102344361117e565b166001600160401b038316818111156101cf5750506001600160401b038116801561028c5750600161026861019393610d1e565b0180546001600160c01b031660c09290921b6001600160c01b031916919091179055565b60849060405190632c648cf160e01b825260406004830152601860448301527f626173654665654368616e676544656e6f6d696e61746f72000000000000000060648301526024820152fd5b346100f75760203660031901126100f7576103016102fc6102f7610c79565b611145565b6110f1565b8054600182015460028301546003840154600490940154604080516001600160401b03808716825286831c81166020830152608087811c82168385015260c097881c60608401528187168184015286841c821660a084015286901c8116878301529490951c60e08601526101008501929092526101208401949094528184166101408401529290921c90911661016082015261018090f35b346100f75760403660031901126100f7576103b2610c79565b6103ba610c8f565b61101333036101e4576001600160401b036103d44361117e565b166001600160401b038316818111156101cf5750506001600160401b038116801561040b575061040661019392610d1e565b610ca5565b60849060405190632c648cf160e01b825260406004830152600b60448301526a0cae0dec6d098cadccee8d60ab1b60648301526024820152fd5b346100f75760403660031901126100f75761045e610c79565b610466610c8f565b61101333036101e4576001600160401b036104804361117e565b166001600160401b038316818111156101cf57610193836104a086610d1e565b80546001600160c01b031660c09290921b6001600160c01b031916919091179055565b346100f75760203660031901126100f75760206100e66102f7610c79565b346100f75760203660031901126100f7576103016102fc610c79565b346100f75760403660031901126100f757610516610c79565b61051e610c8f565b61101333036101e4576001600160401b036105384361117e565b166001600160401b038316818111156101cf5750506001600160401b03811680156105725750600161056c61019393610d1e565b01610ca5565b60849060405190632c648cf160e01b82526040600483015260116044830152701b585e14995c5d595cdd151a5b595bdd5d607a1b60648301526024820152fd5b346100f7576101603660031901126100f7576105cc610c79565b6105d4610c8f565b906044356001600160401b03811681036100f7576064356001600160401b03811681036100f757608435906001600160401b03821682036100f75760a435906001600160401b03821682036100f75760c435926001600160401b03841684036100f75761012435916001600160401b03831683036100f75761014435956001600160401b0387168097036100f7575f5160206112225f395f51905f5254986001600160401b0360ff8b60401c16159a16801590816108cc575b60011490816108c2575b1590816108b9575b506108aa578960016001600160401b03195f5160206112225f395f51905f525416175f5160206112225f395f51905f525561087a575b41330361086b573a61085c57604051986106ee8a610ccd565b5f8a5260208a015f905260408a015f905260608a015f905260808a015f905260a08a015f905260c08a015f905260e08a015f90526101008a015f90526101208a015f90526101408a015f90526101608a015f90526101e0604051906107538183610cfd565b3682376101808b01526107654361117e565b6001600160401b03168a526001600160401b031660208a01526001600160401b031660408901526001600160401b031660608801526001600160401b031660808701526001600160401b031660a08601526001600160401b03166101408501526001600160401b031660c08401526001600160401b031660e083015260e4356101008301526101043561012083015261016082015261080390610e98565b5061080a57005b60ff60401b195f5160206112225f395f51905f5254165f5160206112225f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b600160401b60ff60401b195f5160206112225f395f51905f525416175f5160206112225f395f51905f52556106d5565b63f92ee8a960e01b5f5260045ffd5b9050158b61069f565b303b159150610697565b8b915061068d565b346100f75760403660031901126100f7576108ed610c79565b61101333036101e4576001600160401b036109074361117e565b166001600160401b038216818111156101cf57602435600261092885610d1e565b0155005b346100f75760403660031901126100f757610945610c79565b61094d610c8f565b61101333036101e4576001600160401b036109674361117e565b166001600160401b038316818111156101cf576001600160401b03831684811515806109ef575b6109b25761099d600491610d1e565b01805467ffffffffffffffff19169091179055005b60848260405190632c648cf160e01b825260406004830152600e60448301526d70726f706f736572506f6c69637960901b60648301526024820152fd5b50600182141561098e565b346100f75760403660031901126100f757610a13610c79565b61101333036101e4576001600160401b03610a2d4361117e565b166001600160401b038216818111156101cf57602435600361092885610d1e565b346100f75760403660031901126100f757610a67610c79565b610a6f610c8f565b61101333036101e4576001600160401b03610a894361117e565b166001600160401b038316818111156101cf5750506001600160401b0381168015610abd5750600461056c61019393610d1e565b60849060405190632c648cf160e01b825260406004830152600860448301526719d85cd31a5b5a5d60c21b60648301526024820152fd5b346100f75760403660031901126100f757610b0d610c79565b610b15610c8f565b61101333036101e4576001600160401b03610b2f4361117e565b166001600160401b038316818111156101cf576001600160401b038316848115610b5e5761099d600191610d1e565b60848260405190632c648cf160e01b825260406004830152600e60448301526d1c995c5d595cdd151a5b595bdd5d60921b60648301526024820152fd5b346100f75760403660031901126100f757610bb4610c79565b610bbc610c8f565b61101333036101e4576001600160401b03610bd64361117e565b166001600160401b038316818111156101cf5750506001600160401b0381168015610c3657506001610c0a61019393610d1e565b01805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b60849060405190632c648cf160e01b825260406004830152601460448301527332b630b9ba34b1b4ba3ca6bab63a34b83634b2b960611b60648301526024820152fd5b600435906001600160401b03821682036100f757565b602435906001600160401b03821682036100f757565b9067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b6101a081019081106001600160401b03821117610ce957604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610ce957604052565b906001600160401b03821691825f525f602052826001600160401b0360405f20541614610e5d576102fc610d5191611145565b9160405190610d5f82610ccd565b83546001600160401b03811683526001600160401b038160401c1660208401526001600160401b038160801c16604084015260c01c606083015260018401546001600160401b03811660808401526001600160401b038160401c1660a08401526001600160401b038160801c1660c084015260c01c60e0830152600284015461010083015260038401546101208301526001600160401b03600485015481811661014085015260401c166101608301526040516005819501945f905b600f8210610e47575050610e44939450610e376101e082610cfd565b6101808301528152610e98565b90565b6001602081928954815201970191019095610e1b565b50905f525f60205260405f2090565b600154811015610e845760015f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b906001600160401b0382511691825f525f60205261018060405f20916001600160401b0380825116166001600160401b0319845416178355610ee76001600160401b0360208301511684610ca5565b6040810151835460608301516001600160c01b031960c09190911b166fffffffffffffffffffffffffffffffff90911667ffffffffffffffff60801b608093841b161717845581015160018401805467ffffffffffffffff19166001600160401b0392831617815560a0830151610fb392610f63911682610ca5565b60c0830151815467ffffffffffffffff60801b191660809190911b67ffffffffffffffff60801b1617815560e083015181546001600160c01b031660c09190911b6001600160c01b031916179055565b6101008101516002840155610120810151600384015561014081015160048401805467ffffffffffffffff19166001600160401b03928316178155610160830151610fff921690610ca5565b01515f5b600f81106110da5750505060015491600160401b831015610ce957916001810160015561102f81610e6c565b50505b801515806110b7575b1561109a575f19810190808211611086576110819061106961105c84610e6c565b90549060031b1c91610e6c565b90919082549060031b91821b915f19901b1916179055565b611032565b634e487b7160e01b5f52601160045260245ffd5b826110696110aa92949394610e6c565b5f525f60205260405f2090565b505f198101818111611086576110cd8491610e6c565b90549060031b1c1161103b565b600190602083519301926005828601015501611003565b6001600160401b03600154911690811015611136577fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601545f525f60205260405f2090565b634e23d03560e01b5f5260045ffd5b60015415611179576001600160401b0361115f91166111af565b8015611136575f19810190811161108657610e449061117e565b505f90565b6001600160401b038111611198576001600160401b031690565b6306dfcc6560e41b5f52604060045260245260445ffd5b5f90600154801561121a575b8083106111c757505090565b8083169080841860011c82018092116110865760015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf682015483101561120f57506111bb565b9250600101916111bb565b5050505f9056fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220842da6ab4db80473534586212053ac4c926b302fa7385c0422d5d546a6abcfa664736f6c634300081c0033",
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
// the contract method with ID 0x91e96f70.
//
// Solidity: function getParamIndex(uint64 timepoint) view returns(uint64 idx)
func (istanbulParam *IstanbulParam) PackGetParamIndex(timepoint uint64) []byte {
	enc, err := istanbulParam.abi.Pack("getParamIndex", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetParamIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91e96f70.
//
// Solidity: function getParamIndex(uint64 timepoint) view returns(uint64 idx)
func (istanbulParam *IstanbulParam) UnpackGetParamIndex(data []byte) (uint64, error) {
	out, err := istanbulParam.abi.Unpack("getParamIndex", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackGetParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xace30dbc.
//
// Solidity: function getParams(uint64 timepoint_) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, uint64 proposerPolicy, uint64 gasLimit)
func (istanbulParam *IstanbulParam) PackGetParams(timepoint uint64) []byte {
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
	ProposerPolicy           uint64
	GasLimit                 uint64
}

// UnpackGetParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xace30dbc.
//
// Solidity: function getParams(uint64 timepoint_) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, uint64 proposerPolicy, uint64 gasLimit)
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
	outstruct.ProposerPolicy = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.GasLimit = *abi.ConvertType(out[11], new(uint64)).(*uint64)
	return *outstruct, err

}

// PackGetParamsByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82779267.
//
// Solidity: function getParamsByIndex(uint64 idx) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, uint64 proposerPolicy, uint64 gasLimit)
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
	ProposerPolicy           uint64
	GasLimit                 uint64
}

// UnpackGetParamsByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82779267.
//
// Solidity: function getParamsByIndex(uint64 idx) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, uint64 proposerPolicy, uint64 gasLimit)
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
	outstruct.ProposerPolicy = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.GasLimit = *abi.ConvertType(out[11], new(uint64)).(*uint64)
	return *outstruct, err

}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5078b472.
//
// Solidity: function initialize(uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, uint64 proposerPolicy, uint64 gasLimit) returns()
func (istanbulParam *IstanbulParam) PackInitialize(epochLength uint64, blockPeriod uint64, emptyBlockPeriod uint64, requestTimeout uint64, maxRequestTimeout uint64, elasticityMultiplier uint64, baseFeeChangeDenominator uint64, maxBaseFee *big.Int, minBaseFee *big.Int, proposerPolicy uint64, gasLimit uint64) []byte {
	enc, err := istanbulParam.abi.Pack("initialize", epochLength, blockPeriod, emptyBlockPeriod, requestTimeout, maxRequestTimeout, elasticityMultiplier, baseFeeChangeDenominator, maxBaseFee, minBaseFee, proposerPolicy, gasLimit)
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
// the contract method with ID 0xc848f595.
//
// Solidity: function setBaseFeeChangeDenominator(uint64 timepoint, uint64 baseFeeChangeDenominator) returns()
func (istanbulParam *IstanbulParam) PackSetBaseFeeChangeDenominator(timepoint uint64, baseFeeChangeDenominator uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setBaseFeeChangeDenominator", timepoint, baseFeeChangeDenominator)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetBlockPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc8a37198.
//
// Solidity: function setBlockPeriod(uint64 timepoint, uint64 blockPeriod) returns()
func (istanbulParam *IstanbulParam) PackSetBlockPeriod(timepoint uint64, blockPeriod uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setBlockPeriod", timepoint, blockPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetElasticityMultiplier is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0cc94d1a.
//
// Solidity: function setElasticityMultiplier(uint64 timepoint, uint64 elasticityMultiplier) returns()
func (istanbulParam *IstanbulParam) PackSetElasticityMultiplier(timepoint uint64, elasticityMultiplier uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setElasticityMultiplier", timepoint, elasticityMultiplier)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetEmptyBlockPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9bc75400.
//
// Solidity: function setEmptyBlockPeriod(uint64 timepoint, uint64 emptyBlockPeriod) returns()
func (istanbulParam *IstanbulParam) PackSetEmptyBlockPeriod(timepoint uint64, emptyBlockPeriod uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setEmptyBlockPeriod", timepoint, emptyBlockPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetEpochLength is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4ba5b8b.
//
// Solidity: function setEpochLength(uint64 timepoint, uint64 epochLength) returns()
func (istanbulParam *IstanbulParam) PackSetEpochLength(timepoint uint64, epochLength uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setEpochLength", timepoint, epochLength)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetGasLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x114be79f.
//
// Solidity: function setGasLimit(uint64 timepoint, uint64 gasLimit) returns()
func (istanbulParam *IstanbulParam) PackSetGasLimit(timepoint uint64, gasLimit uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setGasLimit", timepoint, gasLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMaxBaseFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e099f57.
//
// Solidity: function setMaxBaseFee(uint64 timepoint, uint256 maxBaseFee) returns()
func (istanbulParam *IstanbulParam) PackSetMaxBaseFee(timepoint uint64, maxBaseFee *big.Int) []byte {
	enc, err := istanbulParam.abi.Pack("setMaxBaseFee", timepoint, maxBaseFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMaxRequestTimeout is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6eac0108.
//
// Solidity: function setMaxRequestTimeout(uint64 timepoint, uint64 maxRequestTimeout) returns()
func (istanbulParam *IstanbulParam) PackSetMaxRequestTimeout(timepoint uint64, maxRequestTimeout uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setMaxRequestTimeout", timepoint, maxRequestTimeout)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinBaseFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2267e331.
//
// Solidity: function setMinBaseFee(uint64 timepoint, uint256 minBaseFee) returns()
func (istanbulParam *IstanbulParam) PackSetMinBaseFee(timepoint uint64, minBaseFee *big.Int) []byte {
	enc, err := istanbulParam.abi.Pack("setMinBaseFee", timepoint, minBaseFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetProposerPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x436acf0b.
//
// Solidity: function setProposerPolicy(uint64 timepoint, uint64 proposerPolicy) returns()
func (istanbulParam *IstanbulParam) PackSetProposerPolicy(timepoint uint64, proposerPolicy uint64) []byte {
	enc, err := istanbulParam.abi.Pack("setProposerPolicy", timepoint, proposerPolicy)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetRequestTimeout is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e24e724.
//
// Solidity: function setRequestTimeout(uint64 timepoint, uint64 requestTimeout) returns()
func (istanbulParam *IstanbulParam) PackSetRequestTimeout(timepoint uint64, requestTimeout uint64) []byte {
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
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["IndexOutOfBounds"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackIndexOutOfBoundsError(raw[4:])
	}
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

// IstanbulParamIndexOutOfBounds represents a IndexOutOfBounds error raised by the IstanbulParam contract.
type IstanbulParamIndexOutOfBounds struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error IndexOutOfBounds()
func IstanbulParamIndexOutOfBoundsErrorID() common.Hash {
	return common.HexToHash("0x4e23d035f7b33ce60795ca9f0e8e6031d7a267562e8e7756f8c197bb3b4ef58b")
}

// UnpackIndexOutOfBoundsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error IndexOutOfBounds()
func (istanbulParam *IstanbulParam) UnpackIndexOutOfBoundsError(raw []byte) (*IstanbulParamIndexOutOfBounds, error) {
	out := new(IstanbulParamIndexOutOfBounds)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "IndexOutOfBounds", raw); err != nil {
		return nil, err
	}
	return out, nil
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
	Timepoint uint64
	Current   uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoPastUpdate(uint64 timepoint, uint64 current)
func IstanbulParamNoPastUpdateErrorID() common.Hash {
	return common.HexToHash("0xcb0261366b6742b0af8c864efa3e79048d223d92432c8b8966fcb8862f8bd3e2")
}

// UnpackNoPastUpdateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoPastUpdate(uint64 timepoint, uint64 current)
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
