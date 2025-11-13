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

// CheckpointsCheckpoint208 is an auto generated low-level Go binding around an user-defined struct.
type CheckpointsCheckpoint208 struct {
	Key   *big.Int
	Value *big.Int
}

// GovernanceTokenMetaData contains all meta data concerning the GovernanceToken contract.
var GovernanceTokenMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pos\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCheckpoints.Checkpoint208\",\"components\":[{\"name\":\"_key\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"_value\",\"type\":\"uint208\",\"internalType\":\"uint208\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"delegateBySig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"delegateFrom\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegates\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastTotalSupply\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sync\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateChanged\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fromDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateVotesChanged\",\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"previousVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ApproveNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DelegateNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20ExceededSafeSupply\",\"inputs\":[{\"name\":\"increasedSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC5805FutureLookup\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clock\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"ERC6372InconsistentClock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TransferNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VotesExpiredSignature\",\"inputs\":[{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "GovernanceToken",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c806306fdde0314610194578063095ea7b31461018f57806318160ddd1461018a57806323b872dd146101855780632949582d14610180578063313ce5671461017b5780633a46b1a8146101765780634bf5d7e914610171578063587cde1e1461016c5780635c19a95c146101675780636fcfff451461016257806370a082311461015d5780637ecebe00146101585780638129fc1c1461015357806384b0196e1461014e5780638e539e8c1461014957806391ddadf41461014457806395d89b411461013f5780639ab24eb01461013a578063a584119414610135578063a9059cbb14610130578063c3cda5201461012b578063dd62ed3e146101265763f1127ed814610121575f80fd5b610dfc565b610dc4565b610d8f565b610d6c565b610c9c565b610c65565b610ba8565b610b86565b610a97565b6109c7565b6107e4565b61078d565b61072e565b6106d9565b6106b0565b610651565b610601565b610528565b61050d565b61040c565b610345565b61031c565b6102f3565b6101d1565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9060206101ce928181520190610199565b90565b346102c3575f3660031901126102c3576040515f5f516020611e515f395f51905f52546101fd81610ea0565b808452906001811690811561029f5750600114610235575b6102318361022581850382610f0d565b604051918291826101bd565b0390f35b5f516020611e515f395f51905f525f9081527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0939250905b80821061028557509091508101602001610225610215565b91926001816020925483858801015201910190929161026d565b60ff191660208086019190915291151560051b840190910191506102259050610215565b5f80fd5b600435906001600160a01b03821682036102c357565b602435906001600160a01b03821682036102c357565b346102c35760403660031901126102c35761030c6102c7565b50632028747160e01b5f5260045ffd5b346102c3575f3660031901126102c35760205f516020611e915f395f51905f5254604051908152f35b346102c35760603660031901126102c35761035e6102c7565b6103666102dd565b6044359161037433826113e1565b5f198110610383575b5061140a565b8381106103f1576001600160a01b038216156103de5733156103cb578390036103c4336103af846110f7565b9060018060a01b03165f5260205260405f2090565b555f61037d565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8390637dc7a0d960e11b5f523360045260245260445260645ffd5b346102c35760403660031901126102c3576104256102c7565b61042d6102dd565b9061100233036104f8576001600160a01b038181165f8181527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d006020526040812080548685166001600160a01b0319821681179092556104f6969416946104f09390928691907f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9080a46001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b91611657565b005b630f22c43960e41b5f5261100260045260245ffd5b346102c3575f3660031901126102c357602060405160128152f35b346102c35760403660031901126102c3576105416102c7565b6105566105506024359261112f565b9161145e565b8154905f8291600584116105a9575b6105709350846118d7565b908161058e57505060205f5b6040516001600160d01b039091168152f35b6105996020926113c6565b905f52815f20015460301c61057c565b91926105b481611779565b81039081116105fc5761057093855f5265ffffffffffff8260205f2001541665ffffffffffff8516105f146105ea575091610565565b9291506105f6906114b5565b90610565565b6113b2565b346102c3575f3660031901126102c357610231604051610622604082610f0d565b600e81526d06d6f64653d74696d657374616d760941b6020820152604051918291602083526020830190610199565b346102c35760203660031901126102c3576001600160a01b036106726102c7565b165f527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d00602052602060018060a01b0360405f205416604051908152f35b346102c35760203660031901126102c3576106c96102c7565b50633c3d00b360e21b5f5260045ffd5b346102c35760203660031901126102c3576106fa6106f56102c7565b61112f565b5463ffffffff81116107175760209063ffffffff60405191168152f35b6306dfcc6560e41b5f52602060045260245260445ffd5b346102c35760203660031901126102c357602061078561074c6102c7565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b604051908152f35b346102c35760203660031901126102c3576001600160a01b036107ae6102c7565b165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052602060405f2054604051908152f35b346102c3575f3660031901126102c3575f516020611ed15f395f51905f525467ffffffffffffffff61082660ff604084901c16159267ffffffffffffffff1690565b168015908161093c575b6001149081610932575b159081610929575b5061091a5780610879600167ffffffffffffffff195f516020611ed15f395f51905f525416175f516020611ed15f395f51905f5255565b6108e5575b6108866111d8565b61088c57005b6108b660ff60401b195f516020611ed15f395f51905f5254165f516020611ed15f395f51905f5255565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b610915600160401b60ff60401b195f516020611ed15f395f51905f525416175f516020611ed15f395f51905f5255565b61087e565b63f92ee8a960e01b5f5260045ffd5b9050155f610842565b303b15915061083a565b829150610830565b92939161096661097492600f60f81b865260e0602087015260e0860190610199565b908482036040860152610199565b92606083015260018060a01b031660808201525f60a082015260c0818303910152602080835192838152019201905f5b8181106109b15750505090565b82518452602093840193909201916001016109a4565b346102c3575f3660031901126102c3577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100541580610a6e575b15610a3157610a0d610f2f565b610a15611024565b90610231610a2161137d565b6040519384933091469186610944565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015415610a00565b346102c35760203660031901126102c357610ab360043561145e565b5f516020611eb15f395f51905f5254905f829160058411610b26575b610ae793505f516020611eb15f395f51905f526118d7565b80610af857506040515f8152602090f35b610b21610b066020926113c6565b5f516020611eb15f395f51905f525f52825f20015460301c90565b61057c565b9192610b3181611779565b81039081116105fc57610ae7935f516020611eb15f395f51905f525f5265ffffffffffff8260205f2001541665ffffffffffff8516105f14610b74575091610acf565b929150610b80906114b5565b90610acf565b346102c3575f3660031901126102c357602060405165ffffffffffff42168152f35b346102c3575f3660031901126102c3576040515f5f516020611e715f395f51905f5254610bd481610ea0565b808452906001811690811561029f5750600114610bfb576102318361022581850382610f0d565b5f516020611e715f395f51905f525f9081527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa939250905b808210610c4b57509091508101602001610225610215565b919260018160209254838588010152019101909291610c33565b346102c35760203660031901126102c35760206001600160d01b03610c93610c8e6106f56102c7565b6114c3565b16604051908152f35b346102c35760203660031901126102c357610cb56102c7565b61100233036104f857604051630c2eb40360e01b81526001600160a01b0382166004820152906020826024816110025afa918215610d67575f92610d36575b50610cfe81611167565b5482811015610d1a57610d14906104f6936113d4565b9061159d565b828111610d2357005b6104f692610d30916113d4565b906114ed565b610d5991925060203d602011610d60575b610d518183610f0d565b810190611398565b905f610cf4565b503d610d47565b6113a7565b346102c35760403660031901126102c357610d856102c7565b602435903361140a565b346102c35760c03660031901126102c357610da86102c7565b5060643560ff8116036102c357633c3d00b360e21b5f5260045ffd5b346102c35760403660031901126102c3576020610df3610de26102c7565b6103af610ded6102dd565b916110f7565b54604051908152f35b346102c35760403660031901126102c357610e156102c7565b6024359063ffffffff821682036102c35761023191610e48610e5692610e396113f2565b50610e426113f2565b5061112f565b610e506113f2565b50611b25565b5060405190610e6482610eec565b5465ffffffffffff811680835260309190911c60209283019081526040805192835290516001600160d01b031692820192909252918291820190565b90600182811c92168015610ece575b6020831014610eba57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610eaf565b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff821117610f0857604052565b610ed8565b90601f8019910116810190811067ffffffffffffffff821117610f0857604052565b604051905f827fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025491610f6183610ea0565b80835292600181169081156110055750600114610f87575b610f8592500383610f0d565b565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025f90815290917f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5b818310610fe9575050906020610f8592820101610f79565b6020919350806001915483858901015201910190918492610fd1565b60209250610f8594915060ff191682840152151560051b820101610f79565b604051905f827fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103549161105683610ea0565b8083529260018116908115611005575060011461107957610f8592500383610f0d565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b8183106110db575050906020610f8592820101610f79565b60209193508060019154838589010152019101909184926110c3565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b6001600160a01b03165f9081527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d016020526040902090565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006020526040902090565b60405190610f85604083610f0d565b604051906111bd604083610f0d565b60168252565b604051906111d2604083610f0d565b60068252565b41330361136e573a61135f576111ec6111ae565b7521b937b9b99023b7bb32b93730b731b2902a37b5b2b760511b60208201526112136111c3565b90656743524f535360d01b602083015261122b61193b565b61123361193b565b80519067ffffffffffffffff8211610f0857611265826112605f516020611e515f395f51905f5254610ea0565b611966565b602090601f83116001146112c0579180611299926112ad95945f926112b5575b50508160011b915f199060031b1c19161790565b5f516020611e515f395f51905f5255611a11565b610f8561193b565b015190505f80611285565b5f516020611e515f395f51905f525f52601f19831691907f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0925f5b81811061134757509160019391856112ad9796941061132f575b505050811b015f516020611e515f395f51905f5255611a11565b01515f1960f88460031b161c191690555f8080611315565b929360206001819287860151815501950193016112fb565b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b6040519061138c602083610f0d565b5f808352366020840137565b908160209103126102c3575190565b6040513d5f823e3d90fd5b634e487b7160e01b5f52601160045260245ffd5b5f198101919082116105fc57565b919082039182116105fc57565b906103af6113ee926110f7565b5490565b604051906113ff82610eec565b5f6020838281520152565b9091506001600160a01b03161561144b576001600160a01b03161561143857638cd22d1960e01b5f5260045ffd5b63ec442f0560e01b5f525f60045260245ffd5b634b637e8f60e11b5f525f60045260245ffd5b65ffffffffffff42168082101561149f575065ffffffffffff81116114885765ffffffffffff1690565b6306dfcc6560e41b5f52603060045260245260445ffd5b90637669fc0f60e11b5f5260045260245260445ffd5b90600182018092116105fc57565b805490816114d15750505f90565b815f198101116105fc575f525f199060205f2001015460301c90565b6001600160a01b038116929190831561144b5761150981611167565b5482811061158257905f610f8594958482940361152584611167565b55845f516020611e915f395f51905f5254035f516020611e915f395f51905f52557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6040518061157a88829190602083019252565b0390a3611c35565b90508363391434e360e21b5f5260045260245260445260645ffd5b91906001600160a01b0383168015611438575f516020611e915f395f51905f52548281018091116105fc575f516020611e915f395f51905f52556115e084611167565b8054830190556040518281525f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602090a35f516020611e915f395f51905f5254926001600160d01b038085116116405750610f859293505f611c35565b630e58ae9360e11b5f52600485905260245260445ffd5b6001600160a01b03808316939291908116908185141580611752575b61167f575b5050505050565b816116f4575b505082611694575b8080611678565b7fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724916116cb6116c56116d19361112f565b91611b4e565b90611bb6565b604080516001600160d01b039384168152919092166020820152a25f808061168d565b6117306117217fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a7249261112f565b61172a86611b4e565b90611b7f565b604080516001600160d01b039384168152919092166020820152a25f80611685565b50831515611673565b8115611765570490565b634e487b7160e01b5f52601260045260245ffd5b60018111156101ce57806001600160801b82101561189a575b61184061183661182c61182261181861180e6117fd6118479760048a600160401b61184c9c101561188d575b640100000000811015611880575b62010000811015611873575b610100811015611866575b6010811015611859575b1015611851575b60030260011c90565b611807818b61175b565b0160011c90565b611807818a61175b565b611807818961175b565b611807818861175b565b611807818761175b565b611807818661175b565b809361175b565b821190565b900390565b60011b6117f4565b60041c9160021b916117ed565b60081c9160041b916117e3565b60101c9160081b916117d8565b60201c9160101b916117cc565b60401c9160201b916117be565b505061184c61184761184061183661182c61182261181861180e6117fd6118c18a60801c90565b9850600160401b97506117929650505050505050565b91905b8382106118e75750505090565b9091928083169080841860011c82018092116105fc57845f5265ffffffffffff8260205f2001541665ffffffffffff8416105f146119295750925b91906118da565b939250611935906114b5565b91611922565b60ff5f516020611ed15f395f51905f525460401c161561195757565b631afcd79f60e31b5f5260045ffd5b601f8111611972575050565b5f516020611e515f395f51905f525f5260205f20906020601f840160051c830193106119b8575b601f0160051c01905b8181106119ad575050565b5f81556001016119a2565b9091508190611999565b601f82116119cf57505050565b5f5260205f20906020601f840160051c83019310611a07575b601f0160051c01905b8181106119fc575050565b5f81556001016119f1565b90915081906119e8565b90815167ffffffffffffffff8111610f0857611a5181611a3e5f516020611e715f395f51905f5254610ea0565b5f516020611e715f395f51905f526119c2565b602092601f8211600114611a9157611a80929382915f926112b55750508160011b915f199060031b1c19161790565b5f516020611e715f395f51905f5255565b5f516020611e715f395f51905f525f52601f198216937f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa915f5b868110611b0d5750836001959610611af5575b505050811b015f516020611e715f395f51905f5255565b01515f1960f88460031b161c191690555f8080611ade565b91926020600181928685015181550194019201611acb565b8054821015611b3a575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b6001600160d01b038111611b68576001600160d01b031690565b6306dfcc6560e41b5f5260d060045260245260445ffd5b90611b89826114c3565b6001600160d01b03918216908216039081116105fc57611bb29165ffffffffffff421690611d7a565b9091565b90611bc0826114c3565b6001600160d01b03918216908216019081116105fc57611bb29165ffffffffffff421690611d7a565b611bff5f516020611eb15f395f51905f526114c3565b6001600160d01b03918216908216039081116105fc57611bb2904265ffffffffffff165f516020611eb15f395f51905f52611d7a565b9091906001600160a01b03168015611cbb575b610f85926001600160a01b0316908115611ca3575b5f9081527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d006020526040808220549282529020546001600160a01b039081169116611657565b611cb4611caf84611b4e565b611be9565b5050611c5d565b611cc482611b4e565b611cda5f516020611eb15f395f51905f526114c3565b6001600160d01b0391821690821601939084116105fc57610f8593611d15904265ffffffffffff165f516020611eb15f395f51905f52611d7a565b9050509250611c48565b8054600160401b811015610f0857611d3c91600182018155611b25565b611d6757815160209092015160301b65ffffffffffff191665ffffffffffff92909216919091179055565b634e487b7160e01b5f525f60045260245ffd5b80549293928015611e2657611d91611d9c916113c6565b825f5260205f200190565b8054603081901c9365ffffffffffff91821692918116808411611e1757879303611de35750611ddf92509065ffffffffffff82549181199060301b169116179055565b9190565b915050611ddf91611e03611df561119f565b65ffffffffffff9093168352565b6001600160d01b0386166020830152611d1f565b632520601d60e01b5f5260045ffd5b5090611e4b91611e37611df561119f565b6001600160d01b0385166020830152611d1f565b5f919056fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0452c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02e8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d02f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220444a6b6c2f14bc06c36aa31566f2a29daee3bc45528205ff3562d5fd65b323ae64736f6c634300081c0033",
}

