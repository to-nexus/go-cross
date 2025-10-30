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
	ABI:        "[{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pos\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCheckpoints.Checkpoint208\",\"components\":[{\"name\":\"_key\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"_value\",\"type\":\"uint208\",\"internalType\":\"uint208\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateBySig\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegates\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastTotalSupply\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sync\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateChanged\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fromDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateVotesChanged\",\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"previousVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ApproveNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20ExceededSafeSupply\",\"inputs\":[{\"name\":\"increasedSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC5805FutureLookup\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clock\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"ERC6372InconsistentClock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TransferNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VotesExpiredSignature\",\"inputs\":[{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "GovernanceToken",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c806306fdde0314610184578063095ea7b31461017f57806318160ddd1461017a57806323b872dd14610175578063313ce567146101705780633a46b1a81461016b5780634bf5d7e914610166578063587cde1e146101615780635c19a95c1461015c5780636fcfff451461015757806370a08231146101525780637ecebe001461014d5780638129fc1c1461014857806384b0196e146101435780638e539e8c1461013e57806391ddadf41461013957806395d89b41146101345780639ab24eb01461012f578063a58411941461012a578063a9059cbb14610125578063c3cda52014610120578063dd62ed3e1461011b5763f1127ed814610116575f80fd5b610e06565b610dce565b610c8e565b610c6b565b610b86565b610b4f565b610a92565b610a70565b610981565b6108b1565b6106ce565b610677565b610618565b6105c3565b61059f565b610540565b6104f0565b610417565b6103fc565b610335565b61030c565b6102e3565b6101c1565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9060206101be928181520190610189565b90565b346102b3575f3660031901126102b3576040515f5f51602061215c5f395f51905f52546101ed81610eaa565b808452906001811690811561028f5750600114610225575b6102218361021581850382610f17565b604051918291826101ad565b0390f35b5f51602061215c5f395f51905f525f9081527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0939250905b80821061027557509091508101602001610215610205565b91926001816020925483858801015201910190929161025d565b60ff191660208086019190915291151560051b840190910191506102159050610205565b5f80fd5b600435906001600160a01b03821682036102b357565b602435906001600160a01b03821682036102b357565b346102b35760403660031901126102b3576102fc6102b7565b50632028747160e01b5f5260045ffd5b346102b3575f3660031901126102b35760205f51602061219c5f395f51905f5254604051908152f35b346102b35760603660031901126102b35761034e6102b7565b6103566102cd565b6044359161036433826113eb565b5f198110610373575b50611414565b8381106103e1576001600160a01b038216156103ce5733156103bb578390036103b43361039f84611101565b9060018060a01b03165f5260205260405f2090565b555f61036d565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8390637dc7a0d960e11b5f523360045260245260445260645ffd5b346102b3575f3660031901126102b357602060405160128152f35b346102b35760403660031901126102b3576104306102b7565b61044561043f60243592611139565b91611468565b8154905f829160058411610498575b61045f935084611913565b908161047d57505060205f5b6040516001600160d01b039091168152f35b6104886020926113d0565b905f52815f20015460301c61046b565b91926104a3816117b5565b81039081116104eb5761045f93855f5265ffffffffffff8260205f2001541665ffffffffffff8516105f146104d9575091610454565b9291506104e5906114bf565b90610454565b6113bc565b346102b3575f3660031901126102b357610221604051610511604082610f17565b600e81526d06d6f64653d74696d657374616d760941b6020820152604051918291602083526020830190610189565b346102b35760203660031901126102b3576001600160a01b036105616102b7565b165f527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d00602052602060018060a01b0360405f205416604051908152f35b346102b35760203660031901126102b3576105c16105bb6102b7565b336114cd565b005b346102b35760203660031901126102b3576105e46105df6102b7565b611139565b5463ffffffff81116106015760209063ffffffff60405191168152f35b6306dfcc6560e41b5f52602060045260245260445ffd5b346102b35760203660031901126102b357602061066f6106366102b7565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b604051908152f35b346102b35760203660031901126102b3576001600160a01b036106986102b7565b165f527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00602052602060405f2054604051908152f35b346102b3575f3660031901126102b3575f5160206121dc5f395f51905f525467ffffffffffffffff61071060ff604084901c16159267ffffffffffffffff1690565b1680159081610826575b600114908161081c575b159081610813575b506108045780610763600167ffffffffffffffff195f5160206121dc5f395f51905f525416175f5160206121dc5f395f51905f5255565b6107cf575b6107706111e2565b61077657005b6107a060ff60401b195f5160206121dc5f395f51905f5254165f5160206121dc5f395f51905f5255565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b6107ff600160401b60ff60401b195f5160206121dc5f395f51905f525416175f5160206121dc5f395f51905f5255565b610768565b63f92ee8a960e01b5f5260045ffd5b9050155f61072c565b303b159150610724565b82915061071a565b92939161085061085e92600f60f81b865260e0602087015260e0860190610189565b908482036040860152610189565b92606083015260018060a01b031660808201525f60a082015260c0818303910152602080835192838152019201905f5b81811061089b5750505090565b825184526020938401939092019160010161088e565b346102b3575f3660031901126102b3577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100541580610958575b1561091b576108f7610f39565b6108ff61102e565b9061022161090b611387565b604051938493309146918661082e565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10154156108ea565b346102b35760203660031901126102b35761099d600435611468565b5f5160206121bc5f395f51905f5254905f829160058411610a10575b6109d193505f5160206121bc5f395f51905f52611913565b806109e257506040515f8152602090f35b610a0b6109f06020926113d0565b5f5160206121bc5f395f51905f525f52825f20015460301c90565b61046b565b9192610a1b816117b5565b81039081116104eb576109d1935f5160206121bc5f395f51905f525f5265ffffffffffff8260205f2001541665ffffffffffff8516105f14610a5e5750916109b9565b929150610a6a906114bf565b906109b9565b346102b3575f3660031901126102b357602060405165ffffffffffff42168152f35b346102b3575f3660031901126102b3576040515f5f51602061217c5f395f51905f5254610abe81610eaa565b808452906001811690811561028f5750600114610ae5576102218361021581850382610f17565b5f51602061217c5f395f51905f525f9081527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa939250905b808210610b3557509091508101602001610215610205565b919260018160209254838588010152019101909291610b1d565b346102b35760203660031901126102b35760206001600160d01b03610b7d610b786105df6102b7565b61158c565b16604051908152f35b346102b35760203660031901126102b357610b9f6102b7565b6110023303610c5657604051630c2eb40360e01b81526001600160a01b0382166004820152906020826024816110025afa918215610c51575f92610c20575b50610be881611171565b5482811015610c0457610bfe906105c1936113de565b90611666565b828111610c0d57005b6105c192610c1a916113de565b906115b6565b610c4391925060203d602011610c4a575b610c3b8183610f17565b8101906113a2565b905f610bde565b503d610c31565b6113b1565b630f22c43960e41b5f5261100260045260245ffd5b346102b35760403660031901126102b357610c846102b7565b6024359033611414565b346102b35760c03660031901126102b357610ca76102b7565b6044359060243560643560ff811681036102b35760843560a43591854211610dbb5791610daf91610db69360426105c19860405160208101917fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf835260018060a01b038b166040830152896060830152608082015260808152610d2b60a082610f17565b519020610d36611f7b565b610d3e611fe5565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a08152610d8f60c082610f17565b519020906040519161190160f01b83526002830152602282015220611720565b9182611738565b6114cd565b85632341d78760e11b5f5260045260245ffd5b346102b35760403660031901126102b3576020610dfd610dec6102b7565b61039f610df76102cd565b91611101565b54604051908152f35b346102b35760403660031901126102b357610e1f6102b7565b6024359063ffffffff821682036102b35761022191610e52610e6092610e436113fc565b50610e4c6113fc565b50611139565b610e5a6113fc565b50611d81565b5060405190610e6e82610ef6565b5465ffffffffffff811680835260309190911c60209283019081526040805192835290516001600160d01b031692820192909252918291820190565b90600182811c92168015610ed8575b6020831014610ec457565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610eb9565b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff821117610f1257604052565b610ee2565b90601f8019910116810190811067ffffffffffffffff821117610f1257604052565b604051905f827fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025491610f6b83610eaa565b808352926001811690811561100f5750600114610f91575b610f8f92500383610f17565b565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025f90815290917f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5b818310610ff3575050906020610f8f92820101610f83565b6020919350806001915483858901015201910190918492610fdb565b60209250610f8f94915060ff191682840152151560051b820101610f83565b604051905f827fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103549161106083610eaa565b808352926001811690811561100f575060011461108357610f8f92500383610f17565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b8183106110e5575050906020610f8f92820101610f83565b60209193508060019154838589010152019101909184926110cd565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b6001600160a01b03165f9081527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d016020526040902090565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006020526040902090565b60405190610f8f604083610f17565b604051906111c7604083610f17565b60168252565b604051906111dc604083610f17565b60068252565b413303611378573a611369576111f66111b8565b7521b937b9b99023b7bb32b93730b731b2902a37b5b2b760511b602082015261121d6111cd565b90656743524f535360d01b6020830152611235611a7b565b61123d611a7b565b80519067ffffffffffffffff8211610f125761126f8261126a5f51602061215c5f395f51905f5254610eaa565b611aa6565b602090601f83116001146112ca5791806112a3926112b795945f926112bf575b50508160011b915f199060031b1c19161790565b5f51602061215c5f395f51905f5255611b51565b610f8f611a7b565b015190505f8061128f565b5f51602061215c5f395f51905f525f52601f19831691907f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0925f5b81811061135157509160019391856112b797969410611339575b505050811b015f51602061215c5f395f51905f5255611b51565b01515f1960f88460031b161c191690555f808061131f565b92936020600181928786015181550195019301611305565b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b60405190611396602083610f17565b5f808352366020840137565b908160209103126102b3575190565b6040513d5f823e3d90fd5b634e487b7160e01b5f52601160045260245ffd5b5f198101919082116104eb57565b919082039182116104eb57565b9061039f6113f892611101565b5490565b6040519061140982610ef6565b5f6020838281520152565b9091506001600160a01b031615611455576001600160a01b03161561144257638cd22d1960e01b5f5260045ffd5b63ec442f0560e01b5f525f60045260245ffd5b634b637e8f60e11b5f525f60045260245ffd5b65ffffffffffff4216808210156114a9575065ffffffffffff81116114925765ffffffffffff1690565b6306dfcc6560e41b5f52603060045260245260445ffd5b90637669fc0f60e11b5f5260045260245260445ffd5b90600182018092116104eb57565b6001600160a01b038181165f8181527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d006020526040812080548685166001600160a01b031982168117909255610f8f969416946115869390928691907f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9080a46001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b91611977565b8054908161159a5750505f90565b815f198101116104eb575f525f199060205f2001015460301c90565b6001600160a01b0381169291908315611455576115d281611171565b5482811061164b57905f610f8f9495848294036115ee84611171565b55845f51602061219c5f395f51905f5254035f51602061219c5f395f51905f52557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6040518061164388829190602083019252565b0390a3611e91565b90508363391434e360e21b5f5260045260245260445260645ffd5b91906001600160a01b0383168015611442575f51602061219c5f395f51905f52548281018091116104eb575f51602061219c5f395f51905f52556116a984611171565b8054830190556040518281525f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602090a35f51602061219c5f395f51905f5254926001600160d01b038085116117095750610f8f9293505f611e91565b630e58ae9360e11b5f52600485905260245260445ffd5b916101be939161172f93611c65565b90929192611d05565b6001600160a01b03165f8181527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915591829003611782575050565b6301d4b62360e61b5f5260045260245260445ffd5b81156117a1570490565b634e487b7160e01b5f52601260045260245ffd5b60018111156101be57806001600160801b8210156118d6575b61187c61187261186861185e61185461184a6118396118839760048a600160401b6118889c10156118c9575b6401000000008110156118bc575b620100008110156118af575b6101008110156118a2575b6010811015611895575b101561188d575b60030260011c90565b611843818b611797565b0160011c90565b611843818a611797565b6118438189611797565b6118438188611797565b6118438187611797565b6118438186611797565b8093611797565b821190565b900390565b60011b611830565b60041c9160021b91611829565b60081c9160041b9161181f565b60101c9160081b91611814565b60201c9160101b91611808565b60401c9160201b916117fa565b505061188861188361187c61187261186861185e61185461184a6118396118fd8a60801c90565b9850600160401b97506117ce9650505050505050565b91905b8382106119235750505090565b9091928083169080841860011c82018092116104eb57845f5265ffffffffffff8260205f2001541665ffffffffffff8416105f146119655750925b9190611916565b939250611971906114bf565b9161195e565b6001600160a01b03808316939291908116908185141580611a72575b61199f575b5050505050565b81611a14575b5050826119b4575b8080611998565b7fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724916119eb6119e56119f193611139565b91611daa565b90611e12565b604080516001600160d01b039384168152919092166020820152a25f80806119ad565b611a50611a417fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a72492611139565b611a4a86611daa565b90611ddb565b604080516001600160d01b039384168152919092166020820152a25f806119a5565b50831515611993565b60ff5f5160206121dc5f395f51905f525460401c1615611a9757565b631afcd79f60e31b5f5260045ffd5b601f8111611ab2575050565b5f51602061215c5f395f51905f525f5260205f20906020601f840160051c83019310611af8575b601f0160051c01905b818110611aed575050565b5f8155600101611ae2565b9091508190611ad9565b601f8211611b0f57505050565b5f5260205f20906020601f840160051c83019310611b47575b601f0160051c01905b818110611b3c575050565b5f8155600101611b31565b9091508190611b28565b90815167ffffffffffffffff8111610f1257611b9181611b7e5f51602061217c5f395f51905f5254610eaa565b5f51602061217c5f395f51905f52611b02565b602092601f8211600114611bd157611bc0929382915f926112bf5750508160011b915f199060031b1c19161790565b5f51602061217c5f395f51905f5255565b5f51602061217c5f395f51905f525f52601f198216937f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa915f5b868110611c4d5750836001959610611c35575b505050811b015f51602061217c5f395f51905f5255565b01515f1960f88460031b161c191690555f8080611c1e565b91926020600181928685015181550194019201611c0b565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411611cdc579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15610c51575f516001600160a01b03811615611cd257905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611cf157565b634e487b7160e01b5f52602160045260245ffd5b611d0e81611ce7565b80611d17575050565b611d2081611ce7565b60018103611d375763f645eedf60e01b5f5260045ffd5b611d4081611ce7565b60028103611d5b575063fce698f760e01b5f5260045260245ffd5b80611d67600392611ce7565b14611d6f5750565b6335e2f38360e21b5f5260045260245ffd5b8054821015611d96575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b6001600160d01b038111611dc4576001600160d01b031690565b6306dfcc6560e41b5f5260d060045260245260445ffd5b90611de58261158c565b6001600160d01b03918216908216039081116104eb57611e0e9165ffffffffffff421690612085565b9091565b90611e1c8261158c565b6001600160d01b03918216908216019081116104eb57611e0e9165ffffffffffff421690612085565b611e5b5f5160206121bc5f395f51905f5261158c565b6001600160d01b03918216908216039081116104eb57611e0e904265ffffffffffff165f5160206121bc5f395f51905f52612085565b9091906001600160a01b03168015611f17575b610f8f926001600160a01b0316908115611eff575b5f9081527fe8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d006020526040808220549282529020546001600160a01b039081169116611977565b611f10611f0b84611daa565b611e45565b5050611eb9565b611f2082611daa565b611f365f5160206121bc5f395f51905f5261158c565b6001600160d01b0391821690821601939084116104eb57610f8f93611f71904265ffffffffffff165f5160206121bc5f395f51905f52612085565b9050509250611ea4565b611f83610f39565b8051908115611f93576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100548015611fc05790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b611fed61102e565b8051908115611ffd576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101548015611fc05790565b8054600160401b811015610f125761204791600182018155611d81565b61207257815160209092015160301b65ffffffffffff191665ffffffffffff92909216919091179055565b634e487b7160e01b5f525f60045260245ffd5b805492939280156121315761209c6120a7916113d0565b825f5260205f200190565b8054603081901c9365ffffffffffff91821692918116808411612122578793036120ee57506120ea92509065ffffffffffff82549181199060301b169116179055565b9190565b9150506120ea9161210e6121006111a9565b65ffffffffffff9093168352565b6001600160d01b038616602083015261202a565b632520601d60e01b5f5260045ffd5b5090612156916121426121006111a9565b6001600160d01b038516602083015261202a565b5f919056fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0452c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02e8b26c30fad74198956032a3533d903385d56dd795af560196f9c78d4af40d02f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212208872a8057672e98a0141b5a21026674cab44479765363918ca6d163b7e5152ef64736f6c634300081c0033",
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
// Solidity: function delegate(address delegatee) returns()
func (governanceToken *GovernanceToken) PackDelegate(delegatee common.Address) []byte {
	enc, err := governanceToken.abi.Pack("delegate", delegatee)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDelegateBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (governanceToken *GovernanceToken) PackDelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := governanceToken.abi.Pack("delegateBySig", delegatee, nonce, expiry, v, r, s)
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
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return governanceToken.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return governanceToken.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceToken.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return governanceToken.UnpackECDSAInvalidSignatureSError(raw[4:])
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

// GovernanceTokenECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the GovernanceToken contract.
type GovernanceTokenECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func GovernanceTokenECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (governanceToken *GovernanceToken) UnpackECDSAInvalidSignatureError(raw []byte) (*GovernanceTokenECDSAInvalidSignature, error) {
	out := new(GovernanceTokenECDSAInvalidSignature)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the GovernanceToken contract.
type GovernanceTokenECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func GovernanceTokenECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (governanceToken *GovernanceToken) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*GovernanceTokenECDSAInvalidSignatureLength, error) {
	out := new(GovernanceTokenECDSAInvalidSignatureLength)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTokenECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the GovernanceToken contract.
type GovernanceTokenECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func GovernanceTokenECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (governanceToken *GovernanceToken) UnpackECDSAInvalidSignatureSError(raw []byte) (*GovernanceTokenECDSAInvalidSignatureS, error) {
	out := new(GovernanceTokenECDSAInvalidSignatureS)
	if err := governanceToken.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
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
