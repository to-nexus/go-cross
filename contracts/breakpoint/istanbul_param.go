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
	ABI:        "[{\"type\":\"function\",\"name\":\"getParamIndex\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"idx\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParams\",\"inputs\":[{\"name\":\"timepoint_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParamsByIndex\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"timepoint\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBaseFeeChangeDenominator\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setElasticityMultiplier\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmptyBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEpochLength\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasLimit\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProposerPolicy\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"IndexOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoPastUpdate\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "IstanbulParam",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c806308a4f07214610bcc578063131b23b514610b8e5780634195c17714610af85780634708026a14610a885780635078b472146107665780635941158a146106c15780637b765eb1146105ff5780637eaf5a08146105e1578063827792671461052857806383f93bae1461045a5780639d28f9b0146103be578063aa26095714610325578063cf22854b1461023e578063ea821efc146101e6578063f9d260d3146101b85763ff035345146100c9575f80fd5b346101b45760403660031901126101b4576004356100e5610c01565b90611013330361019f576001600160401b03610100436110f5565b169182821115610189576001600160401b0316908115158061017e575b6101415761012c600491610ca6565b01805467ffffffffffffffff19169091179055005b60848260405190632c648cf160e01b825260406004830152600e60448301526d70726f706f736572506f6c69637960901b60648301526024820152fd5b50600182141561011d565b5063497653a360e11b5f5260045260245260445ffd5b630f22c43960e41b5f5261101360045260245ffd5b5f80fd5b346101b4575f3660031901126101b45760206101d56001546110f5565b6001600160401b0360405191168152f35b346101b4576101f436610beb565b90611013330361019f576001600160401b0361020f436110f5565b16808211156102285750610224600391610ca6565b0155005b9063497653a360e11b5f5260045260245260445ffd5b346101b45760403660031901126101b45760043561025a610c01565b611013330361019f576001600160401b03610274436110f5565b168083111561030f57506001600160401b03811680156102c35750600161029d6102c193610ca6565b0180546001600160c01b031660c09290921b6001600160c01b031916919091179055565b005b60849060405190632c648cf160e01b825260406004830152601860448301527f626173654665654368616e676544656e6f6d696e61746f72000000000000000060648301526024820152fd5b8263497653a360e11b5f5260045260245260445ffd5b346101b45760403660031901126101b457600435610341610c01565b90611013330361019f576001600160401b0361035c436110f5565b169182821115610189576001600160401b03169081156103815761012c600191610ca6565b60848260405190632c648cf160e01b825260406004830152600e60448301526d1c995c5d595cdd151a5b595bdd5d60921b60648301526024820152fd5b346101b45760403660031901126101b4576004356103da610c01565b611013330361019f576001600160401b036103f4436110f5565b168083111561030f57506001600160401b0381168015610420575061041b6102c192610ca6565b610c2d565b60849060405190632c648cf160e01b825260406004830152600b60448301526a0cae0dec6d098cadccee8d60ab1b60648301526024820152fd5b346101b45760403660031901126101b457600435610476610c01565b611013330361019f576001600160401b03610490436110f5565b168083111561030f57506001600160401b03811680156104e5575060016104b96102c193610ca6565b01805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b60849060405190632c648cf160e01b825260406004830152601460448301527332b630b9ba34b1b4ba3ca6bab63a34b83634b2b960611b60648301526024820152fd5b346101b45760203660031901126101b457610549610544610c17565b6110b0565b8054600182015460028301546003840154600490940154604080516001600160401b03808716825286831c81166020830152608087811c82168385015260c097881c60608401528187168184015286841c821660a084015286901c8116878301529490951c60e08601526101008501929092526101208401949094528184166101408401529290921c90911661016082015261018090f35b346101b45760203660031901126101b45760206101d5600435611071565b346101b45760403660031901126101b45760043561061b610c01565b611013330361019f576001600160401b03610635436110f5565b168083111561030f57506001600160401b0381168015610687575061065c6102c192610ca6565b805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b60849060405190632c648cf160e01b825260406004830152600b60448301526a189b1bd8dad4195c9a5bd960aa1b60648301526024820152fd5b346101b45760403660031901126101b4576004356106dd610c01565b611013330361019f576001600160401b036106f7436110f5565b168083111561030f57506001600160401b0381168015610726575060016107206102c193610ca6565b01610c2d565b60849060405190632c648cf160e01b82526040600483015260116044830152701b585e14995c5d595cdd151a5b595bdd5d607a1b60648301526024820152fd5b346101b4576101603660031901126101b457610780610c17565b610788610c01565b906044356001600160401b03811681036101b4576064356001600160401b03811681036101b457608435906001600160401b03821682036101b45760a435906001600160401b03821682036101b45760c435926001600160401b03841684036101b45761012435916001600160401b03831683036101b45761014435956001600160401b0387168097036101b4575f5160206111995f395f51905f5254986001600160401b0360ff8b60401c16159a1680159081610a80575b6001149081610a76575b159081610a6d575b50610a5e578960016001600160401b03195f5160206111995f395f51905f525416175f5160206111995f395f51905f5255610a2e575b413303610a1f573a610a1057604051986108a28a610c55565b5f8a5260208a015f905260408a015f905260608a015f905260808a015f905260a08a015f905260c08a015f905260e08a015f90526101008a015f90526101208a015f90526101408a015f90526101608a015f90526101e0604051906109078183610c85565b3682376101808b0152610919436110f5565b6001600160401b03168a526001600160401b031660208a01526001600160401b031660408901526001600160401b031660608801526001600160401b031660808701526001600160401b031660a08601526001600160401b03166101408501526001600160401b031660c08401526001600160401b031660e083015260e435610100830152610104356101208301526101608201526109b790610e25565b506109be57005b60ff60401b195f5160206111995f395f51905f5254165f5160206111995f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b600160401b60ff60401b195f5160206111995f395f51905f525416175f5160206111995f395f51905f5255610889565b63f92ee8a960e01b5f5260045ffd5b9050158b610853565b303b15915061084b565b8b9150610841565b346101b45760403660031901126101b457600435610aa4610c01565b90611013330361019f576001600160401b03610abf436110f5565b1680821115610228576102c183610ad584610ca6565b80546001600160c01b031660c09290921b6001600160c01b031916919091179055565b346101b45760403660031901126101b457600435610b14610c01565b611013330361019f576001600160401b03610b2e436110f5565b168083111561030f57506001600160401b0381168015610b57575060046107206102c193610ca6565b60849060405190632c648cf160e01b825260406004830152600860448301526719d85cd31a5b5a5d60c21b60648301526024820152fd5b346101b457610b9c36610beb565b90611013330361019f576001600160401b03610bb7436110f5565b16808211156102285750610224600291610ca6565b346101b45760203660031901126101b457610549610544600435611071565b60409060031901126101b4576004359060243590565b602435906001600160401b03821682036101b457565b600435906001600160401b03821682036101b457565b9067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b6101a081019081106001600160401b03821117610c7157604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610c7157604052565b805f525f602052806001600160401b0360405f20541614610dec57610ccd61054482611071565b9060405191610cdb83610c55565b80546001600160401b03811684526001600160401b038160401c1660208501526001600160401b038160801c16604085015260c01c606084015260018101546001600160401b03811660808501526001600160401b038160401c1660a08501526001600160401b038160801c1660c085015260c01c60e0840152600281015461010084015260038101546101208401526001600160401b03600482015481811661014086015260401c16610160840152604051906005829101905f905b600f8210610dd657505050610dcb610dd3939282610dc06101e06001600160401b0395610c85565b6101808501526110f5565b168152610e25565b90565b6001602081928554815201930191019091610d98565b5f525f60205260405f2090565b600154811015610e115760015f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b906001600160401b0382511691825f525f60205261018060405f20916001600160401b0380825116166001600160401b0319845416178355610e746001600160401b0360208301511684610c2d565b6040810151835460608301516001600160c01b031960c09190911b166fffffffffffffffffffffffffffffffff90911667ffffffffffffffff60801b608093841b161717845581015160018401805467ffffffffffffffff19166001600160401b0392831617815560a0830151610f4092610ef0911682610c2d565b60c0830151815467ffffffffffffffff60801b191660809190911b67ffffffffffffffff60801b1617815560e083015181546001600160c01b031660c09190911b6001600160c01b031916179055565b6101008101516002840155610120810151600384015561014081015160048401805467ffffffffffffffff19166001600160401b03928316178155610160830151610f8c921690610c2d565b01515f5b600f811061105a5750505060015491600160401b831015610c71579160018101600155610fbc81610df9565b50505b80151580611037575b15611027575f198101908082116110135761100e90610ff6610fe984610df9565b90549060031b1c91610df9565b90919082549060031b91821b915f19901b1916179055565b610fbf565b634e487b7160e01b5f52601160045260245ffd5b82610ff6610dec92949394610df9565b505f1981018181116110135761104d8491610df9565b90549060031b1c11610fc8565b600190602083519301926005828601015501610f90565b600154156110ab5761108290611126565b801561109c575f19810190811161101357610dd3906110f5565b634e23d03560e01b5f5260045ffd5b505f90565b6001600160401b0360015491169081101561109c577fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601545f525f60205260405f2090565b6001600160401b03811161110f576001600160401b031690565b6306dfcc6560e41b5f52604060045260245260445ffd5b5f906001548015611191575b80831061113e57505090565b8083169080841860011c82018092116110135760015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf68201548310156111865750611132565b925060010191611132565b5050505f9056fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220724d873de259ca8811f5c3e51ddf9978c959bce0803965ea6f92d8892a9d131064736f6c634300081c0033",
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
