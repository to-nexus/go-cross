// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package predeploy

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IBaseBridgeBridgeTokenArguments is an auto generated low-level Go binding around an user-defined struct.
type IBaseBridgeBridgeTokenArguments struct {
	ToChainID *big.Int
	FromToken common.Address
	To        common.Address
	Value     *big.Int
	GasFee    *big.Int
	ExFee     *big.Int
	ExtraData []byte
}

// IBaseBridgePermitArguments is an auto generated low-level Go binding around an user-defined struct.
type IBaseBridgePermitArguments struct {
	Token    common.Address
	Account  common.Address
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// IBridgeRegistryFinalizeArguments is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryFinalizeArguments struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	ExtraData   []byte
}

// IBridgeRegistryPendingData is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryPendingData struct {
	Args            IBridgeRegistryFinalizeArguments
	DelayExpiration *big.Int
	Reason          []byte
}

// IBridgeRegistryTokenPair is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryTokenPair struct {
	LocalToken                  common.Address
	RemoteToken                 common.Address
	Paused                      bool
	IsOrigin                    bool
	VerificationAmountThreshold *big.Int
	Deposited                   *big.Int
	PendingAmount               *big.Int
}

// CrossBridgeMetaData contains all meta data concerning the CrossBridge contract.
var CrossBridgeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeManager\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"clearPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"conversionRatio\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasExpiredPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualProcessPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"conversionRatio\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"releaseExpiredPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"resetRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeManager\",\"name\":\"_bridgeFeeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"setVerificationAmountThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RoleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"SafetyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"conversionRatio\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"name\":\"RegistryBalanceLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPauseNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPauseNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleAlreadyGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"25bd5666": "bridgeFeeManager()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"df6e87dc": "calculateFee(uint256,address,uint256)",
		"b7f3358d": "changeThreshold(uint8)",
		"0b43c02c": "clearPending(uint256)",
		"d016d625": "createToken(uint256,address,int256,uint256,string,uint8)",
		"4be13f83": "crossMintableERC20Code()",
		"91cca3db": "dev()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"1938e0f2": "finalizeBridge((uint256,uint256,address,address,uint256,bytes),uint8[],bytes32[],bytes32[])",
		"88d67d6d": "finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[],uint8[][],bytes32[][],bytes32[][])",
		"d5717fc5": "getNextFinalizeIndex(uint256)",
		"ae6893f8": "getNextInitiateIndex(uint256)",
		"b33eb36e": "getPendingArguments(uint256,uint256)",
		"a3246ad3": "getRoleMembers(bytes32)",
		"814914b5": "getTokenPair(uint256,address)",
		"3d507c5e": "hasExpiredPending()",
		"91d14854": "hasRole(bytes32,address)",
		"5187599d": "initialize(uint8,address)",
		"91cf6d3e": "initializedAt()",
		"7f4ab9f5": "manualProcessPending(uint256,uint256)",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"4d5d0056": "permitBridgeToken(uint256,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"d605665b": "permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
		"52d1902d": "proxiableUUID()",
		"1e7bf215": "registerToken(uint256,bool,address,address,int256,uint256)",
		"c1cb05dd": "releaseExpiredPending(uint256)",
		"4ee078ba": "releasePending(uint256,uint256)",
		"715018a6": "renounceOwnership()",
		"2d87b7ee": "resetRole(bytes32,address[])",
		"e2af6cd7": "setCrossMintableERC20Code(address)",
		"d477f05f": "setDev(address)",
		"472d35b9": "setFeeManager(address)",
		"bedb86fb": "setPause(bool)",
		"6160751f": "setPauseChain(uint256,bool)",
		"4d3f0da9": "setPauseToken(uint256,address,bool)",
		"d4bf502a": "setRole(bytes32,address[],bool)",
		"3b96cf5a": "setVerificationAmountThreshold(uint256,address,uint256)",
		"aa1bd0bc": "setVerificationDelay(uint256)",
		"4f86d44b": "setVerificationDelay(uint256,uint256,uint256)",
		"42cde4e8": "threshold()",
		"f2fde38b": "transferOwnership(address)",
		"f4509643": "unregisterToken(uint256,address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615ef16100eb5f395f6139b40152615ef15ff3fe6080604052600436106102c2575f3560e01c806384b0196e1161016f578063bedb86fb116100d8578063d5717fc511610092578063e2af6cd71161006d578063e2af6cd714610973578063f2fde38b14610992578063f4509643146109b1578063f698da25146109d0575f5ffd5b8063d5717fc514610907578063d605665b14610926578063df6e87dc14610939575f5ffd5b8063bedb86fb14610858578063c1cb05dd14610877578063cf56118e14610896578063d016d625146108aa578063d477f05f146108c9578063d4bf502a146108e8575f5ffd5b8063a3246ad311610129578063a3246ad314610766578063aa1bd0bc14610792578063ad3cb1cc146107b1578063ae6893f8146107ee578063b33eb36e1461080d578063b7f3358d14610839575f5ffd5b806384b0196e146106c857806388d67d6d146106ef5780638da5cb5b1461070257806391cca3db1461071657806391cf6d3e1461073357806391d1485414610747575f5ffd5b80634ee078ba1161022b5780635c975abb116101e5578063715018a6116101c0578063715018a61461058357806379214874146105975780637f4ab9f5146105c3578063814914b5146105e2575f5ffd5b80635c975abb1461052e5780635fd262de146105515780636160751f14610564575f5ffd5b80634ee078ba146104705780634f1ef2861461048f5780634f86d44b146104a25780635187599d146104c157806352d1902d146104e05780635b605f5c14610502575f5ffd5b80633d507c5e1161027c5780633d507c5e146103c857806342cde4e8146103df578063472d35b9146104005780634be13f831461041f5780634d3f0da91461043e5780634d5d00561461045d575f5ffd5b80630b43c02c146102ed5780631938e0f21461030c5780631e7bf2151461033457806325bd5666146103535780632d87b7ee1461038a5780633b96cf5a146103a9575f5ffd5b366102e957345f036102e7576040516304a4cdd960e51b815260040160405180910390fd5b005b5f5ffd5b3480156102f8575f5ffd5b506102e7610307366004614c7d565b6109e4565b61031f61031a366004614e03565b610a40565b60405190151581526020015b60405180910390f35b34801561033f575f5ffd5b506102e761034e366004614edd565b610e5b565b34801561035e575f5ffd5b50609654610372906001600160a01b031681565b6040516001600160a01b03909116815260200161032b565b348015610395575f5ffd5b506102e76103a4366004614fa2565b61110b565b3480156103b4575f5ffd5b506102e76103c3366004614fe5565b61113d565b3480156103d3575f5ffd5b5060335442111561031f565b3480156103ea575f5ffd5b5060645460405160ff909116815260200161032b565b34801561040b575f5ffd5b506102e761041a36600461501a565b611230565b34801561042a575f5ffd5b50603254610372906001600160a01b031681565b348015610449575f5ffd5b506102e7610458366004615035565b6112a1565b61031f61046b3660046150b8565b6113a5565b34801561047b575f5ffd5b5061031f61048a36600461515a565b6115c7565b6102e761049d3660046151f0565b6118d2565b3480156104ad575f5ffd5b506102e76104bc366004615232565b6118ed565b3480156104cc575f5ffd5b506102e76104db36600461525b565b6119d0565b3480156104eb575f5ffd5b506104f4611ade565b60405190815260200161032b565b34801561050d575f5ffd5b5061052161051c366004614c7d565b611af9565b60405161032b91906152e3565b348015610539575f5ffd5b505f516020615e7c5f395f51905f525460ff1661031f565b61031f61055f366004615330565b611c7e565b34801561056f575f5ffd5b506102e761057e3660046153b8565b611ce7565b34801561058e575f5ffd5b506102e7611dbe565b3480156105a2575f5ffd5b506105b66105b1366004614c7d565b611dd1565b60405161032b9190615415565b3480156105ce575f5ffd5b5061031f6105dd36600461515a565b611dea565b3480156105ed575f5ffd5b506106bb6105fc366004615427565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152505f9182526039602090815260408084206001600160a01b039384168552825292839020835160e08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260028101546080830152600381015460a08301526004015460c082015290565b60405161032b919061544a565b3480156106d3575f5ffd5b506106dc611e85565b60405161032b9796959493929190615486565b61031f6106fd3660046155b3565b611f2e565b34801561070d575f5ffd5b50610372611fc9565b348015610721575f5ffd5b506097546001600160a01b0316610372565b34801561073e575f5ffd5b506098546104f4565b348015610752575f5ffd5b5061031f610761366004615427565b611ff7565b348015610771575f5ffd5b50610785610780366004614c7d565b61200e565b60405161032b91906156e3565b34801561079d575f5ffd5b506102e76107ac366004614c7d565b612027565b3480156107bc575f5ffd5b506107e1604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161032b9190615723565b3480156107f9575f5ffd5b506104f4610808366004614c7d565b61207b565b348015610818575f5ffd5b5061082c61082736600461515a565b612097565b60405161032b9190615735565b348015610844575f5ffd5b506102e76108533660046157c7565b612241565b348015610863575f5ffd5b506102e76108723660046157e0565b612292565b348015610882575f5ffd5b506102e7610891366004614c7d565b6122b3565b3480156108a1575f5ffd5b506105b661259b565b3480156108b5575f5ffd5b506103726108c43660046157fb565b6125ac565b3480156108d4575f5ffd5b506102e76108e336600461501a565b612667565b3480156108f3575f5ffd5b506102e7610902366004615884565b6126b8565b348015610912575f5ffd5b506104f4610921366004614c7d565b6126ef565b6102e76109343660046158cf565b61270b565b348015610944575f5ffd5b50610958610953366004614fe5565b61289e565b6040805193845260208401929092529082015260600161032b565b34801561097e575f5ffd5b506102e761098d36600461501a565b61292d565b34801561099d575f5ffd5b506102e76109ac36600461501a565b6129dd565b3480156109bc575f5ffd5b506102e76109cb366004615427565b612a17565b3480156109db575f5ffd5b506104f4612afc565b6109ec612b05565b5f818152603b60205260408120610a0290612b37565b90505f5b8151811015610a3b57610a3283838381518110610a2557610a25615968565b6020026020010151612b43565b50600101610a06565b505050565b5f610a49612d50565b610a51612d80565b610a79610a64606087016040880161501a565b86355f90815260386020526040902090612db7565b610a89606087016040880161501a565b90610ab857604051633142cb1160e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f348015610ae3576040516362624a5160e11b815260048101929092526024820152604401610aaf565b505084355f9081526037602052604081206002018054600101908190559050806020870135808214610b315760405163a6ab65c560e01b815260048101929092526024820152604401610aaf565b50610bc390507fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b77219146020880135610b6d60608a0160408b0161501a565b610b7d60808b0160608c0161501a565b60808b0135610b8f60a08d018d61597c565b604051602001610ba597969594939291906159e6565b60405160208183030381529060405280519060200120868686612dd8565b85355f90815260396020526040808220606091839182918290610beb908d8701908e0161501a565b6001600160a01b03908116825260208083019390935260409182015f20825160e08101845281548316815260018201549283169481019490945260ff600160a01b8304811615801594860194909452600160a81b9092049091161515606084015260028101546080840152600381015460a08401526004015460c0830152909150610c9c576040518060400160405280600c81526020016b1d1bdad95b881c185d5cd95960a21b8152509250610d25565b608081015115801590610cb6575089608001358160800151105b15610cf057604051806040016040528060118152602001701bdd995c881cd859995d1e481b1a5b5a5d607a1b815250925060019150610d25565b610d1f8a35610d0560608d0160408e0161501a565b610d1560808e0160608f0161501a565b8d60800135613028565b90945092505b508215610dad57610d3c60608a0160408b0161501a565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b610d7d60808e0160608f0161501a565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4610e35565b610db8898383613116565b610dc860608a0160408b0161501a565b6001600160a01b031660208a01358a357fd3122beb443f4b66faef07d45daf96faf0ffc23e050550cf1bbcbe6105f6bee7610e0960808e0160608f0161501a565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a45b6001945050505050610e5360015f516020615e9c5f395f51905f5255565b949350505050565b610e63612b05565b610e6e6035876132d6565b15610ecb57604080516080810182528781525f6020808301828152838501838152606085018481528c855260379093529490922092518355905160018301559151600282015590516003909101805460ff19169115159190911790555b6001600160a01b038416610ef257604051636ca1fdd760e01b815260040160405180910390fd5b5f868152603860205260409020610f0990856132e1565b8490610f34576040516311dd05f360e31b81526001600160a01b039091166004820152602401610aaf565b506040518060e00160405280856001600160a01b03168152602001846001600160a01b031681526020015f1515815260200186151581526020018281526020015f81526020015f81525060395f8881526020019081526020015f205f866001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160010160156101000a81548160ff0219169083151502179055506080820151816002015560a0820151816003015560c08201518160040155905050600182138061107657505f1982125b156110a0575f868152603a602090815260408083206001600160a01b038816845290915290208290555b826001600160a01b0316846001600160a01b0316877fffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a6685858a6040516110fb9392919092835260208301919091521515604082015260600190565b60405180910390a4505050505050565b61112d826111275f5f8681526020019081526020015f20612b37565b5f6126b8565b611139828260016126b8565b5050565b61114f6420a226a4a760d91b33611ff7565b6420a226a4a760d91b33909161118a5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f8381526038602052604090206111a39083612db7565b82906111ce5760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f8381526039602090815260408083206001600160a01b038616808552908352928190206002018490555183815285917fd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c891015b60405180910390a3505050565b6112426420a226a4a760d91b33611ff7565b6420a226a4a760d91b33909161127d5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b5050609680546001600160a01b0319166001600160a01b0392909216919091179055565b6112b36420a226a4a760d91b33611ff7565b6420a226a4a760d91b3390916112ee5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f8381526038602052604090206113079083612db7565b82906113325760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f8381526039602090815260408083206001600160a01b0386168085529252918290206001018054841515600160a01b0260ff60a01b19909116179055905184907f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea9061122390851515815260200190565b5f6113ae612d50565b89896113ba82826132f5565b6113c2612d80565b6113cf602085018561501a565b6001600160a01b038c81169116148b6113eb602087018761501a565b909161141d576040516313f7c32b60e31b81526001600160a01b03928316600482015291166024820152604401610aaf565b505061142c8c8c8b8b8b6133cb565b90985096508661143c898b615a47565b6114469190615a47565b60408501351015876114588a8c615a47565b6114629190615a47565b8560400135909161148f5760405163943892eb60e01b815260048101929092526024820152604401610aaf565b505f90506114a3604086016020870161501a565b30604087013560608801356114be60a08a0160808b016157c7565b6040516001600160a01b0395861660248201529490931660448501526064840191909152608483015260ff1660a482015260a086013560c482015260c086013560e48201526101040160408051601f19818403018152919052602080820180516001600160e01b031663d505accf60e01b1790529091505f908190611550906115499089018961501a565b5f856134af565b915091508181906115755760405163e8b5c4bb60e01b8152600401610aaf9190615723565b5050505061159e8c8c866020016020810190611591919061501a565b8d8d8d8d60018e8e61362c565b600192506115b860015f516020615e9c5f395f51905f5255565b50509998505050505050505050565b5f6115d0612d50565b6115d8612d80565b5f838152603b602052604090206115ef908361370e565b8290611611576040516373a1390160e11b8152600401610aaf91815260200190565b505f838152603c60209081526040808320858452909152808220815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e084015260058201805484929184916101008501919061168d90615a5a565b80601f01602080910402602001604051908101604052809291908181526020018280546116b990615a5a565b80156117045780601f106116db57610100808354040283529160200191611704565b820191905f5260205f20905b8154815290600101906020018083116116e757829003601f168201915b50505050508152505081526020016006820154815260200160078201805461172b90615a5a565b80601f016020809104026020016040519081016040528092919081815260200182805461175790615a5a565b80156117a25780601f10611779576101008083540402835291602001916117a2565b820191905f5260205f20905b81548152906001019060200180831161178557829003601f168201915b505050919092525050505f85815260396020908152604080832084518201516001600160a01b03908116855290835292819020815160e08101835281548516815260018201549485169381019390935260ff600160a01b8504811615801585850152600160a81b909504161515606084015260028101546080840152600381015460a08401526004015460c0830152835101519293509190611863576040516338384f6f60e11b81526001600160a01b039091166004820152602401610aaf565b50602082015115806118785750428260200151105b82602001514290916118a657604051637a88099160e11b815260048101929092526024820152604401610aaf565b50506118b28585613725565b925050506118cc60015f516020615e9c5f395f51905f5255565b92915050565b6118da6137f7565b6118e38261385d565b6111398282613865565b6118ff6420a226a4a760d91b33611ff7565b6420a226a4a760d91b33909161193a5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f838152603b60205260409020611953908361370e565b8290611975576040516373a1390160e11b8152600401610aaf91815260200190565b506119808142615a47565b5f848152603c602090815260408083208684528252918290206006019290925551828152839185917f530924d6189918da2c072b6eb46ef24e51151e8cb52c31d17482d8959b0d4c5b9101611223565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015611a145750825b90505f826001600160401b03166001148015611a2f5750303b155b905081158015611a3d575080155b15611a5b5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611a8557845460ff60401b1916600160401b1785555b611a8f8787613921565b8315611ad557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b5f611ae76139a9565b505f516020615e5c5f395f51905f5290565b5f81815260386020526040812060609190611b1390612b37565b90505f81516001600160401b03811115611b2f57611b2f614c94565b604051908082528060200260200182016040528015611b9457816020015b6040805160e0810182525f8082526020808301829052928201819052606082018190526080820181905260a0820181905260c082015282525f19909201910181611b4d5790505b5090505f5b8251811015611c765760395f8681526020019081526020015f205f848381518110611bc657611bc6615968565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160e08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b9093049091161515606082015260028201546080820152600382015460a082015260049091015460c08201528251839083908110611c6357611c63615968565b6020908102919091010152600101611b99565b509392505050565b5f611c87612d50565b8888611c9382826132f5565b611c9b612d80565b611ca88b8b8a8a8a6133cb565b9097509550611cbf8b8b338c8c8c8c5f8d8d61362c565b60019250611cd960015f516020615e9c5f395f51905f5255565b505098975050505050505050565b611cf96420a226a4a760d91b33611ff7565b6420a226a4a760d91b339091611d345760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50611d42905060358361370e565b8290611d645760405163ac7dbbfd60e01b8152600401610aaf91815260200190565b505f82815260376020908152604091829020600301805460ff1916841515908117909155915191825283917f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675910160405180910390a25050565b611dc6612b05565b611dcf5f6139f2565b565b5f818152603b602052604090206060906118cc90612b37565b5f611dfd6420a226a4a760d91b33611ff7565b6420a226a4a760d91b339091611e385760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f838152603b60205260409020611e51908361370e565b8290611e73576040516373a1390160e11b8152600401610aaf91815260200190565b50611e7e8383613725565b9392505050565b5f60608082808083815f516020615e3c5f395f51905f528054909150158015611eb057506001810154155b611ef45760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610aaf565b611efc613a62565b611f04613b22565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f805b85811015611fbc57611fb3878783818110611f4e57611f4e615968565b9050602002810190611f609190615a92565b868381518110611f7257611f72615968565b6020026020010151868481518110611f8c57611f8c615968565b6020026020010151868581518110611fa657611fa6615968565b6020026020010151610a40565b50600101611f31565b5060019695505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b5f828152602081905260408120611e7e9083612db7565b5f8181526020819052604090206060906118cc90612b37565b6120396420a226a4a760d91b33611ff7565b6420a226a4a760d91b3390916120745760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b5050603455565b5f8181526037602052604081206001908101546118cc91615a47565b61209f614bca565b5f838152603c6020908152604080832085845290915290819020815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e084015260058201805484929184916101008501919061211b90615a5a565b80601f016020809104026020016040519081016040528092919081815260200182805461214790615a5a565b80156121925780601f1061216957610100808354040283529160200191612192565b820191905f5260205f20905b81548152906001019060200180831161217557829003601f168201915b5050505050815250508152602001600682015481526020016007820180546121b990615a5a565b80601f01602080910402602001604051908101604052809291908181526020018280546121e590615a5a565b80156122305780601f1061220757610100808354040283529160200191612230565b820191905f5260205f20905b81548152906001019060200180831161221357829003601f168201915b505050505081525050905092915050565b612249612b05565b6064805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b61229a612b05565b80156122ab576122a8613b60565b50565b6122a8613bbc565b4260335410156122c05750565b5f806122cc6035612b37565b90505f5b8151811015612595575f8282815181106122ec576122ec615968565b6020908102919091018101515f818152603790925260409091206003015490915060ff161561231b575061258d565b5f818152603b6020526040812061233190612b37565b90505f5b8151811015612589575f82828151811061235157612351615968565b6020908102919091018101515f868152603c83526040808220838352909352828120835161012081019094528054606085019081526001820154608086015260028201546001600160a01b0390811660a087015260038301541660c0860152600482015460e08601526005820180549496509294939192849284916101008501916123db90615a5a565b80601f016020809104026020016040519081016040528092919081815260200182805461240790615a5a565b80156124525780601f1061242957610100808354040283529160200191612452565b820191905f5260205f20905b81548152906001019060200180831161243557829003601f168201915b50505050508152505081526020016006820154815260200160078201805461247990615a5a565b80601f01602080910402602001604051908101604052809291908181526020018280546124a590615a5a565b80156124f05780601f106124c7576101008083540402835291602001916124f0565b820191905f5260205f20905b8154815290600101906020018083116124d357829003601f168201915b5050509190925250508151515f90815260396020908152604080832085518201516001600160a01b0316845290915290206001015491925050600160a01b900460ff161580156125435750602081015115155b80156125525750428160200151105b1561257f576125618583613725565b508861256c89615ab0565b9850881061257f57505050505050505050565b5050600101612335565b5050505b6001016122d0565b50505050565b60606125a76035612b37565b905090565b6032545f906001600160a01b03166125d7576040516315aeca0d60e11b815260040160405180910390fd5b603254604051637c469ea160e11b81526001600160a01b039091169063f88d3d429061260d908a908a9088908890600401615ac8565b6020604051808303815f875af1158015612629573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061264d9190615b05565b905061265d875f83898989610e5b565b9695505050505050565b61266f612b05565b6001600160a01b03811661269657604051638219abe360e01b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b0392909216919091179055565b5f5b8251811015612595576126e7848483815181106126d9576126d9615968565b602002602001015184613c01565b6001016126ba565b5f818152603760205260408120600201546118cc906001615a47565b82811461272b576040516329f517fd60e01b815260040160405180910390fd5b5f5b838110156128975761288e85858381811061274a5761274a615968565b905060200281019061275c9190615b20565b3586868481811061276f5761276f615968565b90506020028101906127819190615b20565b61279290604081019060200161501a565b8787858181106127a4576127a4615968565b90506020028101906127b69190615b20565b6127c790606081019060400161501a565b8888868181106127d9576127d9615968565b90506020028101906127eb9190615b20565b6060013589898781811061280157612801615968565b90506020028101906128139190615b20565b608001358a8a8881811061282957612829615968565b905060200281019061283b9190615b20565b60a001358b8b8981811061285157612851615968565b90506020028101906128639190615b20565b6128719060c081019061597c565b8b8b8b81811061288357612883615968565b905060e002016113a5565b5060010161272d565b5050505050565b6096546040516337dba1f760e21b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063df6e87dc90606401606060405180830381865afa1580156128fa573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061291e9190615b34565b91989097509095509350505050565b612935612b05565b6032546001600160a01b0316801561296c5760405163f6c75f8f60e01b81526001600160a01b039091166004820152602401610aaf565b506001600160a01b03811661299457604051636ca1fdd760e01b815260040160405180910390fd5b603280546001600160a01b0319166001600160a01b0383169081179091556040517fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee905f90a250565b6129e5612b05565b6001600160a01b038116612a0e57604051631e4fbdf760e01b81525f6004820152602401610aaf565b6122a8816139f2565b612a1f612b05565b5f828152603860205260409020612a369082613d30565b8190612a615760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f8281526039602090815260408083206001600160a01b03851680855290835281842080546001600160a01b03191681556001810180546001600160b01b03191690556002810185905560038101859055600401849055858452603a835281842081855290925280832083905551909184917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d9190a35050565b5f6125a7613d44565b33612b0e611fc9565b6001600160a01b031614611dcf5760405163118cdaa760e01b8152336004820152602401610aaf565b60605f611e7e83613d4d565b612b4b614bf0565b5f838152603b60205260409020612b629083613da6565b8290612b8457604051637f11bea960e01b8152600401610aaf91815260200190565b505f838152603c60209081526040808320858452825291829020825160c0810184528154815260018201549281019290925260028101546001600160a01b0390811693830193909352600381015490921660608201526004820154608082015260058201805491929160a084019190612bfc90615a5a565b80601f0160208091040260200160405190810160405280929190818152602001828054612c2890615a5a565b8015612c735780601f10612c4a57610100808354040283529160200191612c73565b820191905f5260205f20905b815481529060010190602001808311612c5657829003601f168201915b505050919092525050505f848152603960209081526040808320848201516001600160a01b031684529091529020600181015491925090600160a81b900460ff1615612cd6578160800151816004015f828254612cd09190615b5f565b90915550505b5f848152603c602090815260408083208684529091528120818155600181018290556002810180546001600160a01b0319908116909155600382018054909116905560048101829055908181612d2f6005830182614c33565b5050600682015f9055600782015f612d479190614c33565b50505092915050565b5f516020615e7c5f395f51905f525460ff1615611dcf5760405163d93c066560e01b815260040160405180910390fd5b5f516020615e9c5f395f51905f52805460011901612db157604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600160a01b0381165f9081526001830160205260408120541515611e7e565b8251825181148015612dea5750815181145b835183518392612e1e576040516324227ae160e21b8152600481019390935260248301919091526044820152606401610aaf565b505060645482915060ff16811015612e4c57604051631aebd17960e11b8152600401610aaf91815260200190565b505f80826001600160401b03811115612e6757612e67614c94565b604051908082528060200260200182016040528015612e90578160200160208202803683370190505b5090505f5b83811015612ff2575f612f00888381518110612eb357612eb3615968565b6020026020010151888481518110612ecd57612ecd615968565b6020026020010151888581518110612ee757612ee7615968565b6020026020010151612ef88d613db1565b929190613ddd565b9050612f18682b20a624a220aa27a960b91b82611ff7565b682b20a624a220aa27a960b91b829091612f575760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b505f9050805b85811015612fa857848181518110612f7757612f77615968565b60200260200101516001600160a01b0316836001600160a01b031603612fa05760019150612fa8565b600101612f5d565b5080612fe85781848681518110612fc157612fc1615968565b60200260200101906001600160a01b031690816001600160a01b0316815250508460010194505b5050600101612e95565b50606454829060ff1681101561301e57604051631aebd17960e11b8152600401610aaf91815260200190565b5050505050505050565b5f848152603a602090815260408083206001600160a01b03871684529091528120546060906001811315613067576130608185615b72565b9350613086565b5f198112156130865761307981615b89565b6130839085615bb7565b93505b5f196001600160a01b038716016130ba576130b0858560405180602001604052805f8152506134af565b925092505061310d565b831561310b575f8781526039602090815260408083206001600160a01b038a168452909152902060010154600160a81b900460ff1615613100576130b087878787613e09565b6130b0868686613eff565b505b94509492505050565b82355f908152603b60209081526040909120613134918501356132d6565b83602001359061315a576040516307a5c91d60e31b8152600401610aaf91815260200190565b505f81156131835760345461316f9042615a47565b90506034544261317f9190615a47565b6033555b60405180606001604052808561319890615bca565b81526020808201849052604091820186905286355f908152603c82528281208883013582528252829020835180518255918201516001820155918101516002830180546001600160a01b039283166001600160a01b03199182161790915560608301516003850180549190931691161790556080810151600483015560a08101518290600582019061322a9082615c94565b505050602082015160068201556040820151600782019061324b9082615c94565b50505083355f90815260396020526040808220908290613271906060890190890161501a565b6001600160a01b0316815260208101919091526040015f206001810154909150600160a81b900460ff1615612897578460800135816004015f8282546132b79190615a47565b90915550505050505050565b60015f516020615e9c5f395f51905f5255565b5f611e7e8383613fe0565b5f611e7e836001600160a01b038416613fe0565b5f82815260386020526040902061330c9082612db7565b81906133375760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f82815260376020526040902060030154829060ff161561336f57604051636fc24b7960e11b8152600401610aaf91815260200190565b505f8281526039602090815260408083206001600160a01b03851684529091529020600101548190600160a01b900460ff1615610a3b576040516338384f6f60e11b81526001600160a01b039091166004820152602401610aaf565b6096545f9081906001600160a01b03166133e957505f9050806134a5565b6096546040516337dba1f760e21b8152600481018990526001600160a01b038881166024830152604482018890525f92169063df6e87dc90606401606060405180830381865afa15801561343f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134639190615b34565b9094509250905080861080159061347a5750828510155b80156134865750818410155b6134a3576040516358e8878560e01b815260040160405180910390fd5b505b9550959350505050565b5f60608347478211156134de57604051632f2a246160e21b815260048101929092526024820152604401610aaf565b50506060856001600160a01b031685856040516134fb9190615d4e565b5f6040518083038185875af1925050503d805f8114613535576040519150601f19603f3d011682016040523d82523d5f602084013e61353a565b606091505b5090935090508261357f57805115613556575f92509050613624565b50506040805180820190915260088152671c995d995c9d195960c21b60208201525f9150613624565b845f0361360f5780515f036135cd57856001600160a01b03163b5f036135c85750506040805180820190915260088152676e6f7420636f646560c01b60208201525f9150613624565b61360f565b602081015160018115151461360d575f6040518060400160405280600c81526020016b72657475726e2066616c736560a01b815250935093505050613624565b505b505060408051602081019091525f8152600191505b935093915050565b6136428a8a8a8961363d898b615a47565b61402c565b5f8a815260376020526040812060019081018054909101908190558190915060395f8d81526020019081526020015f205f8c6001600160a01b03166001600160a01b031681526020019081526020015f206001015f9054906101000a90046001600160a01b03169050896001600160a01b0316828d7f732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db508e858e8e8e8e428f8f8f6040516136f89a99989796959493929190615d64565b60405180910390a4505050505050505050505050565b5f8181526001830160205260408120541515611e7e565b5f5f6137318484612b43565b90505f5f613750835f0151846040015185606001518660800151613028565b915091508181906137745760405162461bcd60e51b8152600401610aaf9190615723565b5082604001516001600160a01b03168360200151845f01517f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b86606001518760800151426040516137e3939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a450600195945050505050565b3061c0de625c0eb760891b01148061383f575061c0de625c0eb760891b016138335f516020615e5c5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15611dcf5760405163703e46dd60e11b815260040160405180910390fd5b6122a8612b05565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156138bf575060408051601f3d908101601f191682019092526138bc91810190615dce565b60015b6138e757604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610aaf565b5f516020615e5c5f395f51905f52811461391757604051632a87526960e21b815260048101829052602401610aaf565b610a3b83836142ab565b613929614300565b61393233614349565b61393a61435a565b613942614362565b61394a614372565b61395382614382565b61395b6143e8565b6001600160a01b03811661398257604051638219abe360e01b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b03929092169190911790555043609855565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611dcf5760405163703e46dd60e11b815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020615e3c5f395f51905f5291613aa090615a5a565b80601f0160208091040260200160405190810160405280929190818152602001828054613acc90615a5a565b8015613b175780601f10613aee57610100808354040283529160200191613b17565b820191905f5260205f20905b815481529060010190602001808311613afa57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020615e3c5f395f51905f5291613aa090615a5a565b613b68612d50565b5f516020615e7c5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b039091168152602001612287565b613bc4614414565b5f516020615e7c5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613ba4565b613c09612b05565b6001600160a01b038216613c3057604051635989b9d360e01b815260040160405180910390fd5b82613c4e57604051630e1b39f960e31b815260040160405180910390fd5b8015613ca5575f838152602081905260409020613c6b90836132e1565b83839091613c9e57604051631b753c1b60e21b815260048101929092526001600160a01b03166024820152604401610aaf565b5050613cf2565b5f838152602081905260409020613cbc9083613d30565b83839091613cef57604051633a401ef360e21b815260048101929092526001600160a01b03166024820152604401610aaf565b50505b82826001600160a01b03167f8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f383604051611223911515815260200190565b5f611e7e836001600160a01b038416614443565b5f6125a7614526565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613d9a57602002820191905f5260205f20905b815481526020019060010190808311613d86575b50505050509050919050565b5f611e7e8383614443565b5f6118cc613dbd613d44565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613ded88888888614599565b925092509250613dfd8282614661565b50909695505050505050565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613e78575060408051601f3d908101601f19168201909252613e7591810190615de5565b60015b613eb2573d808015613ea5576040519150601f19603f3d011682016040523d82523d5f602084013e613eaa565b606091505b50905061310d565b8092508015613ecb57613ec6878786614719565b61310b565b6040518060400160405280600f81526020016e1d1c985b9cd9995c8819985a5b1959608a1b81525091505094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015613f6e575060408051601f3d908101601f19168201909252613f6b91810190615de5565b60015b613fa8573d808015613f9b576040519150601f19603f3d011682016040523d82523d5f602084013e613fa0565b606091505b509050613624565b80925080613fd7576040518060400160405280600b81526020016a1b5a5b9d0819985a5b195960aa1b81525091505b50935093915050565b5f81815260018301602052604081205461402557508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556118cc565b505f6118cc565b5f858152603a602090815260408083206001600160a01b03881684529091529020546001811315614097576140618184615e00565b8590849015614094576040516305330e5560e11b81526001600160a01b0390921660048301526024820152604401610aaf565b50505b5f196001600160a01b03861601614147576140b28284615a47565b34146140be8385615a47565b3490916140e7576040516362624a5160e11b815260048101929092526024820152604401610aaf565b505081156141425760975460408051602081019091525f808252918291614119916001600160a01b03169086906134af565b9150915081819061413e5760405163e8b5c4bb60e01b8152600401610aaf9190615723565b5050505b6142a3565b5f348015614171576040516362624a5160e11b815260048101929092526024820152604401610aaf565b50614195905084306141838587615a47565b6001600160a01b0389169291906147c4565b81156141b5576097546141b5906001600160a01b0387811691168461482b565b5f8681526039602090815260408083206001600160a01b0389168452909152902060010154600160a81b900460ff16156141f45761414286868561485c565b604051632770a7eb60e21b8152306004820152602481018490526001600160a01b03861690639dc29fac906044016020604051808303815f875af115801561423e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142629190615de5565b85858590919261429f57604051639ac2f56d60e01b81526001600160a01b0393841660048201529290911660248301526044820152606401610aaf565b5050505b505050505050565b6142b48261489a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156142f857610a3b82826148fd565b61113961496f565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16611dcf57604051631afcd79f60e31b815260040160405180910390fd5b614351614300565b6122a88161498e565b611dcf614300565b61436a614300565b611dcf614996565b61437a614300565b611dcf6149b6565b61438a614300565b6143d2604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b8152506149be565b6064805460ff191660ff92909216919091179055565b6143f0614300565b61440b6420a226a4a760d91b614404611fc9565b6001613c01565b62015180603455565b5f516020615e7c5f395f51905f525460ff16611dcf57604051638dfc202b60e01b815260040160405180910390fd5b5f818152600183016020526040812054801561451d575f614465600183615b5f565b85549091505f9061447890600190615b5f565b90508082146144d7575f865f01828154811061449657614496615968565b905f5260205f200154905080875f0184815481106144b6576144b6615968565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806144e8576144e8615e13565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506118cc565b5f9150506118cc565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6145506149d0565b614558614a38565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156145d257505f91506003905082614657565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614623573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661464e57505f925060019150829050614657565b92505f91508190505b9450945094915050565b5f82600381111561467457614674615e27565b0361467d575050565b600182600381111561469157614691615e27565b036146af5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156146c3576146c3615e27565b036146e45760405163fce698f760e01b815260048101829052602401610aaf565b60038260038111156146f8576146f8615e27565b03611139576040516335e2f38360e21b815260048101829052602401610aaf565b5f8381526039602090815260408083206001600160a01b0386168452909152902060048101546147499083615a47565b6003820154600483015486928692869290918210156147a1576040516352c2db3360e01b815260048101959095526001600160a01b03909316602485015260448401919091526064830152608482015260a401610aaf565b505050505081816003015f8282546147b99190615b5f565b909155505050505050565b6040516001600160a01b0384811660248301528381166044830152606482018390526125959186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614a7a565b6040516001600160a01b03838116602483015260448201839052610a3b91859182169063a9059cbb906064016147f9565b5f8381526039602090815260408083206001600160a01b038616845290915281206003018054839290614890908490615a47565b9091555050505050565b806001600160a01b03163b5f036148cf57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610aaf565b5f516020615e5c5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516149199190615d4e565b5f60405180830381855af49150503d805f8114614951576040519150601f19603f3d011682016040523d82523d5f602084013e614956565b606091505b5091509150614966858383614ae6565b95945050505050565b3415611dcf5760405163b398979f60e01b815260040160405180910390fd5b6129e5614300565b61499e614300565b5f516020615e7c5f395f51905f52805460ff19169055565b6132c3614300565b6149c6614300565b6111398282614b42565b5f5f516020615e3c5f395f51905f52816149e8613a62565b805190915015614a0057805160209091012092915050565b81548015614a0f579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020615e3c5f395f51905f5281614a50613b22565b805190915015614a6857805160209091012092915050565b60018201548015614a0f579392505050565b5f5f60205f8451602086015f885af180614a99576040513d5f823e3d81fd5b50505f513d91508115614ab0578060011415614abd565b6001600160a01b0384163b155b1561259557604051635274afe760e01b81526001600160a01b0385166004820152602401610aaf565b606082614afb57614af682614ba1565b611e7e565b8151158015614b1257506001600160a01b0384163b155b15614b3b57604051639996b31560e01b81526001600160a01b0385166004820152602401610aaf565b5092915050565b614b4a614300565b5f516020615e3c5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102614b838482615c94565b5060038101614b928382615c94565b505f8082556001909101555050565b805115614bb15780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6040518060600160405280614bdd614bf0565b81526020015f8152602001606081525090565b6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b508054614c3f90615a5a565b5f825580601f10614c4e575050565b601f0160209004905f5260205f20908101906122a891905b80821115614c79575f8155600101614c66565b5090565b5f60208284031215614c8d575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b0381118282101715614cca57614cca614c94565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614cf857614cf8614c94565b604052919050565b5f6001600160401b03821115614d1857614d18614c94565b5060051b60200190565b803560ff81168114614d32575f5ffd5b919050565b5f82601f830112614d46575f5ffd5b8135614d59614d5482614d00565b614cd0565b8082825260208201915060208360051b860101925085831115614d7a575f5ffd5b602085015b83811015614d9e57614d9081614d22565b835260209283019201614d7f565b5095945050505050565b5f82601f830112614db7575f5ffd5b8135614dc5614d5482614d00565b8082825260208201915060208360051b860101925085831115614de6575f5ffd5b602085015b83811015614d9e578035835260209283019201614deb565b5f5f5f5f60808587031215614e16575f5ffd5b84356001600160401b03811115614e2b575f5ffd5b850160c08188031215614e3c575f5ffd5b935060208501356001600160401b03811115614e56575f5ffd5b614e6287828801614d37565b93505060408501356001600160401b03811115614e7d575f5ffd5b614e8987828801614da8565b92505060608501356001600160401b03811115614ea4575f5ffd5b614eb087828801614da8565b91505092959194509250565b80151581146122a8575f5ffd5b6001600160a01b03811681146122a8575f5ffd5b5f5f5f5f5f5f60c08789031215614ef2575f5ffd5b863595506020870135614f0481614ebc565b94506040870135614f1481614ec9565b93506060870135614f2481614ec9565b9598949750929560808101359460a0909101359350915050565b5f82601f830112614f4d575f5ffd5b8135614f5b614d5482614d00565b8082825260208201915060208360051b860101925085831115614f7c575f5ffd5b602085015b83811015614d9e578035614f9481614ec9565b835260209283019201614f81565b5f5f60408385031215614fb3575f5ffd5b8235915060208301356001600160401b03811115614fcf575f5ffd5b614fdb85828601614f3e565b9150509250929050565b5f5f5f60608486031215614ff7575f5ffd5b83359250602084013561500981614ec9565b929592945050506040919091013590565b5f6020828403121561502a575f5ffd5b8135611e7e81614ec9565b5f5f5f60608486031215615047575f5ffd5b83359250602084013561505981614ec9565b9150604084013561506981614ebc565b809150509250925092565b5f5f83601f840112615084575f5ffd5b5081356001600160401b0381111561509a575f5ffd5b6020830191508360208285010111156150b1575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f898b036101c08112156150d2575f5ffd5b8a35995060208b01356150e481614ec9565b985060408b01356150f481614ec9565b975060608b0135965060808b0135955060a08b0135945060c08b01356001600160401b03811115615123575f5ffd5b61512f8d828e01615074565b90955093505060e060df1982011215615146575f5ffd5b5060e08a0190509295985092959850929598565b5f5f6040838503121561516b575f5ffd5b50508035926020909101359150565b5f5f6001600160401b0384111561519357615193614c94565b50601f8301601f19166020016151a881614cd0565b9150508281528383830111156151bc575f5ffd5b828260208301375f602084830101529392505050565b5f82601f8301126151e1575f5ffd5b611e7e8383356020850161517a565b5f5f60408385031215615201575f5ffd5b823561520c81614ec9565b915060208301356001600160401b03811115615226575f5ffd5b614fdb858286016151d2565b5f5f5f60608486031215615244575f5ffd5b505081359360208301359350604090920135919050565b5f5f6040838503121561526c575f5ffd5b61527583614d22565b9150602083013561528581614ec9565b809150509250929050565b80516001600160a01b039081168352602080830151909116908301526040808201511515908301526060808201511515908301526080808201519083015260a0818101519083015260c090810151910152565b602080825282518282018190525f918401906040840190835b818110156153255761530f838551615290565b6020939093019260e092909201916001016152fc565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b031215615347575f5ffd5b88359750602089013561535981614ec9565b9650604089013561536981614ec9565b9550606089013594506080890135935060a0890135925060c08901356001600160401b03811115615398575f5ffd5b6153a48b828c01615074565b999c989b5096995094979396929594505050565b5f5f604083850312156153c9575f5ffd5b82359150602083013561528581614ebc565b5f8151808452602084019350602083015f5b8281101561540b5781518652602095860195909101906001016153ed565b5093949350505050565b602081525f611e7e60208301846153db565b5f5f60408385031215615438575f5ffd5b82359150602083013561528581614ec9565b60e081016118cc8284615290565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff60f81b8816815260e060208201525f6154a460e0830189615458565b82810360408401526154b68189615458565b606084018890526001600160a01b038716608085015260a0840186905283810360c085015290506154e781856153db565b9a9950505050505050505050565b5f5f83601f840112615505575f5ffd5b5081356001600160401b0381111561551b575f5ffd5b6020830191508360208260051b85010111156150b1575f5ffd5b5f82601f830112615544575f5ffd5b8135615552614d5482614d00565b8082825260208201915060208360051b860101925085831115615573575f5ffd5b602085015b83811015614d9e5780356001600160401b03811115615595575f5ffd5b6155a4886020838a0101614da8565b84525060209283019201615578565b5f5f5f5f5f608086880312156155c7575f5ffd5b85356001600160401b038111156155dc575f5ffd5b6155e8888289016154f5565b90965094505060208601356001600160401b03811115615606575f5ffd5b8601601f81018813615616575f5ffd5b8035615624614d5482614d00565b8082825260208201915060208360051b85010192508a831115615645575f5ffd5b602084015b838110156156855780356001600160401b03811115615667575f5ffd5b6156768d602083890101614d37565b8452506020928301920161564a565b50955050505060408601356001600160401b038111156156a3575f5ffd5b6156af88828901615535565b92505060608601356001600160401b038111156156ca575f5ffd5b6156d688828901615535565b9150509295509295909350565b602080825282518282018190525f918401906040840190835b818110156153255783516001600160a01b03168352602093840193909201916001016156fc565b602081525f611e7e6020830184615458565b602081525f82516060602084015280516080840152602081015160a084015260018060a01b0360408201511660c084015260018060a01b0360608201511660e0840152608081015161010084015260a0810151905060c06101208401526157a0610140840182615458565b9050602084015160408401526040840151601f198483030160608501526149668282615458565b5f602082840312156157d7575f5ffd5b611e7e82614d22565b5f602082840312156157f0575f5ffd5b8135611e7e81614ebc565b5f5f5f5f5f5f60c08789031215615810575f5ffd5b86359550602087013561582281614ec9565b9450604087013593506060870135925060808701356001600160401b0381111561584a575f5ffd5b8701601f8101891361585a575f5ffd5b6158698982356020840161517a565b92505061587860a08801614d22565b90509295509295509295565b5f5f5f60608486031215615896575f5ffd5b8335925060208401356001600160401b038111156158b2575f5ffd5b6158be86828701614f3e565b925050604084013561506981614ebc565b5f5f5f5f604085870312156158e2575f5ffd5b84356001600160401b038111156158f7575f5ffd5b615903878288016154f5565b90955093505060208501356001600160401b03811115615921575f5ffd5b8501601f81018713615931575f5ffd5b80356001600160401b03811115615946575f5ffd5b87602060e08302840101111561595a575f5ffd5b949793965060200194505050565b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112615991575f5ffd5b8301803591506001600160401b038211156159aa575f5ffd5b6020019150368190038213156150b1575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90615a2690830184866159be565b9998505050505050505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156118cc576118cc615a33565b600181811c90821680615a6e57607f821691505b602082108103615a8c57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f823560be19833603018112615aa6575f5ffd5b9190910192915050565b5f60018201615ac157615ac1615a33565b5060010190565b8481526001600160a01b03841660208201526080604082018190525f90615af190830185615458565b905060ff8316606083015295945050505050565b5f60208284031215615b15575f5ffd5b8151611e7e81614ec9565b5f823560de19833603018112615aa6575f5ffd5b5f5f5f60608486031215615b46575f5ffd5b5050815160208301516040909301519094929350919050565b818103818111156118cc576118cc615a33565b80820281158282048414176118cc576118cc615a33565b5f600160ff1b8201615b9d57615b9d615a33565b505f0390565b634e487b7160e01b5f52601260045260245ffd5b5f82615bc557615bc5615ba3565b500490565b5f60c08236031215615bda575f5ffd5b615be2614ca8565b82358152602080840135908201526040830135615bfe81614ec9565b60408201526060830135615c1181614ec9565b60608201526080838101359082015260a08301356001600160401b03811115615c38575f5ffd5b615c44368286016151d2565b60a08301525092915050565b601f821115610a3b57805f5260205f20601f840160051c81016020851015615c755750805b601f840160051c820191505b81811015612897575f8155600101615c81565b81516001600160401b03811115615cad57615cad614c94565b615cc181615cbb8454615a5a565b84615c50565b6020601f821160018114615cf3575f8315615cdc5750848201515b5f19600385901b1c1916600184901b178455612897565b5f84815260208120601f198516915b82811015615d225787850151825560209485019460019092019101615d02565b5084821015615d3f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f82518060208501845e5f920191825250919050565b6001600160a01b038b811682528a8116602083015289166040820152606081018890526080810187905260a0810186905260c0810185905283151560e082015261012061010082018190525f90615dbe90830184866159be565b9c9b505050505050505050505050565b5f60208284031215615dde575f5ffd5b5051919050565b5f60208284031215615df5575f5ffd5b8151611e7e81614ebc565b5f82615e0e57615e0e615ba3565b500690565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212204a03319c766578113fe68a534ff2048534893f19cb929138828caf9809e0476564736f6c634300081c0033",
}

// CrossBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossBridgeMetaData.ABI instead.
var CrossBridgeABI = CrossBridgeMetaData.ABI

// CrossBridgeBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossBridgeBinRuntime = "6080604052600436106102c2575f3560e01c806384b0196e1161016f578063bedb86fb116100d8578063d5717fc511610092578063e2af6cd71161006d578063e2af6cd714610973578063f2fde38b14610992578063f4509643146109b1578063f698da25146109d0575f5ffd5b8063d5717fc514610907578063d605665b14610926578063df6e87dc14610939575f5ffd5b8063bedb86fb14610858578063c1cb05dd14610877578063cf56118e14610896578063d016d625146108aa578063d477f05f146108c9578063d4bf502a146108e8575f5ffd5b8063a3246ad311610129578063a3246ad314610766578063aa1bd0bc14610792578063ad3cb1cc146107b1578063ae6893f8146107ee578063b33eb36e1461080d578063b7f3358d14610839575f5ffd5b806384b0196e146106c857806388d67d6d146106ef5780638da5cb5b1461070257806391cca3db1461071657806391cf6d3e1461073357806391d1485414610747575f5ffd5b80634ee078ba1161022b5780635c975abb116101e5578063715018a6116101c0578063715018a61461058357806379214874146105975780637f4ab9f5146105c3578063814914b5146105e2575f5ffd5b80635c975abb1461052e5780635fd262de146105515780636160751f14610564575f5ffd5b80634ee078ba146104705780634f1ef2861461048f5780634f86d44b146104a25780635187599d146104c157806352d1902d146104e05780635b605f5c14610502575f5ffd5b80633d507c5e1161027c5780633d507c5e146103c857806342cde4e8146103df578063472d35b9146104005780634be13f831461041f5780634d3f0da91461043e5780634d5d00561461045d575f5ffd5b80630b43c02c146102ed5780631938e0f21461030c5780631e7bf2151461033457806325bd5666146103535780632d87b7ee1461038a5780633b96cf5a146103a9575f5ffd5b366102e957345f036102e7576040516304a4cdd960e51b815260040160405180910390fd5b005b5f5ffd5b3480156102f8575f5ffd5b506102e7610307366004614c7d565b6109e4565b61031f61031a366004614e03565b610a40565b60405190151581526020015b60405180910390f35b34801561033f575f5ffd5b506102e761034e366004614edd565b610e5b565b34801561035e575f5ffd5b50609654610372906001600160a01b031681565b6040516001600160a01b03909116815260200161032b565b348015610395575f5ffd5b506102e76103a4366004614fa2565b61110b565b3480156103b4575f5ffd5b506102e76103c3366004614fe5565b61113d565b3480156103d3575f5ffd5b5060335442111561031f565b3480156103ea575f5ffd5b5060645460405160ff909116815260200161032b565b34801561040b575f5ffd5b506102e761041a36600461501a565b611230565b34801561042a575f5ffd5b50603254610372906001600160a01b031681565b348015610449575f5ffd5b506102e7610458366004615035565b6112a1565b61031f61046b3660046150b8565b6113a5565b34801561047b575f5ffd5b5061031f61048a36600461515a565b6115c7565b6102e761049d3660046151f0565b6118d2565b3480156104ad575f5ffd5b506102e76104bc366004615232565b6118ed565b3480156104cc575f5ffd5b506102e76104db36600461525b565b6119d0565b3480156104eb575f5ffd5b506104f4611ade565b60405190815260200161032b565b34801561050d575f5ffd5b5061052161051c366004614c7d565b611af9565b60405161032b91906152e3565b348015610539575f5ffd5b505f516020615e7c5f395f51905f525460ff1661031f565b61031f61055f366004615330565b611c7e565b34801561056f575f5ffd5b506102e761057e3660046153b8565b611ce7565b34801561058e575f5ffd5b506102e7611dbe565b3480156105a2575f5ffd5b506105b66105b1366004614c7d565b611dd1565b60405161032b9190615415565b3480156105ce575f5ffd5b5061031f6105dd36600461515a565b611dea565b3480156105ed575f5ffd5b506106bb6105fc366004615427565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152505f9182526039602090815260408084206001600160a01b039384168552825292839020835160e08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260028101546080830152600381015460a08301526004015460c082015290565b60405161032b919061544a565b3480156106d3575f5ffd5b506106dc611e85565b60405161032b9796959493929190615486565b61031f6106fd3660046155b3565b611f2e565b34801561070d575f5ffd5b50610372611fc9565b348015610721575f5ffd5b506097546001600160a01b0316610372565b34801561073e575f5ffd5b506098546104f4565b348015610752575f5ffd5b5061031f610761366004615427565b611ff7565b348015610771575f5ffd5b50610785610780366004614c7d565b61200e565b60405161032b91906156e3565b34801561079d575f5ffd5b506102e76107ac366004614c7d565b612027565b3480156107bc575f5ffd5b506107e1604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161032b9190615723565b3480156107f9575f5ffd5b506104f4610808366004614c7d565b61207b565b348015610818575f5ffd5b5061082c61082736600461515a565b612097565b60405161032b9190615735565b348015610844575f5ffd5b506102e76108533660046157c7565b612241565b348015610863575f5ffd5b506102e76108723660046157e0565b612292565b348015610882575f5ffd5b506102e7610891366004614c7d565b6122b3565b3480156108a1575f5ffd5b506105b661259b565b3480156108b5575f5ffd5b506103726108c43660046157fb565b6125ac565b3480156108d4575f5ffd5b506102e76108e336600461501a565b612667565b3480156108f3575f5ffd5b506102e7610902366004615884565b6126b8565b348015610912575f5ffd5b506104f4610921366004614c7d565b6126ef565b6102e76109343660046158cf565b61270b565b348015610944575f5ffd5b50610958610953366004614fe5565b61289e565b6040805193845260208401929092529082015260600161032b565b34801561097e575f5ffd5b506102e761098d36600461501a565b61292d565b34801561099d575f5ffd5b506102e76109ac36600461501a565b6129dd565b3480156109bc575f5ffd5b506102e76109cb366004615427565b612a17565b3480156109db575f5ffd5b506104f4612afc565b6109ec612b05565b5f818152603b60205260408120610a0290612b37565b90505f5b8151811015610a3b57610a3283838381518110610a2557610a25615968565b6020026020010151612b43565b50600101610a06565b505050565b5f610a49612d50565b610a51612d80565b610a79610a64606087016040880161501a565b86355f90815260386020526040902090612db7565b610a89606087016040880161501a565b90610ab857604051633142cb1160e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f348015610ae3576040516362624a5160e11b815260048101929092526024820152604401610aaf565b505084355f9081526037602052604081206002018054600101908190559050806020870135808214610b315760405163a6ab65c560e01b815260048101929092526024820152604401610aaf565b50610bc390507fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b77219146020880135610b6d60608a0160408b0161501a565b610b7d60808b0160608c0161501a565b60808b0135610b8f60a08d018d61597c565b604051602001610ba597969594939291906159e6565b60405160208183030381529060405280519060200120868686612dd8565b85355f90815260396020526040808220606091839182918290610beb908d8701908e0161501a565b6001600160a01b03908116825260208083019390935260409182015f20825160e08101845281548316815260018201549283169481019490945260ff600160a01b8304811615801594860194909452600160a81b9092049091161515606084015260028101546080840152600381015460a08401526004015460c0830152909150610c9c576040518060400160405280600c81526020016b1d1bdad95b881c185d5cd95960a21b8152509250610d25565b608081015115801590610cb6575089608001358160800151105b15610cf057604051806040016040528060118152602001701bdd995c881cd859995d1e481b1a5b5a5d607a1b815250925060019150610d25565b610d1f8a35610d0560608d0160408e0161501a565b610d1560808e0160608f0161501a565b8d60800135613028565b90945092505b508215610dad57610d3c60608a0160408b0161501a565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b610d7d60808e0160608f0161501a565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4610e35565b610db8898383613116565b610dc860608a0160408b0161501a565b6001600160a01b031660208a01358a357fd3122beb443f4b66faef07d45daf96faf0ffc23e050550cf1bbcbe6105f6bee7610e0960808e0160608f0161501a565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a45b6001945050505050610e5360015f516020615e9c5f395f51905f5255565b949350505050565b610e63612b05565b610e6e6035876132d6565b15610ecb57604080516080810182528781525f6020808301828152838501838152606085018481528c855260379093529490922092518355905160018301559151600282015590516003909101805460ff19169115159190911790555b6001600160a01b038416610ef257604051636ca1fdd760e01b815260040160405180910390fd5b5f868152603860205260409020610f0990856132e1565b8490610f34576040516311dd05f360e31b81526001600160a01b039091166004820152602401610aaf565b506040518060e00160405280856001600160a01b03168152602001846001600160a01b031681526020015f1515815260200186151581526020018281526020015f81526020015f81525060395f8881526020019081526020015f205f866001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160010160156101000a81548160ff0219169083151502179055506080820151816002015560a0820151816003015560c08201518160040155905050600182138061107657505f1982125b156110a0575f868152603a602090815260408083206001600160a01b038816845290915290208290555b826001600160a01b0316846001600160a01b0316877fffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a6685858a6040516110fb9392919092835260208301919091521515604082015260600190565b60405180910390a4505050505050565b61112d826111275f5f8681526020019081526020015f20612b37565b5f6126b8565b611139828260016126b8565b5050565b61114f6420a226a4a760d91b33611ff7565b6420a226a4a760d91b33909161118a5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f8381526038602052604090206111a39083612db7565b82906111ce5760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f8381526039602090815260408083206001600160a01b038616808552908352928190206002018490555183815285917fd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c891015b60405180910390a3505050565b6112426420a226a4a760d91b33611ff7565b6420a226a4a760d91b33909161127d5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b5050609680546001600160a01b0319166001600160a01b0392909216919091179055565b6112b36420a226a4a760d91b33611ff7565b6420a226a4a760d91b3390916112ee5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f8381526038602052604090206113079083612db7565b82906113325760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f8381526039602090815260408083206001600160a01b0386168085529252918290206001018054841515600160a01b0260ff60a01b19909116179055905184907f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea9061122390851515815260200190565b5f6113ae612d50565b89896113ba82826132f5565b6113c2612d80565b6113cf602085018561501a565b6001600160a01b038c81169116148b6113eb602087018761501a565b909161141d576040516313f7c32b60e31b81526001600160a01b03928316600482015291166024820152604401610aaf565b505061142c8c8c8b8b8b6133cb565b90985096508661143c898b615a47565b6114469190615a47565b60408501351015876114588a8c615a47565b6114629190615a47565b8560400135909161148f5760405163943892eb60e01b815260048101929092526024820152604401610aaf565b505f90506114a3604086016020870161501a565b30604087013560608801356114be60a08a0160808b016157c7565b6040516001600160a01b0395861660248201529490931660448501526064840191909152608483015260ff1660a482015260a086013560c482015260c086013560e48201526101040160408051601f19818403018152919052602080820180516001600160e01b031663d505accf60e01b1790529091505f908190611550906115499089018961501a565b5f856134af565b915091508181906115755760405163e8b5c4bb60e01b8152600401610aaf9190615723565b5050505061159e8c8c866020016020810190611591919061501a565b8d8d8d8d60018e8e61362c565b600192506115b860015f516020615e9c5f395f51905f5255565b50509998505050505050505050565b5f6115d0612d50565b6115d8612d80565b5f838152603b602052604090206115ef908361370e565b8290611611576040516373a1390160e11b8152600401610aaf91815260200190565b505f838152603c60209081526040808320858452909152808220815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e084015260058201805484929184916101008501919061168d90615a5a565b80601f01602080910402602001604051908101604052809291908181526020018280546116b990615a5a565b80156117045780601f106116db57610100808354040283529160200191611704565b820191905f5260205f20905b8154815290600101906020018083116116e757829003601f168201915b50505050508152505081526020016006820154815260200160078201805461172b90615a5a565b80601f016020809104026020016040519081016040528092919081815260200182805461175790615a5a565b80156117a25780601f10611779576101008083540402835291602001916117a2565b820191905f5260205f20905b81548152906001019060200180831161178557829003601f168201915b505050919092525050505f85815260396020908152604080832084518201516001600160a01b03908116855290835292819020815160e08101835281548516815260018201549485169381019390935260ff600160a01b8504811615801585850152600160a81b909504161515606084015260028101546080840152600381015460a08401526004015460c0830152835101519293509190611863576040516338384f6f60e11b81526001600160a01b039091166004820152602401610aaf565b50602082015115806118785750428260200151105b82602001514290916118a657604051637a88099160e11b815260048101929092526024820152604401610aaf565b50506118b28585613725565b925050506118cc60015f516020615e9c5f395f51905f5255565b92915050565b6118da6137f7565b6118e38261385d565b6111398282613865565b6118ff6420a226a4a760d91b33611ff7565b6420a226a4a760d91b33909161193a5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f838152603b60205260409020611953908361370e565b8290611975576040516373a1390160e11b8152600401610aaf91815260200190565b506119808142615a47565b5f848152603c602090815260408083208684528252918290206006019290925551828152839185917f530924d6189918da2c072b6eb46ef24e51151e8cb52c31d17482d8959b0d4c5b9101611223565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015611a145750825b90505f826001600160401b03166001148015611a2f5750303b155b905081158015611a3d575080155b15611a5b5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611a8557845460ff60401b1916600160401b1785555b611a8f8787613921565b8315611ad557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b5f611ae76139a9565b505f516020615e5c5f395f51905f5290565b5f81815260386020526040812060609190611b1390612b37565b90505f81516001600160401b03811115611b2f57611b2f614c94565b604051908082528060200260200182016040528015611b9457816020015b6040805160e0810182525f8082526020808301829052928201819052606082018190526080820181905260a0820181905260c082015282525f19909201910181611b4d5790505b5090505f5b8251811015611c765760395f8681526020019081526020015f205f848381518110611bc657611bc6615968565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160e08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b9093049091161515606082015260028201546080820152600382015460a082015260049091015460c08201528251839083908110611c6357611c63615968565b6020908102919091010152600101611b99565b509392505050565b5f611c87612d50565b8888611c9382826132f5565b611c9b612d80565b611ca88b8b8a8a8a6133cb565b9097509550611cbf8b8b338c8c8c8c5f8d8d61362c565b60019250611cd960015f516020615e9c5f395f51905f5255565b505098975050505050505050565b611cf96420a226a4a760d91b33611ff7565b6420a226a4a760d91b339091611d345760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50611d42905060358361370e565b8290611d645760405163ac7dbbfd60e01b8152600401610aaf91815260200190565b505f82815260376020908152604091829020600301805460ff1916841515908117909155915191825283917f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675910160405180910390a25050565b611dc6612b05565b611dcf5f6139f2565b565b5f818152603b602052604090206060906118cc90612b37565b5f611dfd6420a226a4a760d91b33611ff7565b6420a226a4a760d91b339091611e385760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f838152603b60205260409020611e51908361370e565b8290611e73576040516373a1390160e11b8152600401610aaf91815260200190565b50611e7e8383613725565b9392505050565b5f60608082808083815f516020615e3c5f395f51905f528054909150158015611eb057506001810154155b611ef45760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610aaf565b611efc613a62565b611f04613b22565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f805b85811015611fbc57611fb3878783818110611f4e57611f4e615968565b9050602002810190611f609190615a92565b868381518110611f7257611f72615968565b6020026020010151868481518110611f8c57611f8c615968565b6020026020010151868581518110611fa657611fa6615968565b6020026020010151610a40565b50600101611f31565b5060019695505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b5f828152602081905260408120611e7e9083612db7565b5f8181526020819052604090206060906118cc90612b37565b6120396420a226a4a760d91b33611ff7565b6420a226a4a760d91b3390916120745760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b5050603455565b5f8181526037602052604081206001908101546118cc91615a47565b61209f614bca565b5f838152603c6020908152604080832085845290915290819020815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e084015260058201805484929184916101008501919061211b90615a5a565b80601f016020809104026020016040519081016040528092919081815260200182805461214790615a5a565b80156121925780601f1061216957610100808354040283529160200191612192565b820191905f5260205f20905b81548152906001019060200180831161217557829003601f168201915b5050505050815250508152602001600682015481526020016007820180546121b990615a5a565b80601f01602080910402602001604051908101604052809291908181526020018280546121e590615a5a565b80156122305780601f1061220757610100808354040283529160200191612230565b820191905f5260205f20905b81548152906001019060200180831161221357829003601f168201915b505050505081525050905092915050565b612249612b05565b6064805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b61229a612b05565b80156122ab576122a8613b60565b50565b6122a8613bbc565b4260335410156122c05750565b5f806122cc6035612b37565b90505f5b8151811015612595575f8282815181106122ec576122ec615968565b6020908102919091018101515f818152603790925260409091206003015490915060ff161561231b575061258d565b5f818152603b6020526040812061233190612b37565b90505f5b8151811015612589575f82828151811061235157612351615968565b6020908102919091018101515f868152603c83526040808220838352909352828120835161012081019094528054606085019081526001820154608086015260028201546001600160a01b0390811660a087015260038301541660c0860152600482015460e08601526005820180549496509294939192849284916101008501916123db90615a5a565b80601f016020809104026020016040519081016040528092919081815260200182805461240790615a5a565b80156124525780601f1061242957610100808354040283529160200191612452565b820191905f5260205f20905b81548152906001019060200180831161243557829003601f168201915b50505050508152505081526020016006820154815260200160078201805461247990615a5a565b80601f01602080910402602001604051908101604052809291908181526020018280546124a590615a5a565b80156124f05780601f106124c7576101008083540402835291602001916124f0565b820191905f5260205f20905b8154815290600101906020018083116124d357829003601f168201915b5050509190925250508151515f90815260396020908152604080832085518201516001600160a01b0316845290915290206001015491925050600160a01b900460ff161580156125435750602081015115155b80156125525750428160200151105b1561257f576125618583613725565b508861256c89615ab0565b9850881061257f57505050505050505050565b5050600101612335565b5050505b6001016122d0565b50505050565b60606125a76035612b37565b905090565b6032545f906001600160a01b03166125d7576040516315aeca0d60e11b815260040160405180910390fd5b603254604051637c469ea160e11b81526001600160a01b039091169063f88d3d429061260d908a908a9088908890600401615ac8565b6020604051808303815f875af1158015612629573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061264d9190615b05565b905061265d875f83898989610e5b565b9695505050505050565b61266f612b05565b6001600160a01b03811661269657604051638219abe360e01b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b0392909216919091179055565b5f5b8251811015612595576126e7848483815181106126d9576126d9615968565b602002602001015184613c01565b6001016126ba565b5f818152603760205260408120600201546118cc906001615a47565b82811461272b576040516329f517fd60e01b815260040160405180910390fd5b5f5b838110156128975761288e85858381811061274a5761274a615968565b905060200281019061275c9190615b20565b3586868481811061276f5761276f615968565b90506020028101906127819190615b20565b61279290604081019060200161501a565b8787858181106127a4576127a4615968565b90506020028101906127b69190615b20565b6127c790606081019060400161501a565b8888868181106127d9576127d9615968565b90506020028101906127eb9190615b20565b6060013589898781811061280157612801615968565b90506020028101906128139190615b20565b608001358a8a8881811061282957612829615968565b905060200281019061283b9190615b20565b60a001358b8b8981811061285157612851615968565b90506020028101906128639190615b20565b6128719060c081019061597c565b8b8b8b81811061288357612883615968565b905060e002016113a5565b5060010161272d565b5050505050565b6096546040516337dba1f760e21b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063df6e87dc90606401606060405180830381865afa1580156128fa573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061291e9190615b34565b91989097509095509350505050565b612935612b05565b6032546001600160a01b0316801561296c5760405163f6c75f8f60e01b81526001600160a01b039091166004820152602401610aaf565b506001600160a01b03811661299457604051636ca1fdd760e01b815260040160405180910390fd5b603280546001600160a01b0319166001600160a01b0383169081179091556040517fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee905f90a250565b6129e5612b05565b6001600160a01b038116612a0e57604051631e4fbdf760e01b81525f6004820152602401610aaf565b6122a8816139f2565b612a1f612b05565b5f828152603860205260409020612a369082613d30565b8190612a615760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f8281526039602090815260408083206001600160a01b03851680855290835281842080546001600160a01b03191681556001810180546001600160b01b03191690556002810185905560038101859055600401849055858452603a835281842081855290925280832083905551909184917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d9190a35050565b5f6125a7613d44565b33612b0e611fc9565b6001600160a01b031614611dcf5760405163118cdaa760e01b8152336004820152602401610aaf565b60605f611e7e83613d4d565b612b4b614bf0565b5f838152603b60205260409020612b629083613da6565b8290612b8457604051637f11bea960e01b8152600401610aaf91815260200190565b505f838152603c60209081526040808320858452825291829020825160c0810184528154815260018201549281019290925260028101546001600160a01b0390811693830193909352600381015490921660608201526004820154608082015260058201805491929160a084019190612bfc90615a5a565b80601f0160208091040260200160405190810160405280929190818152602001828054612c2890615a5a565b8015612c735780601f10612c4a57610100808354040283529160200191612c73565b820191905f5260205f20905b815481529060010190602001808311612c5657829003601f168201915b505050919092525050505f848152603960209081526040808320848201516001600160a01b031684529091529020600181015491925090600160a81b900460ff1615612cd6578160800151816004015f828254612cd09190615b5f565b90915550505b5f848152603c602090815260408083208684529091528120818155600181018290556002810180546001600160a01b0319908116909155600382018054909116905560048101829055908181612d2f6005830182614c33565b5050600682015f9055600782015f612d479190614c33565b50505092915050565b5f516020615e7c5f395f51905f525460ff1615611dcf5760405163d93c066560e01b815260040160405180910390fd5b5f516020615e9c5f395f51905f52805460011901612db157604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600160a01b0381165f9081526001830160205260408120541515611e7e565b8251825181148015612dea5750815181145b835183518392612e1e576040516324227ae160e21b8152600481019390935260248301919091526044820152606401610aaf565b505060645482915060ff16811015612e4c57604051631aebd17960e11b8152600401610aaf91815260200190565b505f80826001600160401b03811115612e6757612e67614c94565b604051908082528060200260200182016040528015612e90578160200160208202803683370190505b5090505f5b83811015612ff2575f612f00888381518110612eb357612eb3615968565b6020026020010151888481518110612ecd57612ecd615968565b6020026020010151888581518110612ee757612ee7615968565b6020026020010151612ef88d613db1565b929190613ddd565b9050612f18682b20a624a220aa27a960b91b82611ff7565b682b20a624a220aa27a960b91b829091612f575760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b505f9050805b85811015612fa857848181518110612f7757612f77615968565b60200260200101516001600160a01b0316836001600160a01b031603612fa05760019150612fa8565b600101612f5d565b5080612fe85781848681518110612fc157612fc1615968565b60200260200101906001600160a01b031690816001600160a01b0316815250508460010194505b5050600101612e95565b50606454829060ff1681101561301e57604051631aebd17960e11b8152600401610aaf91815260200190565b5050505050505050565b5f848152603a602090815260408083206001600160a01b03871684529091528120546060906001811315613067576130608185615b72565b9350613086565b5f198112156130865761307981615b89565b6130839085615bb7565b93505b5f196001600160a01b038716016130ba576130b0858560405180602001604052805f8152506134af565b925092505061310d565b831561310b575f8781526039602090815260408083206001600160a01b038a168452909152902060010154600160a81b900460ff1615613100576130b087878787613e09565b6130b0868686613eff565b505b94509492505050565b82355f908152603b60209081526040909120613134918501356132d6565b83602001359061315a576040516307a5c91d60e31b8152600401610aaf91815260200190565b505f81156131835760345461316f9042615a47565b90506034544261317f9190615a47565b6033555b60405180606001604052808561319890615bca565b81526020808201849052604091820186905286355f908152603c82528281208883013582528252829020835180518255918201516001820155918101516002830180546001600160a01b039283166001600160a01b03199182161790915560608301516003850180549190931691161790556080810151600483015560a08101518290600582019061322a9082615c94565b505050602082015160068201556040820151600782019061324b9082615c94565b50505083355f90815260396020526040808220908290613271906060890190890161501a565b6001600160a01b0316815260208101919091526040015f206001810154909150600160a81b900460ff1615612897578460800135816004015f8282546132b79190615a47565b90915550505050505050565b60015f516020615e9c5f395f51905f5255565b5f611e7e8383613fe0565b5f611e7e836001600160a01b038416613fe0565b5f82815260386020526040902061330c9082612db7565b81906133375760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f82815260376020526040902060030154829060ff161561336f57604051636fc24b7960e11b8152600401610aaf91815260200190565b505f8281526039602090815260408083206001600160a01b03851684529091529020600101548190600160a01b900460ff1615610a3b576040516338384f6f60e11b81526001600160a01b039091166004820152602401610aaf565b6096545f9081906001600160a01b03166133e957505f9050806134a5565b6096546040516337dba1f760e21b8152600481018990526001600160a01b038881166024830152604482018890525f92169063df6e87dc90606401606060405180830381865afa15801561343f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134639190615b34565b9094509250905080861080159061347a5750828510155b80156134865750818410155b6134a3576040516358e8878560e01b815260040160405180910390fd5b505b9550959350505050565b5f60608347478211156134de57604051632f2a246160e21b815260048101929092526024820152604401610aaf565b50506060856001600160a01b031685856040516134fb9190615d4e565b5f6040518083038185875af1925050503d805f8114613535576040519150601f19603f3d011682016040523d82523d5f602084013e61353a565b606091505b5090935090508261357f57805115613556575f92509050613624565b50506040805180820190915260088152671c995d995c9d195960c21b60208201525f9150613624565b845f0361360f5780515f036135cd57856001600160a01b03163b5f036135c85750506040805180820190915260088152676e6f7420636f646560c01b60208201525f9150613624565b61360f565b602081015160018115151461360d575f6040518060400160405280600c81526020016b72657475726e2066616c736560a01b815250935093505050613624565b505b505060408051602081019091525f8152600191505b935093915050565b6136428a8a8a8961363d898b615a47565b61402c565b5f8a815260376020526040812060019081018054909101908190558190915060395f8d81526020019081526020015f205f8c6001600160a01b03166001600160a01b031681526020019081526020015f206001015f9054906101000a90046001600160a01b03169050896001600160a01b0316828d7f732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db508e858e8e8e8e428f8f8f6040516136f89a99989796959493929190615d64565b60405180910390a4505050505050505050505050565b5f8181526001830160205260408120541515611e7e565b5f5f6137318484612b43565b90505f5f613750835f0151846040015185606001518660800151613028565b915091508181906137745760405162461bcd60e51b8152600401610aaf9190615723565b5082604001516001600160a01b03168360200151845f01517f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b86606001518760800151426040516137e3939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a450600195945050505050565b3061c0de625c0eb760891b01148061383f575061c0de625c0eb760891b016138335f516020615e5c5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15611dcf5760405163703e46dd60e11b815260040160405180910390fd5b6122a8612b05565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156138bf575060408051601f3d908101601f191682019092526138bc91810190615dce565b60015b6138e757604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610aaf565b5f516020615e5c5f395f51905f52811461391757604051632a87526960e21b815260048101829052602401610aaf565b610a3b83836142ab565b613929614300565b61393233614349565b61393a61435a565b613942614362565b61394a614372565b61395382614382565b61395b6143e8565b6001600160a01b03811661398257604051638219abe360e01b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b03929092169190911790555043609855565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611dcf5760405163703e46dd60e11b815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020615e3c5f395f51905f5291613aa090615a5a565b80601f0160208091040260200160405190810160405280929190818152602001828054613acc90615a5a565b8015613b175780601f10613aee57610100808354040283529160200191613b17565b820191905f5260205f20905b815481529060010190602001808311613afa57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020615e3c5f395f51905f5291613aa090615a5a565b613b68612d50565b5f516020615e7c5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b039091168152602001612287565b613bc4614414565b5f516020615e7c5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613ba4565b613c09612b05565b6001600160a01b038216613c3057604051635989b9d360e01b815260040160405180910390fd5b82613c4e57604051630e1b39f960e31b815260040160405180910390fd5b8015613ca5575f838152602081905260409020613c6b90836132e1565b83839091613c9e57604051631b753c1b60e21b815260048101929092526001600160a01b03166024820152604401610aaf565b5050613cf2565b5f838152602081905260409020613cbc9083613d30565b83839091613cef57604051633a401ef360e21b815260048101929092526001600160a01b03166024820152604401610aaf565b50505b82826001600160a01b03167f8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f383604051611223911515815260200190565b5f611e7e836001600160a01b038416614443565b5f6125a7614526565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613d9a57602002820191905f5260205f20905b815481526020019060010190808311613d86575b50505050509050919050565b5f611e7e8383614443565b5f6118cc613dbd613d44565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613ded88888888614599565b925092509250613dfd8282614661565b50909695505050505050565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613e78575060408051601f3d908101601f19168201909252613e7591810190615de5565b60015b613eb2573d808015613ea5576040519150601f19603f3d011682016040523d82523d5f602084013e613eaa565b606091505b50905061310d565b8092508015613ecb57613ec6878786614719565b61310b565b6040518060400160405280600f81526020016e1d1c985b9cd9995c8819985a5b1959608a1b81525091505094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015613f6e575060408051601f3d908101601f19168201909252613f6b91810190615de5565b60015b613fa8573d808015613f9b576040519150601f19603f3d011682016040523d82523d5f602084013e613fa0565b606091505b509050613624565b80925080613fd7576040518060400160405280600b81526020016a1b5a5b9d0819985a5b195960aa1b81525091505b50935093915050565b5f81815260018301602052604081205461402557508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556118cc565b505f6118cc565b5f858152603a602090815260408083206001600160a01b03881684529091529020546001811315614097576140618184615e00565b8590849015614094576040516305330e5560e11b81526001600160a01b0390921660048301526024820152604401610aaf565b50505b5f196001600160a01b03861601614147576140b28284615a47565b34146140be8385615a47565b3490916140e7576040516362624a5160e11b815260048101929092526024820152604401610aaf565b505081156141425760975460408051602081019091525f808252918291614119916001600160a01b03169086906134af565b9150915081819061413e5760405163e8b5c4bb60e01b8152600401610aaf9190615723565b5050505b6142a3565b5f348015614171576040516362624a5160e11b815260048101929092526024820152604401610aaf565b50614195905084306141838587615a47565b6001600160a01b0389169291906147c4565b81156141b5576097546141b5906001600160a01b0387811691168461482b565b5f8681526039602090815260408083206001600160a01b0389168452909152902060010154600160a81b900460ff16156141f45761414286868561485c565b604051632770a7eb60e21b8152306004820152602481018490526001600160a01b03861690639dc29fac906044016020604051808303815f875af115801561423e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142629190615de5565b85858590919261429f57604051639ac2f56d60e01b81526001600160a01b0393841660048201529290911660248301526044820152606401610aaf565b5050505b505050505050565b6142b48261489a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156142f857610a3b82826148fd565b61113961496f565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16611dcf57604051631afcd79f60e31b815260040160405180910390fd5b614351614300565b6122a88161498e565b611dcf614300565b61436a614300565b611dcf614996565b61437a614300565b611dcf6149b6565b61438a614300565b6143d2604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b8152506149be565b6064805460ff191660ff92909216919091179055565b6143f0614300565b61440b6420a226a4a760d91b614404611fc9565b6001613c01565b62015180603455565b5f516020615e7c5f395f51905f525460ff16611dcf57604051638dfc202b60e01b815260040160405180910390fd5b5f818152600183016020526040812054801561451d575f614465600183615b5f565b85549091505f9061447890600190615b5f565b90508082146144d7575f865f01828154811061449657614496615968565b905f5260205f200154905080875f0184815481106144b6576144b6615968565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806144e8576144e8615e13565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506118cc565b5f9150506118cc565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6145506149d0565b614558614a38565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156145d257505f91506003905082614657565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614623573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661464e57505f925060019150829050614657565b92505f91508190505b9450945094915050565b5f82600381111561467457614674615e27565b0361467d575050565b600182600381111561469157614691615e27565b036146af5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156146c3576146c3615e27565b036146e45760405163fce698f760e01b815260048101829052602401610aaf565b60038260038111156146f8576146f8615e27565b03611139576040516335e2f38360e21b815260048101829052602401610aaf565b5f8381526039602090815260408083206001600160a01b0386168452909152902060048101546147499083615a47565b6003820154600483015486928692869290918210156147a1576040516352c2db3360e01b815260048101959095526001600160a01b03909316602485015260448401919091526064830152608482015260a401610aaf565b505050505081816003015f8282546147b99190615b5f565b909155505050505050565b6040516001600160a01b0384811660248301528381166044830152606482018390526125959186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614a7a565b6040516001600160a01b03838116602483015260448201839052610a3b91859182169063a9059cbb906064016147f9565b5f8381526039602090815260408083206001600160a01b038616845290915281206003018054839290614890908490615a47565b9091555050505050565b806001600160a01b03163b5f036148cf57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610aaf565b5f516020615e5c5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516149199190615d4e565b5f60405180830381855af49150503d805f8114614951576040519150601f19603f3d011682016040523d82523d5f602084013e614956565b606091505b5091509150614966858383614ae6565b95945050505050565b3415611dcf5760405163b398979f60e01b815260040160405180910390fd5b6129e5614300565b61499e614300565b5f516020615e7c5f395f51905f52805460ff19169055565b6132c3614300565b6149c6614300565b6111398282614b42565b5f5f516020615e3c5f395f51905f52816149e8613a62565b805190915015614a0057805160209091012092915050565b81548015614a0f579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020615e3c5f395f51905f5281614a50613b22565b805190915015614a6857805160209091012092915050565b60018201548015614a0f579392505050565b5f5f60205f8451602086015f885af180614a99576040513d5f823e3d81fd5b50505f513d91508115614ab0578060011415614abd565b6001600160a01b0384163b155b1561259557604051635274afe760e01b81526001600160a01b0385166004820152602401610aaf565b606082614afb57614af682614ba1565b611e7e565b8151158015614b1257506001600160a01b0384163b155b15614b3b57604051639996b31560e01b81526001600160a01b0385166004820152602401610aaf565b5092915050565b614b4a614300565b5f516020615e3c5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102614b838482615c94565b5060038101614b928382615c94565b505f8082556001909101555050565b805115614bb15780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6040518060600160405280614bdd614bf0565b81526020015f8152602001606081525090565b6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b508054614c3f90615a5a565b5f825580601f10614c4e575050565b601f0160209004905f5260205f20908101906122a891905b80821115614c79575f8155600101614c66565b5090565b5f60208284031215614c8d575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b0381118282101715614cca57614cca614c94565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614cf857614cf8614c94565b604052919050565b5f6001600160401b03821115614d1857614d18614c94565b5060051b60200190565b803560ff81168114614d32575f5ffd5b919050565b5f82601f830112614d46575f5ffd5b8135614d59614d5482614d00565b614cd0565b8082825260208201915060208360051b860101925085831115614d7a575f5ffd5b602085015b83811015614d9e57614d9081614d22565b835260209283019201614d7f565b5095945050505050565b5f82601f830112614db7575f5ffd5b8135614dc5614d5482614d00565b8082825260208201915060208360051b860101925085831115614de6575f5ffd5b602085015b83811015614d9e578035835260209283019201614deb565b5f5f5f5f60808587031215614e16575f5ffd5b84356001600160401b03811115614e2b575f5ffd5b850160c08188031215614e3c575f5ffd5b935060208501356001600160401b03811115614e56575f5ffd5b614e6287828801614d37565b93505060408501356001600160401b03811115614e7d575f5ffd5b614e8987828801614da8565b92505060608501356001600160401b03811115614ea4575f5ffd5b614eb087828801614da8565b91505092959194509250565b80151581146122a8575f5ffd5b6001600160a01b03811681146122a8575f5ffd5b5f5f5f5f5f5f60c08789031215614ef2575f5ffd5b863595506020870135614f0481614ebc565b94506040870135614f1481614ec9565b93506060870135614f2481614ec9565b9598949750929560808101359460a0909101359350915050565b5f82601f830112614f4d575f5ffd5b8135614f5b614d5482614d00565b8082825260208201915060208360051b860101925085831115614f7c575f5ffd5b602085015b83811015614d9e578035614f9481614ec9565b835260209283019201614f81565b5f5f60408385031215614fb3575f5ffd5b8235915060208301356001600160401b03811115614fcf575f5ffd5b614fdb85828601614f3e565b9150509250929050565b5f5f5f60608486031215614ff7575f5ffd5b83359250602084013561500981614ec9565b929592945050506040919091013590565b5f6020828403121561502a575f5ffd5b8135611e7e81614ec9565b5f5f5f60608486031215615047575f5ffd5b83359250602084013561505981614ec9565b9150604084013561506981614ebc565b809150509250925092565b5f5f83601f840112615084575f5ffd5b5081356001600160401b0381111561509a575f5ffd5b6020830191508360208285010111156150b1575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f898b036101c08112156150d2575f5ffd5b8a35995060208b01356150e481614ec9565b985060408b01356150f481614ec9565b975060608b0135965060808b0135955060a08b0135945060c08b01356001600160401b03811115615123575f5ffd5b61512f8d828e01615074565b90955093505060e060df1982011215615146575f5ffd5b5060e08a0190509295985092959850929598565b5f5f6040838503121561516b575f5ffd5b50508035926020909101359150565b5f5f6001600160401b0384111561519357615193614c94565b50601f8301601f19166020016151a881614cd0565b9150508281528383830111156151bc575f5ffd5b828260208301375f602084830101529392505050565b5f82601f8301126151e1575f5ffd5b611e7e8383356020850161517a565b5f5f60408385031215615201575f5ffd5b823561520c81614ec9565b915060208301356001600160401b03811115615226575f5ffd5b614fdb858286016151d2565b5f5f5f60608486031215615244575f5ffd5b505081359360208301359350604090920135919050565b5f5f6040838503121561526c575f5ffd5b61527583614d22565b9150602083013561528581614ec9565b809150509250929050565b80516001600160a01b039081168352602080830151909116908301526040808201511515908301526060808201511515908301526080808201519083015260a0818101519083015260c090810151910152565b602080825282518282018190525f918401906040840190835b818110156153255761530f838551615290565b6020939093019260e092909201916001016152fc565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b031215615347575f5ffd5b88359750602089013561535981614ec9565b9650604089013561536981614ec9565b9550606089013594506080890135935060a0890135925060c08901356001600160401b03811115615398575f5ffd5b6153a48b828c01615074565b999c989b5096995094979396929594505050565b5f5f604083850312156153c9575f5ffd5b82359150602083013561528581614ebc565b5f8151808452602084019350602083015f5b8281101561540b5781518652602095860195909101906001016153ed565b5093949350505050565b602081525f611e7e60208301846153db565b5f5f60408385031215615438575f5ffd5b82359150602083013561528581614ec9565b60e081016118cc8284615290565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff60f81b8816815260e060208201525f6154a460e0830189615458565b82810360408401526154b68189615458565b606084018890526001600160a01b038716608085015260a0840186905283810360c085015290506154e781856153db565b9a9950505050505050505050565b5f5f83601f840112615505575f5ffd5b5081356001600160401b0381111561551b575f5ffd5b6020830191508360208260051b85010111156150b1575f5ffd5b5f82601f830112615544575f5ffd5b8135615552614d5482614d00565b8082825260208201915060208360051b860101925085831115615573575f5ffd5b602085015b83811015614d9e5780356001600160401b03811115615595575f5ffd5b6155a4886020838a0101614da8565b84525060209283019201615578565b5f5f5f5f5f608086880312156155c7575f5ffd5b85356001600160401b038111156155dc575f5ffd5b6155e8888289016154f5565b90965094505060208601356001600160401b03811115615606575f5ffd5b8601601f81018813615616575f5ffd5b8035615624614d5482614d00565b8082825260208201915060208360051b85010192508a831115615645575f5ffd5b602084015b838110156156855780356001600160401b03811115615667575f5ffd5b6156768d602083890101614d37565b8452506020928301920161564a565b50955050505060408601356001600160401b038111156156a3575f5ffd5b6156af88828901615535565b92505060608601356001600160401b038111156156ca575f5ffd5b6156d688828901615535565b9150509295509295909350565b602080825282518282018190525f918401906040840190835b818110156153255783516001600160a01b03168352602093840193909201916001016156fc565b602081525f611e7e6020830184615458565b602081525f82516060602084015280516080840152602081015160a084015260018060a01b0360408201511660c084015260018060a01b0360608201511660e0840152608081015161010084015260a0810151905060c06101208401526157a0610140840182615458565b9050602084015160408401526040840151601f198483030160608501526149668282615458565b5f602082840312156157d7575f5ffd5b611e7e82614d22565b5f602082840312156157f0575f5ffd5b8135611e7e81614ebc565b5f5f5f5f5f5f60c08789031215615810575f5ffd5b86359550602087013561582281614ec9565b9450604087013593506060870135925060808701356001600160401b0381111561584a575f5ffd5b8701601f8101891361585a575f5ffd5b6158698982356020840161517a565b92505061587860a08801614d22565b90509295509295509295565b5f5f5f60608486031215615896575f5ffd5b8335925060208401356001600160401b038111156158b2575f5ffd5b6158be86828701614f3e565b925050604084013561506981614ebc565b5f5f5f5f604085870312156158e2575f5ffd5b84356001600160401b038111156158f7575f5ffd5b615903878288016154f5565b90955093505060208501356001600160401b03811115615921575f5ffd5b8501601f81018713615931575f5ffd5b80356001600160401b03811115615946575f5ffd5b87602060e08302840101111561595a575f5ffd5b949793965060200194505050565b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112615991575f5ffd5b8301803591506001600160401b038211156159aa575f5ffd5b6020019150368190038213156150b1575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90615a2690830184866159be565b9998505050505050505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156118cc576118cc615a33565b600181811c90821680615a6e57607f821691505b602082108103615a8c57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f823560be19833603018112615aa6575f5ffd5b9190910192915050565b5f60018201615ac157615ac1615a33565b5060010190565b8481526001600160a01b03841660208201526080604082018190525f90615af190830185615458565b905060ff8316606083015295945050505050565b5f60208284031215615b15575f5ffd5b8151611e7e81614ec9565b5f823560de19833603018112615aa6575f5ffd5b5f5f5f60608486031215615b46575f5ffd5b5050815160208301516040909301519094929350919050565b818103818111156118cc576118cc615a33565b80820281158282048414176118cc576118cc615a33565b5f600160ff1b8201615b9d57615b9d615a33565b505f0390565b634e487b7160e01b5f52601260045260245ffd5b5f82615bc557615bc5615ba3565b500490565b5f60c08236031215615bda575f5ffd5b615be2614ca8565b82358152602080840135908201526040830135615bfe81614ec9565b60408201526060830135615c1181614ec9565b60608201526080838101359082015260a08301356001600160401b03811115615c38575f5ffd5b615c44368286016151d2565b60a08301525092915050565b601f821115610a3b57805f5260205f20601f840160051c81016020851015615c755750805b601f840160051c820191505b81811015612897575f8155600101615c81565b81516001600160401b03811115615cad57615cad614c94565b615cc181615cbb8454615a5a565b84615c50565b6020601f821160018114615cf3575f8315615cdc5750848201515b5f19600385901b1c1916600184901b178455612897565b5f84815260208120601f198516915b82811015615d225787850151825560209485019460019092019101615d02565b5084821015615d3f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f82518060208501845e5f920191825250919050565b6001600160a01b038b811682528a8116602083015289166040820152606081018890526080810187905260a0810186905260c0810185905283151560e082015261012061010082018190525f90615dbe90830184866159be565b9c9b505050505050505050505050565b5f60208284031215615dde575f5ffd5b5051919050565b5f60208284031215615df5575f5ffd5b8151611e7e81614ebc565b5f82615e0e57615e0e615ba3565b500690565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212204a03319c766578113fe68a534ff2048534893f19cb929138828caf9809e0476564736f6c634300081c0033"