// GovernanceToken is an auto generated Go binding around an Ethereum contract.
type GovernanceToken struct {
	abi abi.ABI
}

// NewGovernanceToken creates a new instance of GovernanceToken.
func NewGovernanceToken() *GovernanceToken {
	parsed, err := GovernanceTokenMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &GovernanceToken{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *GovernanceToken) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCLOCKMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (governanceToken *GovernanceToken) PackCLOCKMODE() []byte {
	enc, err := governanceToken.abi.Pack("CLOCK_MODE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCLOCKMODE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (governanceToken *GovernanceToken) UnpackCLOCKMODE(data []byte) (string, error) {
	out, err := governanceToken.abi.Unpack("CLOCK_MODE", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (governanceToken *GovernanceToken) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := governanceToken.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (governanceToken *GovernanceToken) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := governanceToken.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns(bool)
func (governanceToken *GovernanceToken) PackApprove(arg0 common.Address, arg1 *big.Int) []byte {
	enc, err := governanceToken.abi.Pack("approve", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns(bool)
func (governanceToken *GovernanceToken) UnpackApprove(data []byte) (bool, error) {
	out, err := governanceToken.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (governanceToken *GovernanceToken) PackBalanceOf(account common.Address) []byte {
	enc, err := governanceToken.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (governanceToken *GovernanceToken) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := governanceToken.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCheckpoints is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (governanceToken *GovernanceToken) PackCheckpoints(account common.Address, pos uint32) []byte {
	enc, err := governanceToken.abi.Pack("checkpoints", account, pos)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCheckpoints is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (governanceToken *GovernanceToken) UnpackCheckpoints(data []byte) (CheckpointsCheckpoint208, error) {
	out, err := governanceToken.abi.Unpack("checkpoints", data)
	if err != nil {
		return *new(CheckpointsCheckpoint208), err
	}
	out0 := *abi.ConvertType(out[0], new(CheckpointsCheckpoint208)).(*CheckpointsCheckpoint208)
	return out0, err
}

// PackClock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (governanceToken *GovernanceToken) PackClock() []byte {
	enc, err := governanceToken.abi.Pack("clock")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackClock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (governanceToken *GovernanceToken) UnpackClock(data []byte) (*big.Int, error) {
	out, err := governanceToken.abi.Unpack("clock", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (governanceToken *GovernanceToken) PackDecimals() []byte {
	enc, err := governanceToken.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (governanceToken *GovernanceToken) UnpackDecimals(data []byte) (uint8, error) {
	out, err := governanceToken.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackDelegate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c19a95c.
//
// Solidity: function delegate(address ) pure returns()
func (governanceToken *GovernanceToken) PackDelegate(arg0 common.Address) []byte {
	enc, err := governanceToken.abi.Pack("delegate", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDelegateBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3cda520.
//
// Solidity: function delegateBySig(address , uint256 , uint256 , uint8 , bytes32 , bytes32 ) pure returns()
func (governanceToken *GovernanceToken) PackDelegateBySig(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 uint8, arg4 [32]byte, arg5 [32]byte) []byte {
	enc, err := governanceToken.abi.Pack("delegateBySig", arg0, arg1, arg2, arg3, arg4, arg5)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDelegateFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2949582d.
//
// Solidity: function delegateFrom(address delegator, address delegatee) returns()
func (governanceToken *GovernanceToken) PackDelegateFrom(delegator common.Address, delegatee common.Address) []byte {
	enc, err := governanceToken.abi.Pack("delegateFrom", delegator, delegatee)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDelegates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (governanceToken *GovernanceToken) PackDelegates(account common.Address) []byte {
	enc, err := governanceToken.abi.Pack("delegates", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDelegates is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (governanceToken *GovernanceToken) UnpackDelegates(data []byte) (common.Address, error) {
	out, err := governanceToken.abi.Unpack("delegates", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (governanceToken *GovernanceToken) PackEip712Domain() []byte {
	enc, err := governanceToken.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (governanceToken *GovernanceToken) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := governanceToken.abi.Unpack("eip712Domain", data)
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

// PackGetPastTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (governanceToken *GovernanceToken) PackGetPastTotalSupply(timepoint *big.Int) []byte {
	enc, err := governanceToken.abi.Pack("getPastTotalSupply", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPastTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (governanceToken *GovernanceToken) UnpackGetPastTotalSupply(data []byte) (*big.Int, error) {
	out, err := governanceToken.abi.Unpack("getPastTotalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetPastVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (governanceToken *GovernanceToken) PackGetPastVotes(account common.Address, timepoint *big.Int) []byte {
	enc, err := governanceToken.abi.Pack("getPastVotes", account, timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPastVotes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (governanceToken *GovernanceToken) UnpackGetPastVotes(data []byte) (*big.Int, error) {
	out, err := governanceToken.abi.Unpack("getPastVotes", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (governanceToken *GovernanceToken) PackGetVotes(account common.Address) []byte {
	enc, err := governanceToken.abi.Pack("getVotes", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetVotes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (governanceToken *GovernanceToken) UnpackGetVotes(data []byte) (*big.Int, error) {
	out, err := governanceToken.abi.Unpack("getVotes", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (governanceToken *GovernanceToken) PackInitialize() []byte {
	enc, err := governanceToken.abi.Pack("initialize")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (governanceToken *GovernanceToken) PackName() []byte {
	enc, err := governanceToken.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (governanceToken *GovernanceToken) UnpackName(data []byte) (string, error) {
	out, err := governanceToken.abi.Unpack("name", data)
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
func (governanceToken *GovernanceToken) PackNonces(owner common.Address) []byte {
	enc, err := governanceToken.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (governanceToken *GovernanceToken) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := governanceToken.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackNumCheckpoints is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (governanceToken *GovernanceToken) PackNumCheckpoints(account common.Address) []byte {
	enc, err := governanceToken.abi.Pack("numCheckpoints", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNumCheckpoints is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (governanceToken *GovernanceToken) UnpackNumCheckpoints(data []byte) (uint32, error) {
	out, err := governanceToken.abi.Unpack("numCheckpoints", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (governanceToken *GovernanceToken) PackSymbol() []byte {
	enc, err := governanceToken.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (governanceToken *GovernanceToken) UnpackSymbol(data []byte) (string, error) {
	out, err := governanceToken.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackSync is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5841194.
//
// Solidity: function sync(address account) returns()
func (governanceToken *GovernanceToken) PackSync(account common.Address) []byte {
	enc, err := governanceToken.abi.Pack("sync", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (governanceToken *GovernanceToken) PackTotalSupply() []byte {
	enc, err := governanceToken.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (governanceToken *GovernanceToken) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := governanceToken.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (governanceToken *GovernanceToken) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := governanceToken.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (governanceToken *GovernanceToken) UnpackTransfer(data []byte) (bool, error) {
	out, err := governanceToken.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (governanceToken *GovernanceToken) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := governanceToken.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (governanceToken *GovernanceToken) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := governanceToken.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// GovernanceTokenApproval represents a Approval event raised by the GovernanceToken contract.
type GovernanceTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const GovernanceTokenApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (GovernanceTokenApproval) ContractEventName() string {
	return GovernanceTokenApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (governanceToken *GovernanceToken) UnpackApprovalEvent(log *types.Log) (*GovernanceTokenApproval, error) {
	event := "Approval"
	if log.Topics[0] != governanceToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTokenApproval)
	if len(log.Data) > 0 {
		if err := governanceToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceToken.abi.Events[event].Inputs {
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

// GovernanceTokenDelegateChanged represents a DelegateChanged event raised by the GovernanceToken contract.
type GovernanceTokenDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const GovernanceTokenDelegateChangedEventName = "DelegateChanged"

// ContractEventName returns the user-defined event name.
func (GovernanceTokenDelegateChanged) ContractEventName() string {
	return GovernanceTokenDelegateChangedEventName
}

// UnpackDelegateChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (governanceToken *GovernanceToken) UnpackDelegateChangedEvent(log *types.Log) (*GovernanceTokenDelegateChanged, error) {
	event := "DelegateChanged"
	if log.Topics[0] != governanceToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTokenDelegateChanged)
	if len(log.Data) > 0 {
		if err := governanceToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceToken.abi.Events[event].Inputs {
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

// GovernanceTokenDelegateVotesChanged represents a DelegateVotesChanged event raised by the GovernanceToken contract.
type GovernanceTokenDelegateVotesChanged struct {
	Delegate      common.Address
	PreviousVotes *big.Int
	NewVotes      *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const GovernanceTokenDelegateVotesChangedEventName = "DelegateVotesChanged"

// ContractEventName returns the user-defined event name.
func (GovernanceTokenDelegateVotesChanged) ContractEventName() string {
	return GovernanceTokenDelegateVotesChangedEventName
}

// UnpackDelegateVotesChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (governanceToken *GovernanceToken) UnpackDelegateVotesChangedEvent(log *types.Log) (*GovernanceTokenDelegateVotesChanged, error) {
	event := "DelegateVotesChanged"
	if log.Topics[0] != governanceToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTokenDelegateVotesChanged)
	if len(log.Data) > 0 {
		if err := governanceToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceToken.abi.Events[event].Inputs {
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

// GovernanceTokenEIP712DomainChanged represents a EIP712DomainChanged event raised by the GovernanceToken contract.
type GovernanceTokenEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const GovernanceTokenEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (GovernanceTokenEIP712DomainChanged) ContractEventName() string {
	return GovernanceTokenEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (governanceToken *GovernanceToken) UnpackEIP712DomainChangedEvent(log *types.Log) (*GovernanceTokenEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != governanceToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTokenEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := governanceToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceToken.abi.Events[event].Inputs {
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

// GovernanceTokenInitialized represents a Initialized event raised by the GovernanceToken contract.
type GovernanceTokenInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const GovernanceTokenInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (GovernanceTokenInitialized) ContractEventName() string {
	return GovernanceTokenInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (governanceToken *GovernanceToken) UnpackInitializedEvent(log *types.Log) (*GovernanceTokenInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != governanceToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTokenInitialized)
	if len(log.Data) > 0 {
		if err := governanceToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceToken.abi.Events[event].Inputs {
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

// GovernanceTokenTransfer represents a Transfer event raised by the GovernanceToken contract.
type GovernanceTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const GovernanceTokenTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (GovernanceTokenTransfer) ContractEventName() string {
	return GovernanceTokenTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (governanceToken *GovernanceToken) UnpackTransferEvent(log *types.Log) (*GovernanceTokenTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != governanceToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTokenTransfer)
	if len(log.Data) > 0 {
		if err := governanceToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceToken.abi.Events[event].Inputs {
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
func (governanceToken *GovernanceToken) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ApproveNotAllowed"].ID.Bytes()[:4]) {
		return governanceToken.UnpackApproveNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["CheckpointUnorderedInsertion"].ID.Bytes()[:4]) {
		return governanceToken.UnpackCheckpointUnorderedInsertionError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["DelegateNotAllowed"].ID.Bytes()[:4]) {
		return governanceToken.UnpackDelegateNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ERC20ExceededSafeSupply"].ID.Bytes()[:4]) {
		return governanceToken.UnpackERC20ExceededSafeSupplyError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return governanceToken.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return governanceToken.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return governanceToken.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return governanceToken.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return governanceToken.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return governanceToken.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ERC5805FutureLookup"].ID.Bytes()[:4]) {
		return governanceToken.UnpackERC5805FutureLookupError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ERC6372InconsistentClock"].ID.Bytes()[:4]) {
		return governanceToken.UnpackERC6372InconsistentClockError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return governanceToken.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return governanceToken.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return governanceToken.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return governanceToken.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return governanceToken.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return governanceToken.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return governanceToken.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return governanceToken.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["TransferNotAllowed"].ID.Bytes()[:4]) {
		return governanceToken.UnpackTransferNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["VotesExpiredSignature"].ID.Bytes()[:4]) {
		return governanceToken.UnpackVotesExpiredSignatureError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// GovernanceTokenApproveNotAllowed represents a ApproveNotAllowed error raised by the GovernanceToken contract.
type GovernanceTokenApproveNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ApproveNotAllowed()
func GovernanceTokenApproveNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x20287471febb94d2a4e7be5d3662f51a83f2bee180c5c2ea14db4432e1cd3e48")
}

// UnpackApproveNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ApproveNotAllowed()
func (governanceToken *GovernanceToken) UnpackApproveNotAllowedError(raw []byte) (*GovernanceTokenApproveNotAllowed, error) {
	out := new(GovernanceTokenApproveNotAllowed)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ApproveNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenCheckpointUnorderedInsertion represents a CheckpointUnorderedInsertion error raised by the GovernanceToken contract.
type GovernanceTokenCheckpointUnorderedInsertion struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CheckpointUnorderedInsertion()
func GovernanceTokenCheckpointUnorderedInsertionErrorID() common.Hash {
	return common.HexToHash("0x2520601d9d60b717c34a36ad270857824c5a1ebbfd08ff39aba6930089495cfa")
}

// UnpackCheckpointUnorderedInsertionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CheckpointUnorderedInsertion()
func (governanceToken *GovernanceToken) UnpackCheckpointUnorderedInsertionError(raw []byte) (*GovernanceTokenCheckpointUnorderedInsertion, error) {
	out := new(GovernanceTokenCheckpointUnorderedInsertion)
	if err := governanceToken.abi.UnpackIntoInterface(out, "CheckpointUnorderedInsertion", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenDelegateNotAllowed represents a DelegateNotAllowed error raised by the GovernanceToken contract.
type GovernanceTokenDelegateNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DelegateNotAllowed()
func GovernanceTokenDelegateNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xf0f402cc756bb93e22baa284e14eb921151c876d4ebc2a22846b7510f46b4686")
}

// UnpackDelegateNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DelegateNotAllowed()
func (governanceToken *GovernanceToken) UnpackDelegateNotAllowedError(raw []byte) (*GovernanceTokenDelegateNotAllowed, error) {
	out := new(GovernanceTokenDelegateNotAllowed)
	if err := governanceToken.abi.UnpackIntoInterface(out, "DelegateNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenERC20ExceededSafeSupply represents a ERC20ExceededSafeSupply error raised by the GovernanceToken contract.
type GovernanceTokenERC20ExceededSafeSupply struct {
	IncreasedSupply *big.Int
	Cap             *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20ExceededSafeSupply(uint256 increasedSupply, uint256 cap)
func GovernanceTokenERC20ExceededSafeSupplyErrorID() common.Hash {
	return common.HexToHash("0x1cb15d26dea6ae78228522d00b5965f950275bea9a67abeec04cb99806defd4c")
}

// UnpackERC20ExceededSafeSupplyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20ExceededSafeSupply(uint256 increasedSupply, uint256 cap)
func (governanceToken *GovernanceToken) UnpackERC20ExceededSafeSupplyError(raw []byte) (*GovernanceTokenERC20ExceededSafeSupply, error) {
	out := new(GovernanceTokenERC20ExceededSafeSupply)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ERC20ExceededSafeSupply", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the GovernanceToken contract.
type GovernanceTokenERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func GovernanceTokenERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (governanceToken *GovernanceToken) UnpackERC20InsufficientAllowanceError(raw []byte) (*GovernanceTokenERC20InsufficientAllowance, error) {
	out := new(GovernanceTokenERC20InsufficientAllowance)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the GovernanceToken contract.
type GovernanceTokenERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func GovernanceTokenERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (governanceToken *GovernanceToken) UnpackERC20InsufficientBalanceError(raw []byte) (*GovernanceTokenERC20InsufficientBalance, error) {
	out := new(GovernanceTokenERC20InsufficientBalance)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenERC20InvalidApprover represents a ERC20InvalidApprover error raised by the GovernanceToken contract.
type GovernanceTokenERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func GovernanceTokenERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (governanceToken *GovernanceToken) UnpackERC20InvalidApproverError(raw []byte) (*GovernanceTokenERC20InvalidApprover, error) {
	out := new(GovernanceTokenERC20InvalidApprover)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the GovernanceToken contract.
type GovernanceTokenERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func GovernanceTokenERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (governanceToken *GovernanceToken) UnpackERC20InvalidReceiverError(raw []byte) (*GovernanceTokenERC20InvalidReceiver, error) {
	out := new(GovernanceTokenERC20InvalidReceiver)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenERC20InvalidSender represents a ERC20InvalidSender error raised by the GovernanceToken contract.
type GovernanceTokenERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func GovernanceTokenERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (governanceToken *GovernanceToken) UnpackERC20InvalidSenderError(raw []byte) (*GovernanceTokenERC20InvalidSender, error) {
	out := new(GovernanceTokenERC20InvalidSender)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenERC20InvalidSpender represents a ERC20InvalidSpender error raised by the GovernanceToken contract.
type GovernanceTokenERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func GovernanceTokenERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (governanceToken *GovernanceToken) UnpackERC20InvalidSpenderError(raw []byte) (*GovernanceTokenERC20InvalidSpender, error) {
	out := new(GovernanceTokenERC20InvalidSpender)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenERC5805FutureLookup represents a ERC5805FutureLookup error raised by the GovernanceToken contract.
type GovernanceTokenERC5805FutureLookup struct {
	Timepoint *big.Int
	Clock     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC5805FutureLookup(uint256 timepoint, uint48 clock)
func GovernanceTokenERC5805FutureLookupErrorID() common.Hash {
	return common.HexToHash("0xecd3f81ef0a2e77b4e11bb288f0a162344834f778af7e47bdbee611607f14457")
}

// UnpackERC5805FutureLookupError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC5805FutureLookup(uint256 timepoint, uint48 clock)
func (governanceToken *GovernanceToken) UnpackERC5805FutureLookupError(raw []byte) (*GovernanceTokenERC5805FutureLookup, error) {
	out := new(GovernanceTokenERC5805FutureLookup)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ERC5805FutureLookup", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenERC6372InconsistentClock represents a ERC6372InconsistentClock error raised by the GovernanceToken contract.
type GovernanceTokenERC6372InconsistentClock struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC6372InconsistentClock()
func GovernanceTokenERC6372InconsistentClockErrorID() common.Hash {
	return common.HexToHash("0x6ff07140ae905cb0f8e72fb38f5dd8e756c387b0065be0e1bf85bd0621a43dc7")
}

// UnpackERC6372InconsistentClockError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC6372InconsistentClock()
func (governanceToken *GovernanceToken) UnpackERC6372InconsistentClockError(raw []byte) (*GovernanceTokenERC6372InconsistentClock, error) {
	out := new(GovernanceTokenERC6372InconsistentClock)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ERC6372InconsistentClock", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenInvalidAccountNonce represents a InvalidAccountNonce error raised by the GovernanceToken contract.
type GovernanceTokenInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func GovernanceTokenInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (governanceToken *GovernanceToken) UnpackInvalidAccountNonceError(raw []byte) (*GovernanceTokenInvalidAccountNonce, error) {
	out := new(GovernanceTokenInvalidAccountNonce)
	if err := governanceToken.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenInvalidInitialization represents a InvalidInitialization error raised by the GovernanceToken contract.
type GovernanceTokenInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func GovernanceTokenInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (governanceToken *GovernanceToken) UnpackInvalidInitializationError(raw []byte) (*GovernanceTokenInvalidInitialization, error) {
	out := new(GovernanceTokenInvalidInitialization)
	if err := governanceToken.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenInvalidValue represents a InvalidValue error raised by the GovernanceToken contract.
type GovernanceTokenInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func GovernanceTokenInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (governanceToken *GovernanceToken) UnpackInvalidValueError(raw []byte) (*GovernanceTokenInvalidValue, error) {
	out := new(GovernanceTokenInvalidValue)
	if err := governanceToken.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenNotInitializing represents a NotInitializing error raised by the GovernanceToken contract.
type GovernanceTokenNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func GovernanceTokenNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (governanceToken *GovernanceToken) UnpackNotInitializingError(raw []byte) (*GovernanceTokenNotInitializing, error) {
	out := new(GovernanceTokenNotInitializing)
	if err := governanceToken.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenOnlyCoinbase represents a OnlyCoinbase error raised by the GovernanceToken contract.
type GovernanceTokenOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func GovernanceTokenOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (governanceToken *GovernanceToken) UnpackOnlyCoinbaseError(raw []byte) (*GovernanceTokenOnlyCoinbase, error) {
	out := new(GovernanceTokenOnlyCoinbase)
	if err := governanceToken.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenOnlySystemContract represents a OnlySystemContract error raised by the GovernanceToken contract.
type GovernanceTokenOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func GovernanceTokenOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (governanceToken *GovernanceToken) UnpackOnlySystemContractError(raw []byte) (*GovernanceTokenOnlySystemContract, error) {
	out := new(GovernanceTokenOnlySystemContract)
	if err := governanceToken.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the GovernanceToken contract.
type GovernanceTokenOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func GovernanceTokenOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (governanceToken *GovernanceToken) UnpackOnlyZeroGasPriceError(raw []byte) (*GovernanceTokenOnlyZeroGasPrice, error) {
	out := new(GovernanceTokenOnlyZeroGasPrice)
	if err := governanceToken.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the GovernanceToken contract.
type GovernanceTokenSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func GovernanceTokenSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (governanceToken *GovernanceToken) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*GovernanceTokenSafeCastOverflowedUintDowncast, error) {
	out := new(GovernanceTokenSafeCastOverflowedUintDowncast)
	if err := governanceToken.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenTransferNotAllowed represents a TransferNotAllowed error raised by the GovernanceToken contract.
type GovernanceTokenTransferNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferNotAllowed()
func GovernanceTokenTransferNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x8cd22d191543fbb893cff76797482bc41fda8a261404bdcd20ff77b9dec06ff5")
}

// UnpackTransferNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferNotAllowed()
func (governanceToken *GovernanceToken) UnpackTransferNotAllowedError(raw []byte) (*GovernanceTokenTransferNotAllowed, error) {
	out := new(GovernanceTokenTransferNotAllowed)
	if err := governanceToken.abi.UnpackIntoInterface(out, "TransferNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenVotesExpiredSignature represents a VotesExpiredSignature error raised by the GovernanceToken contract.
type GovernanceTokenVotesExpiredSignature struct {
	Expiry *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VotesExpiredSignature(uint256 expiry)
func GovernanceTokenVotesExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x4683af0ecae671986a1b991272ba6e7bcb633f179b33ccfb3beb636962b1efde")
}

// UnpackVotesExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VotesExpiredSignature(uint256 expiry)
func (governanceToken *GovernanceToken) UnpackVotesExpiredSignatureError(raw []byte) (*GovernanceTokenVotesExpiredSignature, error) {
	out := new(GovernanceTokenVotesExpiredSignature)
	if err := governanceToken.abi.UnpackIntoInterface(out, "VotesExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}
