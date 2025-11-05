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
	ABI:        "[{\"type\":\"function\",\"name\":\"getParamIndex\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"idx\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParams\",\"inputs\":[{\"name\":\"timepoint_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParamsByIndex\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBaseFeeChangeDenominator\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setElasticityMultiplier\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmptyBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEpochLength\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasLimit\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProposerPolicy\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoPastUpdate\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "IstanbulParam",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c806308a4f07214610b80578063131b23b514610b425780634195c17714610aac5780634708026a14610a3c5780635078b4721461071a5780635941158a146106755780637b765eb1146105d95780637eaf5a08146105bb578063827792671461050157806383f93bae146104595780639d28f9b0146103bd578063aa26095714610324578063cf22854b1461023d578063ea821efc146101e5578063f9d260d3146101b85763ff035345146100c9575f80fd5b346101b45760403660031901126101b4576004356100e5610bb5565b90611008330361019f576001600160401b03610100436111f6565b169182821115610189576001600160401b0316908115158061017e575b6101415761012c600491610cd1565b01805467ffffffffffffffff19169091179055005b60848260405190632c648cf160e01b825260406004830152600e60448301526d70726f706f736572506f6c69637960901b60648301526024820152fd5b50600182141561011d565b5063497653a360e11b5f5260045260245260445ffd5b630f22c43960e41b5f5261100860045260245ffd5b5f80fd5b346101b4575f3660031901126101b45760206101d45f546111f6565b6001600160401b0360405191168152f35b346101b4576101f336610b9f565b90611008330361019f576001600160401b0361020e436111f6565b16808211156102275750610223600391610cd1565b0155005b9063497653a360e11b5f5260045260245260445ffd5b346101b45760403660031901126101b457600435610259610bb5565b611008330361019f576001600160401b03610273436111f6565b168083111561030e57506001600160401b03811680156102c25750600161029c6102c093610cd1565b0180546001600160c01b031660c09290921b6001600160c01b031916919091179055565b005b60849060405190632c648cf160e01b825260406004830152601860448301527f626173654665654368616e676544656e6f6d696e61746f72000000000000000060648301526024820152fd5b8263497653a360e11b5f5260045260245260445ffd5b346101b45760403660031901126101b457600435610340610bb5565b90611008330361019f576001600160401b0361035b436111f6565b169182821115610189576001600160401b03169081156103805761012c600191610cd1565b60848260405190632c648cf160e01b825260406004830152600e60448301526d1c995c5d595cdd151a5b595bdd5d60921b60648301526024820152fd5b346101b45760403660031901126101b4576004356103d9610bb5565b611008330361019f576001600160401b036103f3436111f6565b168083111561030e57506001600160401b038116801561041f575061041a6102c092610cd1565b610be1565b60849060405190632c648cf160e01b825260406004830152600b60448301526a0cae0dec6d098cadccee8d60ab1b60648301526024820152fd5b346101b45760403660031901126101b457600435610475610bb5565b611008330361019f576001600160401b0361048f436111f6565b168083111561030e57506001600160401b03811680156104be575060016104b86102c093610cd1565b01610c5a565b60849060405190632c648cf160e01b825260406004830152601460448301527332b630b9ba34b1b4ba3ca6bab63a34b83634b2b960611b60648301526024820152fd5b346101b45760203660031901126101b45761052261051d610bcb565b610c85565b508054600182015460028301546003840154600490940154604080516001600160401b03808716825286831c81166020830152608087811c82168385015260c097881c60608401528187168184015286841c821660a084015286901c8116878301529490951c60e08601526101008501929092526101208401949094528184166101408401529290921c90911661016082015261018090f35b346101b45760203660031901126101b45760206101d46004356111ae565b346101b45760403660031901126101b4576004356105f5610bb5565b611008330361019f576001600160401b0361060f436111f6565b168083111561030e57506001600160401b038116801561063b57506106366102c092610cd1565b610c5a565b60849060405190632c648cf160e01b825260406004830152600b60448301526a189b1bd8dad4195c9a5bd960aa1b60648301526024820152fd5b346101b45760403660031901126101b457600435610691610bb5565b611008330361019f576001600160401b036106ab436111f6565b168083111561030e57506001600160401b03811680156106da575060016106d46102c093610cd1565b01610be1565b60849060405190632c648cf160e01b82526040600483015260116044830152701b585e14995c5d595cdd151a5b595bdd5d607a1b60648301526024820152fd5b346101b4576101603660031901126101b457610734610bcb565b61073c610bb5565b906044356001600160401b03811681036101b4576064356001600160401b03811681036101b457608435906001600160401b03821682036101b45760a435906001600160401b03821682036101b45760c435926001600160401b03841684036101b45761012435916001600160401b03831683036101b45761014435956001600160401b0387168097036101b4575f5160206112285f395f51905f5254986001600160401b0360ff8b60401c16159a1680159081610a34575b6001149081610a2a575b159081610a21575b50610a12578960016001600160401b03195f5160206112285f395f51905f525416175f5160206112285f395f51905f52556109e2575b4133036109d3573a6109c457604051986108568a610c09565b5f8a5260208a015f905260408a015f905260608a015f905260808a015f905260a08a015f905260c08a015f905260e08a015f90526101008a015f90526101208a015f90526101408a015f90526101608a015f90526101e0604051906108bb8183610c39565b3682376101808b01526108cd436111f6565b6001600160401b03168a526001600160401b031660208a01526001600160401b031660408901526001600160401b031660608801526001600160401b031660808701526001600160401b031660a08601526001600160401b03166101408501526001600160401b031660c08401526001600160401b031660e083015260e4356101008301526101043561012083015261016082015261096b90610e44565b5061097257005b60ff60401b195f5160206112285f395f51905f5254165f5160206112285f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b600160401b60ff60401b195f5160206112285f395f51905f525416175f5160206112285f395f51905f525561083d565b63f92ee8a960e01b5f5260045ffd5b9050158b610807565b303b1591506107ff565b8b91506107f5565b346101b45760403660031901126101b457600435610a58610bb5565b90611008330361019f576001600160401b03610a73436111f6565b1680821115610227576102c083610a8984610cd1565b80546001600160c01b031660c09290921b6001600160c01b031916919091179055565b346101b45760403660031901126101b457600435610ac8610bb5565b611008330361019f576001600160401b03610ae2436111f6565b168083111561030e57506001600160401b0381168015610b0b575060046106d46102c093610cd1565b60849060405190632c648cf160e01b825260406004830152600860448301526719d85cd31a5b5a5d60c21b60648301526024820152fd5b346101b457610b5036610b9f565b90611008330361019f576001600160401b03610b6b436111f6565b16808211156102275750610223600291610cd1565b346101b45760203660031901126101b45761052261051d6004356111ae565b60409060031901126101b4576004359060243590565b602435906001600160401b03821682036101b457565b600435906001600160401b03821682036101b457565b9067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b6101a081019081106001600160401b03821117610c2557604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610c2557604052565b805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b5f54811015610cbd575f8080526014919091027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630191565b634e487b7160e01b5f52603260045260245ffd5b610cdd61051d826111ae565b50908154916001600160401b03831692828414610e065760405193610d0185610c09565b84526001600160401b038160401c1660208501526001600160401b038160801c16604085015260c01c606084015260018101546001600160401b03811660808501526001600160401b038160401c1660a08501526001600160401b038160801c1660c085015260c01c60e0840152600281015461010084015260038101546101208401526001600160401b03600482015481811661014086015260401c16610160840152604051906005829101905f905b600f8210610df057505050610de5610ded939282610dda6101e06001600160401b0395610c39565b6101808501526111f6565b168152610e44565b90565b6001602081928554815201930191019091610db2565b5091505090565b90808214610e4057908154905f915b600f8310610e2a5750505050565b6001809194019283549481840155019192610e1c565b5050565b5f54600160401b811015610c255760018101805f55610e6282610c85565b50508111610fec575b8015158061117d575b15611000575f19810190808211610fec57610e98610e9183610c85565b5091610c85565b610fd957818103610eab575b5050610e6b565b600580836001600160401b0380610fd2965416166001600160401b0319855416178455610ee56001600160401b03825460401c1685610be1565b610efc6001600160401b03825460801c1685610c5a565b805484546001600160c01b03166001600160c01b0319909116178455610f8a60018501600183016001600160401b0380825416166001600160401b0319835416178255610f566001600160401b03825460401c1683610be1565b610f6d6001600160401b03825460801c1683610c5a565b5481546001600160c01b03166001600160c01b0319909116179055565b6002810154600285015560038101546003850155610fca600485016001600160401b036004840181808254161682198454161783555460401c1690610be1565b019101610e0d565b5f80610ea4565b634e487b7160e01b5f525f60045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b91909161100c81610c85565b939093610fd957806001600160401b0380610180935116166001600160401b031986541617855561104a6001600160401b0360208301511686610be1565b6110616001600160401b0360408301511686610c5a565b606081015185546001600160c01b031660c09190911b6001600160c01b0319161785556110fd600186016001600160401b0380608085015116166001600160401b03198254161781556110c16001600160401b0360a08501511682610be1565b6110d86001600160401b0360c08501511682610c5a565b60e083015181546001600160c01b031660c09190911b6001600160c01b031916179055565b6101008101516002860155610120810151600386015561014081015160048601805467ffffffffffffffff19166001600160401b03928316178155610160830151611149921690610be1565b01515f5b600f8110611166575050611162919250610c85565b5090565b60019060208351930192600582880101550161114d565b505f198101818111610fec5761119a6001600160401b0391610c85565b5054166001600160401b0383511610610e74565b5f5480915f905b8382106111c9575b5050610ded91506111f6565b9091816001600160401b036111dd85610c85565b505416116111f0575060018201906111b5565b916111bd565b6001600160401b038111611210576001600160401b031690565b6306dfcc6560e41b5f52604060045260245260445ffdfef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220c783f1624ace9a1b6daa26c2ffbb9736dd75e49df5756a57901219f6f9ff1a1064736f6c634300081c0033",
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
// Solidity: function getParams(uint256 timepoint_) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, uint64 proposerPolicy, uint64 gasLimit)
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
	ProposerPolicy           uint64
	GasLimit                 uint64
}

// UnpackGetParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x08a4f072.
//
// Solidity: function getParams(uint256 timepoint_) view returns(uint64 timepoint, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, uint64 proposerPolicy, uint64 gasLimit)
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