// Deprecated: Use CrossBridgeMetaData.Sigs instead.
// CrossBridgeFuncSigs maps the 4-byte function signature to its string representation.
var CrossBridgeFuncSigs = CrossBridgeMetaData.Sigs

// CrossBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossBridgeMetaData.Bin instead.
var CrossBridgeBin = CrossBridgeMetaData.Bin

// DeployCrossBridge deploys a new Ethereum contract, binding an instance of CrossBridge to it.
func DeployCrossBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossBridge, error) {
	parsed, err := CrossBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossBridge{CrossBridgeCaller: CrossBridgeCaller{contract: contract}, CrossBridgeTransactor: CrossBridgeTransactor{contract: contract}, CrossBridgeFilterer: CrossBridgeFilterer{contract: contract}}, nil
}

// CrossBridge is an auto generated Go binding around an Ethereum contract.
type CrossBridge struct {
	CrossBridgeCaller     // Read-only binding to the contract
	CrossBridgeTransactor // Write-only binding to the contract
	CrossBridgeFilterer   // Log filterer for contract events
}

// CrossBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossBridgeSession struct {
	Contract     *CrossBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossBridgeCallerSession struct {
	Contract *CrossBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CrossBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossBridgeTransactorSession struct {
	Contract     *CrossBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CrossBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossBridgeRaw struct {
	Contract *CrossBridge // Generic contract binding to access the raw methods on
}

// CrossBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossBridgeCallerRaw struct {
	Contract *CrossBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// CrossBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossBridgeTransactorRaw struct {
	Contract *CrossBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossBridge creates a new instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridge(address common.Address, backend bind.ContractBackend) (*CrossBridge, error) {
	contract, err := bindCrossBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossBridge{CrossBridgeCaller: CrossBridgeCaller{contract: contract}, CrossBridgeTransactor: CrossBridgeTransactor{contract: contract}, CrossBridgeFilterer: CrossBridgeFilterer{contract: contract}}, nil
}

// NewCrossBridgeCaller creates a new read-only instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeCaller(address common.Address, caller bind.ContractCaller) (*CrossBridgeCaller, error) {
	contract, err := bindCrossBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCaller{contract: contract}, nil
}

// NewCrossBridgeTransactor creates a new write-only instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossBridgeTransactor, error) {
	contract, err := bindCrossBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTransactor{contract: contract}, nil
}

