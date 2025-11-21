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

// ValidatorShareUnbondingInfo is an auto generated low-level Go binding around an user-defined struct.
type ValidatorShareUnbondingInfo struct {
	Assets     *big.Int
	Shares     *big.Int
	UnbondTime *big.Int
	Reserved   [7]*big.Int
}

// ValidatorShareMetaData contains all meta data concerning the ValidatorShare contract.
var ValidatorShareMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"COMMISSION_RATE_BASE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[{\"name\":\"totalClaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertToAssets\",\"inputs\":[{\"name\":\"share\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertToShares\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"distributeReward\",\"inputs\":[{\"name\":\"commissionRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getDelegators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnbondingInfo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structValidatorShare.UnbondingInfo[]\",\"components\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unbondTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_reserved\",\"type\":\"uint256[7]\",\"internalType\":\"uint256[7]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"validator_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"maxDeposit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"slashAssets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"slashedAssets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorClaimAll\",\"inputs\":[],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyUnbonding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApproveNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAssets\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDelegator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughAssets\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotUnbonding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StillUnbonding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferNotAllowed\",\"inputs\":[]}]",
	ID:         "ValidatorShare",
	BinRuntime: "0x6080806040526004361015610049575b50361561001a575f80fd5b61100233036100345761002f3460015461145f565b600155005b630f22c43960e41b5f5261100260045260245ffd5b5f3560e01c90816301e1d114146112d85750806306fdde031461121b57806307a2d13a14610d2d578063095ea7b3146111f257806318160ddd146111c95780631e83409a14610f6f57806323b872dd14610e9f578063313ce56714610e8457806338d52e0f14610e6a5780633a5381b514610e43578063402d267d14610e1d57806345bc4d1014610d325780634cdad50614610d2d5780635e607d7614610cb15780636e553f6514610c3e57806370a08231146106fe5780637c78024814610c1b57806383a96e1614610b5c57806395d89b4114610a6a578063a9059cbb14610a47578063ba08765214610836578063bee8380e1461081a578063c6e6f592146106b2578063ce96cb77146107d6578063d85b203914610729578063d905777e146106fe578063dd62ed3e146106b7578063ef8b30f7146106b25763f399e22e14610194575f61000f565b6040366003190112610457576101a861133a565b6024359067ffffffffffffffff821161045757366023830112156104575781600401359167ffffffffffffffff83116104575760248101906024843692010111610457575f516020611c4a5f395f51905f52549060ff8260401c16159167ffffffffffffffff8116801590816106aa575b60011490816106a0575b159081610697575b506106885767ffffffffffffffff1981166001175f516020611c4a5f395f51905f52558261065c575b506110023303610034576102759061026d3686836117ab565b9436916117ab565b9161027e611a5f565b610286611a5f565b835167ffffffffffffffff8111610568576102ae5f516020611bca5f395f51905f5254611384565b601f8111610602575b50602094601f8211600114610587579481929394955f9261057c575b50508160011b915f199060031b1c1916175f516020611bca5f395f51905f52555b825167ffffffffffffffff81116105685761031c5f516020611c0a5f395f51905f5254611384565b601f8111610503575b506020601f821160011461048857819293945f9261047d575b50508160011b915f199060031b1c1916175f516020611c0a5f395f51905f52555b610367611a5f565b6001600160a01b0316801561046e575f80546001600160a01b03191682179055604051633c621da160e21b81526020816004816110025afa908115610463575f9161042d575b50341061041e576103c1903490349061187a565b6103c757005b68ff0000000000000000195f516020611c4a5f395f51905f5254165f516020611c4a5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b630de1bf7560e21b5f5260045ffd5b90506020813d60201161045b575b81610448602093836113d8565b8101031261045757515f6103ad565b5f80fd5b3d915061043b565b6040513d5f823e3d90fd5b63b9f0f17160e01b5f5260045ffd5b015190505f8061033e565b601f198216905f516020611c0a5f395f51905f525f52805f20915f5b8181106104eb575095836001959697106104d3575b505050811b015f516020611c0a5f395f51905f525561035f565b01515f1960f88460031b161c191690555f80806104b9565b9192602060018192868b0151815501940192016104a4565b5f516020611c0a5f395f51905f525f52610558907f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa601f840160051c8101916020851061055e575b601f0160051c019061146c565b5f610325565b909150819061054b565b634e487b7160e01b5f52604160045260245ffd5b015190505f806102d3565b601f198216955f516020611bca5f395f51905f525f52805f20915f5b8881106105ea575083600195969798106105d2575b505050811b015f516020611bca5f395f51905f52556102f4565b01515f1960f88460031b161c191690555f80806105b8565b919260206001819286850151815501940192016105a3565b5f516020611bca5f395f51905f525f52610656907f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0601f840160051c8101916020851061055e57601f0160051c019061146c565b5f6102b7565b68ffffffffffffffffff191668010000000000000001175f516020611c4a5f395f51905f52555f610254565b63f92ee8a960e01b5f5260045ffd5b9050155f61022b565b303b159150610223565b849150610219565b611366565b34610457576040366003190112610457576106d061133a565b6106e16106db611350565b916113fa565b9060018060a01b03165f52602052602060405f2054604051908152f35b3461045757602036600319011261045757602061072161071c61133a565b61150b565b604051908152f35b346104575760203660031901126104575761074a61074561133a565b61163c565b6040518091602082016020835281518091526020604084019201905f5b818110610775575050500390f35b9193509160608451805183526020810151602084015260408101516040840152015160608201905f915b600783106107c0575050506020610140600192019401910191849392610767565b602080600192845181520192019201919061079f565b34610457576020366003190112610457576001600160a01b036107f761133a565b165f525f516020611bea5f395f51905f52602052602061072160405f20546117e1565b34610457575f3660031901126104575760206040516127108152f35b3461045757606036600319011261045757600435610852611350565b50604435906001600160a01b0382169081830361045757611002330361003457610887825f52600360205260405f2054151590565b1561046e578015610a3857815f525f516020611bea5f395f51905f5260205260405f20548111610a29576108ba816117e1565b928215610a1657816108cb91611933565b6001548311610a0d5782600154036001555b815f525f516020611bea5f395f51905f5260205260405f2054156109fe575b604051636cf6d67560e01b81526020816004816110025afa8015610463575f906109ca575b61092c91504261145f565b916109356115d9565b9084825260208201928352604082019384525f52600460205260405f208054600160401b8110156105685761096f91600182018155611432565b9390936109b7576060928251855551600185015551600284015501515f5b600781106109a057602084604051908152f35b60019060208351930192600382860101550161098d565b634e487b7160e01b5f525f60045260245ffd5b506020813d6020116109f6575b816109e4602093836113d8565b810103126104575761092c9051610921565b3d91506109d7565b610a0782611a9f565b506108fc565b5f6001556108dd565b634b637e8f60e11b5f525f60045260245ffd5b633c57b48560e21b5f5260045ffd5b636edcc52360e01b5f5260045ffd5b3461045757604036600319011261045757610a6061133a565b6024359033611814565b34610457575f366003190112610457576040515f5f516020611c0a5f395f51905f5254610a9681611384565b8084529060018116908115610b385750600114610ace575b610aca83610abe818503826113d8565b604051918291826112f2565b0390f35b5f516020611c0a5f395f51905f525f9081527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa939250905b808210610b1e57509091508101602001610abe610aae565b919260018160209254838588010152019101909291610b06565b60ff191660208086019190915291151560051b84019091019150610abe9050610aae565b34610457575f366003190112610457576040518060206002549283815201809260025f527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace905f5b818110610c055750505081610bba9103826113d8565b604051918291602083019060208452518091526040830191905f5b818110610be3575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610bd5565b8254845260209093019260019283019201610ba4565b34610457575f366003190112610457576110023303610034576020610721611531565b604036600319011261045757600435610c55611350565b906110023303610034576001600160a01b0382161561046e5780158015610ca7575b610c9857610c8481611855565b908115610a2957816107219160209461187a565b6318374fd160e21b5f5260045ffd5b5034811415610c77565b60203660031901126104575760043567ffffffffffffffff8116809103610457576110023303610034573415610c9857610cee6127109134611482565b04610d04610cfc82346114e0565b60015461145f565b600155610d1081611855565b8015610a29575f54610d2b92906001600160a01b031661187a565b005b61131c565b34610457576020366003190112610457576004356110023303610034578015610c98575f546001600160a01b0316610d698161150b565b91610d7381611855565b9083821115610e16575050610d87826117e1565b915b8115610a1657610d999082611933565b6001548211610e0d5781600154036001555b805f525f516020611bea5f395f51905f5260205260405f205415610dfd575b505f808080846110055af1610ddd6114b1565b5015610dee57602090604051908152f35b6312171d8360e31b5f5260045ffd5b610e0690611a9f565b5081610dca565b5f600155610dab565b9250610d89565b3461045757602036600319011261045757610e3661133a565b5060206040516001198152f35b34610457575f366003190112610457575f546040516001600160a01b039091168152602090f35b34610457575f3660031901126104575760206040515f8152f35b34610457575f36600319011261045757602060405160128152f35b3461045757606036600319011261045757610eb861133a565b610ec0611350565b60443591610ecd816113fa565b335f90815260209190915260409020545f198110610eec575b50611814565b838110610f54576001600160a01b03821615610f41573315610f2e57610f11826113fa565b60018060a01b0333165f526020528360405f209103905583610ee6565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8390637dc7a0d960e11b5f523360045260245260445260645ffd5b34610457576020366003190112610457576004356001600160a01b03811690819003610457576110023303610034577f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c6111ba5760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d5f90801561046e57805f52600460205260405f2090805f52600560205260405f2054825492838210156111ab5781905b848210611132575b5050831561112357839281036111095750805f52600460205260405f208054905f81558161109e575b50505f808093838294835260056020528260408120555b5af161106a6114b1565b5015610dee576020905f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d604051908152f35b9080929350600a0291600a8304036110f557908392915f5260205f20908101905b8181101561104957600a919293505f81555f60018201555f60028201556110eb8282016003830161146c565b01908392916110bf565b634e487b7160e01b5f52601160045260245ffd5b5f8093838294938394845260056020526040842055611060565b633fc3ac2b60e01b5f5260045ffd5b90919460026111418484611432565b50015442106111a457611160906111588484611432565b50549061145f565b9461116b8383611432565b6109b75761118f905f81555f60018201555f60028201556003600a8201910161146c565b5f1981146110f5576001809101920190611018565b9491611020565b630107608560e01b5f5260045ffd5b633ee5aeb560e01b5f5260045ffd5b34610457575f3660031901126104575760205f516020611c2a5f395f51905f5254604051908152f35b346104575760403660031901126104575761120b61133a565b50632028747160e01b5f5260045ffd5b34610457575f366003190112610457576040515f5f516020611bca5f395f51905f525461124781611384565b8084529060018116908115610b38575060011461126e57610aca83610abe818503826113d8565b5f516020611bca5f395f51905f525f9081527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0939250905b8082106112be57509091508101602001610abe610aae565b9192600181602092548385880101520191019092916112a6565b34610457575f366003190112610457576020906001548152f35b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b346104575760203660031901126104575760206107216004356117e1565b600435906001600160a01b038216820361045757565b602435906001600160a01b038216820361045757565b34610457576020366003190112610457576020610721600435611855565b90600182811c921680156113b2575b602083101461139e57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691611393565b6080810190811067ffffffffffffffff82111761056857604052565b90601f8019910116810190811067ffffffffffffffff82111761056857604052565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b805482101561144b575f52600a60205f20910201905f90565b634e487b7160e01b5f52603260045260245ffd5b919082018092116110f557565b818110611477575050565b5f815560010161146c565b818102929181159184041417156110f557565b67ffffffffffffffff811161056857601f01601f191660200190565b3d156114db573d906114c282611495565b916114d060405193846113d8565b82523d5f602084013e565b606090565b919082039182116110f557565b81156114f7570490565b634e487b7160e01b5f52601260045260245ffd5b6001600160a01b03165f9081525f516020611bea5f395f51905f52602052604090205490565b5f546001600160a01b0316906115468261150b565b9182156115d357611556836117e1565b928115610a16576115679082611933565b60015483116115ca5782600154036001555b805f525f516020611bea5f395f51905f5260205260405f2054156115ba575b505f8080808560018060a01b038254165af16115b26114b1565b5015610dee57565b6115c390611a9f565b505f611598565b5f600155611579565b505f9150565b604051906115e6826113bc565b815f81525f60208201525f604082015260606040519161160760e0846113d8565b60e03684370152565b67ffffffffffffffff81116105685760051b60200190565b805182101561144b5760209160051b010190565b60018060a01b031690815f52600460205260405f20915f52600560205260405f20549180548084101561176f5761167d8461167781846114e0565b926114e0565b9161168783611610565b9261169560405194856113d8565b8084526116a4601f1991611610565b015f5b8181106117585750505f5b8281106116c157509193505050565b6116d46116ce828861145f565b83611432565b50604051906116e2826113bc565b8054825260018101546020830152600281015460408301526040519060038291015f915b6007831061174257505050908161172160e0600195946113d8565b60608201526117308287611628565b5261173b8186611628565b50016116b2565b6001602081928454815201920192019190611706565b6020906117636115d9565b828288010152016116a7565b50909150506040516117826020826113d8565b5f81525f805b81811061179457505090565b60209061179f6115d9565b82828601015201611788565b9291926117b782611495565b916117c560405193846113d8565b829481845281830111610457578281602093845f960137010152565b5f516020611c2a5f395f51905f5254801561180e5761180661180b9260015490611482565b6114ed565b90565b50505f90565b9091506001600160a01b031615610a16576001600160a01b03161561184257638cd22d1960e01b5f5260045ffd5b63ec442f0560e01b5f525f60045260245ffd5b600154801561180e5761180661180b925f516020611c2a5f395f51905f525490611482565b6001600160a01b03169190821561184257825f7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020611907956118cc815f516020611c2a5f395f51905f525461145f565b5f516020611c2a5f395f51905f52558484525f516020611bea5f395f51905f52825260408420818154019055604051908152a360015461145f565b60015561191f815f52600360205260405f2054151590565b156119275750565b61193090611b74565b50565b9091906001600160a01b03168015801580611a58575b611a4957156119c9577fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020846119905f95965f516020611c2a5f395f51905f525461145f565b5f516020611c2a5f395f51905f52555b805f516020611c2a5f395f51905f5254035f516020611c2a5f395f51905f5255604051908152a3565b805f525f516020611bea5f395f51905f5260205260405f2054838110611a2f576020845f94957fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef938587525f516020611bea5f395f51905f5284520360408620556119a0565b915063391434e360e21b5f5260045260245260445260645ffd5b638cd22d1960e01b5f5260045ffd5b505f611949565b60ff5f516020611c4a5f395f51905f525460401c1615611a7b57565b631afcd79f60e31b5f5260045ffd5b805482101561144b575f5260205f2001905f90565b5f81815260036020526040902054801561180e575f1981018181116110f5576002545f198101919082116110f557818103611b26575b5050506002548015611b12575f1901611aef816002611a8a565b8154905f199060031b1b191690556002555f5260036020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b611b5e611b37611b48936002611a8a565b90549060031b1c9283926002611a8a565b819391549060031b91821b915f19901b19161790565b90555f52600360205260405f20555f8080611ad5565b805f52600360205260405f2054155f14611bc457600254600160401b81101561056857611bad611b488260018594016002556002611a8a565b9055600254905f52600360205260405f2055600190565b505f9056fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0052c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0452c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220c9b4303ca24daa1d5d12e18e7bbccdb8c480ed02ef9e78e7810bfd11764dfee664736f6c634300081c0033",
}

