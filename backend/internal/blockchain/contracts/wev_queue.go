// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IWEVQueueRedemptionRequest is an auto generated low-level Go binding around an user-defined struct.
type IWEVQueueRedemptionRequest struct {
	RequestId   [32]byte
	User        common.Address
	SWattAmount *big.Int
	PriorityFee *big.Int
	RequestedAt *big.Int
	Status      uint8
}

// WEVQueueMetaData contains all meta data concerning the WEVQueue contract.
var WEVQueueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientPriorityFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIWEVQueue.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"NotQueued\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"NotRequestOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToProcess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"RequestNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"BatchProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RedemptionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wattAmount\",\"type\":\"uint256\"}],\"name\":\"RedemptionFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sWattAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPriority\",\"type\":\"bool\"}],\"name\":\"RedemptionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRIORITY_FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRIORITY_WAIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROCESSOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STANDARD_WAIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQueueDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"getRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sWattAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedAt\",\"type\":\"uint256\"},{\"internalType\":\"enumIWEVQueue.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIWEVQueue.RedemptionRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserRequests\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sWattUSD_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wattUSD_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextProcessingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"requestIds\",\"type\":\"bytes32[]\"}],\"name\":\"processBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sWattAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priorityFee\",\"type\":\"uint256\"}],\"name\":\"requestPriorityRedeem\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sWattAmount\",\"type\":\"uint256\"}],\"name\":\"requestRedeem\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sWattUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wattUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WEVQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use WEVQueueMetaData.ABI instead.
var WEVQueueABI = WEVQueueMetaData.ABI

// WEVQueue is an auto generated Go binding around an Ethereum contract.
type WEVQueue struct {
	WEVQueueCaller     // Read-only binding to the contract
	WEVQueueTransactor // Write-only binding to the contract
	WEVQueueFilterer   // Log filterer for contract events
}

// WEVQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type WEVQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WEVQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WEVQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WEVQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WEVQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WEVQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WEVQueueSession struct {
	Contract     *WEVQueue         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WEVQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WEVQueueCallerSession struct {
	Contract *WEVQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WEVQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WEVQueueTransactorSession struct {
	Contract     *WEVQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WEVQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type WEVQueueRaw struct {
	Contract *WEVQueue // Generic contract binding to access the raw methods on
}

// WEVQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WEVQueueCallerRaw struct {
	Contract *WEVQueueCaller // Generic read-only contract binding to access the raw methods on
}

// WEVQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WEVQueueTransactorRaw struct {
	Contract *WEVQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWEVQueue creates a new instance of WEVQueue, bound to a specific deployed contract.
func NewWEVQueue(address common.Address, backend bind.ContractBackend) (*WEVQueue, error) {
	contract, err := bindWEVQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WEVQueue{WEVQueueCaller: WEVQueueCaller{contract: contract}, WEVQueueTransactor: WEVQueueTransactor{contract: contract}, WEVQueueFilterer: WEVQueueFilterer{contract: contract}}, nil
}

// NewWEVQueueCaller creates a new read-only instance of WEVQueue, bound to a specific deployed contract.
func NewWEVQueueCaller(address common.Address, caller bind.ContractCaller) (*WEVQueueCaller, error) {
	contract, err := bindWEVQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WEVQueueCaller{contract: contract}, nil
}

// NewWEVQueueTransactor creates a new write-only instance of WEVQueue, bound to a specific deployed contract.
func NewWEVQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*WEVQueueTransactor, error) {
	contract, err := bindWEVQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WEVQueueTransactor{contract: contract}, nil
}

// NewWEVQueueFilterer creates a new log filterer instance of WEVQueue, bound to a specific deployed contract.
func NewWEVQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*WEVQueueFilterer, error) {
	contract, err := bindWEVQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WEVQueueFilterer{contract: contract}, nil
}