// NewCrossBridgeFilterer creates a new log filterer instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossBridgeFilterer, error) {
	contract, err := bindCrossBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeFilterer{contract: contract}, nil
}

// bindCrossBridge binds a generic wrapper to an already deployed contract.
func bindCrossBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridge *CrossBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridge.Contract.CrossBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridge *CrossBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.Contract.CrossBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridge *CrossBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridge.Contract.CrossBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridge *CrossBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridge *CrossBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridge *CrossBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridge.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridge.Contract.UPGRADEINTERFACEVERSION(&_CrossBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridge.Contract.UPGRADEINTERFACEVERSION(&_CrossBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridge.Contract.AllChainIDs(&_CrossBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridge.Contract.AllChainIDs(&_CrossBridge.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridge.Contract.AllPendingIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridge.Contract.AllPendingIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.AllTokenPairs(&_CrossBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.AllTokenPairs(&_CrossBridge.CallOpts, remoteChainID)
}

// BridgeFeeManager is a free data retrieval call binding the contract method 0x25bd5666.
//
// Solidity: function bridgeFeeManager() view returns(address)
func (_CrossBridge *CrossBridgeCaller) BridgeFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "bridgeFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeFeeManager is a free data retrieval call binding the contract method 0x25bd5666.
//
// Solidity: function bridgeFeeManager() view returns(address)
func (_CrossBridge *CrossBridgeSession) BridgeFeeManager() (common.Address, error) {
	return _CrossBridge.Contract.BridgeFeeManager(&_CrossBridge.CallOpts)
}

// BridgeFeeManager is a free data retrieval call binding the contract method 0x25bd5666.
//
// Solidity: function bridgeFeeManager() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) BridgeFeeManager() (common.Address, error) {
	return _CrossBridge.Contract.BridgeFeeManager(&_CrossBridge.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_CrossBridge *CrossBridgeCaller) CalculateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "calculateFee", remoteChainID, token, value)

	outstruct := new(struct {
		MinimumValue *big.Int
		GasFee       *big.Int
		ExFee        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_CrossBridge *CrossBridgeSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	return _CrossBridge.Contract.CalculateFee(&_CrossBridge.CallOpts, remoteChainID, token, value)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_CrossBridge *CrossBridgeCallerSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	return _CrossBridge.Contract.CalculateFee(&_CrossBridge.CallOpts, remoteChainID, token, value)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridge *CrossBridgeCaller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridge *CrossBridgeSession) CrossMintableERC20Code() (common.Address, error) {
	return _CrossBridge.Contract.CrossMintableERC20Code(&_CrossBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _CrossBridge.Contract.CrossMintableERC20Code(&_CrossBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridge *CrossBridgeCaller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridge *CrossBridgeSession) Dev() (common.Address, error) {
	return _CrossBridge.Contract.Dev(&_CrossBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) Dev() (common.Address, error) {
	return _CrossBridge.Contract.Dev(&_CrossBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) DomainSeparator() ([32]byte, error) {
	return _CrossBridge.Contract.DomainSeparator(&_CrossBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _CrossBridge.Contract.DomainSeparator(&_CrossBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridge *CrossBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridge *CrossBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridge.Contract.Eip712Domain(&_CrossBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridge *CrossBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridge.Contract.Eip712Domain(&_CrossBridge.CallOpts)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextFinalizeIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextFinalizeIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextInitiateIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextInitiateIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_CrossBridge *CrossBridgeCaller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_CrossBridge *CrossBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridge.Contract.GetPendingArguments(&_CrossBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_CrossBridge *CrossBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridge.Contract.GetPendingArguments(&_CrossBridge.CallOpts, remoteChainID, index)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridge.Contract.GetRoleMembers(&_CrossBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridge.Contract.GetRoleMembers(&_CrossBridge.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.GetTokenPair(&_CrossBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.GetTokenPair(&_CrossBridge.CallOpts, remoteChainID, token)
}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_CrossBridge *CrossBridgeCaller) HasExpiredPending(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "hasExpiredPending")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_CrossBridge *CrossBridgeSession) HasExpiredPending() (bool, error) {
	return _CrossBridge.Contract.HasExpiredPending(&_CrossBridge.CallOpts)
}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) HasExpiredPending() (bool, error) {
	return _CrossBridge.Contract.HasExpiredPending(&_CrossBridge.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridge.Contract.HasRole(&_CrossBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridge.Contract.HasRole(&_CrossBridge.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) InitializedAt() (*big.Int, error) {
	return _CrossBridge.Contract.InitializedAt(&_CrossBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _CrossBridge.Contract.InitializedAt(&_CrossBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossBridge *CrossBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossBridge *CrossBridgeSession) Owner() (common.Address, error) {
	return _CrossBridge.Contract.Owner(&_CrossBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) Owner() (common.Address, error) {
	return _CrossBridge.Contract.Owner(&_CrossBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeSession) Paused() (bool, error) {
	return _CrossBridge.Contract.Paused(&_CrossBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) Paused() (bool, error) {
	return _CrossBridge.Contract.Paused(&_CrossBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _CrossBridge.Contract.ProxiableUUID(&_CrossBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossBridge.Contract.ProxiableUUID(&_CrossBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeSession) Threshold() (uint8, error) {
	return _CrossBridge.Contract.Threshold(&_CrossBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeCallerSession) Threshold() (uint8, error) {
	return _CrossBridge.Contract.Threshold(&_CrossBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, gasFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.ChangeThreshold(&_CrossBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.ChangeThreshold(&_CrossBridge.TransactOpts, threshold_)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_CrossBridge *CrossBridgeTransactor) ClearPending(opts *bind.TransactOpts, remoteChainID *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "clearPending", remoteChainID)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_CrossBridge *CrossBridgeSession) ClearPending(remoteChainID *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ClearPending(&_CrossBridge.TransactOpts, remoteChainID)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ClearPending(remoteChainID *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ClearPending(&_CrossBridge.TransactOpts, remoteChainID)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, conversionRatio, verificationAmountThreshold, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.CreateToken(&_CrossBridge.TransactOpts, remoteChainID, remoteToken, conversionRatio, verificationAmountThreshold, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.CreateToken(&_CrossBridge.TransactOpts, remoteChainID, remoteToken, conversionRatio, verificationAmountThreshold, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridge(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridge(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridgeBatch(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridgeBatch(&_CrossBridge.TransactOpts, args, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address dev_) returns()
func (_CrossBridge *CrossBridgeTransactor) Initialize(opts *bind.TransactOpts, _threshold uint8, dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "initialize", _threshold, dev_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address dev_) returns()
func (_CrossBridge *CrossBridgeSession) Initialize(_threshold uint8, dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.Initialize(&_CrossBridge.TransactOpts, _threshold, dev_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address dev_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) Initialize(_threshold uint8, dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.Initialize(&_CrossBridge.TransactOpts, _threshold, dev_)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeTransactor) ManualProcessPending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "manualProcessPending", remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeSession) ManualProcessPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualProcessPending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) ManualProcessPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualProcessPending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "permitBridgeToken", toChainID, fromToken, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeTokenBatch(&_CrossBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeTokenBatch(&_CrossBridge.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold) returns()
func (_CrossBridge *CrossBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken, conversionRatio, verificationAmountThreshold)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold) returns()
func (_CrossBridge *CrossBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RegisterToken(&_CrossBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, conversionRatio, verificationAmountThreshold)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RegisterToken(&_CrossBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, conversionRatio, verificationAmountThreshold)
}

// ReleaseExpiredPending is a paid mutator transaction binding the contract method 0xc1cb05dd.
//
// Solidity: function releaseExpiredPending(uint256 maxCount) returns()
func (_CrossBridge *CrossBridgeTransactor) ReleaseExpiredPending(opts *bind.TransactOpts, maxCount *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "releaseExpiredPending", maxCount)
}

// ReleaseExpiredPending is a paid mutator transaction binding the contract method 0xc1cb05dd.
//
// Solidity: function releaseExpiredPending(uint256 maxCount) returns()
func (_CrossBridge *CrossBridgeSession) ReleaseExpiredPending(maxCount *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleaseExpiredPending(&_CrossBridge.TransactOpts, maxCount)
}

// ReleaseExpiredPending is a paid mutator transaction binding the contract method 0xc1cb05dd.
//
// Solidity: function releaseExpiredPending(uint256 maxCount) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ReleaseExpiredPending(maxCount *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleaseExpiredPending(&_CrossBridge.TransactOpts, maxCount)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeTransactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossBridge *CrossBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossBridge *CrossBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossBridge.Contract.RenounceOwnership(&_CrossBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossBridge *CrossBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossBridge.Contract.RenounceOwnership(&_CrossBridge.TransactOpts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_CrossBridge *CrossBridgeTransactor) ResetRole(opts *bind.TransactOpts, role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "resetRole", role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_CrossBridge *CrossBridgeSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.ResetRole(&_CrossBridge.TransactOpts, role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.ResetRole(&_CrossBridge.TransactOpts, role, newAccounts)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridge *CrossBridgeTransactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridge *CrossBridgeSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossMintableERC20Code(&_CrossBridge.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossMintableERC20Code(&_CrossBridge.TransactOpts, _crossMintableERC20Code)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridge *CrossBridgeTransactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridge *CrossBridgeSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetDev(&_CrossBridge.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetDev(&_CrossBridge.TransactOpts, dev_)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _bridgeFeeManager) returns()
func (_CrossBridge *CrossBridgeTransactor) SetFeeManager(opts *bind.TransactOpts, _bridgeFeeManager common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setFeeManager", _bridgeFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _bridgeFeeManager) returns()
func (_CrossBridge *CrossBridgeSession) SetFeeManager(_bridgeFeeManager common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetFeeManager(&_CrossBridge.TransactOpts, _bridgeFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _bridgeFeeManager) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetFeeManager(_bridgeFeeManager common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetFeeManager(&_CrossBridge.TransactOpts, _bridgeFeeManager)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPause(&_CrossBridge.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPause(&_CrossBridge.TransactOpts, set)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactor) SetPauseChain(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setPauseChain", remoteChainID, pause)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeSession) SetPauseChain(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPauseChain(&_CrossBridge.TransactOpts, remoteChainID, pause)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetPauseChain(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPauseChain(&_CrossBridge.TransactOpts, remoteChainID, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactor) SetPauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setPauseToken", remoteChainID, token, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_CrossBridge *CrossBridgeSession) SetPauseToken(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPauseToken(&_CrossBridge.TransactOpts, remoteChainID, token, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetPauseToken(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPauseToken(&_CrossBridge.TransactOpts, remoteChainID, token, pause)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_CrossBridge *CrossBridgeTransactor) SetRole(opts *bind.TransactOpts, role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setRole", role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_CrossBridge *CrossBridgeSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetRole(&_CrossBridge.TransactOpts, role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetRole(&_CrossBridge.TransactOpts, role, accounts, set)
}

// SetVerificationAmountThreshold is a paid mutator transaction binding the contract method 0x3b96cf5a.
//
// Solidity: function setVerificationAmountThreshold(uint256 remoteChainID, address token, uint256 verificationAmountThreshold) returns()
func (_CrossBridge *CrossBridgeTransactor) SetVerificationAmountThreshold(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setVerificationAmountThreshold", remoteChainID, token, verificationAmountThreshold)
}

// SetVerificationAmountThreshold is a paid mutator transaction binding the contract method 0x3b96cf5a.
//
// Solidity: function setVerificationAmountThreshold(uint256 remoteChainID, address token, uint256 verificationAmountThreshold) returns()
func (_CrossBridge *CrossBridgeSession) SetVerificationAmountThreshold(remoteChainID *big.Int, token common.Address, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationAmountThreshold(&_CrossBridge.TransactOpts, remoteChainID, token, verificationAmountThreshold)
}

// SetVerificationAmountThreshold is a paid mutator transaction binding the contract method 0x3b96cf5a.
//
// Solidity: function setVerificationAmountThreshold(uint256 remoteChainID, address token, uint256 verificationAmountThreshold) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetVerificationAmountThreshold(remoteChainID *big.Int, token common.Address, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationAmountThreshold(&_CrossBridge.TransactOpts, remoteChainID, token, verificationAmountThreshold)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0x4f86d44b.
//
// Solidity: function setVerificationDelay(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactor) SetVerificationDelay(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setVerificationDelay", remoteChainID, index, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0x4f86d44b.
//
// Solidity: function setVerificationDelay(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridge *CrossBridgeSession) SetVerificationDelay(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelay(&_CrossBridge.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0x4f86d44b.
//
// Solidity: function setVerificationDelay(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetVerificationDelay(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelay(&_CrossBridge.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelay0 is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactor) SetVerificationDelay0(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setVerificationDelay0", delay)
}

// SetVerificationDelay0 is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridge *CrossBridgeSession) SetVerificationDelay0(delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelay0(&_CrossBridge.TransactOpts, delay)
}

// SetVerificationDelay0 is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetVerificationDelay0(delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelay0(&_CrossBridge.TransactOpts, delay)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossBridge *CrossBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossBridge *CrossBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.TransferOwnership(&_CrossBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossBridge *CrossBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.TransferOwnership(&_CrossBridge.TransactOpts, newOwner)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_CrossBridge *CrossBridgeTransactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_CrossBridge *CrossBridgeSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.UnregisterToken(&_CrossBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_CrossBridge *CrossBridgeTransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.UnregisterToken(&_CrossBridge.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.UpgradeToAndCall(&_CrossBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.UpgradeToAndCall(&_CrossBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrossBridge *CrossBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrossBridge *CrossBridgeSession) Receive() (*types.Transaction, error) {
	return _CrossBridge.Contract.Receive(&_CrossBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _CrossBridge.Contract.Receive(&_CrossBridge.TransactOpts)
}

// CrossBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the CrossBridge contract.
type CrossBridgeBridgeFinalizedIterator struct {
	Event *CrossBridgeBridgeFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeBridgeFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeFinalized represents a BridgeFinalized event raised by the CrossBridge contract.
type CrossBridgeBridgeFinalized struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalized is a free log retrieval operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*CrossBridgeBridgeFinalizedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeFinalizedIterator{contract: _CrossBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeFinalized)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeFinalized is a log parse operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeFinalized(log types.Log) (*CrossBridgeBridgeFinalized, error) {
	event := new(CrossBridgeBridgeFinalized)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the CrossBridge contract.
type CrossBridgeBridgeInitiatedIterator struct {
	Event *CrossBridgeBridgeInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeBridgeInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeInitiated represents a BridgeInitiated event raised by the CrossBridge contract.
type CrossBridgeBridgeInitiated struct {
	ToChainID *big.Int
	Index     *big.Int
	FromToken common.Address
	ToToken   common.Address
	From      common.Address
	To        common.Address
	Value     *big.Int
	GasFee    *big.Int
	ExFee     *big.Int
	Time      *big.Int
	Permit    bool
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiated is a free log retrieval operation binding the contract event 0x732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db50.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 exFee, uint256 time, bool permit, bytes extraData)
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*CrossBridgeBridgeInitiatedIterator, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeInitiatedIterator{contract: _CrossBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db50.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 exFee, uint256 time, bool permit, bytes extraData)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeInitiated)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeInitiated is a log parse operation binding the contract event 0x732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db50.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 exFee, uint256 time, bool permit, bytes extraData)
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeInitiated(log types.Log) (*CrossBridgeBridgeInitiated, error) {
	event := new(CrossBridgeBridgeInitiated)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the CrossBridge contract.
type CrossBridgeBridgePendingIterator struct {
	Event *CrossBridgeBridgePending // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeBridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgePending)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeBridgePending)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeBridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgePending represents a BridgePending event raised by the CrossBridge contract.
type CrossBridgeBridgePending struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgePending is a free log retrieval operation binding the contract event 0xd3122beb443f4b66faef07d45daf96faf0ffc23e050550cf1bbcbe6105f6bee7.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*CrossBridgeBridgePendingIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgePendingIterator{contract: _CrossBridge.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0xd3122beb443f4b66faef07d45daf96faf0ffc23e050550cf1bbcbe6105f6bee7.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgePending)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgePending is a log parse operation binding the contract event 0xd3122beb443f4b66faef07d45daf96faf0ffc23e050550cf1bbcbe6105f6bee7.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) ParseBridgePending(log types.Log) (*CrossBridgeBridgePending, error) {
	event := new(CrossBridgeBridgePending)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the CrossBridge contract.
type CrossBridgeChainPauseSetIterator struct {
	Event *CrossBridgeChainPauseSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeChainPauseSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeChainPauseSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeChainPauseSet represents a ChainPauseSet event raised by the CrossBridge contract.
type CrossBridgeChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*CrossBridgeChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeChainPauseSetIterator{contract: _CrossBridge.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeChainPauseSet)
				if err := _CrossBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainPauseSet is a log parse operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) ParseChainPauseSet(log types.Log) (*CrossBridgeChainPauseSet, error) {
	event := new(CrossBridgeChainPauseSet)
	if err := _CrossBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeCrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the CrossBridge contract.
type CrossBridgeCrossMintableERC20CodeSetIterator struct {
	Event *CrossBridgeCrossMintableERC20CodeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeCrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeCrossMintableERC20CodeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeCrossMintableERC20CodeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeCrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeCrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeCrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the CrossBridge contract.
type CrossBridgeCrossMintableERC20CodeSet struct {
	Code common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_CrossBridge *CrossBridgeFilterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, code []common.Address) (*CrossBridgeCrossMintableERC20CodeSetIterator, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCrossMintableERC20CodeSetIterator{contract: _CrossBridge.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_CrossBridge *CrossBridgeFilterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeCrossMintableERC20CodeSet, code []common.Address) (event.Subscription, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeCrossMintableERC20CodeSet)
				if err := _CrossBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossMintableERC20CodeSet is a log parse operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_CrossBridge *CrossBridgeFilterer) ParseCrossMintableERC20CodeSet(log types.Log) (*CrossBridgeCrossMintableERC20CodeSet, error) {
	event := new(CrossBridgeCrossMintableERC20CodeSet)
	if err := _CrossBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the CrossBridge contract.
type CrossBridgeEIP712DomainChangedIterator struct {
	Event *CrossBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the CrossBridge contract.
type CrossBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridge *CrossBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*CrossBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeEIP712DomainChangedIterator{contract: _CrossBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridge *CrossBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeEIP712DomainChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridge *CrossBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*CrossBridgeEIP712DomainChanged, error) {
	event := new(CrossBridgeEIP712DomainChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossBridge contract.
type CrossBridgeInitializedIterator struct {
	Event *CrossBridgeInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeInitialized represents a Initialized event raised by the CrossBridge contract.
type CrossBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridge *CrossBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossBridgeInitializedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeInitializedIterator{contract: _CrossBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridge *CrossBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeInitialized)
				if err := _CrossBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridge *CrossBridgeFilterer) ParseInitialized(log types.Log) (*CrossBridgeInitialized, error) {
	event := new(CrossBridgeInitialized)
	if err := _CrossBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CrossBridge contract.
type CrossBridgeOwnershipTransferredIterator struct {
	Event *CrossBridgeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the CrossBridge contract.
type CrossBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossBridge *CrossBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrossBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeOwnershipTransferredIterator{contract: _CrossBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossBridge *CrossBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrossBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeOwnershipTransferred)
				if err := _CrossBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossBridge *CrossBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*CrossBridgeOwnershipTransferred, error) {
	event := new(CrossBridgeOwnershipTransferred)
	if err := _CrossBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CrossBridge contract.
type CrossBridgePausedIterator struct {
	Event *CrossBridgePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgePaused represents a Paused event raised by the CrossBridge contract.
type CrossBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridge *CrossBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*CrossBridgePausedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgePausedIterator{contract: _CrossBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridge *CrossBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossBridgePaused) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgePaused)
				if err := _CrossBridge.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridge *CrossBridgeFilterer) ParsePaused(log types.Log) (*CrossBridgePaused, error) {
	event := new(CrossBridgePaused)
	if err := _CrossBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeRoleUpdatedIterator is returned from FilterRoleUpdated and is used to iterate over the raw logs and unpacked data for RoleUpdated events raised by the CrossBridge contract.
type CrossBridgeRoleUpdatedIterator struct {
	Event *CrossBridgeRoleUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeRoleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeRoleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeRoleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeRoleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeRoleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeRoleUpdated represents a RoleUpdated event raised by the CrossBridge contract.
type CrossBridgeRoleUpdated struct {
	Account common.Address
	Role    [32]byte
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleUpdated is a free log retrieval operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool status)
func (_CrossBridge *CrossBridgeFilterer) FilterRoleUpdated(opts *bind.FilterOpts, account []common.Address, role [][32]byte) (*CrossBridgeRoleUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "RoleUpdated", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeRoleUpdatedIterator{contract: _CrossBridge.contract, event: "RoleUpdated", logs: logs, sub: sub}, nil
}

// WatchRoleUpdated is a free log subscription operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool status)
func (_CrossBridge *CrossBridgeFilterer) WatchRoleUpdated(opts *bind.WatchOpts, sink chan<- *CrossBridgeRoleUpdated, account []common.Address, role [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "RoleUpdated", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeRoleUpdated)
				if err := _CrossBridge.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleUpdated is a log parse operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool status)
func (_CrossBridge *CrossBridgeFilterer) ParseRoleUpdated(log types.Log) (*CrossBridgeRoleUpdated, error) {
	event := new(CrossBridgeRoleUpdated)
	if err := _CrossBridge.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeSafetyLimitSetIterator is returned from FilterSafetyLimitSet and is used to iterate over the raw logs and unpacked data for SafetyLimitSet events raised by the CrossBridge contract.
type CrossBridgeSafetyLimitSetIterator struct {
	Event *CrossBridgeSafetyLimitSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeSafetyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeSafetyLimitSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeSafetyLimitSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeSafetyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeSafetyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeSafetyLimitSet represents a SafetyLimitSet event raised by the CrossBridge contract.
type CrossBridgeSafetyLimitSet struct {
	RemoteChainID               *big.Int
	Token                       common.Address
	VerificationAmountThreshold *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSafetyLimitSet is a free log retrieval operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 verificationAmountThreshold)
func (_CrossBridge *CrossBridgeFilterer) FilterSafetyLimitSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*CrossBridgeSafetyLimitSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "SafetyLimitSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeSafetyLimitSetIterator{contract: _CrossBridge.contract, event: "SafetyLimitSet", logs: logs, sub: sub}, nil
}

// WatchSafetyLimitSet is a free log subscription operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 verificationAmountThreshold)
func (_CrossBridge *CrossBridgeFilterer) WatchSafetyLimitSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeSafetyLimitSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "SafetyLimitSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeSafetyLimitSet)
				if err := _CrossBridge.contract.UnpackLog(event, "SafetyLimitSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSafetyLimitSet is a log parse operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 verificationAmountThreshold)
func (_CrossBridge *CrossBridgeFilterer) ParseSafetyLimitSet(log types.Log) (*CrossBridgeSafetyLimitSet, error) {
	event := new(CrossBridgeSafetyLimitSet)
	if err := _CrossBridge.contract.UnpackLog(event, "SafetyLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the CrossBridge contract.
type CrossBridgeThresholdChangedIterator struct {
	Event *CrossBridgeThresholdChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeThresholdChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeThresholdChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeThresholdChanged represents a ThresholdChanged event raised by the CrossBridge contract.
type CrossBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridge *CrossBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*CrossBridgeThresholdChangedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeThresholdChangedIterator{contract: _CrossBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridge *CrossBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeThresholdChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseThresholdChanged is a log parse operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridge *CrossBridgeFilterer) ParseThresholdChanged(log types.Log) (*CrossBridgeThresholdChanged, error) {
	event := new(CrossBridgeThresholdChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the CrossBridge contract.
type CrossBridgeTokenPairRegisteredIterator struct {
	Event *CrossBridgeTokenPairRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPairRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeTokenPairRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the CrossBridge contract.
type CrossBridgeTokenPairRegistered struct {
	RemoteChainID               *big.Int
	LocalToken                  common.Address
	RemoteToken                 common.Address
	ConversionRatio             *big.Int
	VerificationAmountThreshold *big.Int
	IsOrigin                    bool
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*CrossBridgeTokenPairRegisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPairRegisteredIterator{contract: _CrossBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPairRegistered)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPairRegistered is a log parse operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*CrossBridgeTokenPairRegistered, error) {
	event := new(CrossBridgeTokenPairRegistered)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the CrossBridge contract.
type CrossBridgeTokenPairUnregisteredIterator struct {
	Event *CrossBridgeTokenPairUnregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeTokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPairUnregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeTokenPairUnregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeTokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPairUnregistered represents a TokenPairUnregistered event raised by the CrossBridge contract.
type CrossBridgeTokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*CrossBridgeTokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPairUnregisteredIterator{contract: _CrossBridge.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPairUnregistered)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPairUnregistered is a log parse operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPairUnregistered(log types.Log) (*CrossBridgeTokenPairUnregistered, error) {
	event := new(CrossBridgeTokenPairUnregistered)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the CrossBridge contract.
type CrossBridgeTokenPauseSetIterator struct {
	Event *CrossBridgeTokenPauseSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeTokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPauseSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeTokenPauseSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeTokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPauseSet represents a TokenPauseSet event raised by the CrossBridge contract.
type CrossBridgeTokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*CrossBridgeTokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPauseSetIterator{contract: _CrossBridge.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPauseSet)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPauseSet is a log parse operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPauseSet(log types.Log) (*CrossBridgeTokenPauseSet, error) {
	event := new(CrossBridgeTokenPauseSet)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CrossBridge contract.
type CrossBridgeUnpausedIterator struct {
	Event *CrossBridgeUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeUnpaused represents a Unpaused event raised by the CrossBridge contract.
type CrossBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridge *CrossBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CrossBridgeUnpausedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeUnpausedIterator{contract: _CrossBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridge *CrossBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeUnpaused)
				if err := _CrossBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridge *CrossBridgeFilterer) ParseUnpaused(log types.Log) (*CrossBridgeUnpaused, error) {
	event := new(CrossBridgeUnpaused)
	if err := _CrossBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossBridge contract.
type CrossBridgeUpgradedIterator struct {
	Event *CrossBridgeUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeUpgraded represents a Upgraded event raised by the CrossBridge contract.
type CrossBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridge *CrossBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeUpgradedIterator{contract: _CrossBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridge *CrossBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeUpgraded)
				if err := _CrossBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridge *CrossBridgeFilterer) ParseUpgraded(log types.Log) (*CrossBridgeUpgraded, error) {
	event := new(CrossBridgeUpgraded)
	if err := _CrossBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeVerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the CrossBridge contract.
type CrossBridgeVerificationDelaySetIterator struct {
	Event *CrossBridgeVerificationDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossBridgeVerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeVerificationDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossBridgeVerificationDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossBridgeVerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeVerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeVerificationDelaySet represents a VerificationDelaySet event raised by the CrossBridge contract.
type CrossBridgeVerificationDelaySet struct {
	FromChainID *big.Int
	Index       *big.Int
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0x530924d6189918da2c072b6eb46ef24e51151e8cb52c31d17482d8959b0d4c5b.
//
// Solidity: event VerificationDelaySet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_CrossBridge *CrossBridgeFilterer) FilterVerificationDelaySet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*CrossBridgeVerificationDelaySetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "VerificationDelaySet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeVerificationDelaySetIterator{contract: _CrossBridge.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0x530924d6189918da2c072b6eb46ef24e51151e8cb52c31d17482d8959b0d4c5b.
//
// Solidity: event VerificationDelaySet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_CrossBridge *CrossBridgeFilterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *CrossBridgeVerificationDelaySet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "VerificationDelaySet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeVerificationDelaySet)
				if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerificationDelaySet is a log parse operation binding the contract event 0x530924d6189918da2c072b6eb46ef24e51151e8cb52c31d17482d8959b0d4c5b.
//
// Solidity: event VerificationDelaySet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_CrossBridge *CrossBridgeFilterer) ParseVerificationDelaySet(log types.Log) (*CrossBridgeVerificationDelaySet, error) {
	event := new(CrossBridgeVerificationDelaySet)
	if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
