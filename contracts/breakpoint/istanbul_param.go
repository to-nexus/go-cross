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
	ABI:        "[{\"type\":\"function\",\"name\":\"getParams\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParamsByIndex\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"beneficiary_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBaseFeeChangeDenominator\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseFeeChangeDenominator\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBeneficiary\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setElasticityMultiplier\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmptyBlockPeriod\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"emptyBlockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEpochLength\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasLimit\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBaseFee\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProposerPolicy\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposerPolicy\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequestTimeout\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestTimeout\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "IstanbulParam",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c806308a4f07214610114578063131b23b51461010f5780631c846d591461010a5780634195c177146101055780634708026a146101005780635941158a146100fb5780637b765eb1146100f657806383f93bae146100f15780639d28f9b0146100ec578063aa260957146100e7578063c4d66de8146100e2578063cf22854b146100dd578063ea821efc146100d8578063ebcac817146100d3578063f9d260d3146100ce5763ff035345146100c9575f80fd5b610af3565b610ad7565b6109db565b6109af565b610909565b6107a2565b61072f565b610699565b6105f4565b61055b565b6104bc565b610476565b6103ed565b610327565b6102b6565b6101e2565b9a98969492909d9c9b99979593916101a08c019e6001600160401b03168c526001600160401b031660208c01526001600160401b031660408b01526001600160401b031660608a015260808901610177916001600160401b03169052565b6001600160401b031660a08801526001600160401b031660c08701526001600160401b031660e08601526101008501526101208401526001600160a01b03166101408301526001600160401b0316610160820152610180016101e0916001600160401b03169052565b565b346102b25760203660031901126102b2576101fe600435610d41565b80546102ae60018301549261021d846001600160401b039060801c1690565b906102288560c01c90565b6002820154600383015460048401549290919061026c600561025d6001600160a01b0387169660a01c6001600160401b031690565b9601546001600160401b031690565b95604051998a996001600160401b03808360401c169216908060c01c906001600160401b038160801c16908d6001600160401b03808360401c16921690610119565b0390f35b5f80fd5b346102b25760403660031901126102b2576024356004356110071933016102e6576102e2600291610f1c565b0155005b630f22c43960e41b5f5261100860045260245ffd5b602435906001600160a01b03821682036102b257565b600435906001600160a01b03821682036102b257565b346102b25760403660031901126102b2576004356103436102fb565b61100833036102e6576001600160a01b038116801561038d5750600461036b61038b93610f1c565b0180546001600160a01b0319166001600160a01b03909216919091179055565b005b60849060405190632c648cf160e01b825260406004830152600b60448301526a62656e656669636961727960a81b60648301526024820152fd5b60409060031901126102b257600435906024356001600160401b03811681036102b25790565b346102b2576103fb366103c7565b61100833036102e6576001600160401b038116801561043f5750600561042361038b93610f1c565b01906001600160401b03166001600160401b0319825416179055565b60849060405190632c648cf160e01b825260406004830152600860448301526719d85cd31a5b5a5d60c21b60648301526024820152fd5b346102b257610484366103c7565b61100833036102e65761049961038b92610f1c565b80546001600160c01b031660c09290921b6001600160c01b031916919091179055565b346102b2576104ca366103c7565b61100833036102e6576001600160401b038116801561051b575060016104f261038b93610f1c565b019067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b60849060405190632c648cf160e01b82526040600483015260116044830152701b585e14995c5d595cdd151a5b595bdd5d607a1b60648301526024820152fd5b346102b257610569366103c7565b61100833036102e6576001600160401b03811680156105ba575061058f61038b92610f1c565b805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b60849060405190632c648cf160e01b825260406004830152600b60448301526a189b1bd8dad4195c9a5bd960aa1b60648301526024820152fd5b346102b257610602366103c7565b61100833036102e6576001600160401b03811680156106565750600161062a61038b93610f1c565b01805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b60849060405190632c648cf160e01b825260406004830152601460448301527332b630b9ba34b1b4ba3ca6bab63a34b83634b2b960611b60648301526024820152fd5b346102b2576106a7366103c7565b61100833036102e6576001600160401b03811680156106f557506106cd61038b92610f1c565b9067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b60849060405190632c648cf160e01b825260406004830152600b60448301526a0cae0dec6d098cadccee8d60ab1b60648301526024820152fd5b346102b25761073d366103c7565b61100833036102e6576001600160401b03811680156107655750600161042361038b93610f1c565b60849060405190632c648cf160e01b825260406004830152600e60448301526d1c995c5d595cdd151a5b595bdd5d60921b60648301526024820152fd5b346102b25760203660031901126102b2576107bb610311565b5f5160206116945f395f51905f5254906001600160401b036107ec60ff604085901c1615936001600160401b031690565b1680159081610901575b60011490816108f7575b1590816108ee575b506108df5761084b908261084260016001600160401b03195f5160206116945f395f51905f525416175f5160206116945f395f51905f5255565b6108aa57610ba4565b61085157005b61087b60ff60401b195f5160206116945f395f51905f5254165f5160206116945f395f51905f5255565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b6108da600160401b60ff60401b195f5160206116945f395f51905f525416175f5160206116945f395f51905f5255565b610ba4565b63f92ee8a960e01b5f5260045ffd5b9050155f610808565b303b159150610800565b8391506107f6565b346102b257610917366103c7565b61100833036102e6576001600160401b03811680156109635750600161093f61038b93610f1c565b0180546001600160c01b031660c09290921b6001600160c01b031916919091179055565b60849060405190632c648cf160e01b825260406004830152601860448301527f626173654665654368616e676544656e6f6d696e61746f72000000000000000060648301526024820152fd5b346102b25760403660031901126102b2576024356004356110071933016102e6576102e2600391610f1c565b346102b25760203660031901126102b2576004356001600160a01b038111610ac0576001600160a01b03165f90815260016020526040902080546102ae6001600160401b03821691604081901c6001600160401b031693610a4a608083901c6001600160401b03169260c01c90565b600182015460028301546003840154600485015493909260c081901c916001600160401b03608083901c811692604081901c821692911690610ab36005610aa46001600160a01b038b169a60a01c6001600160401b031690565b9a01546001600160401b031690565b996040519d8e9d8e610119565b6306dfcc6560e41b5f5260a060045260245260445ffd5b346102b2575f3660031901126102b25760205f54604051908152f35b346102b257610b01366103c7565b61100833036102e6576001600160401b03811680151580610b99575b610b5c57506004610b3061038b93610f1c565b01805467ffffffffffffffff60a01b191660a09290921b67ffffffffffffffff60a01b16919091179055565b60849060405190632c648cf160e01b825260406004830152600e60448301526d70726f706f736572506f6c69637960901b60648301526024820152fd5b506001811415610b1d565b413303610c44573a610c3557610c3290610c1d610bbf610cc2565b5f815262015180602082015291600160408401525f60608401526113886080840152603c60a08401525f610160840152600360c0840152600860e08401525f61010084018190526101208401526001600160a01b0316610140830152565b630c8458806101808201526110e0565b6110e0565b50565b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610c8857604052565b610c53565b604051906101e06101c083610c67565b604051906101e0604083610c67565b604051906101c0610cbd8184610c67565b368337565b604051906101c082018281106001600160401b03821117610c8857604052815f81525f60208201525f60408201525f60608201525f60808201525f60a08201525f60c08201525f60e08201525f6101008201525f6101208201525f6101408201525f6101608201525f6101808201526101a0610d3c610cac565b910152565b610d4a9061134e565b5f54905f918060058211610e6c575b6001600160601b039150929116905b828110610e0057505080610da657505f805260016020527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4990565b90565b610de7610de0610db8610da39361137f565b5f80527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630190565b5460601c90565b6001600160a01b03165f90815260016020526040902090565b90918082169080831860011c8201809211610e67575f80527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5638201546001600160601b0316841015610e555750915b90610d68565b929150610e619061138d565b90610e4f565b6110a8565b610e75826113ea565b8203918211610e67575f80526001600160601b03827f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630154166001600160601b038416105f14610ece57506001600160601b0390610d59565b9250610ee16001600160601b039161138d565b92610d59565b60405191905f835b600e8210610f06575050506101e06101c083610c67565b6001602081928554815201930191019091610eef565b610f2581610d41565b90815490806001600160401b038316146110a357611096610c2d9161108b6006610da396610fa1610f91610f57610c8d565b6001600160401b038a168152986001600160401b03604082901c1660208b01526001600160401b03608082901c1660408b01525b60c01c90565b6001600160401b03166060890152565b61100a610ffa6001830154610fcf610fbf826001600160401b031690565b6001600160401b031660808c0152565b6001600160401b03604082901c1660a08b01526001600160401b03608082901c1660c08b0152610f8b565b6001600160401b031660e0890152565b6002810154610100880152600381015461012088015260048101546001600160a01b03811661014089015261105d9061104c9060a01c6001600160401b031690565b6001600160401b0316610160890152565b61108561107460058301546001600160401b031690565b6001600160401b0316610180890152565b01610ee7565b6101a085015261139b565b6001600160401b03168252565b505090565b634e487b7160e01b5f52601160045260245ffd5b905f5b600e81106110cc57505050565b6001906020835193019281850155016110bf565b61131d611333610da39261132e61132960018060a01b035f5416948593845f52600160205261131060405f2061113d61112084516001600160401b031690565b825467ffffffffffffffff19166001600160401b03909116178255565b61118161115460208501516001600160401b031690565b82546fffffffffffffffff0000000000000000191660409190911b67ffffffffffffffff60401b16178255565b6111c061119860408501516001600160401b031690565b825467ffffffffffffffff60801b191660809190911b67ffffffffffffffff60801b16178255565b6111f76111d760608501516001600160401b031690565b82546001600160c01b031660c09190911b6001600160c01b031916178255565b6112716001820161121561112060808701516001600160401b031690565b61122c61115460a08701516001600160401b031690565b61124361119860c08701516001600160401b031690565b60e08501516001600160401b031681546001600160c01b031660c09190911b6001600160c01b031916179055565b610100830151600282015561012083015160038201556101408301516004820180546101608601516001600160a01b039093166001600160e01b03199091161760a09290921b67ffffffffffffffff60a01b169190911790556113016112e26101808501516001600160401b031690565b60058301906001600160401b03166001600160401b0319825416179055565b60066101a084015191016110bc565b516001600160401b031690565b6001600160401b031690565b61134e565b6115b1565b50506001600160a01b03165f90815260016020526040902090565b6001600160601b038111611368576001600160601b031690565b6306dfcc6560e41b5f52606060045260245260445ffd5b5f19810191908211610e6757565b9060018201809211610e6757565b6001600160401b0381116113b5576001600160401b031690565b6306dfcc6560e41b5f52604060045260245260445ffd5b81156113d6570490565b634e487b7160e01b5f52601260045260245ffd5b6001811115610da357806001600160801b82101561150b575b6114b16114a761149d61149361148961147f61146e6114b89760048a600160401b6114bd9c10156114fe575b6401000000008110156114f1575b620100008110156114e4575b6101008110156114d7575b60108110156114ca575b10156114c2575b60030260011c90565b611478818b6113cc565b0160011c90565b611478818a6113cc565b61147881896113cc565b61147881886113cc565b61147881876113cc565b61147881866113cc565b80936113cc565b821190565b900390565b60011b611465565b60041c9160021b9161145e565b60081c9160041b91611454565b60101c9160081b91611449565b60201c9160101b9161143d565b60401c9160201b9161142f565b50506114bd6114b86114b16114a761149d61149361148961147f61146e6115328a60801c90565b9850600160401b97506114039650505050505050565b908154600160401b811015610c88576001810180845581101561159d575f9283526020928390208251929093015160601b6bffffffffffffffffffffffff19166001600160601b039290921691909117910155565b634e487b7160e01b5f52603260045260245ffd5b5f54919291801561166a57610db86115c89161137f565b9081546115e56115de826001600160601b031690565b9160601c90565b926001600160601b0380841692169180831161165b57869203611625576116219250906001600160601b0382549181199060601b169116179055565b9190565b505061162190611645611636610c9d565b6001600160601b039092168252565b6001600160a01b03851660208201525b5f611548565b632520601d60e01b5f5260045ffd5b5061168e9061167a611636610c9d565b6001600160a01b0384166020820152611655565b5f919056fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212203037d2b53736aa585eabb13eca632a11e80cf3ead235d3566a3a04b56529bd6b64736f6c634300081c0033",
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

// PackGetParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08a4f072.
//
// Solidity: function getParams(uint256 timepoint) view returns(uint64 timestamp, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, address beneficiary, uint64 proposerPolicy, uint64 gasLimit)
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
	Timestamp                uint64
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
// Solidity: function getParams(uint256 timepoint) view returns(uint64 timestamp, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, address beneficiary, uint64 proposerPolicy, uint64 gasLimit)
func (istanbulParam *IstanbulParam) UnpackGetParams(data []byte) (GetParamsOutput, error) {
	out, err := istanbulParam.abi.Unpack("getParams", data)
	outstruct := new(GetParamsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
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
// the contract method with ID 0xebcac817.
//
// Solidity: function getParamsByIndex(uint256 idx) view returns(uint64 timestamp, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, address beneficiary, uint64 proposerPolicy, uint64 gasLimit)
func (istanbulParam *IstanbulParam) PackGetParamsByIndex(idx *big.Int) []byte {
	enc, err := istanbulParam.abi.Pack("getParamsByIndex", idx)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetParamsByIndexOutput serves as a container for the return parameters of contract
// method GetParamsByIndex.
type GetParamsByIndexOutput struct {
	Timestamp                uint64
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
// from invoking the contract method with ID 0xebcac817.
//
// Solidity: function getParamsByIndex(uint256 idx) view returns(uint64 timestamp, uint64 epochLength, uint64 blockPeriod, uint64 emptyBlockPeriod, uint64 requestTimeout, uint64 maxRequestTimeout, uint64 elasticityMultiplier, uint64 baseFeeChangeDenominator, uint256 maxBaseFee, uint256 minBaseFee, address beneficiary, uint64 proposerPolicy, uint64 gasLimit)
func (istanbulParam *IstanbulParam) UnpackGetParamsByIndex(data []byte) (GetParamsByIndexOutput, error) {
	out, err := istanbulParam.abi.Unpack("getParamsByIndex", data)
	outstruct := new(GetParamsByIndexOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
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
// Solidity: function numCheckpoints() view returns(uint256)
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
// Solidity: function numCheckpoints() view returns(uint256)
func (istanbulParam *IstanbulParam) UnpackNumCheckpoints(data []byte) (*big.Int, error) {
	out, err := istanbulParam.abi.Unpack("numCheckpoints", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
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
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["CheckpointUnorderedInsertion"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackCheckpointUnorderedInsertionError(raw[4:])
	}
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], istanbulParam.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return istanbulParam.UnpackInvalidValueError(raw[4:])
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

// IstanbulParamCheckpointUnorderedInsertion represents a CheckpointUnorderedInsertion error raised by the IstanbulParam contract.
type IstanbulParamCheckpointUnorderedInsertion struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CheckpointUnorderedInsertion()
func IstanbulParamCheckpointUnorderedInsertionErrorID() common.Hash {
	return common.HexToHash("0x2520601d9d60b717c34a36ad270857824c5a1ebbfd08ff39aba6930089495cfa")
}

// UnpackCheckpointUnorderedInsertionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CheckpointUnorderedInsertion()
func (istanbulParam *IstanbulParam) UnpackCheckpointUnorderedInsertionError(raw []byte) (*IstanbulParamCheckpointUnorderedInsertion, error) {
	out := new(IstanbulParamCheckpointUnorderedInsertion)
	if err := istanbulParam.abi.UnpackIntoInterface(out, "CheckpointUnorderedInsertion", raw); err != nil {
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