// ValidatorShare is an auto generated Go binding around an Ethereum contract.
type ValidatorShare struct {
	abi abi.ABI
}

// NewValidatorShare creates a new instance of ValidatorShare.
func NewValidatorShare() *ValidatorShare {
	parsed, err := ValidatorShareMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ValidatorShare{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ValidatorShare) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCOMMISSIONRATEBASE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbee8380e.
//
// Solidity: function COMMISSION_RATE_BASE() view returns(uint256)
func (validatorShare *ValidatorShare) PackCOMMISSIONRATEBASE() []byte {
	enc, err := validatorShare.abi.Pack("COMMISSION_RATE_BASE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCOMMISSIONRATEBASE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbee8380e.
//
// Solidity: function COMMISSION_RATE_BASE() view returns(uint256)
func (validatorShare *ValidatorShare) UnpackCOMMISSIONRATEBASE(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("COMMISSION_RATE_BASE", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (validatorShare *ValidatorShare) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := validatorShare.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (validatorShare *ValidatorShare) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("allowance", data)
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
func (validatorShare *ValidatorShare) PackApprove(arg0 common.Address, arg1 *big.Int) []byte {
	enc, err := validatorShare.abi.Pack("approve", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns(bool)
func (validatorShare *ValidatorShare) UnpackApprove(data []byte) (bool, error) {
	out, err := validatorShare.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackAsset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38d52e0f.
//
// Solidity: function asset() pure returns(address)
func (validatorShare *ValidatorShare) PackAsset() []byte {
	enc, err := validatorShare.abi.Pack("asset")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAsset is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x38d52e0f.
//
// Solidity: function asset() pure returns(address)
func (validatorShare *ValidatorShare) UnpackAsset(data []byte) (common.Address, error) {
	out, err := validatorShare.abi.Unpack("asset", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (validatorShare *ValidatorShare) PackBalanceOf(account common.Address) []byte {
	enc, err := validatorShare.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (validatorShare *ValidatorShare) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackClaim is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e83409a.
//
// Solidity: function claim(address owner) returns(uint256 totalClaimed)
func (validatorShare *ValidatorShare) PackClaim(owner common.Address) []byte {
	enc, err := validatorShare.abi.Pack("claim", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackClaim is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1e83409a.
//
// Solidity: function claim(address owner) returns(uint256 totalClaimed)
func (validatorShare *ValidatorShare) UnpackClaim(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("claim", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackConvertToAssets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 share) view returns(uint256)
func (validatorShare *ValidatorShare) PackConvertToAssets(share *big.Int) []byte {
	enc, err := validatorShare.abi.Pack("convertToAssets", share)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackConvertToAssets is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 share) view returns(uint256)
func (validatorShare *ValidatorShare) UnpackConvertToAssets(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("convertToAssets", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackConvertToShares is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (validatorShare *ValidatorShare) PackConvertToShares(assets *big.Int) []byte {
	enc, err := validatorShare.abi.Pack("convertToShares", assets)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackConvertToShares is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (validatorShare *ValidatorShare) UnpackConvertToShares(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("convertToShares", data)
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
func (validatorShare *ValidatorShare) PackDecimals() []byte {
	enc, err := validatorShare.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (validatorShare *ValidatorShare) UnpackDecimals(data []byte) (uint8, error) {
	out, err := validatorShare.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address delegator) payable returns(uint256 shares)
func (validatorShare *ValidatorShare) PackDeposit(assets *big.Int, delegator common.Address) []byte {
	enc, err := validatorShare.abi.Pack("deposit", assets, delegator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDeposit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address delegator) payable returns(uint256 shares)
func (validatorShare *ValidatorShare) UnpackDeposit(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("deposit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDistributeReward is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e607d76.
//
// Solidity: function distributeReward(uint64 commissionRate) payable returns()
func (validatorShare *ValidatorShare) PackDistributeReward(commissionRate uint64) []byte {
	enc, err := validatorShare.abi.Pack("distributeReward", commissionRate)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetDelegators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83a96e16.
//
// Solidity: function getDelegators() view returns(address[])
func (validatorShare *ValidatorShare) PackGetDelegators() []byte {
	enc, err := validatorShare.abi.Pack("getDelegators")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDelegators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x83a96e16.
//
// Solidity: function getDelegators() view returns(address[])
func (validatorShare *ValidatorShare) UnpackGetDelegators(data []byte) ([]common.Address, error) {
	out, err := validatorShare.abi.Unpack("getDelegators", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackGetUnbondingInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd85b2039.
//
// Solidity: function getUnbondingInfo(address account) view returns((uint256,uint256,uint256,uint256[7])[])
func (validatorShare *ValidatorShare) PackGetUnbondingInfo(account common.Address) []byte {
	enc, err := validatorShare.abi.Pack("getUnbondingInfo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUnbondingInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd85b2039.
//
// Solidity: function getUnbondingInfo(address account) view returns((uint256,uint256,uint256,uint256[7])[])
func (validatorShare *ValidatorShare) UnpackGetUnbondingInfo(data []byte) ([]ValidatorShareUnbondingInfo, error) {
	out, err := validatorShare.abi.Unpack("getUnbondingInfo", data)
	if err != nil {
		return *new([]ValidatorShareUnbondingInfo), err
	}
	out0 := *abi.ConvertType(out[0], new([]ValidatorShareUnbondingInfo)).(*[]ValidatorShareUnbondingInfo)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf399e22e.
//
// Solidity: function initialize(address validator_, string id_) payable returns()
func (validatorShare *ValidatorShare) PackInitialize(validator common.Address, id string) []byte {
	enc, err := validatorShare.abi.Pack("initialize", validator, id)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMaxDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (validatorShare *ValidatorShare) PackMaxDeposit(arg0 common.Address) []byte {
	enc, err := validatorShare.abi.Pack("maxDeposit", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxDeposit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (validatorShare *ValidatorShare) UnpackMaxDeposit(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("maxDeposit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackMaxRedeem is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (validatorShare *ValidatorShare) PackMaxRedeem(owner common.Address) []byte {
	enc, err := validatorShare.abi.Pack("maxRedeem", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxRedeem is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (validatorShare *ValidatorShare) UnpackMaxRedeem(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("maxRedeem", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackMaxWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (validatorShare *ValidatorShare) PackMaxWithdraw(owner common.Address) []byte {
	enc, err := validatorShare.abi.Pack("maxWithdraw", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxWithdraw is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (validatorShare *ValidatorShare) UnpackMaxWithdraw(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("maxWithdraw", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (validatorShare *ValidatorShare) PackName() []byte {
	enc, err := validatorShare.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (validatorShare *ValidatorShare) UnpackName(data []byte) (string, error) {
	out, err := validatorShare.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackPreviewDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (validatorShare *ValidatorShare) PackPreviewDeposit(assets *big.Int) []byte {
	enc, err := validatorShare.abi.Pack("previewDeposit", assets)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPreviewDeposit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (validatorShare *ValidatorShare) UnpackPreviewDeposit(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("previewDeposit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPreviewRedeem is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (validatorShare *ValidatorShare) PackPreviewRedeem(shares *big.Int) []byte {
	enc, err := validatorShare.abi.Pack("previewRedeem", shares)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPreviewRedeem is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (validatorShare *ValidatorShare) UnpackPreviewRedeem(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("previewRedeem", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackRedeem is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (validatorShare *ValidatorShare) PackRedeem(shares *big.Int, receiver common.Address, owner common.Address) []byte {
	enc, err := validatorShare.abi.Pack("redeem", shares, receiver, owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRedeem is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (validatorShare *ValidatorShare) UnpackRedeem(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("redeem", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackSlash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x45bc4d10.
//
// Solidity: function slash(uint256 slashAssets) returns(uint256 slashedAssets)
func (validatorShare *ValidatorShare) PackSlash(slashAssets *big.Int) []byte {
	enc, err := validatorShare.abi.Pack("slash", slashAssets)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSlash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x45bc4d10.
//
// Solidity: function slash(uint256 slashAssets) returns(uint256 slashedAssets)
func (validatorShare *ValidatorShare) UnpackSlash(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("slash", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (validatorShare *ValidatorShare) PackSymbol() []byte {
	enc, err := validatorShare.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (validatorShare *ValidatorShare) UnpackSymbol(data []byte) (string, error) {
	out, err := validatorShare.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTotalAssets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (validatorShare *ValidatorShare) PackTotalAssets() []byte {
	enc, err := validatorShare.abi.Pack("totalAssets")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalAssets is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (validatorShare *ValidatorShare) UnpackTotalAssets(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("totalAssets", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (validatorShare *ValidatorShare) PackTotalSupply() []byte {
	enc, err := validatorShare.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (validatorShare *ValidatorShare) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("totalSupply", data)
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
func (validatorShare *ValidatorShare) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := validatorShare.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (validatorShare *ValidatorShare) UnpackTransfer(data []byte) (bool, error) {
	out, err := validatorShare.abi.Unpack("transfer", data)
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
func (validatorShare *ValidatorShare) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := validatorShare.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (validatorShare *ValidatorShare) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := validatorShare.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (validatorShare *ValidatorShare) PackValidator() []byte {
	enc, err := validatorShare.abi.Pack("validator")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (validatorShare *ValidatorShare) UnpackValidator(data []byte) (common.Address, error) {
	out, err := validatorShare.abi.Unpack("validator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackValidatorClaimAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c780248.
//
// Solidity: function validatorClaimAll() returns(uint256 assets)
func (validatorShare *ValidatorShare) PackValidatorClaimAll() []byte {
	enc, err := validatorShare.abi.Pack("validatorClaimAll")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidatorClaimAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7c780248.
//
// Solidity: function validatorClaimAll() returns(uint256 assets)
func (validatorShare *ValidatorShare) UnpackValidatorClaimAll(data []byte) (*big.Int, error) {
	out, err := validatorShare.abi.Unpack("validatorClaimAll", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// ValidatorShareApproval represents a Approval event raised by the ValidatorShare contract.
type ValidatorShareApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ValidatorShareApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ValidatorShareApproval) ContractEventName() string {
	return ValidatorShareApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (validatorShare *ValidatorShare) UnpackApprovalEvent(log *types.Log) (*ValidatorShareApproval, error) {
	event := "Approval"
	if log.Topics[0] != validatorShare.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ValidatorShareApproval)
	if len(log.Data) > 0 {
		if err := validatorShare.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range validatorShare.abi.Events[event].Inputs {
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

// ValidatorShareInitialized represents a Initialized event raised by the ValidatorShare contract.
type ValidatorShareInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ValidatorShareInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ValidatorShareInitialized) ContractEventName() string {
	return ValidatorShareInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (validatorShare *ValidatorShare) UnpackInitializedEvent(log *types.Log) (*ValidatorShareInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != validatorShare.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ValidatorShareInitialized)
	if len(log.Data) > 0 {
		if err := validatorShare.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range validatorShare.abi.Events[event].Inputs {
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

// ValidatorShareTransfer represents a Transfer event raised by the ValidatorShare contract.
type ValidatorShareTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ValidatorShareTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ValidatorShareTransfer) ContractEventName() string {
	return ValidatorShareTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (validatorShare *ValidatorShare) UnpackTransferEvent(log *types.Log) (*ValidatorShareTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != validatorShare.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ValidatorShareTransfer)
	if len(log.Data) > 0 {
		if err := validatorShare.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range validatorShare.abi.Events[event].Inputs {
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
func (validatorShare *ValidatorShare) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["AlreadyUnbonding"].ID.Bytes()[:4]) {
		return validatorShare.UnpackAlreadyUnbondingError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["ApproveNotAllowed"].ID.Bytes()[:4]) {
		return validatorShare.UnpackApproveNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return validatorShare.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return validatorShare.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return validatorShare.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return validatorShare.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return validatorShare.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return validatorShare.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return validatorShare.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["InvalidAssets"].ID.Bytes()[:4]) {
		return validatorShare.UnpackInvalidAssetsError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["InvalidDelegator"].ID.Bytes()[:4]) {
		return validatorShare.UnpackInvalidDelegatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return validatorShare.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["InvalidShares"].ID.Bytes()[:4]) {
		return validatorShare.UnpackInvalidSharesError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return validatorShare.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["NotEnoughAssets"].ID.Bytes()[:4]) {
		return validatorShare.UnpackNotEnoughAssetsError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["NotEnoughShares"].ID.Bytes()[:4]) {
		return validatorShare.UnpackNotEnoughSharesError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return validatorShare.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["NotUnbonding"].ID.Bytes()[:4]) {
		return validatorShare.UnpackNotUnbondingError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return validatorShare.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return validatorShare.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return validatorShare.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return validatorShare.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["StillUnbonding"].ID.Bytes()[:4]) {
		return validatorShare.UnpackStillUnbondingError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["TransferFailed"].ID.Bytes()[:4]) {
		return validatorShare.UnpackTransferFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorShare.abi.Errors["TransferNotAllowed"].ID.Bytes()[:4]) {
		return validatorShare.UnpackTransferNotAllowedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ValidatorShareAlreadyUnbonding represents a AlreadyUnbonding error raised by the ValidatorShare contract.
type ValidatorShareAlreadyUnbonding struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyUnbonding()
func ValidatorShareAlreadyUnbondingErrorID() common.Hash {
	return common.HexToHash("0x9f064e927c2ba9921d737748f0e02adc3e3deb3f729ec7e04f3ec799a688eb90")
}

// UnpackAlreadyUnbondingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyUnbonding()
func (validatorShare *ValidatorShare) UnpackAlreadyUnbondingError(raw []byte) (*ValidatorShareAlreadyUnbonding, error) {
	out := new(ValidatorShareAlreadyUnbonding)
	if err := validatorShare.abi.UnpackIntoInterface(out, "AlreadyUnbonding", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareApproveNotAllowed represents a ApproveNotAllowed error raised by the ValidatorShare contract.
type ValidatorShareApproveNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ApproveNotAllowed()
func ValidatorShareApproveNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x20287471febb94d2a4e7be5d3662f51a83f2bee180c5c2ea14db4432e1cd3e48")
}

// UnpackApproveNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ApproveNotAllowed()
func (validatorShare *ValidatorShare) UnpackApproveNotAllowedError(raw []byte) (*ValidatorShareApproveNotAllowed, error) {
	out := new(ValidatorShareApproveNotAllowed)
	if err := validatorShare.abi.UnpackIntoInterface(out, "ApproveNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ValidatorShare contract.
type ValidatorShareERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ValidatorShareERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (validatorShare *ValidatorShare) UnpackERC20InsufficientAllowanceError(raw []byte) (*ValidatorShareERC20InsufficientAllowance, error) {
	out := new(ValidatorShareERC20InsufficientAllowance)
	if err := validatorShare.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ValidatorShare contract.
type ValidatorShareERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ValidatorShareERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (validatorShare *ValidatorShare) UnpackERC20InsufficientBalanceError(raw []byte) (*ValidatorShareERC20InsufficientBalance, error) {
	out := new(ValidatorShareERC20InsufficientBalance)
	if err := validatorShare.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ValidatorShare contract.
type ValidatorShareERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ValidatorShareERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (validatorShare *ValidatorShare) UnpackERC20InvalidApproverError(raw []byte) (*ValidatorShareERC20InvalidApprover, error) {
	out := new(ValidatorShareERC20InvalidApprover)
	if err := validatorShare.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ValidatorShare contract.
type ValidatorShareERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ValidatorShareERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (validatorShare *ValidatorShare) UnpackERC20InvalidReceiverError(raw []byte) (*ValidatorShareERC20InvalidReceiver, error) {
	out := new(ValidatorShareERC20InvalidReceiver)
	if err := validatorShare.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareERC20InvalidSender represents a ERC20InvalidSender error raised by the ValidatorShare contract.
type ValidatorShareERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ValidatorShareERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (validatorShare *ValidatorShare) UnpackERC20InvalidSenderError(raw []byte) (*ValidatorShareERC20InvalidSender, error) {
	out := new(ValidatorShareERC20InvalidSender)
	if err := validatorShare.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ValidatorShare contract.
type ValidatorShareERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ValidatorShareERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (validatorShare *ValidatorShare) UnpackERC20InvalidSpenderError(raw []byte) (*ValidatorShareERC20InvalidSpender, error) {
	out := new(ValidatorShareERC20InvalidSpender)
	if err := validatorShare.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareInsufficientBalance represents a InsufficientBalance error raised by the ValidatorShare contract.
type ValidatorShareInsufficientBalance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance()
func ValidatorShareInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xf4d678b8ce6b5157126b1484a53523762a93571537a7d5ae97d8014a44715c94")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance()
func (validatorShare *ValidatorShare) UnpackInsufficientBalanceError(raw []byte) (*ValidatorShareInsufficientBalance, error) {
	out := new(ValidatorShareInsufficientBalance)
	if err := validatorShare.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareInvalidAssets represents a InvalidAssets error raised by the ValidatorShare contract.
type ValidatorShareInvalidAssets struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAssets()
func ValidatorShareInvalidAssetsErrorID() common.Hash {
	return common.HexToHash("0x60dd3f448144d4e0eb37fdff3e4d839dfbcb63882f1685952bdd4da4746acf9e")
}

// UnpackInvalidAssetsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAssets()
func (validatorShare *ValidatorShare) UnpackInvalidAssetsError(raw []byte) (*ValidatorShareInvalidAssets, error) {
	out := new(ValidatorShareInvalidAssets)
	if err := validatorShare.abi.UnpackIntoInterface(out, "InvalidAssets", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareInvalidDelegator represents a InvalidDelegator error raised by the ValidatorShare contract.
type ValidatorShareInvalidDelegator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDelegator()
func ValidatorShareInvalidDelegatorErrorID() common.Hash {
	return common.HexToHash("0xb9f0f171cd404c8c5cc9943d60966d64de0cd7b803c221dea8d1f1b85443edad")
}

// UnpackInvalidDelegatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDelegator()
func (validatorShare *ValidatorShare) UnpackInvalidDelegatorError(raw []byte) (*ValidatorShareInvalidDelegator, error) {
	out := new(ValidatorShareInvalidDelegator)
	if err := validatorShare.abi.UnpackIntoInterface(out, "InvalidDelegator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareInvalidInitialization represents a InvalidInitialization error raised by the ValidatorShare contract.
type ValidatorShareInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ValidatorShareInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (validatorShare *ValidatorShare) UnpackInvalidInitializationError(raw []byte) (*ValidatorShareInvalidInitialization, error) {
	out := new(ValidatorShareInvalidInitialization)
	if err := validatorShare.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareInvalidShares represents a InvalidShares error raised by the ValidatorShare contract.
type ValidatorShareInvalidShares struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShares()
func ValidatorShareInvalidSharesErrorID() common.Hash {
	return common.HexToHash("0x6edcc5232bad22a9fc844f62349d66192e45664a1e24b557e8ed043ba3c3dfd4")
}

// UnpackInvalidSharesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShares()
func (validatorShare *ValidatorShare) UnpackInvalidSharesError(raw []byte) (*ValidatorShareInvalidShares, error) {
	out := new(ValidatorShareInvalidShares)
	if err := validatorShare.abi.UnpackIntoInterface(out, "InvalidShares", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareInvalidValue represents a InvalidValue error raised by the ValidatorShare contract.
type ValidatorShareInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func ValidatorShareInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (validatorShare *ValidatorShare) UnpackInvalidValueError(raw []byte) (*ValidatorShareInvalidValue, error) {
	out := new(ValidatorShareInvalidValue)
	if err := validatorShare.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareNotEnoughAssets represents a NotEnoughAssets error raised by the ValidatorShare contract.
type ValidatorShareNotEnoughAssets struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEnoughAssets()
func ValidatorShareNotEnoughAssetsErrorID() common.Hash {
	return common.HexToHash("0x3786fdd444eef70b91a201761543d6da1f00472777ad07400c3534ae31b39c5d")
}

// UnpackNotEnoughAssetsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEnoughAssets()
func (validatorShare *ValidatorShare) UnpackNotEnoughAssetsError(raw []byte) (*ValidatorShareNotEnoughAssets, error) {
	out := new(ValidatorShareNotEnoughAssets)
	if err := validatorShare.abi.UnpackIntoInterface(out, "NotEnoughAssets", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareNotEnoughShares represents a NotEnoughShares error raised by the ValidatorShare contract.
type ValidatorShareNotEnoughShares struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEnoughShares()
func ValidatorShareNotEnoughSharesErrorID() common.Hash {
	return common.HexToHash("0xf15ed2146de0776d0c3238708a746d883c49ffc423a1269fe2022bfaf59c7446")
}

// UnpackNotEnoughSharesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEnoughShares()
func (validatorShare *ValidatorShare) UnpackNotEnoughSharesError(raw []byte) (*ValidatorShareNotEnoughShares, error) {
	out := new(ValidatorShareNotEnoughShares)
	if err := validatorShare.abi.UnpackIntoInterface(out, "NotEnoughShares", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareNotInitializing represents a NotInitializing error raised by the ValidatorShare contract.
type ValidatorShareNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ValidatorShareNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (validatorShare *ValidatorShare) UnpackNotInitializingError(raw []byte) (*ValidatorShareNotInitializing, error) {
	out := new(ValidatorShareNotInitializing)
	if err := validatorShare.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareNotUnbonding represents a NotUnbonding error raised by the ValidatorShare contract.
type ValidatorShareNotUnbonding struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotUnbonding()
func ValidatorShareNotUnbondingErrorID() common.Hash {
	return common.HexToHash("0x010760850e21b4b279a419fb7d6b977871bae131b80149964565307c2558b467")
}

// UnpackNotUnbondingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotUnbonding()
func (validatorShare *ValidatorShare) UnpackNotUnbondingError(raw []byte) (*ValidatorShareNotUnbonding, error) {
	out := new(ValidatorShareNotUnbonding)
	if err := validatorShare.abi.UnpackIntoInterface(out, "NotUnbonding", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareOnlyCoinbase represents a OnlyCoinbase error raised by the ValidatorShare contract.
type ValidatorShareOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func ValidatorShareOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (validatorShare *ValidatorShare) UnpackOnlyCoinbaseError(raw []byte) (*ValidatorShareOnlyCoinbase, error) {
	out := new(ValidatorShareOnlyCoinbase)
	if err := validatorShare.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareOnlySystemContract represents a OnlySystemContract error raised by the ValidatorShare contract.
type ValidatorShareOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func ValidatorShareOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (validatorShare *ValidatorShare) UnpackOnlySystemContractError(raw []byte) (*ValidatorShareOnlySystemContract, error) {
	out := new(ValidatorShareOnlySystemContract)
	if err := validatorShare.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the ValidatorShare contract.
type ValidatorShareOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func ValidatorShareOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (validatorShare *ValidatorShare) UnpackOnlyZeroGasPriceError(raw []byte) (*ValidatorShareOnlyZeroGasPrice, error) {
	out := new(ValidatorShareOnlyZeroGasPrice)
	if err := validatorShare.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the ValidatorShare contract.
type ValidatorShareReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func ValidatorShareReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (validatorShare *ValidatorShare) UnpackReentrancyGuardReentrantCallError(raw []byte) (*ValidatorShareReentrancyGuardReentrantCall, error) {
	out := new(ValidatorShareReentrancyGuardReentrantCall)
	if err := validatorShare.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareStillUnbonding represents a StillUnbonding error raised by the ValidatorShare contract.
type ValidatorShareStillUnbonding struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StillUnbonding()
func ValidatorShareStillUnbondingErrorID() common.Hash {
	return common.HexToHash("0x3fc3ac2b71b9d6db4ffce6c8f98f80546ec314fba8589bd22e61e2789b2a5974")
}

// UnpackStillUnbondingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StillUnbonding()
func (validatorShare *ValidatorShare) UnpackStillUnbondingError(raw []byte) (*ValidatorShareStillUnbonding, error) {
	out := new(ValidatorShareStillUnbonding)
	if err := validatorShare.abi.UnpackIntoInterface(out, "StillUnbonding", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareTransferFailed represents a TransferFailed error raised by the ValidatorShare contract.
type ValidatorShareTransferFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferFailed()
func ValidatorShareTransferFailedErrorID() common.Hash {
	return common.HexToHash("0x90b8ec1877afffd816d05d9b13947f3ff18ec5851c38bad15ec2b710f92391b1")
}

// UnpackTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferFailed()
func (validatorShare *ValidatorShare) UnpackTransferFailedError(raw []byte) (*ValidatorShareTransferFailed, error) {
	out := new(ValidatorShareTransferFailed)
	if err := validatorShare.abi.UnpackIntoInterface(out, "TransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorShareTransferNotAllowed represents a TransferNotAllowed error raised by the ValidatorShare contract.
type ValidatorShareTransferNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferNotAllowed()
func ValidatorShareTransferNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x8cd22d191543fbb893cff76797482bc41fda8a261404bdcd20ff77b9dec06ff5")
}

// UnpackTransferNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferNotAllowed()
func (validatorShare *ValidatorShare) UnpackTransferNotAllowedError(raw []byte) (*ValidatorShareTransferNotAllowed, error) {
	out := new(ValidatorShareTransferNotAllowed)
	if err := validatorShare.abi.UnpackIntoInterface(out, "TransferNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}
