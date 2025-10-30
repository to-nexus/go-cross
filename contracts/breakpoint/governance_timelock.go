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

// GovernanceTimelockMetaData contains all meta data concerning the GovernanceTimelock contract.
var GovernanceTimelockMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"CANCELLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXECUTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROPOSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getMinDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperationState\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumTimelockControllerUpgradeable.OperationState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimestamp\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOperation\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hashOperationBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"minDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"executors\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperation\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationDone\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationPending\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationReady\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"schedule\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"scheduleBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallExecuted\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSalt\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallScheduled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Cancelled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinDelayChange\",\"inputs\":[{\"name\":\"oldDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimelockInsufficientDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TimelockInvalidOperationLength\",\"inputs\":[{\"name\":\"targets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payloads\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"values\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TimelockUnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TimelockUnexecutedPredecessor\",\"inputs\":[{\"name\":\"predecessorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TimelockUnexpectedOperationState\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedStates\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	ID:         "GovernanceTimelock",
	BinRuntime: "0x6080604052600436101561001a575b3615610018575f80fd5b005b5f5f3560e01c806301d5062a146110f157806301ffc9a71461108157806307bd02651461105a578063134008d314610fad57806313bc9f2014610f8f578063150b7a0214610f3a578063248a9ca314610f035780632ab0f52914610ee55780632f2ff15d14610e9b57806331d5075014610e7d57806336568abe14610e39578063584b153e14610e1157806364d6235314610d945780637958004c14610d515780638065657f14610d325780638129fc1c14610ae65780638f2a0bb0146109615780638f61f4f51461092757806391d14854146108d2578063a217fddf146108b8578063b08e51c014610891578063b1c5f42714610867578063ba29482f1461079b578063bc197c8114610705578063c4c4c7b31461049f578063c4d252f5146103c5578063d45c44351461038f578063d547741f14610339578063e38335e514610200578063f23a6e61146101a65763f27a0c921461017a575061000e565b346101a357806003193601126101a35760205f516020611f795f395f51905f5254604051908152f35b80fd5b50346101a35760a03660031901126101a3576101c061119c565b506101c96111b2565b506084356001600160401b0381116101fc576101e990369060040161128c565b5060405163f23a6e6160e01b8152602090f35b5080fd5b5061020a36611311565b5f516020611f995f395f51905f5289989792989695939496525f516020611f595f395f51905f52602052604089205f805260205260ff60405f2054161561032b575b828214801590610321575b6103065761026e61027591898689878a888e611686565b97886119a0565b875b81811061028b578861028889611a5b565b80f35b8080897fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b5889896102fd6102e4868f6102dd828f928f908f60019f6102d2916102d793611621565b611631565b97611621565b3595611645565b906102f1828287876119f8565b604051948594856114b0565b0390a301610277565b60648984868563ffb0321160e01b8452600452602452604452fd5b5083821415610257565b610334336118d8565b61024c565b50346101a35760403660031901126101a35761038b6004356103596111b2565b9061038661037f825f525f516020611f595f395f51905f52602052600160405f20015490565b339061194b565b611e31565b5080f35b50346101a35760203660031901126101a357604060209160043581525f516020611f195f395f51905f5283522054604051908152f35b50346101a35760203660031901126101a3575f516020611fb95f395f51905f5281525f516020611f595f395f51905f5260209081526040808320335f90815292529020546004359060ff161561047b5761041e8161151e565b15610464578082525f516020611f195f395f51905f526020528160408120557fbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb708280a280f35b635ead8eb560e01b82526004526006602452604490fd5b63e2517d3f60e01b8252336004525f516020611fb95f395f51905f52602452604482fd5b50346101a35760803660031901126101a357600435906024356001600160401b0381116101fc576104d49036906004016113f3565b916044356001600160401b038111610701576104f49036906004016113f3565b6064356001600160a01b03811693919291908481036101fc575f516020611fd95f395f51905f5254946001600160401b0360ff8760401c16159616801590816106f9575b60011490816106ef575b1590816106e6575b506106d757856105586115e0565b6106a7575b610565611ecd565b61056d611ecd565b61057630611a85565b50610697575b50805b85518110156105c9576001906105a76001600160a01b036105a0838a61160d565b5116611b9b565b506105c2828060a01b036105bb838a61160d565b5116611c54565b500161057f565b5083815b84518110156105fc576001906105f56001600160a01b036105ee838961160d565b5116611cfa565b50016105cd565b507f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5604084805f516020611f795f395f51905f52558151908582526020820152a16106445780f35b60ff60401b195f516020611fd95f395f51905f5254165f516020611fd95f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6106a090611a85565b505f61057c565b600160401b60ff60401b195f516020611fd95f395f51905f525416175f516020611fd95f395f51905f525561055d565b63f92ee8a960e01b8352600483fd5b9050155f61054a565b303b159150610542565b879150610538565b8280fd5b50346101a35760a03660031901126101a35761071f61119c565b506107286111b2565b506044356001600160401b0381116101fc57610748903690600401611396565b506064356001600160401b0381116101fc57610768903690600401611396565b506084356001600160401b0381116101fc5761078890369060040161128c565b5060405163bc197c8160e01b8152602090f35b503461080e57602036600319011261080e5760043561100833036108525780158015610846575b61081257303b1561080e57604051906364d6235360e01b825260048201525f8160248183305af18015610803576107f7575080f35b61001891505f90611257565b6040513d5f823e3d90fd5b5f80fd5b60849060405190632c648cf160e01b825260406004830152600560448301526464656c617960d81b60648301526024820152fd5b506212750081116107c2565b630f22c43960e41b5f5261100860045260245ffd5b3461080e57602061088961087a36611311565b96959095949194939293611686565b604051908152f35b3461080e575f36600319011261080e5760206040515f516020611fb95f395f51905f528152f35b3461080e575f36600319011261080e5760206040515f8152f35b3461080e57604036600319011261080e576108eb6111b2565b6004355f525f516020611f595f395f51905f5260205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b3461080e575f36600319011261080e5760206040517fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc18152f35b3461080e5760c036600319011261080e576004356001600160401b03811161080e576109919036906004016112e1565b906024356001600160401b03811161080e576109b19036906004016112e1565b6044929192356001600160401b03811161080e576109d39036906004016112e1565b9390916064356084359560a435926109ea33611852565b808914801590610adc575b610ac257610a0988848489858a8f8e611686565b98610a14858b6117c7565b895f5b828110610a5457508980610a2757005b60207f20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d038791604051908152a2005b806001927f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca8b8b610ab78f8c610aaa8f928e610aa38f8f90610a9d6102d28f8097948195611621565b99611621565b3597611645565b9060405196879687611478565b0390a3018a90610a17565b908863ffb0321160e01b5f5260045260245260445260645ffd5b50818914156109f5565b3461080e575f36600319011261080e575f516020611fd95f395f51905f52546001600160401b0360ff8260401c1615911680159081610d2a575b6001149081610d20575b159081610d17575b50610d085780610b406115e0565b610cd8575b413303610cc9573a610cba576040805191610b608284611257565b6001835260208301601f198301368237835115610ca6576110059052610b84611ecd565b610b8c611ecd565b610b9530611a85565b50610b9e611afb565b505f5b8351811015610bdd57600190610bc26001600160a01b036105a0838861160d565b50610bd6828060a01b036105bb838861160d565b5001610ba1565b50905f5b8351811015610c0957600190610c026001600160a01b036105ee838861160d565b5001610be1565b5090620151805f516020611f795f395f51905f52557f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d58280515f8152620151806020820152a1610c5557005b60207fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29160ff60401b195f516020611fd95f395f51905f5254165f516020611fd95f395f51905f52555160018152a1005b634e487b7160e01b5f52603260045260245ffd5b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b600160401b60ff60401b195f516020611fd95f395f51905f525416175f516020611fd95f395f51905f5255610b45565b63f92ee8a960e01b5f5260045ffd5b90501582610b32565b303b159150610b2a565b829150610b20565b3461080e576020610889610d4536611209565b9493909392919261158b565b3461080e57602036600319011261080e57610d6d600435611547565b6040516004821015610d80576020918152f35b634e487b7160e01b5f52602160045260245ffd5b3461080e57602036600319011261080e57600435303303610dfe577f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d560405f516020611f795f395f51905f52548151908152836020820152a15f516020611f795f395f51905f5255005b63e2850c5960e01b5f523360045260245ffd5b3461080e57602036600319011261080e576020610e2f60043561151e565b6040519015158152f35b3461080e57604036600319011261080e57610e526111b2565b336001600160a01b03821603610e6e5761001890600435611e31565b63334bd91960e11b5f5260045ffd5b3461080e57602036600319011261080e576020610e2f600435611507565b3461080e57604036600319011261080e57610018600435610eba6111b2565b90610ee061037f825f525f516020611f595f395f51905f52602052600160405f20015490565b611da0565b3461080e57602036600319011261080e576020610e2f6004356114ef565b3461080e57602036600319011261080e5760206108896004355f525f516020611f595f395f51905f52602052600160405f20015490565b3461080e57608036600319011261080e57610f5361119c565b50610f5c6111b2565b506064356001600160401b03811161080e57610f7c90369060040161128c565b50604051630a85bd0160e11b8152602090f35b3461080e57602036600319011261080e576020610e2f6004356114d7565b6100186110385f6110447fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b5861102f610fe436611209565b5f516020611f995f395f51905f528a9995979299949394525f516020611f595f395f51905f5260205260408a208a805260205260ff60408b2054161561104c575b888484898961158b565b988997886119a0565b6102f1828287876119f8565b0390a3611a5b565b611055336118d8565b611025565b3461080e575f36600319011261080e5760206040515f516020611f995f395f51905f528152f35b3461080e57602036600319011261080e5760043563ffffffff60e01b811680910361080e57602090630271189760e51b81149081156110c6575b506040519015158152f35b637965db0b60e01b8114915081156110e0575b50826110bb565b6301ffc9a760e01b149050826110d9565b3461080e5760c036600319011261080e5761110a61119c565b602435906044356001600160401b03811161080e577f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca926111505f9236906004016111dc565b949091606435946111926084359660a4359061116b33611852565b61117989828c8a898961158b565b998a97611186848a6117c7565b60405196879687611478565b0390a380610a2757005b600435906001600160a01b038216820361080e57565b602435906001600160a01b038216820361080e57565b35906001600160a01b038216820361080e57565b9181601f8401121561080e578235916001600160401b03831161080e576020838186019501011161080e57565b60a060031982011261080e576004356001600160a01b038116810361080e579160243591604435906001600160401b03821161080e5761124b916004016111dc565b90916064359060843590565b90601f801991011681019081106001600160401b0382111761127857604052565b634e487b7160e01b5f52604160045260245ffd5b81601f8201121561080e578035906001600160401b03821161127857604051926112c0601f8401601f191660200185611257565b8284526020838301011161080e57815f926020809301838601378301015290565b9181601f8401121561080e578235916001600160401b03831161080e576020808501948460051b01011161080e57565b60a060031982011261080e576004356001600160401b03811161080e578161133b916004016112e1565b929092916024356001600160401b03811161080e578161135d916004016112e1565b92909291604435906001600160401b03821161080e5761124b916004016112e1565b6001600160401b0381116112785760051b60200190565b9080601f8301121561080e5781356113ad8161137f565b926113bb6040519485611257565b81845260208085019260051b82010192831161080e57602001905b8282106113e35750505090565b81358152602091820191016113d6565b9080601f8301121561080e57813561140a8161137f565b926114186040519485611257565b81845260208085019260051b82010192831161080e57602001905b8282106114405750505090565b6020809161144d846111c8565b815201910190611433565b908060209392818452848401375f828201840152601f01601f1916010190565b9290936114a6926080959897969860018060a01b03168552602085015260a0604085015260a0840191611458565b9460608201520152565b6114d4949260609260018060a01b0316825260208201528160408201520191611458565b90565b6114e090611547565b6004811015610d805760021490565b6114f890611547565b6004811015610d805760031490565b61151090611547565b6004811015610d8057151590565b61152790611547565b6004811015610d80576001811490811561153f575090565b600291501490565b5f525f516020611f195f395f51905f5260205260405f205480155f1461156c57505f90565b6001810361157a5750600390565b42101561158657600190565b600290565b946115c16115da94959293604051968795602087019960018060a01b03168a52604087015260a0606087015260c0860191611458565b91608084015260a083015203601f198101835282611257565b51902090565b60016001600160401b03195f516020611fd95f395f51905f525416175f516020611fd95f395f51905f5255565b8051821015610ca65760209160051b010190565b9190811015610ca65760051b0190565b356001600160a01b038116810361080e5790565b9190811015610ca65760051b81013590601e198136030182121561080e5701908135916001600160401b03831161080e57602001823603811361080e579190565b9693949190969592956040519660208801988060c08a0160a08c525260e0890192905f5b8181106117a157505050878203601f190160408901528082526001600160fb1b03811161080e579087959394929160051b8092602083013701848103606086015260208101849052600584901b8101604090810194908201915f90889036829003601e1901905b84841061173b575050505050506115da9450608084015260a083015203601f198101835282611257565b91939597909294969850601f19601f198383030101875289358381121561080e57840190602082359201916001600160401b03811161080e57803603831361080e5761178d6020928392600195611458565b9b0197019401918a98969997959391611711565b909193602080600192838060a01b036117b9896111c8565b1681520195019291016116aa565b906117d182611507565b61183a575f516020611f795f395f51905f52548082106118245750420190814211611810575f525f516020611f195f395f51905f5260205260405f2055565b634e487b7160e01b5f52601160045260245ffd5b90635433660960e01b5f5260045260245260445ffd5b50635ead8eb560e01b5f52600452600160245260445ffd5b6001600160a01b0381165f9081527f5a8734c34b98d7c96eb2ea25f298989407e1f25da116ec139bcce0887bcb7cf7602052604090205460ff16156118945750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004527fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc1602452604490fd5b6001600160a01b0381165f9081527f52fce5e8a5d0d9e8d1ea29f4525e512e9c27bf92cae50374d497f918ab48f382602052604090205460ff161561191a5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020611f995f395f51905f52602452604490fd5b90815f525f516020611f595f395f51905f5260205260405f2060018060a01b0382165f5260205260ff60405f20541615611983575050565b63e2517d3f60e01b5f5260018060a01b031660045260245260445ffd5b6119a9816114d7565b156119e15750801515806119d1575b6119bf5750565b63121534c360e31b5f5260045260245ffd5b506119db816114ef565b156119b8565b635ead8eb560e01b5f52600452600460245260445ffd5b91906110083b1561080e57611a3b935f9360405195869485938493631cff79cd60e01b855260018060a01b03166004850152604060248501526044840191611458565b03916110085af1801561080357611a4f5750565b5f611a5991611257565b565b611a64816114d7565b156119e1575f525f516020611f195f395f51905f52602052600160405f2055565b6001600160a01b0381165f9081525f516020611f395f395f51905f52602052604090205460ff16611af6576001600160a01b03165f8181525f516020611f395f395f51905f5260205260408120805460ff191660011790553391905f516020611ef95f395f51905f528180a4600190565b505f90565b6110055f525f516020611f395f395f51905f526020527f4b1200bfdaca04549b6591ccc1e82bab28ea8ccf39522e75b36da363b3bf532f5460ff16611b97576110055f8181525f516020611f395f395f51905f526020527f4b1200bfdaca04549b6591ccc1e82bab28ea8ccf39522e75b36da363b3bf532f805460ff191660011790553391905f516020611ef95f395f51905f528180a4600190565b5f90565b6001600160a01b0381165f9081527f5a8734c34b98d7c96eb2ea25f298989407e1f25da116ec139bcce0887bcb7cf7602052604090205460ff16611af6576001600160a01b03165f8181527f5a8734c34b98d7c96eb2ea25f298989407e1f25da116ec139bcce0887bcb7cf760205260408120805460ff191660011790553391907fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc1905f516020611ef95f395f51905f529080a4600190565b6001600160a01b0381165f9081527ffa71e07f24c4701ef65a970775979de1292cfe909335cd18a32d2b7b73987914602052604090205460ff16611af6576001600160a01b03165f8181527ffa71e07f24c4701ef65a970775979de1292cfe909335cd18a32d2b7b7398791460205260408120805460ff191660011790553391905f516020611fb95f395f51905f52905f516020611ef95f395f51905f529080a4600190565b6001600160a01b0381165f9081527f52fce5e8a5d0d9e8d1ea29f4525e512e9c27bf92cae50374d497f918ab48f382602052604090205460ff16611af6576001600160a01b03165f8181527f52fce5e8a5d0d9e8d1ea29f4525e512e9c27bf92cae50374d497f918ab48f38260205260408120805460ff191660011790553391905f516020611f995f395f51905f52905f516020611ef95f395f51905f529080a4600190565b5f8181525f516020611f595f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16611e2b575f8181525f516020611f595f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291905f516020611ef95f395f51905f529080a4600190565b50505f90565b5f8181525f516020611f595f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff1615611e2b575f8181525f516020611f595f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b60ff5f516020611fd95f395f51905f525460401c1615611ee957565b631afcd79f60e31b5f5260045ffdfe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9a37c2aa9d186a0969ff8a8267bf4e07e864c2f2768f5040949e28a624fb3600b7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268009a37c2aa9d186a0969ff8a8267bf4e07e864c2f2768f5040949e28a624fb3601d8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63fd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f783f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122086d0c88c6dd61cf2cf06f7d47c2df4d819c9754605a54fc767f1abbfd29d70ef64736f6c634300081c0033",
}