// bindWEVQueue binds a generic wrapper to an already deployed contract.
func bindWEVQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WEVQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WEVQueue *WEVQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WEVQueue.Contract.WEVQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WEVQueue *WEVQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEVQueue.Contract.WEVQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WEVQueue *WEVQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WEVQueue.Contract.WEVQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WEVQueue *WEVQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WEVQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WEVQueue *WEVQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEVQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WEVQueue *WEVQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WEVQueue.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueSession) ADMINROLE() ([32]byte, error) {
	return _WEVQueue.Contract.ADMINROLE(&_WEVQueue.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCallerSession) ADMINROLE() ([32]byte, error) {
	return _WEVQueue.Contract.ADMINROLE(&_WEVQueue.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _WEVQueue.Contract.DEFAULTADMINROLE(&_WEVQueue.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _WEVQueue.Contract.DEFAULTADMINROLE(&_WEVQueue.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueSession) PAUSERROLE() ([32]byte, error) {
	return _WEVQueue.Contract.PAUSERROLE(&_WEVQueue.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCallerSession) PAUSERROLE() ([32]byte, error) {
	return _WEVQueue.Contract.PAUSERROLE(&_WEVQueue.CallOpts)
}

// PRIORITYFEEBPS is a free data retrieval call binding the contract method 0x05ee4b32.
//
// Solidity: function PRIORITY_FEE_BPS() view returns(uint256)
func (_WEVQueue *WEVQueueCaller) PRIORITYFEEBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "PRIORITY_FEE_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRIORITYFEEBPS is a free data retrieval call binding the contract method 0x05ee4b32.
//
// Solidity: function PRIORITY_FEE_BPS() view returns(uint256)
func (_WEVQueue *WEVQueueSession) PRIORITYFEEBPS() (*big.Int, error) {
	return _WEVQueue.Contract.PRIORITYFEEBPS(&_WEVQueue.CallOpts)
}

// PRIORITYFEEBPS is a free data retrieval call binding the contract method 0x05ee4b32.
//
// Solidity: function PRIORITY_FEE_BPS() view returns(uint256)
func (_WEVQueue *WEVQueueCallerSession) PRIORITYFEEBPS() (*big.Int, error) {
	return _WEVQueue.Contract.PRIORITYFEEBPS(&_WEVQueue.CallOpts)
}

// PRIORITYWAIT is a free data retrieval call binding the contract method 0x8a2c650d.
//
// Solidity: function PRIORITY_WAIT() view returns(uint256)
func (_WEVQueue *WEVQueueCaller) PRIORITYWAIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "PRIORITY_WAIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRIORITYWAIT is a free data retrieval call binding the contract method 0x8a2c650d.
//
// Solidity: function PRIORITY_WAIT() view returns(uint256)
func (_WEVQueue *WEVQueueSession) PRIORITYWAIT() (*big.Int, error) {
	return _WEVQueue.Contract.PRIORITYWAIT(&_WEVQueue.CallOpts)
}

// PRIORITYWAIT is a free data retrieval call binding the contract method 0x8a2c650d.
//
// Solidity: function PRIORITY_WAIT() view returns(uint256)
func (_WEVQueue *WEVQueueCallerSession) PRIORITYWAIT() (*big.Int, error) {
	return _WEVQueue.Contract.PRIORITYWAIT(&_WEVQueue.CallOpts)
}

// PROCESSORROLE is a free data retrieval call binding the contract method 0x8222bdb2.
//
// Solidity: function PROCESSOR_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCaller) PROCESSORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "PROCESSOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROCESSORROLE is a free data retrieval call binding the contract method 0x8222bdb2.
//
// Solidity: function PROCESSOR_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueSession) PROCESSORROLE() ([32]byte, error) {
	return _WEVQueue.Contract.PROCESSORROLE(&_WEVQueue.CallOpts)
}

// PROCESSORROLE is a free data retrieval call binding the contract method 0x8222bdb2.
//
// Solidity: function PROCESSOR_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCallerSession) PROCESSORROLE() ([32]byte, error) {
	return _WEVQueue.Contract.PROCESSORROLE(&_WEVQueue.CallOpts)
}

// STANDARDWAIT is a free data retrieval call binding the contract method 0x485f8ff3.
//
// Solidity: function STANDARD_WAIT() view returns(uint256)
func (_WEVQueue *WEVQueueCaller) STANDARDWAIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "STANDARD_WAIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STANDARDWAIT is a free data retrieval call binding the contract method 0x485f8ff3.
//
// Solidity: function STANDARD_WAIT() view returns(uint256)
func (_WEVQueue *WEVQueueSession) STANDARDWAIT() (*big.Int, error) {
	return _WEVQueue.Contract.STANDARDWAIT(&_WEVQueue.CallOpts)
}

// STANDARDWAIT is a free data retrieval call binding the contract method 0x485f8ff3.
//
// Solidity: function STANDARD_WAIT() view returns(uint256)
func (_WEVQueue *WEVQueueCallerSession) STANDARDWAIT() (*big.Int, error) {
	return _WEVQueue.Contract.STANDARDWAIT(&_WEVQueue.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueSession) UPGRADERROLE() ([32]byte, error) {
	return _WEVQueue.Contract.UPGRADERROLE(&_WEVQueue.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_WEVQueue *WEVQueueCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _WEVQueue.Contract.UPGRADERROLE(&_WEVQueue.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_WEVQueue *WEVQueueCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_WEVQueue *WEVQueueSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _WEVQueue.Contract.UPGRADEINTERFACEVERSION(&_WEVQueue.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_WEVQueue *WEVQueueCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _WEVQueue.Contract.UPGRADEINTERFACEVERSION(&_WEVQueue.CallOpts)
}

// GetProtocolFees is a free data retrieval call binding the contract method 0xd8dfcea0.
//
// Solidity: function getProtocolFees() view returns(uint256)
func (_WEVQueue *WEVQueueCaller) GetProtocolFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "getProtocolFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFees is a free data retrieval call binding the contract method 0xd8dfcea0.
//
// Solidity: function getProtocolFees() view returns(uint256)
func (_WEVQueue *WEVQueueSession) GetProtocolFees() (*big.Int, error) {
	return _WEVQueue.Contract.GetProtocolFees(&_WEVQueue.CallOpts)
}

// GetProtocolFees is a free data retrieval call binding the contract method 0xd8dfcea0.
//
// Solidity: function getProtocolFees() view returns(uint256)
func (_WEVQueue *WEVQueueCallerSession) GetProtocolFees() (*big.Int, error) {
	return _WEVQueue.Contract.GetProtocolFees(&_WEVQueue.CallOpts)
}

// GetQueueDepth is a free data retrieval call binding the contract method 0x7e58bf1e.
//
// Solidity: function getQueueDepth() view returns(uint256)
func (_WEVQueue *WEVQueueCaller) GetQueueDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "getQueueDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQueueDepth is a free data retrieval call binding the contract method 0x7e58bf1e.
//
// Solidity: function getQueueDepth() view returns(uint256)
func (_WEVQueue *WEVQueueSession) GetQueueDepth() (*big.Int, error) {
	return _WEVQueue.Contract.GetQueueDepth(&_WEVQueue.CallOpts)
}

// GetQueueDepth is a free data retrieval call binding the contract method 0x7e58bf1e.
//
// Solidity: function getQueueDepth() view returns(uint256)
func (_WEVQueue *WEVQueueCallerSession) GetQueueDepth() (*big.Int, error) {
	return _WEVQueue.Contract.GetQueueDepth(&_WEVQueue.CallOpts)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 requestId) view returns((bytes32,address,uint256,uint256,uint256,uint8))
func (_WEVQueue *WEVQueueCaller) GetRequest(opts *bind.CallOpts, requestId [32]byte) (IWEVQueueRedemptionRequest, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "getRequest", requestId)

	if err != nil {
		return *new(IWEVQueueRedemptionRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IWEVQueueRedemptionRequest)).(*IWEVQueueRedemptionRequest)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 requestId) view returns((bytes32,address,uint256,uint256,uint256,uint8))
func (_WEVQueue *WEVQueueSession) GetRequest(requestId [32]byte) (IWEVQueueRedemptionRequest, error) {
	return _WEVQueue.Contract.GetRequest(&_WEVQueue.CallOpts, requestId)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 requestId) view returns((bytes32,address,uint256,uint256,uint256,uint8))
func (_WEVQueue *WEVQueueCallerSession) GetRequest(requestId [32]byte) (IWEVQueueRedemptionRequest, error) {
	return _WEVQueue.Contract.GetRequest(&_WEVQueue.CallOpts, requestId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WEVQueue *WEVQueueCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WEVQueue *WEVQueueSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _WEVQueue.Contract.GetRoleAdmin(&_WEVQueue.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WEVQueue *WEVQueueCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _WEVQueue.Contract.GetRoleAdmin(&_WEVQueue.CallOpts, role)
}

// GetUserRequests is a free data retrieval call binding the contract method 0xb337cf74.
//
// Solidity: function getUserRequests(address user) view returns(bytes32[])
func (_WEVQueue *WEVQueueCaller) GetUserRequests(opts *bind.CallOpts, user common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "getUserRequests", user)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetUserRequests is a free data retrieval call binding the contract method 0xb337cf74.
//
// Solidity: function getUserRequests(address user) view returns(bytes32[])
func (_WEVQueue *WEVQueueSession) GetUserRequests(user common.Address) ([][32]byte, error) {
	return _WEVQueue.Contract.GetUserRequests(&_WEVQueue.CallOpts, user)
}

// GetUserRequests is a free data retrieval call binding the contract method 0xb337cf74.
//
// Solidity: function getUserRequests(address user) view returns(bytes32[])
func (_WEVQueue *WEVQueueCallerSession) GetUserRequests(user common.Address) ([][32]byte, error) {
	return _WEVQueue.Contract.GetUserRequests(&_WEVQueue.CallOpts, user)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WEVQueue *WEVQueueCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WEVQueue *WEVQueueSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _WEVQueue.Contract.HasRole(&_WEVQueue.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WEVQueue *WEVQueueCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _WEVQueue.Contract.HasRole(&_WEVQueue.CallOpts, role, account)
}

// NextProcessingTimestamp is a free data retrieval call binding the contract method 0x50583217.
//
// Solidity: function nextProcessingTimestamp() view returns(uint256)
func (_WEVQueue *WEVQueueCaller) NextProcessingTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "nextProcessingTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextProcessingTimestamp is a free data retrieval call binding the contract method 0x50583217.
//
// Solidity: function nextProcessingTimestamp() view returns(uint256)
func (_WEVQueue *WEVQueueSession) NextProcessingTimestamp() (*big.Int, error) {
	return _WEVQueue.Contract.NextProcessingTimestamp(&_WEVQueue.CallOpts)
}

// NextProcessingTimestamp is a free data retrieval call binding the contract method 0x50583217.
//
// Solidity: function nextProcessingTimestamp() view returns(uint256)
func (_WEVQueue *WEVQueueCallerSession) NextProcessingTimestamp() (*big.Int, error) {
	return _WEVQueue.Contract.NextProcessingTimestamp(&_WEVQueue.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WEVQueue *WEVQueueCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WEVQueue *WEVQueueSession) Paused() (bool, error) {
	return _WEVQueue.Contract.Paused(&_WEVQueue.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WEVQueue *WEVQueueCallerSession) Paused() (bool, error) {
	return _WEVQueue.Contract.Paused(&_WEVQueue.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WEVQueue *WEVQueueCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WEVQueue *WEVQueueSession) ProxiableUUID() ([32]byte, error) {
	return _WEVQueue.Contract.ProxiableUUID(&_WEVQueue.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WEVQueue *WEVQueueCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WEVQueue.Contract.ProxiableUUID(&_WEVQueue.CallOpts)
}

// SWattUSD is a free data retrieval call binding the contract method 0xe02fd053.
//
// Solidity: function sWattUSD() view returns(address)
func (_WEVQueue *WEVQueueCaller) SWattUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "sWattUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SWattUSD is a free data retrieval call binding the contract method 0xe02fd053.
//
// Solidity: function sWattUSD() view returns(address)
func (_WEVQueue *WEVQueueSession) SWattUSD() (common.Address, error) {
	return _WEVQueue.Contract.SWattUSD(&_WEVQueue.CallOpts)
}

// SWattUSD is a free data retrieval call binding the contract method 0xe02fd053.
//
// Solidity: function sWattUSD() view returns(address)
func (_WEVQueue *WEVQueueCallerSession) SWattUSD() (common.Address, error) {
	return _WEVQueue.Contract.SWattUSD(&_WEVQueue.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WEVQueue *WEVQueueCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WEVQueue *WEVQueueSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WEVQueue.Contract.SupportsInterface(&_WEVQueue.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WEVQueue *WEVQueueCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WEVQueue.Contract.SupportsInterface(&_WEVQueue.CallOpts, interfaceId)
}

// WattUSD is a free data retrieval call binding the contract method 0x7aff9566.
//
// Solidity: function wattUSD() view returns(address)
func (_WEVQueue *WEVQueueCaller) WattUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WEVQueue.contract.Call(opts, &out, "wattUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WattUSD is a free data retrieval call binding the contract method 0x7aff9566.
//
// Solidity: function wattUSD() view returns(address)
func (_WEVQueue *WEVQueueSession) WattUSD() (common.Address, error) {
	return _WEVQueue.Contract.WattUSD(&_WEVQueue.CallOpts)
}

// WattUSD is a free data retrieval call binding the contract method 0x7aff9566.
//
// Solidity: function wattUSD() view returns(address)
func (_WEVQueue *WEVQueueCallerSession) WattUSD() (common.Address, error) {
	return _WEVQueue.Contract.WattUSD(&_WEVQueue.CallOpts)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x50125546.
//
// Solidity: function cancelRequest(bytes32 requestId) returns()
func (_WEVQueue *WEVQueueTransactor) CancelRequest(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "cancelRequest", requestId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x50125546.
//
// Solidity: function cancelRequest(bytes32 requestId) returns()
func (_WEVQueue *WEVQueueSession) CancelRequest(requestId [32]byte) (*types.Transaction, error) {
	return _WEVQueue.Contract.CancelRequest(&_WEVQueue.TransactOpts, requestId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x50125546.
//
// Solidity: function cancelRequest(bytes32 requestId) returns()
func (_WEVQueue *WEVQueueTransactorSession) CancelRequest(requestId [32]byte) (*types.Transaction, error) {
	return _WEVQueue.Contract.CancelRequest(&_WEVQueue.TransactOpts, requestId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WEVQueue *WEVQueueTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WEVQueue *WEVQueueSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.GrantRole(&_WEVQueue.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WEVQueue *WEVQueueTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.GrantRole(&_WEVQueue.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address admin, address sWattUSD_, address wattUSD_) returns()
func (_WEVQueue *WEVQueueTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, sWattUSD_ common.Address, wattUSD_ common.Address) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "initialize", admin, sWattUSD_, wattUSD_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address admin, address sWattUSD_, address wattUSD_) returns()
func (_WEVQueue *WEVQueueSession) Initialize(admin common.Address, sWattUSD_ common.Address, wattUSD_ common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.Initialize(&_WEVQueue.TransactOpts, admin, sWattUSD_, wattUSD_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address admin, address sWattUSD_, address wattUSD_) returns()
func (_WEVQueue *WEVQueueTransactorSession) Initialize(admin common.Address, sWattUSD_ common.Address, wattUSD_ common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.Initialize(&_WEVQueue.TransactOpts, admin, sWattUSD_, wattUSD_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WEVQueue *WEVQueueTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WEVQueue *WEVQueueSession) Pause() (*types.Transaction, error) {
	return _WEVQueue.Contract.Pause(&_WEVQueue.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WEVQueue *WEVQueueTransactorSession) Pause() (*types.Transaction, error) {
	return _WEVQueue.Contract.Pause(&_WEVQueue.TransactOpts)
}

// ProcessBatch is a paid mutator transaction binding the contract method 0xea5f34e4.
//
// Solidity: function processBatch(bytes32[] requestIds) returns()
func (_WEVQueue *WEVQueueTransactor) ProcessBatch(opts *bind.TransactOpts, requestIds [][32]byte) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "processBatch", requestIds)
}

// ProcessBatch is a paid mutator transaction binding the contract method 0xea5f34e4.
//
// Solidity: function processBatch(bytes32[] requestIds) returns()
func (_WEVQueue *WEVQueueSession) ProcessBatch(requestIds [][32]byte) (*types.Transaction, error) {
	return _WEVQueue.Contract.ProcessBatch(&_WEVQueue.TransactOpts, requestIds)
}

// ProcessBatch is a paid mutator transaction binding the contract method 0xea5f34e4.
//
// Solidity: function processBatch(bytes32[] requestIds) returns()
func (_WEVQueue *WEVQueueTransactorSession) ProcessBatch(requestIds [][32]byte) (*types.Transaction, error) {
	return _WEVQueue.Contract.ProcessBatch(&_WEVQueue.TransactOpts, requestIds)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_WEVQueue *WEVQueueTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_WEVQueue *WEVQueueSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.RenounceRole(&_WEVQueue.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_WEVQueue *WEVQueueTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.RenounceRole(&_WEVQueue.TransactOpts, role, callerConfirmation)
}

// RequestPriorityRedeem is a paid mutator transaction binding the contract method 0xe07f435c.
//
// Solidity: function requestPriorityRedeem(uint256 sWattAmount, uint256 priorityFee) returns(bytes32 requestId)
func (_WEVQueue *WEVQueueTransactor) RequestPriorityRedeem(opts *bind.TransactOpts, sWattAmount *big.Int, priorityFee *big.Int) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "requestPriorityRedeem", sWattAmount, priorityFee)
}

// RequestPriorityRedeem is a paid mutator transaction binding the contract method 0xe07f435c.
//
// Solidity: function requestPriorityRedeem(uint256 sWattAmount, uint256 priorityFee) returns(bytes32 requestId)
func (_WEVQueue *WEVQueueSession) RequestPriorityRedeem(sWattAmount *big.Int, priorityFee *big.Int) (*types.Transaction, error) {
	return _WEVQueue.Contract.RequestPriorityRedeem(&_WEVQueue.TransactOpts, sWattAmount, priorityFee)
}

// RequestPriorityRedeem is a paid mutator transaction binding the contract method 0xe07f435c.
//
// Solidity: function requestPriorityRedeem(uint256 sWattAmount, uint256 priorityFee) returns(bytes32 requestId)
func (_WEVQueue *WEVQueueTransactorSession) RequestPriorityRedeem(sWattAmount *big.Int, priorityFee *big.Int) (*types.Transaction, error) {
	return _WEVQueue.Contract.RequestPriorityRedeem(&_WEVQueue.TransactOpts, sWattAmount, priorityFee)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0xaa2f892d.
//
// Solidity: function requestRedeem(uint256 sWattAmount) returns(bytes32 requestId)
func (_WEVQueue *WEVQueueTransactor) RequestRedeem(opts *bind.TransactOpts, sWattAmount *big.Int) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "requestRedeem", sWattAmount)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0xaa2f892d.
//
// Solidity: function requestRedeem(uint256 sWattAmount) returns(bytes32 requestId)
func (_WEVQueue *WEVQueueSession) RequestRedeem(sWattAmount *big.Int) (*types.Transaction, error) {
	return _WEVQueue.Contract.RequestRedeem(&_WEVQueue.TransactOpts, sWattAmount)
}

// RequestRedeem is a paid mutator transaction binding the contract method 0xaa2f892d.
//
// Solidity: function requestRedeem(uint256 sWattAmount) returns(bytes32 requestId)
func (_WEVQueue *WEVQueueTransactorSession) RequestRedeem(sWattAmount *big.Int) (*types.Transaction, error) {
	return _WEVQueue.Contract.RequestRedeem(&_WEVQueue.TransactOpts, sWattAmount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WEVQueue *WEVQueueTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WEVQueue *WEVQueueSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.RevokeRole(&_WEVQueue.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WEVQueue *WEVQueueTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.RevokeRole(&_WEVQueue.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WEVQueue *WEVQueueTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WEVQueue *WEVQueueSession) Unpause() (*types.Transaction, error) {
	return _WEVQueue.Contract.Unpause(&_WEVQueue.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WEVQueue *WEVQueueTransactorSession) Unpause() (*types.Transaction, error) {
	return _WEVQueue.Contract.Unpause(&_WEVQueue.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WEVQueue *WEVQueueTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WEVQueue *WEVQueueSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WEVQueue.Contract.UpgradeToAndCall(&_WEVQueue.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WEVQueue *WEVQueueTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WEVQueue.Contract.UpgradeToAndCall(&_WEVQueue.TransactOpts, newImplementation, data)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address to) returns()
func (_WEVQueue *WEVQueueTransactor) WithdrawFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _WEVQueue.contract.Transact(opts, "withdrawFees", to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address to) returns()
func (_WEVQueue *WEVQueueSession) WithdrawFees(to common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.WithdrawFees(&_WEVQueue.TransactOpts, to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address to) returns()
func (_WEVQueue *WEVQueueTransactorSession) WithdrawFees(to common.Address) (*types.Transaction, error) {
	return _WEVQueue.Contract.WithdrawFees(&_WEVQueue.TransactOpts, to)
}

// WEVQueueBatchProcessedIterator is returned from FilterBatchProcessed and is used to iterate over the raw logs and unpacked data for BatchProcessed events raised by the WEVQueue contract.
type WEVQueueBatchProcessedIterator struct {
	Event *WEVQueueBatchProcessed // Event containing the contract specifics and raw log

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
func (it *WEVQueueBatchProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueBatchProcessed)
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
		it.Event = new(WEVQueueBatchProcessed)
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
func (it *WEVQueueBatchProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueBatchProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueBatchProcessed represents a BatchProcessed event raised by the WEVQueue contract.
type WEVQueueBatchProcessed struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBatchProcessed is a free log retrieval operation binding the contract event 0xf563fdb95d8ba3f4669716ecc8559d8995059570a1324ffddf9f8094d2896c08.
//
// Solidity: event BatchProcessed(uint256 count)
func (_WEVQueue *WEVQueueFilterer) FilterBatchProcessed(opts *bind.FilterOpts) (*WEVQueueBatchProcessedIterator, error) {

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "BatchProcessed")
	if err != nil {
		return nil, err
	}
	return &WEVQueueBatchProcessedIterator{contract: _WEVQueue.contract, event: "BatchProcessed", logs: logs, sub: sub}, nil
}

// WatchBatchProcessed is a free log subscription operation binding the contract event 0xf563fdb95d8ba3f4669716ecc8559d8995059570a1324ffddf9f8094d2896c08.
//
// Solidity: event BatchProcessed(uint256 count)
func (_WEVQueue *WEVQueueFilterer) WatchBatchProcessed(opts *bind.WatchOpts, sink chan<- *WEVQueueBatchProcessed) (event.Subscription, error) {

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "BatchProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueBatchProcessed)
				if err := _WEVQueue.contract.UnpackLog(event, "BatchProcessed", log); err != nil {
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

// ParseBatchProcessed is a log parse operation binding the contract event 0xf563fdb95d8ba3f4669716ecc8559d8995059570a1324ffddf9f8094d2896c08.
//
// Solidity: event BatchProcessed(uint256 count)
func (_WEVQueue *WEVQueueFilterer) ParseBatchProcessed(log types.Log) (*WEVQueueBatchProcessed, error) {
	event := new(WEVQueueBatchProcessed)
	if err := _WEVQueue.contract.UnpackLog(event, "BatchProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueFeesWithdrawnIterator is returned from FilterFeesWithdrawn and is used to iterate over the raw logs and unpacked data for FeesWithdrawn events raised by the WEVQueue contract.
type WEVQueueFeesWithdrawnIterator struct {
	Event *WEVQueueFeesWithdrawn // Event containing the contract specifics and raw log

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
func (it *WEVQueueFeesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueFeesWithdrawn)
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
		it.Event = new(WEVQueueFeesWithdrawn)
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
func (it *WEVQueueFeesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueFeesWithdrawn represents a FeesWithdrawn event raised by the WEVQueue contract.
type WEVQueueFeesWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeesWithdrawn is a free log retrieval operation binding the contract event 0xc0819c13be868895eb93e40eaceb96de976442fa1d404e5c55f14bb65a8c489a.
//
// Solidity: event FeesWithdrawn(address indexed to, uint256 amount)
func (_WEVQueue *WEVQueueFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts, to []common.Address) (*WEVQueueFeesWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "FeesWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &WEVQueueFeesWithdrawnIterator{contract: _WEVQueue.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeesWithdrawn is a free log subscription operation binding the contract event 0xc0819c13be868895eb93e40eaceb96de976442fa1d404e5c55f14bb65a8c489a.
//
// Solidity: event FeesWithdrawn(address indexed to, uint256 amount)
func (_WEVQueue *WEVQueueFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *WEVQueueFeesWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "FeesWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueFeesWithdrawn)
				if err := _WEVQueue.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

// ParseFeesWithdrawn is a log parse operation binding the contract event 0xc0819c13be868895eb93e40eaceb96de976442fa1d404e5c55f14bb65a8c489a.
//
// Solidity: event FeesWithdrawn(address indexed to, uint256 amount)
func (_WEVQueue *WEVQueueFilterer) ParseFeesWithdrawn(log types.Log) (*WEVQueueFeesWithdrawn, error) {
	event := new(WEVQueueFeesWithdrawn)
	if err := _WEVQueue.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WEVQueue contract.
type WEVQueueInitializedIterator struct {
	Event *WEVQueueInitialized // Event containing the contract specifics and raw log

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
func (it *WEVQueueInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueInitialized)
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
		it.Event = new(WEVQueueInitialized)
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
func (it *WEVQueueInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueInitialized represents a Initialized event raised by the WEVQueue contract.
type WEVQueueInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_WEVQueue *WEVQueueFilterer) FilterInitialized(opts *bind.FilterOpts) (*WEVQueueInitializedIterator, error) {

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WEVQueueInitializedIterator{contract: _WEVQueue.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_WEVQueue *WEVQueueFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WEVQueueInitialized) (event.Subscription, error) {

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueInitialized)
				if err := _WEVQueue.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WEVQueue *WEVQueueFilterer) ParseInitialized(log types.Log) (*WEVQueueInitialized, error) {
	event := new(WEVQueueInitialized)
	if err := _WEVQueue.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueuePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WEVQueue contract.
type WEVQueuePausedIterator struct {
	Event *WEVQueuePaused // Event containing the contract specifics and raw log

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
func (it *WEVQueuePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueuePaused)
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
		it.Event = new(WEVQueuePaused)
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
func (it *WEVQueuePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueuePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueuePaused represents a Paused event raised by the WEVQueue contract.
type WEVQueuePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WEVQueue *WEVQueueFilterer) FilterPaused(opts *bind.FilterOpts) (*WEVQueuePausedIterator, error) {

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WEVQueuePausedIterator{contract: _WEVQueue.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WEVQueue *WEVQueueFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WEVQueuePaused) (event.Subscription, error) {

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueuePaused)
				if err := _WEVQueue.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_WEVQueue *WEVQueueFilterer) ParsePaused(log types.Log) (*WEVQueuePaused, error) {
	event := new(WEVQueuePaused)
	if err := _WEVQueue.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueRedemptionCancelledIterator is returned from FilterRedemptionCancelled and is used to iterate over the raw logs and unpacked data for RedemptionCancelled events raised by the WEVQueue contract.
type WEVQueueRedemptionCancelledIterator struct {
	Event *WEVQueueRedemptionCancelled // Event containing the contract specifics and raw log

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
func (it *WEVQueueRedemptionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueRedemptionCancelled)
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
		it.Event = new(WEVQueueRedemptionCancelled)
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
func (it *WEVQueueRedemptionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueRedemptionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueRedemptionCancelled represents a RedemptionCancelled event raised by the WEVQueue contract.
type WEVQueueRedemptionCancelled struct {
	RequestId [32]byte
	User      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedemptionCancelled is a free log retrieval operation binding the contract event 0x664d64fc839e6d517a3d20013509b20b03ae2ac797bff6f3845fdb96ed68667b.
//
// Solidity: event RedemptionCancelled(bytes32 indexed requestId, address indexed user)
func (_WEVQueue *WEVQueueFilterer) FilterRedemptionCancelled(opts *bind.FilterOpts, requestId [][32]byte, user []common.Address) (*WEVQueueRedemptionCancelledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "RedemptionCancelled", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &WEVQueueRedemptionCancelledIterator{contract: _WEVQueue.contract, event: "RedemptionCancelled", logs: logs, sub: sub}, nil
}

// WatchRedemptionCancelled is a free log subscription operation binding the contract event 0x664d64fc839e6d517a3d20013509b20b03ae2ac797bff6f3845fdb96ed68667b.
//
// Solidity: event RedemptionCancelled(bytes32 indexed requestId, address indexed user)
func (_WEVQueue *WEVQueueFilterer) WatchRedemptionCancelled(opts *bind.WatchOpts, sink chan<- *WEVQueueRedemptionCancelled, requestId [][32]byte, user []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "RedemptionCancelled", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueRedemptionCancelled)
				if err := _WEVQueue.contract.UnpackLog(event, "RedemptionCancelled", log); err != nil {
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

// ParseRedemptionCancelled is a log parse operation binding the contract event 0x664d64fc839e6d517a3d20013509b20b03ae2ac797bff6f3845fdb96ed68667b.
//
// Solidity: event RedemptionCancelled(bytes32 indexed requestId, address indexed user)
func (_WEVQueue *WEVQueueFilterer) ParseRedemptionCancelled(log types.Log) (*WEVQueueRedemptionCancelled, error) {
	event := new(WEVQueueRedemptionCancelled)
	if err := _WEVQueue.contract.UnpackLog(event, "RedemptionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueRedemptionFulfilledIterator is returned from FilterRedemptionFulfilled and is used to iterate over the raw logs and unpacked data for RedemptionFulfilled events raised by the WEVQueue contract.
type WEVQueueRedemptionFulfilledIterator struct {
	Event *WEVQueueRedemptionFulfilled // Event containing the contract specifics and raw log

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
func (it *WEVQueueRedemptionFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueRedemptionFulfilled)
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
		it.Event = new(WEVQueueRedemptionFulfilled)
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
func (it *WEVQueueRedemptionFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueRedemptionFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueRedemptionFulfilled represents a RedemptionFulfilled event raised by the WEVQueue contract.
type WEVQueueRedemptionFulfilled struct {
	RequestId  [32]byte
	User       common.Address
	WattAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedemptionFulfilled is a free log retrieval operation binding the contract event 0xe78451b7eb8009cda53bb73d9ff26e834d7af4e0cb1107f7ae6d055b07559388.
//
// Solidity: event RedemptionFulfilled(bytes32 indexed requestId, address indexed user, uint256 wattAmount)
func (_WEVQueue *WEVQueueFilterer) FilterRedemptionFulfilled(opts *bind.FilterOpts, requestId [][32]byte, user []common.Address) (*WEVQueueRedemptionFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "RedemptionFulfilled", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &WEVQueueRedemptionFulfilledIterator{contract: _WEVQueue.contract, event: "RedemptionFulfilled", logs: logs, sub: sub}, nil
}

// WatchRedemptionFulfilled is a free log subscription operation binding the contract event 0xe78451b7eb8009cda53bb73d9ff26e834d7af4e0cb1107f7ae6d055b07559388.
//
// Solidity: event RedemptionFulfilled(bytes32 indexed requestId, address indexed user, uint256 wattAmount)
func (_WEVQueue *WEVQueueFilterer) WatchRedemptionFulfilled(opts *bind.WatchOpts, sink chan<- *WEVQueueRedemptionFulfilled, requestId [][32]byte, user []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "RedemptionFulfilled", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueRedemptionFulfilled)
				if err := _WEVQueue.contract.UnpackLog(event, "RedemptionFulfilled", log); err != nil {
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

// ParseRedemptionFulfilled is a log parse operation binding the contract event 0xe78451b7eb8009cda53bb73d9ff26e834d7af4e0cb1107f7ae6d055b07559388.
//
// Solidity: event RedemptionFulfilled(bytes32 indexed requestId, address indexed user, uint256 wattAmount)
func (_WEVQueue *WEVQueueFilterer) ParseRedemptionFulfilled(log types.Log) (*WEVQueueRedemptionFulfilled, error) {
	event := new(WEVQueueRedemptionFulfilled)
	if err := _WEVQueue.contract.UnpackLog(event, "RedemptionFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueRedemptionRequestedIterator is returned from FilterRedemptionRequested and is used to iterate over the raw logs and unpacked data for RedemptionRequested events raised by the WEVQueue contract.
type WEVQueueRedemptionRequestedIterator struct {
	Event *WEVQueueRedemptionRequested // Event containing the contract specifics and raw log

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
func (it *WEVQueueRedemptionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueRedemptionRequested)
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
		it.Event = new(WEVQueueRedemptionRequested)
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
func (it *WEVQueueRedemptionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueRedemptionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueRedemptionRequested represents a RedemptionRequested event raised by the WEVQueue contract.
type WEVQueueRedemptionRequested struct {
	RequestId   [32]byte
	User        common.Address
	SWattAmount *big.Int
	IsPriority  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRedemptionRequested is a free log retrieval operation binding the contract event 0xc6e9579464219b3f93587b813bd7a320bae9f733c5af1e940e35fbc1cbd63928.
//
// Solidity: event RedemptionRequested(bytes32 indexed requestId, address indexed user, uint256 sWattAmount, bool isPriority)
func (_WEVQueue *WEVQueueFilterer) FilterRedemptionRequested(opts *bind.FilterOpts, requestId [][32]byte, user []common.Address) (*WEVQueueRedemptionRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "RedemptionRequested", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &WEVQueueRedemptionRequestedIterator{contract: _WEVQueue.contract, event: "RedemptionRequested", logs: logs, sub: sub}, nil
}

// WatchRedemptionRequested is a free log subscription operation binding the contract event 0xc6e9579464219b3f93587b813bd7a320bae9f733c5af1e940e35fbc1cbd63928.
//
// Solidity: event RedemptionRequested(bytes32 indexed requestId, address indexed user, uint256 sWattAmount, bool isPriority)
func (_WEVQueue *WEVQueueFilterer) WatchRedemptionRequested(opts *bind.WatchOpts, sink chan<- *WEVQueueRedemptionRequested, requestId [][32]byte, user []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "RedemptionRequested", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueRedemptionRequested)
				if err := _WEVQueue.contract.UnpackLog(event, "RedemptionRequested", log); err != nil {
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

// ParseRedemptionRequested is a log parse operation binding the contract event 0xc6e9579464219b3f93587b813bd7a320bae9f733c5af1e940e35fbc1cbd63928.
//
// Solidity: event RedemptionRequested(bytes32 indexed requestId, address indexed user, uint256 sWattAmount, bool isPriority)
func (_WEVQueue *WEVQueueFilterer) ParseRedemptionRequested(log types.Log) (*WEVQueueRedemptionRequested, error) {
	event := new(WEVQueueRedemptionRequested)
	if err := _WEVQueue.contract.UnpackLog(event, "RedemptionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the WEVQueue contract.
type WEVQueueRoleAdminChangedIterator struct {
	Event *WEVQueueRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *WEVQueueRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueRoleAdminChanged)
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
		it.Event = new(WEVQueueRoleAdminChanged)
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
func (it *WEVQueueRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueRoleAdminChanged represents a RoleAdminChanged event raised by the WEVQueue contract.
type WEVQueueRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_WEVQueue *WEVQueueFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*WEVQueueRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &WEVQueueRoleAdminChangedIterator{contract: _WEVQueue.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_WEVQueue *WEVQueueFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *WEVQueueRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueRoleAdminChanged)
				if err := _WEVQueue.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_WEVQueue *WEVQueueFilterer) ParseRoleAdminChanged(log types.Log) (*WEVQueueRoleAdminChanged, error) {
	event := new(WEVQueueRoleAdminChanged)
	if err := _WEVQueue.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the WEVQueue contract.
type WEVQueueRoleGrantedIterator struct {
	Event *WEVQueueRoleGranted // Event containing the contract specifics and raw log

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
func (it *WEVQueueRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueRoleGranted)
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
		it.Event = new(WEVQueueRoleGranted)
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
func (it *WEVQueueRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueRoleGranted represents a RoleGranted event raised by the WEVQueue contract.
type WEVQueueRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_WEVQueue *WEVQueueFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*WEVQueueRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WEVQueueRoleGrantedIterator{contract: _WEVQueue.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_WEVQueue *WEVQueueFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *WEVQueueRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueRoleGranted)
				if err := _WEVQueue.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_WEVQueue *WEVQueueFilterer) ParseRoleGranted(log types.Log) (*WEVQueueRoleGranted, error) {
	event := new(WEVQueueRoleGranted)
	if err := _WEVQueue.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the WEVQueue contract.
type WEVQueueRoleRevokedIterator struct {
	Event *WEVQueueRoleRevoked // Event containing the contract specifics and raw log

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
func (it *WEVQueueRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueRoleRevoked)
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
		it.Event = new(WEVQueueRoleRevoked)
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
func (it *WEVQueueRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueRoleRevoked represents a RoleRevoked event raised by the WEVQueue contract.
type WEVQueueRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_WEVQueue *WEVQueueFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*WEVQueueRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WEVQueueRoleRevokedIterator{contract: _WEVQueue.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_WEVQueue *WEVQueueFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *WEVQueueRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueRoleRevoked)
				if err := _WEVQueue.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_WEVQueue *WEVQueueFilterer) ParseRoleRevoked(log types.Log) (*WEVQueueRoleRevoked, error) {
	event := new(WEVQueueRoleRevoked)
	if err := _WEVQueue.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WEVQueue contract.
type WEVQueueUnpausedIterator struct {
	Event *WEVQueueUnpaused // Event containing the contract specifics and raw log

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
func (it *WEVQueueUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueUnpaused)
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
		it.Event = new(WEVQueueUnpaused)
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
func (it *WEVQueueUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueUnpaused represents a Unpaused event raised by the WEVQueue contract.
type WEVQueueUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WEVQueue *WEVQueueFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WEVQueueUnpausedIterator, error) {

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WEVQueueUnpausedIterator{contract: _WEVQueue.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WEVQueue *WEVQueueFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WEVQueueUnpaused) (event.Subscription, error) {

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueUnpaused)
				if err := _WEVQueue.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_WEVQueue *WEVQueueFilterer) ParseUnpaused(log types.Log) (*WEVQueueUnpaused, error) {
	event := new(WEVQueueUnpaused)
	if err := _WEVQueue.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEVQueueUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WEVQueue contract.
type WEVQueueUpgradedIterator struct {
	Event *WEVQueueUpgraded // Event containing the contract specifics and raw log

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
func (it *WEVQueueUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEVQueueUpgraded)
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
		it.Event = new(WEVQueueUpgraded)
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
func (it *WEVQueueUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEVQueueUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEVQueueUpgraded represents a Upgraded event raised by the WEVQueue contract.
type WEVQueueUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WEVQueue *WEVQueueFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WEVQueueUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WEVQueue.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WEVQueueUpgradedIterator{contract: _WEVQueue.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WEVQueue *WEVQueueFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WEVQueueUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WEVQueue.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEVQueueUpgraded)
				if err := _WEVQueue.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_WEVQueue *WEVQueueFilterer) ParseUpgraded(log types.Log) (*WEVQueueUpgraded, error) {
	event := new(WEVQueueUpgraded)
	if err := _WEVQueue.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
