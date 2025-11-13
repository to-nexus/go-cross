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
	ABI:        "[{\"type\":\"function\",\"name\":\"addToBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyValidator\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getKeeper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLeavingTime\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorAddress\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeAmount\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnstakingAmount\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorAddress\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isJailedValidator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLeavingValidator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"leaveValidator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minStakeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"offlineJailDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"offlineSlashAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFromBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeeper\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStakeAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOfflineJailDuration\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOfflineSlashAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnbondingPeriod\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidatorThreshold\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOffline\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSlashedAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unbondingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unjail\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperChanged\",\"inputs\":[{\"name\":\"previousKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperUpdated\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinStakeAmountUpdated\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineJailDurationUpdated\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineSlashAmountUpdated\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeClaimed\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnBlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnbondingPeriodUpdated\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorApplied\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorJailed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"jailUntil\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorLeft\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"leavingTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSlashed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"jailUntil\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorThresholdUpdated\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUnjailed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyLeaving\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateOperatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JailedValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInLeavingState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotJailedValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyKeeper\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StillInJailPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StillInUnbondingPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "StakeHub",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f5f3560e01c806302c75532146113575780630c2eb40314611304578063114eaf5514611270578063391b6f4e1461123b5780633f4ba83a146111ad57806340550a1c14611171578063417c73a7146110f157806342fce3cd146110d357806343a769d91461103b57806344690e5c14610ebb578063449ecfe614610dcd5780634a49ac4c14610d3d5780634fd101d714610d1f5780635bc0fbc714610c885780635c975abb14610c59578063603eff8b14610aeb5780636cf6d67514610acd5780637071688a14610ab0578063748747e6146109ea5780638456cb591461094a5780639679904a146108b3578063a9cc60f314610895578063b6e5395f14610850578063ba723eda14610832578063bff02e2014610788578063c4d66de81461059c578063d137933214610554578063eaa43d1414610514578063eb3211731461032c578063eb4af04514610284578063f158628d1461023e578063f188768414610220578063f292a643146101e55763fc1332db14610190575f80fd5b346101e25760203660031901126101e257600435906001600160a01b0382168083036101e0576101c16020936117fe565b156101da5781526002825260409020545b604051908152f35b506101d2565b505b80fd5b50346101e25760203660031901126101e257600435906001600160a01b03821682036101e2576020610216836117fe565b6040519015158152f35b50346101e257806003193601126101e2576020600554604051908152f35b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e05790602091815260038252604060018060a01b0391205416604051908152f35b50346101e25760203660031901126101e25760043561100833036103175780156102da576020817f8448c02797b448f4946bc25b3bf925e5556d1df822c944da701c54bab8a3162f92600555604051908152a180f35b60849060405190632c648cf160e01b825260406004830152600e60448301526d1b5a5b94dd185ad9505b5bdd5b9d60921b60648301526024820152fd5b630f22c43960e41b8252611008600452602482fd5b50346101e257806003193601126101e25761034561177e565b60ff61035033611543565b54166105015761036b335f52600160205260405f2054151590565b156104f2577f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c6104e35760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d33815260026020526040812080549081156104d457600481015480156104c55742106104b6576001810154908390556001600160a01b03168280808085335af13d156104b1573d67ffffffffffffffff811161049d5760405190610428601f8201601f19166020018361157b565b81528460203d92013e5b15610489576040519182527f8766f9fdb3d7ff5c4fe7ccf145668fbc991a789c4d655f9535b0e08c4964a04d60203393a3807f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d80f35b63c39ba1a960e01b83526004829052602483fd5b634e487b7160e01b85526041600452602485fd5b610432565b630e8c992760e01b8352600483fd5b63281f6b9160e01b8452600484fd5b633b4f091f60e21b8352600483fd5b633ee5aeb560e01b8152600490fd5b632ec5b44960e01b8152600490fd5b63033788ff60e01b815233600452602490fd5b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e0576040826004926020945260028452200154604051908152f35b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e0579060209181526002825260016040818060a01b039220015416604051908152f35b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e0575f5160206118f35f395f51905f52549060ff8260401c16159167ffffffffffffffff811680159081610780575b6001149081610776575b15908161076d575b5061075e5767ffffffffffffffff1981166001175f5160206118f35f395f51905f525582610732575b50413303610723573a6107145769021e19e0c9bab2400000600555600a60065562093a80600755678ac7230489e800006008556202a30060095561066d611887565b610675611887565b6bffffffffffffffffffffffff60a01b5f5160206118b35f395f51905f525416175f5160206118b35f395f51905f52556106ad611887565b6106b5611887565b6106bc5780f35b68ff0000000000000000195f5160206118f35f395f51905f5254165f5160206118f35f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6383f1b1d360e01b8352600483fd5b63022d8c9560e31b8352600483fd5b68ffffffffffffffffff191668010000000000000001175f5160206118f35f395f51905f52555f61062b565b63f92ee8a960e01b8452600484fd5b9050155f610602565b303b1591506105fa565b8491506105f0565b50346101e25760403660031901126101e2576107a8602435600435611651565b604051606080825284519082018190526080820195919460200190825b81811061081357505050838503602085015260208084519687815201930190945b8086106107fb57505082935060408301520390f35b909260208060019286518152019401950194906107e6565b82516001600160a01b03168852602097880197909201916001016107c5565b50346101e257806003193601126101e2576020600454604051908152f35b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e057600660408360ff936020955260028552200154166040519015158152f35b50346101e257806003193601126101e2576020600854604051908152f35b50346101e25760203660031901126101e2576004356110083303610317578015610909576020817fff6e1245b60ea5b68ddca24a93cb8707677289cbdbdec9ecb7ccfc25541a774992600855604051908152a180f35b60849060405190632c648cf160e01b82526040600483015260126044830152711bd9999b1a5b9954db185cda105b5bdd5b9d60721b60648301526024820152fd5b50346101e257806003193601126101e2575f5160206118b35f395f51905f52546001600160a01b031633036109db5761098161177e565b61098961177e565b600160ff195f5160206118d35f395f51905f525416175f5160206118d35f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b63c60eb33560e01b8152600490fd5b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e0576110083303610317578015610a7b575f5160206118b35f395f51905f5254816001600160a01b0382167f068b48a2fe7f498b57ff6da64f075ae658fde8d77124b092e62b3dc58d91ce358580a36001600160a01b031916175f5160206118b35f395f51905f525580f35b60848260405190632c648cf160e01b825260406004830152600660448301526535b2b2b832b960d11b60648301526024820152fd5b50346101e257806003193601126101e25760209054604051908152f35b50346101e257806003193601126101e2576020600754604051908152f35b50346101e257806003193601126101e257610b0461177e565b60ff610b0f33611543565b541661050157610b2a335f52600160205260405f2054151590565b156104f25733815260026020526040812060048101908154610c4a5760ff600682015416610c3b57610b5e600754426115d2565b8092556001808060a01b03910154166040519182527fe73d82fe262870ff059c299a893937ea90edc32a71143234a83299297ff8966660203393a36110063b156101e257604051632961046560e21b815233600482015281908181602481836110065af18015610c1b57610c26575b50506110063b156101e25780604051632949582d60e01b81523360048201528160248201528181604481836110065af18015610c1b57610c0a5750f35b81610c149161157b565b6101e25780f35b6040513d84823e3d90fd5b81610c309161157b565b6101e257805f610bcd565b633c59519b60e11b8352600483fd5b6333dd977160e01b8352600483fd5b50346101e257806003193601126101e257602060ff5f5160206118d35f395f51905f5254166040519015158152f35b50346101e25760203660031901126101e2576004356110083303610317578015610cde576020817f324143af489ab0ba77b0d3580e2486ee612b673b87db2179c589f61ac532fd5992600655604051908152a180f35b60849060405190632c648cf160e01b82526040600483015260126044830152711d985b1a59185d1bdc951a1c995cda1bdb1960721b60648301526024820152fd5b50346101e257806003193601126101e2576020600654604051908152f35b50346101e25760203660031901126101e2576004356001600160a01b03811690818103610dc9575f5160206118b35f395f51905f52546001600160a01b03163303610dba57610d8b90611543565b805460ff191690557fe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de40118280a280f35b63c60eb33560e01b8352600483fd5b8280fd5b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e057610dfc61177e565b60ff610e0733611543565b5416610ea857610e22335f52600160205260405f2054151590565b15610e9957808252600260205260408220600681019081549060ff821615610e8a57600501544210610e7b5760ff191690557f9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f198280a280f35b63065f425760e21b8452600484fd5b635103560d60e11b8552600485fd5b632ec5b44960e01b8252600482fd5b63033788ff60e01b825233600452602482fd5b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e0576110043303611026578152600360209081526040808320546001600160a01b03165f818152600190935291205415610e995780829182526002602052807f83b04ecf7330997e742429a641e136d9f3698c3e9ac9cb9ce0cc2d6da36a244d604080852080546008548082115f1461101b57610f5e9080926115b1565b8255610f6c816004546115d2565b600455610f7b600954426115d2565b91826005820180548211611013575b505060068101600160ff19825416179055600260018060a01b03910154167feb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e9582580260208551858152a282519182526020820152a26110063b156110105760405190632961046560e21b825260048201528181602481836110065af18015610c1b57610c0a5750f35b50fd5b55825f610f8a565b50610f5e81806115b1565b630f22c43960e41b8252611004600452602482fd5b50346101e25760203660031901126101e2576004356110083303610317578015611091576020817fbdf592a42adc0296ad8a75a5e7b8a0912c37793a1a275d78275bc335911c6bd392600955604051908152a180f35b60849060405190632c648cf160e01b825260406004830152601360448301527237b3333634b732a530b4b6223ab930ba34b7b760691b60648301526024820152fd5b50346101e257806003193601126101e2576020600954604051908152f35b50346101e25760203660031901126101e2576004356001600160a01b03811690818103610dc9575f5160206118b35f395f51905f52546001600160a01b03163303610dba5761113f90611543565b805460ff191660011790557f7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec61941868280a280f35b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e057604082610216926020945260028452206117cd565b50346101e257806003193601126101e2575f5160206118b35f395f51905f52546001600160a01b031633036109db576111e46117a5565b6111ec6117a5565b60ff195f5160206118d35f395f51905f5254165f5160206118d35f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a180f35b50346101e257806003193601126101e2575f5160206118b35f395f51905f52546040516001600160a01b039091168152602090f35b50346101e25760203660031901126101e25760043561100833036103175780156112c6576020817f38a1644f59872db6fd17fdced395495fcaa3ca7d2825a0704db5a3acbd1dd06492600755604051908152a180f35b60849060405190632c648cf160e01b825260406004830152600f60448301526e1d5b989bdb991a5b99d4195c9a5bd9608a1b60648301526024820152fd5b50346101e25760203660031901126101e2576004356001600160a01b038116908190036101e0579060209181526002825260408120906004820154155f14611350575054604051908152f35b90506101d2565b5060203660031901126114f0576004356001600160a01b038116908190036114f05761138161177e565b60ff61138c33611543565b5416611530578015611521576005543410611512576113b6335f52600160205260405f2054151590565b611503575f818152600360205260409020546001600160a01b03166114f4576113de3361182d565b505f8181526003602081815260408084208054336001600160a01b03199182168117909255818652600280855283872060018101805484168a1790559081018054909216831790915534808255429582019590955560040194909455519182529192917f0d83b503e698d0879dc7ad5505a19535ed048371a78d4b5908e85ca253a31a3c91a36110063b156114f057604051632961046560e21b81523360048201525f81602481836110065af180156114e5576114d2575b506110063b156101e25780604051632949582d60e01b81523360048201523360248201528181604481836110065af18015610c1b57610c0a5750f35b6114de91505f9061157b565b5f5f611496565b6040513d5f823e3d90fd5b5f80fd5b63813f326760e01b5f5260045ffd5b639cc20ced60e01b5f5260045ffd5b63e008b5f960e01b5f5260045ffd5b63713ce51160e01b5f5260045ffd5b63033788ff60e01b5f523360045260245ffd5b6001600160a01b03165f9081527f46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e5016020526040902090565b90601f8019910116810190811067ffffffffffffffff82111761159d57604052565b634e487b7160e01b5f52604160045260245ffd5b919082039182116115be57565b634e487b7160e01b5f52601160045260245ffd5b919082018092116115be57565b67ffffffffffffffff811161159d5760051b60200190565b90611601826115df565b61160e604051918261157b565b828152809261161f601f19916115df565b0190602036910137565b805182101561163d5760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b905f5490818310156117495780611740575061166e815b836115d2565b82828211611738575b611683919493946115b1565b9161168d836115f7565b93611697846115f7565b935f915f54925b8281106116ab5750505050565b6116b581836115d2565b908482101561163d577f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101546001600160a01b039081165f9081526002602052604090206001818101549093921661170f838c611629565b52611719816117cd565b1561173157545b61172a828a611629565b520161169e565b505f611720565b829150611677565b61166e90611668565b50905060209160405161175c848261157b565b5f81525f36813760405193611771818661157b565b5f8552505f368137929190565b60ff5f5160206118d35f395f51905f52541661179657565b63d93c066560e01b5f5260045ffd5b60ff5f5160206118d35f395f51905f525416156117be57565b638dfc202b60e01b5f5260045ffd5b60048101541590816117f0575b816117e3575090565b60ff915060060154161590565b8054600554111591506117da565b6001600160a01b03165f908152600260205260409020600401548015159081611825575090565b905042101590565b805f52600160205260405f2054155f14611882575f546801000000000000000081101561159d57600181015f555f5481101561163d5781905f805260205f2001555f54905f52600160205260405f2055600190565b505f90565b60ff5f5160206118f35f395f51905f525460401c16156118a357565b631afcd79f60e31b5f5260045ffdfe46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e500cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220f40b583b5c1114a54275c41d90bd27af5de4b22883aa5a44f9aa92d39582be6264736f6c634300081c0033",
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

// PackGetLeavingTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeaa43d14.
//
// Solidity: function getLeavingTime(address operatorAddr) view returns(uint256)
func (stakeHub *StakeHub) PackGetLeavingTime(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getLeavingTime", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetLeavingTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeaa43d14.
//
// Solidity: function getLeavingTime(address operatorAddr) view returns(uint256)
func (stakeHub *StakeHub) UnpackGetLeavingTime(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("getLeavingTime", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
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

// PackGetUnstakingAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc1332db.
//
// Solidity: function getUnstakingAmount(address operatorAddr) view returns(uint256)
func (stakeHub *StakeHub) PackGetUnstakingAmount(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getUnstakingAmount", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUnstakingAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfc1332db.
//
// Solidity: function getUnstakingAmount(address operatorAddr) view returns(uint256)
func (stakeHub *StakeHub) UnpackGetUnstakingAmount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("getUnstakingAmount", data)
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

// PackIsJailedValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb6e5395f.
//
// Solidity: function isJailedValidator(address operatorAddr) view returns(bool)
func (stakeHub *StakeHub) PackIsJailedValidator(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("isJailedValidator", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsJailedValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb6e5395f.
//
// Solidity: function isJailedValidator(address operatorAddr) view returns(bool)
func (stakeHub *StakeHub) UnpackIsJailedValidator(data []byte) (bool, error) {
	out, err := stakeHub.abi.Unpack("isJailedValidator", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsLeavingValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf292a643.
//
// Solidity: function isLeavingValidator(address operatorAddr) view returns(bool)
func (stakeHub *StakeHub) PackIsLeavingValidator(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("isLeavingValidator", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsLeavingValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf292a643.
//
// Solidity: function isLeavingValidator(address operatorAddr) view returns(bool)
func (stakeHub *StakeHub) UnpackIsLeavingValidator(data []byte) (bool, error) {
	out, err := stakeHub.abi.Unpack("isLeavingValidator", data)
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

// PackOfflineJailDuration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42fce3cd.
//
// Solidity: function offlineJailDuration() view returns(uint256)
func (stakeHub *StakeHub) PackOfflineJailDuration() []byte {
	enc, err := stakeHub.abi.Pack("offlineJailDuration")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOfflineJailDuration is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x42fce3cd.
//
// Solidity: function offlineJailDuration() view returns(uint256)
func (stakeHub *StakeHub) UnpackOfflineJailDuration(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("offlineJailDuration", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOfflineSlashAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9cc60f3.
//
// Solidity: function offlineSlashAmount() view returns(uint256)
func (stakeHub *StakeHub) PackOfflineSlashAmount() []byte {
	enc, err := stakeHub.abi.Pack("offlineSlashAmount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOfflineSlashAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9cc60f3.
//
// Solidity: function offlineSlashAmount() view returns(uint256)
func (stakeHub *StakeHub) UnpackOfflineSlashAmount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("offlineSlashAmount", data)
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

// PackSetOfflineJailDuration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43a769d9.
//
// Solidity: function setOfflineJailDuration(uint256 duration) returns()
func (stakeHub *StakeHub) PackSetOfflineJailDuration(duration *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setOfflineJailDuration", duration)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetOfflineSlashAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9679904a.
//
// Solidity: function setOfflineSlashAmount(uint256 amount) returns()
func (stakeHub *StakeHub) PackSetOfflineSlashAmount(amount *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setOfflineSlashAmount", amount)
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

// PackSetValidatorThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5bc0fbc7.
//
// Solidity: function setValidatorThreshold(uint256 count) returns()
func (stakeHub *StakeHub) PackSetValidatorThreshold(count *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setValidatorThreshold", count)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSlashOffline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44690e5c.
//
// Solidity: function slashOffline(address validatorAddr) returns()
func (stakeHub *StakeHub) PackSlashOffline(validatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("slashOffline", validatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalSlashedAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba723eda.
//
// Solidity: function totalSlashedAmount() view returns(uint256)
func (stakeHub *StakeHub) PackTotalSlashedAmount() []byte {
	enc, err := stakeHub.abi.Pack("totalSlashedAmount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSlashedAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xba723eda.
//
// Solidity: function totalSlashedAmount() view returns(uint256)
func (stakeHub *StakeHub) UnpackTotalSlashedAmount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("totalSlashedAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
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

// PackUnjail is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x449ecfe6.
//
// Solidity: function unjail(address operatorAddr) returns()
func (stakeHub *StakeHub) PackUnjail(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("unjail", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
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

// StakeHubOfflineJailDurationUpdated represents a OfflineJailDurationUpdated event raised by the StakeHub contract.
type StakeHubOfflineJailDurationUpdated struct {
	Duration *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const StakeHubOfflineJailDurationUpdatedEventName = "OfflineJailDurationUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubOfflineJailDurationUpdated) ContractEventName() string {
	return StakeHubOfflineJailDurationUpdatedEventName
}

// UnpackOfflineJailDurationUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OfflineJailDurationUpdated(uint256 duration)
func (stakeHub *StakeHub) UnpackOfflineJailDurationUpdatedEvent(log *types.Log) (*StakeHubOfflineJailDurationUpdated, error) {
	event := "OfflineJailDurationUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubOfflineJailDurationUpdated)
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

// StakeHubOfflineSlashAmountUpdated represents a OfflineSlashAmountUpdated event raised by the StakeHub contract.
type StakeHubOfflineSlashAmountUpdated struct {
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const StakeHubOfflineSlashAmountUpdatedEventName = "OfflineSlashAmountUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubOfflineSlashAmountUpdated) ContractEventName() string {
	return StakeHubOfflineSlashAmountUpdatedEventName
}

// UnpackOfflineSlashAmountUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OfflineSlashAmountUpdated(uint256 amount)
func (stakeHub *StakeHub) UnpackOfflineSlashAmountUpdatedEvent(log *types.Log) (*StakeHubOfflineSlashAmountUpdated, error) {
	event := "OfflineSlashAmountUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubOfflineSlashAmountUpdated)
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

// StakeHubValidatorJailed represents a ValidatorJailed event raised by the StakeHub contract.
type StakeHubValidatorJailed struct {
	OperatorAddr common.Address
	JailUntil    *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorJailedEventName = "ValidatorJailed"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorJailed) ContractEventName() string {
	return StakeHubValidatorJailedEventName
}

// UnpackValidatorJailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorJailed(address indexed operatorAddr, uint256 jailUntil)
func (stakeHub *StakeHub) UnpackValidatorJailedEvent(log *types.Log) (*StakeHubValidatorJailed, error) {
	event := "ValidatorJailed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorJailed)
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

// StakeHubValidatorSlashed represents a ValidatorSlashed event raised by the StakeHub contract.
type StakeHubValidatorSlashed struct {
	OperatorAddr common.Address
	JailUntil    *big.Int
	SlashAmount  *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorSlashedEventName = "ValidatorSlashed"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorSlashed) ContractEventName() string {
	return StakeHubValidatorSlashedEventName
}

// UnpackValidatorSlashedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorSlashed(address indexed operatorAddr, uint256 jailUntil, uint256 slashAmount)
func (stakeHub *StakeHub) UnpackValidatorSlashedEvent(log *types.Log) (*StakeHubValidatorSlashed, error) {
	event := "ValidatorSlashed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorSlashed)
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

// StakeHubValidatorUnjailed represents a ValidatorUnjailed event raised by the StakeHub contract.
type StakeHubValidatorUnjailed struct {
	OperatorAddr common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorUnjailedEventName = "ValidatorUnjailed"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorUnjailed) ContractEventName() string {
	return StakeHubValidatorUnjailedEventName
}

// UnpackValidatorUnjailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUnjailed(address indexed operatorAddr)
func (stakeHub *StakeHub) UnpackValidatorUnjailedEvent(log *types.Log) (*StakeHubValidatorUnjailed, error) {
	event := "ValidatorUnjailed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorUnjailed)
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
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidClaim"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidClaimError(raw[4:])
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
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["JailedValidator"].ID.Bytes()[:4]) {
		return stakeHub.UnpackJailedValidatorError(raw[4:])
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
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotJailedValidator"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotJailedValidatorError(raw[4:])
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
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["StillInJailPeriod"].ID.Bytes()[:4]) {
		return stakeHub.UnpackStillInJailPeriodError(raw[4:])
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

// StakeHubInvalidClaim represents a InvalidClaim error raised by the StakeHub contract.
type StakeHubInvalidClaim struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidClaim()
func StakeHubInvalidClaimErrorID() common.Hash {
	return common.HexToHash("0xed3c247cc0268b6a8ea6ea7af517a687f38b09876966c761867fb56e61e3d0eb")
}

// UnpackInvalidClaimError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidClaim()
func (stakeHub *StakeHub) UnpackInvalidClaimError(raw []byte) (*StakeHubInvalidClaim, error) {
	out := new(StakeHubInvalidClaim)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidClaim", raw); err != nil {
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

// StakeHubJailedValidator represents a JailedValidator error raised by the StakeHub contract.
type StakeHubJailedValidator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error JailedValidator()
func StakeHubJailedValidatorErrorID() common.Hash {
	return common.HexToHash("0x78b2a3366692ee51cd0e427429ed8d526c9090e05446ec218ab6a9eaa846431b")
}

// UnpackJailedValidatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error JailedValidator()
func (stakeHub *StakeHub) UnpackJailedValidatorError(raw []byte) (*StakeHubJailedValidator, error) {
	out := new(StakeHubJailedValidator)
	if err := stakeHub.abi.UnpackIntoInterface(out, "JailedValidator", raw); err != nil {
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

// StakeHubNotJailedValidator represents a NotJailedValidator error raised by the StakeHub contract.
type StakeHubNotJailedValidator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotJailedValidator()
func StakeHubNotJailedValidatorErrorID() common.Hash {
	return common.HexToHash("0xa206ac1a18393ea5c89e8aa5e2a5c0f20bdb5f463ed25133d9f23d36dd49af50")
}

// UnpackNotJailedValidatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotJailedValidator()
func (stakeHub *StakeHub) UnpackNotJailedValidatorError(raw []byte) (*StakeHubNotJailedValidator, error) {
	out := new(StakeHubNotJailedValidator)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotJailedValidator", raw); err != nil {
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

// StakeHubStillInJailPeriod represents a StillInJailPeriod error raised by the StakeHub contract.
type StakeHubStillInJailPeriod struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StillInJailPeriod()
func StakeHubStillInJailPeriodErrorID() common.Hash {
	return common.HexToHash("0x197d095ce12fcbcecb760823b2e9199a8b252bc1aae24c14e24b5670f2738dff")
}

// UnpackStillInJailPeriodError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StillInJailPeriod()
func (stakeHub *StakeHub) UnpackStillInJailPeriodError(raw []byte) (*StakeHubStillInJailPeriod, error) {
	out := new(StakeHubStillInJailPeriod)
	if err := stakeHub.abi.UnpackIntoInterface(out, "StillInJailPeriod", raw); err != nil {
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