// GovernanceTimelock is an auto generated Go binding around an Ethereum contract.
type GovernanceTimelock struct {
	abi abi.ABI
}

// NewGovernanceTimelock creates a new instance of GovernanceTimelock.
func NewGovernanceTimelock() *GovernanceTimelock {
	parsed, err := GovernanceTimelockMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &GovernanceTimelock{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *GovernanceTimelock) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCANCELLERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) PackCANCELLERROLE() []byte {
	enc, err := governanceTimelock.abi.Pack("CANCELLER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCANCELLERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) UnpackCANCELLERROLE(data []byte) ([32]byte, error) {
	out, err := governanceTimelock.abi.Unpack("CANCELLER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) PackDEFAULTADMINROLE() []byte {
	enc, err := governanceTimelock.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := governanceTimelock.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackEXECUTORROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) PackEXECUTORROLE() []byte {
	enc, err := governanceTimelock.abi.Pack("EXECUTOR_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEXECUTORROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) UnpackEXECUTORROLE(data []byte) ([32]byte, error) {
	out, err := governanceTimelock.abi.Unpack("EXECUTOR_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackPROPOSERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) PackPROPOSERROLE() []byte {
	enc, err := governanceTimelock.abi.Pack("PROPOSER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPROPOSERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) UnpackPROPOSERROLE(data []byte) ([32]byte, error) {
	out, err := governanceTimelock.abi.Unpack("PROPOSER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (governanceTimelock *GovernanceTimelock) PackCancel(id [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("cancel", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (governanceTimelock *GovernanceTimelock) PackExecute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("execute", target, value, payload, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExecuteBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (governanceTimelock *GovernanceTimelock) PackExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("executeBatch", targets, values, payloads, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetMinDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (governanceTimelock *GovernanceTimelock) PackGetMinDelay() []byte {
	enc, err := governanceTimelock.abi.Pack("getMinDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMinDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (governanceTimelock *GovernanceTimelock) UnpackGetMinDelay(data []byte) (*big.Int, error) {
	out, err := governanceTimelock.abi.Unpack("getMinDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetOperationState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (governanceTimelock *GovernanceTimelock) PackGetOperationState(id [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("getOperationState", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetOperationState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (governanceTimelock *GovernanceTimelock) UnpackGetOperationState(data []byte) (uint8, error) {
	out, err := governanceTimelock.abi.Unpack("getOperationState", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (governanceTimelock *GovernanceTimelock) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := governanceTimelock.abi.Unpack("getRoleAdmin", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGetTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (governanceTimelock *GovernanceTimelock) PackGetTimestamp(id [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("getTimestamp", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTimestamp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (governanceTimelock *GovernanceTimelock) UnpackGetTimestamp(data []byte) (*big.Int, error) {
	out, err := governanceTimelock.abi.Unpack("getTimestamp", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (governanceTimelock *GovernanceTimelock) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := governanceTimelock.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (governanceTimelock *GovernanceTimelock) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := governanceTimelock.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (governanceTimelock *GovernanceTimelock) UnpackHasRole(data []byte) (bool, error) {
	out, err := governanceTimelock.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackHashOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (governanceTimelock *GovernanceTimelock) PackHashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("hashOperation", target, value, data, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHashOperation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (governanceTimelock *GovernanceTimelock) UnpackHashOperation(data []byte) ([32]byte, error) {
	out, err := governanceTimelock.abi.Unpack("hashOperation", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackHashOperationBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (governanceTimelock *GovernanceTimelock) PackHashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("hashOperationBatch", targets, values, payloads, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHashOperationBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (governanceTimelock *GovernanceTimelock) UnpackHashOperationBatch(data []byte) ([32]byte, error) {
	out, err := governanceTimelock.abi.Unpack("hashOperationBatch", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (governanceTimelock *GovernanceTimelock) PackInitialize() []byte {
	enc, err := governanceTimelock.abi.Pack("initialize")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackInitialize0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4c4c7b3.
//
// Solidity: function initialize(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (governanceTimelock *GovernanceTimelock) PackInitialize0(minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) []byte {
	enc, err := governanceTimelock.abi.Pack("initialize0", minDelay, proposers, executors, admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (governanceTimelock *GovernanceTimelock) PackIsOperation(id [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("isOperation", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsOperation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (governanceTimelock *GovernanceTimelock) UnpackIsOperation(data []byte) (bool, error) {
	out, err := governanceTimelock.abi.Unpack("isOperation", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsOperationDone is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (governanceTimelock *GovernanceTimelock) PackIsOperationDone(id [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("isOperationDone", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsOperationDone is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (governanceTimelock *GovernanceTimelock) UnpackIsOperationDone(data []byte) (bool, error) {
	out, err := governanceTimelock.abi.Unpack("isOperationDone", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsOperationPending is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (governanceTimelock *GovernanceTimelock) PackIsOperationPending(id [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("isOperationPending", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsOperationPending is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (governanceTimelock *GovernanceTimelock) UnpackIsOperationPending(data []byte) (bool, error) {
	out, err := governanceTimelock.abi.Unpack("isOperationPending", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsOperationReady is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (governanceTimelock *GovernanceTimelock) PackIsOperationReady(id [32]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("isOperationReady", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsOperationReady is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (governanceTimelock *GovernanceTimelock) UnpackIsOperationReady(data []byte) (bool, error) {
	out, err := governanceTimelock.abi.Unpack("isOperationReady", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (governanceTimelock *GovernanceTimelock) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := governanceTimelock.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (governanceTimelock *GovernanceTimelock) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := governanceTimelock.abi.Unpack("onERC1155BatchReceived", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (governanceTimelock *GovernanceTimelock) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := governanceTimelock.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (governanceTimelock *GovernanceTimelock) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := governanceTimelock.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (governanceTimelock *GovernanceTimelock) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := governanceTimelock.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (governanceTimelock *GovernanceTimelock) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := governanceTimelock.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (governanceTimelock *GovernanceTimelock) PackRenounceRole(role [32]byte, callerConfirmation common.Address) []byte {
	enc, err := governanceTimelock.abi.Pack("renounceRole", role, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (governanceTimelock *GovernanceTimelock) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := governanceTimelock.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (governanceTimelock *GovernanceTimelock) PackSchedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) []byte {
	enc, err := governanceTimelock.abi.Pack("schedule", target, value, data, predecessor, salt, delay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackScheduleBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (governanceTimelock *GovernanceTimelock) PackScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) []byte {
	enc, err := governanceTimelock.abi.Pack("scheduleBatch", targets, values, payloads, predecessor, salt, delay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba29482f.
//
// Solidity: function setMinDelay(uint256 delay) returns()
func (governanceTimelock *GovernanceTimelock) PackSetMinDelay(delay *big.Int) []byte {
	enc, err := governanceTimelock.abi.Pack("setMinDelay", delay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (governanceTimelock *GovernanceTimelock) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := governanceTimelock.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (governanceTimelock *GovernanceTimelock) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := governanceTimelock.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUpdateDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (governanceTimelock *GovernanceTimelock) PackUpdateDelay(newDelay *big.Int) []byte {
	enc, err := governanceTimelock.abi.Pack("updateDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// GovernanceTimelockCallExecuted represents a CallExecuted event raised by the GovernanceTimelock contract.
type GovernanceTimelockCallExecuted struct {
	Id     [32]byte
	Index  *big.Int
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const GovernanceTimelockCallExecutedEventName = "CallExecuted"

// ContractEventName returns the user-defined event name.
func (GovernanceTimelockCallExecuted) ContractEventName() string {
	return GovernanceTimelockCallExecutedEventName
}

// UnpackCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (governanceTimelock *GovernanceTimelock) UnpackCallExecutedEvent(log *types.Log) (*GovernanceTimelockCallExecuted, error) {
	event := "CallExecuted"
	if log.Topics[0] != governanceTimelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTimelockCallExecuted)
	if len(log.Data) > 0 {
		if err := governanceTimelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceTimelock.abi.Events[event].Inputs {
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

// GovernanceTimelockCallSalt represents a CallSalt event raised by the GovernanceTimelock contract.
type GovernanceTimelockCallSalt struct {
	Id   [32]byte
	Salt [32]byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const GovernanceTimelockCallSaltEventName = "CallSalt"

// ContractEventName returns the user-defined event name.
func (GovernanceTimelockCallSalt) ContractEventName() string {
	return GovernanceTimelockCallSaltEventName
}

// UnpackCallSaltEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (governanceTimelock *GovernanceTimelock) UnpackCallSaltEvent(log *types.Log) (*GovernanceTimelockCallSalt, error) {
	event := "CallSalt"
	if log.Topics[0] != governanceTimelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTimelockCallSalt)
	if len(log.Data) > 0 {
		if err := governanceTimelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceTimelock.abi.Events[event].Inputs {
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

// GovernanceTimelockCallScheduled represents a CallScheduled event raised by the GovernanceTimelock contract.
type GovernanceTimelockCallScheduled struct {
	Id          [32]byte
	Index       *big.Int
	Target      common.Address
	Value       *big.Int
	Data        []byte
	Predecessor [32]byte
	Delay       *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const GovernanceTimelockCallScheduledEventName = "CallScheduled"

// ContractEventName returns the user-defined event name.
func (GovernanceTimelockCallScheduled) ContractEventName() string {
	return GovernanceTimelockCallScheduledEventName
}

// UnpackCallScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (governanceTimelock *GovernanceTimelock) UnpackCallScheduledEvent(log *types.Log) (*GovernanceTimelockCallScheduled, error) {
	event := "CallScheduled"
	if log.Topics[0] != governanceTimelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTimelockCallScheduled)
	if len(log.Data) > 0 {
		if err := governanceTimelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceTimelock.abi.Events[event].Inputs {
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

// GovernanceTimelockCancelled represents a Cancelled event raised by the GovernanceTimelock contract.
type GovernanceTimelockCancelled struct {
	Id  [32]byte
	Raw *types.Log // Blockchain specific contextual infos
}

const GovernanceTimelockCancelledEventName = "Cancelled"

// ContractEventName returns the user-defined event name.
func (GovernanceTimelockCancelled) ContractEventName() string {
	return GovernanceTimelockCancelledEventName
}

// UnpackCancelledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (governanceTimelock *GovernanceTimelock) UnpackCancelledEvent(log *types.Log) (*GovernanceTimelockCancelled, error) {
	event := "Cancelled"
	if log.Topics[0] != governanceTimelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTimelockCancelled)
	if len(log.Data) > 0 {
		if err := governanceTimelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceTimelock.abi.Events[event].Inputs {
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

// GovernanceTimelockInitialized represents a Initialized event raised by the GovernanceTimelock contract.
type GovernanceTimelockInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const GovernanceTimelockInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (GovernanceTimelockInitialized) ContractEventName() string {
	return GovernanceTimelockInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (governanceTimelock *GovernanceTimelock) UnpackInitializedEvent(log *types.Log) (*GovernanceTimelockInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != governanceTimelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTimelockInitialized)
	if len(log.Data) > 0 {
		if err := governanceTimelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceTimelock.abi.Events[event].Inputs {
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

// GovernanceTimelockMinDelayChange represents a MinDelayChange event raised by the GovernanceTimelock contract.
type GovernanceTimelockMinDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const GovernanceTimelockMinDelayChangeEventName = "MinDelayChange"

// ContractEventName returns the user-defined event name.
func (GovernanceTimelockMinDelayChange) ContractEventName() string {
	return GovernanceTimelockMinDelayChangeEventName
}

// UnpackMinDelayChangeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (governanceTimelock *GovernanceTimelock) UnpackMinDelayChangeEvent(log *types.Log) (*GovernanceTimelockMinDelayChange, error) {
	event := "MinDelayChange"
	if log.Topics[0] != governanceTimelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTimelockMinDelayChange)
	if len(log.Data) > 0 {
		if err := governanceTimelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceTimelock.abi.Events[event].Inputs {
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

// GovernanceTimelockRoleAdminChanged represents a RoleAdminChanged event raised by the GovernanceTimelock contract.
type GovernanceTimelockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const GovernanceTimelockRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (GovernanceTimelockRoleAdminChanged) ContractEventName() string {
	return GovernanceTimelockRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (governanceTimelock *GovernanceTimelock) UnpackRoleAdminChangedEvent(log *types.Log) (*GovernanceTimelockRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != governanceTimelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTimelockRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := governanceTimelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceTimelock.abi.Events[event].Inputs {
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

// GovernanceTimelockRoleGranted represents a RoleGranted event raised by the GovernanceTimelock contract.
type GovernanceTimelockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const GovernanceTimelockRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (GovernanceTimelockRoleGranted) ContractEventName() string {
	return GovernanceTimelockRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (governanceTimelock *GovernanceTimelock) UnpackRoleGrantedEvent(log *types.Log) (*GovernanceTimelockRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != governanceTimelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTimelockRoleGranted)
	if len(log.Data) > 0 {
		if err := governanceTimelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceTimelock.abi.Events[event].Inputs {
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

// GovernanceTimelockRoleRevoked represents a RoleRevoked event raised by the GovernanceTimelock contract.
type GovernanceTimelockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const GovernanceTimelockRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (GovernanceTimelockRoleRevoked) ContractEventName() string {
	return GovernanceTimelockRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (governanceTimelock *GovernanceTimelock) UnpackRoleRevokedEvent(log *types.Log) (*GovernanceTimelockRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != governanceTimelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernanceTimelockRoleRevoked)
	if len(log.Data) > 0 {
		if err := governanceTimelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governanceTimelock.abi.Events[event].Inputs {
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
func (governanceTimelock *GovernanceTimelock) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["TimelockInsufficientDelay"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackTimelockInsufficientDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["TimelockInvalidOperationLength"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackTimelockInvalidOperationLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["TimelockUnauthorizedCaller"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackTimelockUnauthorizedCallerError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["TimelockUnexecutedPredecessor"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackTimelockUnexecutedPredecessorError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceTimelock.abi.Errors["TimelockUnexpectedOperationState"].ID.Bytes()[:4]) {
		return governanceTimelock.UnpackTimelockUnexpectedOperationStateError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// GovernanceTimelockAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the GovernanceTimelock contract.
type GovernanceTimelockAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func GovernanceTimelockAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (governanceTimelock *GovernanceTimelock) UnpackAccessControlBadConfirmationError(raw []byte) (*GovernanceTimelockAccessControlBadConfirmation, error) {
	out := new(GovernanceTimelockAccessControlBadConfirmation)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the GovernanceTimelock contract.
type GovernanceTimelockAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func GovernanceTimelockAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (governanceTimelock *GovernanceTimelock) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*GovernanceTimelockAccessControlUnauthorizedAccount, error) {
	out := new(GovernanceTimelockAccessControlUnauthorizedAccount)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockInvalidInitialization represents a InvalidInitialization error raised by the GovernanceTimelock contract.
type GovernanceTimelockInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func GovernanceTimelockInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (governanceTimelock *GovernanceTimelock) UnpackInvalidInitializationError(raw []byte) (*GovernanceTimelockInvalidInitialization, error) {
	out := new(GovernanceTimelockInvalidInitialization)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockInvalidValue represents a InvalidValue error raised by the GovernanceTimelock contract.
type GovernanceTimelockInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func GovernanceTimelockInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (governanceTimelock *GovernanceTimelock) UnpackInvalidValueError(raw []byte) (*GovernanceTimelockInvalidValue, error) {
	out := new(GovernanceTimelockInvalidValue)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockNotInitializing represents a NotInitializing error raised by the GovernanceTimelock contract.
type GovernanceTimelockNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func GovernanceTimelockNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (governanceTimelock *GovernanceTimelock) UnpackNotInitializingError(raw []byte) (*GovernanceTimelockNotInitializing, error) {
	out := new(GovernanceTimelockNotInitializing)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockOnlyCoinbase represents a OnlyCoinbase error raised by the GovernanceTimelock contract.
type GovernanceTimelockOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func GovernanceTimelockOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (governanceTimelock *GovernanceTimelock) UnpackOnlyCoinbaseError(raw []byte) (*GovernanceTimelockOnlyCoinbase, error) {
	out := new(GovernanceTimelockOnlyCoinbase)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockOnlySystemContract represents a OnlySystemContract error raised by the GovernanceTimelock contract.
type GovernanceTimelockOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func GovernanceTimelockOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (governanceTimelock *GovernanceTimelock) UnpackOnlySystemContractError(raw []byte) (*GovernanceTimelockOnlySystemContract, error) {
	out := new(GovernanceTimelockOnlySystemContract)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the GovernanceTimelock contract.
type GovernanceTimelockOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func GovernanceTimelockOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (governanceTimelock *GovernanceTimelock) UnpackOnlyZeroGasPriceError(raw []byte) (*GovernanceTimelockOnlyZeroGasPrice, error) {
	out := new(GovernanceTimelockOnlyZeroGasPrice)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockTimelockInsufficientDelay represents a TimelockInsufficientDelay error raised by the GovernanceTimelock contract.
type GovernanceTimelockTimelockInsufficientDelay struct {
	Delay    *big.Int
	MinDelay *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockInsufficientDelay(uint256 delay, uint256 minDelay)
func GovernanceTimelockTimelockInsufficientDelayErrorID() common.Hash {
	return common.HexToHash("0x543366097fc46a718a4890a59e5216cbbc872f61973a9e2a6666ecccaa1a07ca")
}

// UnpackTimelockInsufficientDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockInsufficientDelay(uint256 delay, uint256 minDelay)
func (governanceTimelock *GovernanceTimelock) UnpackTimelockInsufficientDelayError(raw []byte) (*GovernanceTimelockTimelockInsufficientDelay, error) {
	out := new(GovernanceTimelockTimelockInsufficientDelay)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "TimelockInsufficientDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockTimelockInvalidOperationLength represents a TimelockInvalidOperationLength error raised by the GovernanceTimelock contract.
type GovernanceTimelockTimelockInvalidOperationLength struct {
	Targets  *big.Int
	Payloads *big.Int
	Values   *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockInvalidOperationLength(uint256 targets, uint256 payloads, uint256 values)
func GovernanceTimelockTimelockInvalidOperationLengthErrorID() common.Hash {
	return common.HexToHash("0xffb0321166292613ddbe36e2eb9b9b1e8877aa314505f6d35f0f4ae660e8ada3")
}

// UnpackTimelockInvalidOperationLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockInvalidOperationLength(uint256 targets, uint256 payloads, uint256 values)
func (governanceTimelock *GovernanceTimelock) UnpackTimelockInvalidOperationLengthError(raw []byte) (*GovernanceTimelockTimelockInvalidOperationLength, error) {
	out := new(GovernanceTimelockTimelockInvalidOperationLength)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "TimelockInvalidOperationLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockTimelockUnauthorizedCaller represents a TimelockUnauthorizedCaller error raised by the GovernanceTimelock contract.
type GovernanceTimelockTimelockUnauthorizedCaller struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockUnauthorizedCaller(address caller)
func GovernanceTimelockTimelockUnauthorizedCallerErrorID() common.Hash {
	return common.HexToHash("0xe2850c5900ceb2d1b367e63ffd96510279f191862fece2dde12a1b1bce580ebd")
}

// UnpackTimelockUnauthorizedCallerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockUnauthorizedCaller(address caller)
func (governanceTimelock *GovernanceTimelock) UnpackTimelockUnauthorizedCallerError(raw []byte) (*GovernanceTimelockTimelockUnauthorizedCaller, error) {
	out := new(GovernanceTimelockTimelockUnauthorizedCaller)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "TimelockUnauthorizedCaller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockTimelockUnexecutedPredecessor represents a TimelockUnexecutedPredecessor error raised by the GovernanceTimelock contract.
type GovernanceTimelockTimelockUnexecutedPredecessor struct {
	PredecessorId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockUnexecutedPredecessor(bytes32 predecessorId)
func GovernanceTimelockTimelockUnexecutedPredecessorErrorID() common.Hash {
	return common.HexToHash("0x90a9a618cfd84aabd2505bb50bbc6c95924a5d4f10d2bf500768fcbf91f51cab")
}

// UnpackTimelockUnexecutedPredecessorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockUnexecutedPredecessor(bytes32 predecessorId)
func (governanceTimelock *GovernanceTimelock) UnpackTimelockUnexecutedPredecessorError(raw []byte) (*GovernanceTimelockTimelockUnexecutedPredecessor, error) {
	out := new(GovernanceTimelockTimelockUnexecutedPredecessor)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "TimelockUnexecutedPredecessor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceTimelockTimelockUnexpectedOperationState represents a TimelockUnexpectedOperationState error raised by the GovernanceTimelock contract.
type GovernanceTimelockTimelockUnexpectedOperationState struct {
	OperationId    [32]byte
	ExpectedStates [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockUnexpectedOperationState(bytes32 operationId, bytes32 expectedStates)
func GovernanceTimelockTimelockUnexpectedOperationStateErrorID() common.Hash {
	return common.HexToHash("0x5ead8eb51d1c1b7ef2eb1ef3ec2f009cfba25e924d04ccc853f25803ea40b489")
}

// UnpackTimelockUnexpectedOperationStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockUnexpectedOperationState(bytes32 operationId, bytes32 expectedStates)
func (governanceTimelock *GovernanceTimelock) UnpackTimelockUnexpectedOperationStateError(raw []byte) (*GovernanceTimelockTimelockUnexpectedOperationState, error) {
	out := new(GovernanceTimelockTimelockUnexpectedOperationState)
	if err := governanceTimelock.abi.UnpackIntoInterface(out, "TimelockUnexpectedOperationState", raw); err != nil {
		return nil, err
	}
	return out, nil
}
