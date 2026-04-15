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

// IAssetRegistryAsset is an auto generated low-level Go binding around an user-defined struct.
type IAssetRegistryAsset struct {
	AssetId      [32]byte
	AssetType    uint8
	Borrower     common.Address
	Ltv          uint16
	Status       uint8
	RegisteredAt *big.Int
}

// AssetRegistryMetaData contains all meta data concerning the AssetRegistry contract.
var AssetRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"AssetAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"AssetNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"}],\"name\":\"InvalidLTV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumIAssetRegistry.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"}],\"name\":\"AssetRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newLTV\",\"type\":\"uint16\"}],\"name\":\"LTVUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumIAssetRegistry.AssetStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumIAssetRegistry.AssetStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"StatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LENDINGPOOL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LTV\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRAR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"getAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIAssetRegistry.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"enumIAssetRegistry.AssetStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"registeredAt\",\"type\":\"uint256\"}],\"internalType\":\"structIAssetRegistry.Asset\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIAssetRegistry.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"}],\"name\":\"registerAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"newLTV\",\"type\":\"uint16\"}],\"name\":\"updateLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIAssetRegistry.AssetStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updateStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AssetRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetRegistryMetaData.ABI instead.
var AssetRegistryABI = AssetRegistryMetaData.ABI

// AssetRegistry is an auto generated Go binding around an Ethereum contract.
type AssetRegistry struct {
	AssetRegistryCaller     // Read-only binding to the contract
	AssetRegistryTransactor // Write-only binding to the contract
	AssetRegistryFilterer   // Log filterer for contract events
}

// AssetRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetRegistrySession struct {
	Contract     *AssetRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetRegistryCallerSession struct {
	Contract *AssetRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AssetRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetRegistryTransactorSession struct {
	Contract     *AssetRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AssetRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetRegistryRaw struct {
	Contract *AssetRegistry // Generic contract binding to access the raw methods on
}

// AssetRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetRegistryCallerRaw struct {
	Contract *AssetRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AssetRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetRegistryTransactorRaw struct {
	Contract *AssetRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetRegistry creates a new instance of AssetRegistry, bound to a specific deployed contract.
func NewAssetRegistry(address common.Address, backend bind.ContractBackend) (*AssetRegistry, error) {
	contract, err := bindAssetRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetRegistry{AssetRegistryCaller: AssetRegistryCaller{contract: contract}, AssetRegistryTransactor: AssetRegistryTransactor{contract: contract}, AssetRegistryFilterer: AssetRegistryFilterer{contract: contract}}, nil
}

// NewAssetRegistryCaller creates a new read-only instance of AssetRegistry, bound to a specific deployed contract.
func NewAssetRegistryCaller(address common.Address, caller bind.ContractCaller) (*AssetRegistryCaller, error) {
	contract, err := bindAssetRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryCaller{contract: contract}, nil
}

// NewAssetRegistryTransactor creates a new write-only instance of AssetRegistry, bound to a specific deployed contract.
func NewAssetRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetRegistryTransactor, error) {
	contract, err := bindAssetRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryTransactor{contract: contract}, nil
}

// NewAssetRegistryFilterer creates a new log filterer instance of AssetRegistry, bound to a specific deployed contract.
func NewAssetRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetRegistryFilterer, error) {
	contract, err := bindAssetRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryFilterer{contract: contract}, nil
}

// bindAssetRegistry binds a generic wrapper to an already deployed contract.
func bindAssetRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssetRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetRegistry *AssetRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetRegistry.Contract.AssetRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetRegistry *AssetRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetRegistry.Contract.AssetRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetRegistry *AssetRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetRegistry.Contract.AssetRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetRegistry *AssetRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetRegistry *AssetRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetRegistry *AssetRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetRegistry.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistrySession) ADMINROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.ADMINROLE(&_AssetRegistry.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCallerSession) ADMINROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.ADMINROLE(&_AssetRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.DEFAULTADMINROLE(&_AssetRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.DEFAULTADMINROLE(&_AssetRegistry.CallOpts)
}

// LENDINGPOOLROLE is a free data retrieval call binding the contract method 0xecdf56cd.
//
// Solidity: function LENDINGPOOL_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCaller) LENDINGPOOLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "LENDINGPOOL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LENDINGPOOLROLE is a free data retrieval call binding the contract method 0xecdf56cd.
//
// Solidity: function LENDINGPOOL_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistrySession) LENDINGPOOLROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.LENDINGPOOLROLE(&_AssetRegistry.CallOpts)
}

// LENDINGPOOLROLE is a free data retrieval call binding the contract method 0xecdf56cd.
//
// Solidity: function LENDINGPOOL_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCallerSession) LENDINGPOOLROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.LENDINGPOOLROLE(&_AssetRegistry.CallOpts)
}

// MAXLTV is a free data retrieval call binding the contract method 0x81bf8d3d.
//
// Solidity: function MAX_LTV() view returns(uint16)
func (_AssetRegistry *AssetRegistryCaller) MAXLTV(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "MAX_LTV")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXLTV is a free data retrieval call binding the contract method 0x81bf8d3d.
//
// Solidity: function MAX_LTV() view returns(uint16)
func (_AssetRegistry *AssetRegistrySession) MAXLTV() (uint16, error) {
	return _AssetRegistry.Contract.MAXLTV(&_AssetRegistry.CallOpts)
}

// MAXLTV is a free data retrieval call binding the contract method 0x81bf8d3d.
//
// Solidity: function MAX_LTV() view returns(uint16)
func (_AssetRegistry *AssetRegistryCallerSession) MAXLTV() (uint16, error) {
	return _AssetRegistry.Contract.MAXLTV(&_AssetRegistry.CallOpts)
}

// REGISTRARROLE is a free data retrieval call binding the contract method 0xf68e9553.
//
// Solidity: function REGISTRAR_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCaller) REGISTRARROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "REGISTRAR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRARROLE is a free data retrieval call binding the contract method 0xf68e9553.
//
// Solidity: function REGISTRAR_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistrySession) REGISTRARROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.REGISTRARROLE(&_AssetRegistry.CallOpts)
}

// REGISTRARROLE is a free data retrieval call binding the contract method 0xf68e9553.
//
// Solidity: function REGISTRAR_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCallerSession) REGISTRARROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.REGISTRARROLE(&_AssetRegistry.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistrySession) UPGRADERROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.UPGRADERROLE(&_AssetRegistry.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _AssetRegistry.Contract.UPGRADERROLE(&_AssetRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AssetRegistry *AssetRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AssetRegistry *AssetRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AssetRegistry.Contract.UPGRADEINTERFACEVERSION(&_AssetRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AssetRegistry *AssetRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AssetRegistry.Contract.UPGRADEINTERFACEVERSION(&_AssetRegistry.CallOpts)
}

// GetAsset is a free data retrieval call binding the contract method 0x2cc3ce80.
//
// Solidity: function getAsset(bytes32 assetId) view returns((bytes32,uint8,address,uint16,uint8,uint256))
func (_AssetRegistry *AssetRegistryCaller) GetAsset(opts *bind.CallOpts, assetId [32]byte) (IAssetRegistryAsset, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "getAsset", assetId)

	if err != nil {
		return *new(IAssetRegistryAsset), err
	}

	out0 := *abi.ConvertType(out[0], new(IAssetRegistryAsset)).(*IAssetRegistryAsset)

	return out0, err

}

// GetAsset is a free data retrieval call binding the contract method 0x2cc3ce80.
//
// Solidity: function getAsset(bytes32 assetId) view returns((bytes32,uint8,address,uint16,uint8,uint256))
func (_AssetRegistry *AssetRegistrySession) GetAsset(assetId [32]byte) (IAssetRegistryAsset, error) {
	return _AssetRegistry.Contract.GetAsset(&_AssetRegistry.CallOpts, assetId)
}

// GetAsset is a free data retrieval call binding the contract method 0x2cc3ce80.
//
// Solidity: function getAsset(bytes32 assetId) view returns((bytes32,uint8,address,uint16,uint8,uint256))
func (_AssetRegistry *AssetRegistryCallerSession) GetAsset(assetId [32]byte) (IAssetRegistryAsset, error) {
	return _AssetRegistry.Contract.GetAsset(&_AssetRegistry.CallOpts, assetId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetRegistry *AssetRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetRegistry *AssetRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetRegistry.Contract.GetRoleAdmin(&_AssetRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetRegistry *AssetRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetRegistry.Contract.GetRoleAdmin(&_AssetRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetRegistry *AssetRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetRegistry *AssetRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetRegistry.Contract.HasRole(&_AssetRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetRegistry *AssetRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetRegistry.Contract.HasRole(&_AssetRegistry.CallOpts, role, account)
}

// IsActive is a free data retrieval call binding the contract method 0x5c36901c.
//
// Solidity: function isActive(bytes32 assetId) view returns(bool)
func (_AssetRegistry *AssetRegistryCaller) IsActive(opts *bind.CallOpts, assetId [32]byte) (bool, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "isActive", assetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x5c36901c.
//
// Solidity: function isActive(bytes32 assetId) view returns(bool)
func (_AssetRegistry *AssetRegistrySession) IsActive(assetId [32]byte) (bool, error) {
	return _AssetRegistry.Contract.IsActive(&_AssetRegistry.CallOpts, assetId)
}

// IsActive is a free data retrieval call binding the contract method 0x5c36901c.
//
// Solidity: function isActive(bytes32 assetId) view returns(bool)
func (_AssetRegistry *AssetRegistryCallerSession) IsActive(assetId [32]byte) (bool, error) {
	return _AssetRegistry.Contract.IsActive(&_AssetRegistry.CallOpts, assetId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AssetRegistry *AssetRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _AssetRegistry.Contract.ProxiableUUID(&_AssetRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AssetRegistry *AssetRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AssetRegistry.Contract.ProxiableUUID(&_AssetRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetRegistry *AssetRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AssetRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetRegistry *AssetRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetRegistry.Contract.SupportsInterface(&_AssetRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetRegistry *AssetRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetRegistry.Contract.SupportsInterface(&_AssetRegistry.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetRegistry *AssetRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetRegistry *AssetRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.GrantRole(&_AssetRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.GrantRole(&_AssetRegistry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_AssetRegistry *AssetRegistryTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_AssetRegistry *AssetRegistrySession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.Initialize(&_AssetRegistry.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.Initialize(&_AssetRegistry.TransactOpts, admin)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x823448c3.
//
// Solidity: function registerAsset(bytes32 assetId, uint8 assetType, address borrower, uint16 ltv) returns()
func (_AssetRegistry *AssetRegistryTransactor) RegisterAsset(opts *bind.TransactOpts, assetId [32]byte, assetType uint8, borrower common.Address, ltv uint16) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "registerAsset", assetId, assetType, borrower, ltv)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x823448c3.
//
// Solidity: function registerAsset(bytes32 assetId, uint8 assetType, address borrower, uint16 ltv) returns()
func (_AssetRegistry *AssetRegistrySession) RegisterAsset(assetId [32]byte, assetType uint8, borrower common.Address, ltv uint16) (*types.Transaction, error) {
	return _AssetRegistry.Contract.RegisterAsset(&_AssetRegistry.TransactOpts, assetId, assetType, borrower, ltv)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x823448c3.
//
// Solidity: function registerAsset(bytes32 assetId, uint8 assetType, address borrower, uint16 ltv) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) RegisterAsset(assetId [32]byte, assetType uint8, borrower common.Address, ltv uint16) (*types.Transaction, error) {
	return _AssetRegistry.Contract.RegisterAsset(&_AssetRegistry.TransactOpts, assetId, assetType, borrower, ltv)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AssetRegistry *AssetRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AssetRegistry *AssetRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.RenounceRole(&_AssetRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.RenounceRole(&_AssetRegistry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetRegistry *AssetRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetRegistry *AssetRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.RevokeRole(&_AssetRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetRegistry.Contract.RevokeRole(&_AssetRegistry.TransactOpts, role, account)
}

// UpdateLTV is a paid mutator transaction binding the contract method 0x3becd9b8.
//
// Solidity: function updateLTV(bytes32 assetId, uint16 newLTV) returns()
func (_AssetRegistry *AssetRegistryTransactor) UpdateLTV(opts *bind.TransactOpts, assetId [32]byte, newLTV uint16) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "updateLTV", assetId, newLTV)
}

// UpdateLTV is a paid mutator transaction binding the contract method 0x3becd9b8.
//
// Solidity: function updateLTV(bytes32 assetId, uint16 newLTV) returns()
func (_AssetRegistry *AssetRegistrySession) UpdateLTV(assetId [32]byte, newLTV uint16) (*types.Transaction, error) {
	return _AssetRegistry.Contract.UpdateLTV(&_AssetRegistry.TransactOpts, assetId, newLTV)
}

// UpdateLTV is a paid mutator transaction binding the contract method 0x3becd9b8.
//
// Solidity: function updateLTV(bytes32 assetId, uint16 newLTV) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) UpdateLTV(assetId [32]byte, newLTV uint16) (*types.Transaction, error) {
	return _AssetRegistry.Contract.UpdateLTV(&_AssetRegistry.TransactOpts, assetId, newLTV)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x054372ed.
//
// Solidity: function updateStatus(bytes32 assetId, uint8 newStatus) returns()
func (_AssetRegistry *AssetRegistryTransactor) UpdateStatus(opts *bind.TransactOpts, assetId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "updateStatus", assetId, newStatus)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x054372ed.
//
// Solidity: function updateStatus(bytes32 assetId, uint8 newStatus) returns()
func (_AssetRegistry *AssetRegistrySession) UpdateStatus(assetId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _AssetRegistry.Contract.UpdateStatus(&_AssetRegistry.TransactOpts, assetId, newStatus)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x054372ed.
//
// Solidity: function updateStatus(bytes32 assetId, uint8 newStatus) returns()
func (_AssetRegistry *AssetRegistryTransactorSession) UpdateStatus(assetId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _AssetRegistry.Contract.UpdateStatus(&_AssetRegistry.TransactOpts, assetId, newStatus)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AssetRegistry *AssetRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AssetRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AssetRegistry *AssetRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AssetRegistry.Contract.UpgradeToAndCall(&_AssetRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AssetRegistry *AssetRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AssetRegistry.Contract.UpgradeToAndCall(&_AssetRegistry.TransactOpts, newImplementation, data)
}

// AssetRegistryAssetRegisteredIterator is returned from FilterAssetRegistered and is used to iterate over the raw logs and unpacked data for AssetRegistered events raised by the AssetRegistry contract.
type AssetRegistryAssetRegisteredIterator struct {
	Event *AssetRegistryAssetRegistered // Event containing the contract specifics and raw log

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
func (it *AssetRegistryAssetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryAssetRegistered)
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
		it.Event = new(AssetRegistryAssetRegistered)
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
func (it *AssetRegistryAssetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryAssetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryAssetRegistered represents a AssetRegistered event raised by the AssetRegistry contract.
type AssetRegistryAssetRegistered struct {
	AssetId   [32]byte
	AssetType uint8
	Borrower  common.Address
	Ltv       uint16
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssetRegistered is a free log retrieval operation binding the contract event 0xdcbe5dbb7edec38c564395d08f42871432446fc51682f936bab0fb9714fb32f8.
//
// Solidity: event AssetRegistered(bytes32 indexed assetId, uint8 assetType, address indexed borrower, uint16 ltv)
func (_AssetRegistry *AssetRegistryFilterer) FilterAssetRegistered(opts *bind.FilterOpts, assetId [][32]byte, borrower []common.Address) (*AssetRegistryAssetRegisteredIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "AssetRegistered", assetIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryAssetRegisteredIterator{contract: _AssetRegistry.contract, event: "AssetRegistered", logs: logs, sub: sub}, nil
}

// WatchAssetRegistered is a free log subscription operation binding the contract event 0xdcbe5dbb7edec38c564395d08f42871432446fc51682f936bab0fb9714fb32f8.
//
// Solidity: event AssetRegistered(bytes32 indexed assetId, uint8 assetType, address indexed borrower, uint16 ltv)
func (_AssetRegistry *AssetRegistryFilterer) WatchAssetRegistered(opts *bind.WatchOpts, sink chan<- *AssetRegistryAssetRegistered, assetId [][32]byte, borrower []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "AssetRegistered", assetIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryAssetRegistered)
				if err := _AssetRegistry.contract.UnpackLog(event, "AssetRegistered", log); err != nil {
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

// ParseAssetRegistered is a log parse operation binding the contract event 0xdcbe5dbb7edec38c564395d08f42871432446fc51682f936bab0fb9714fb32f8.
//
// Solidity: event AssetRegistered(bytes32 indexed assetId, uint8 assetType, address indexed borrower, uint16 ltv)
func (_AssetRegistry *AssetRegistryFilterer) ParseAssetRegistered(log types.Log) (*AssetRegistryAssetRegistered, error) {
	event := new(AssetRegistryAssetRegistered)
	if err := _AssetRegistry.contract.UnpackLog(event, "AssetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AssetRegistry contract.
type AssetRegistryInitializedIterator struct {
	Event *AssetRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *AssetRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryInitialized)
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
		it.Event = new(AssetRegistryInitialized)
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
func (it *AssetRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryInitialized represents a Initialized event raised by the AssetRegistry contract.
type AssetRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AssetRegistry *AssetRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*AssetRegistryInitializedIterator, error) {

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AssetRegistryInitializedIterator{contract: _AssetRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AssetRegistry *AssetRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AssetRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryInitialized)
				if err := _AssetRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AssetRegistry *AssetRegistryFilterer) ParseInitialized(log types.Log) (*AssetRegistryInitialized, error) {
	event := new(AssetRegistryInitialized)
	if err := _AssetRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryLTVUpdatedIterator is returned from FilterLTVUpdated and is used to iterate over the raw logs and unpacked data for LTVUpdated events raised by the AssetRegistry contract.
type AssetRegistryLTVUpdatedIterator struct {
	Event *AssetRegistryLTVUpdated // Event containing the contract specifics and raw log

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
func (it *AssetRegistryLTVUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryLTVUpdated)
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
		it.Event = new(AssetRegistryLTVUpdated)
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
func (it *AssetRegistryLTVUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryLTVUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryLTVUpdated represents a LTVUpdated event raised by the AssetRegistry contract.
type AssetRegistryLTVUpdated struct {
	AssetId [32]byte
	OldLTV  uint16
	NewLTV  uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLTVUpdated is a free log retrieval operation binding the contract event 0xfcf41846e25bb604ecaa822df82486fe3a57d8bfb20248cb82389fa2d859bd46.
//
// Solidity: event LTVUpdated(bytes32 indexed assetId, uint16 oldLTV, uint16 newLTV)
func (_AssetRegistry *AssetRegistryFilterer) FilterLTVUpdated(opts *bind.FilterOpts, assetId [][32]byte) (*AssetRegistryLTVUpdatedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "LTVUpdated", assetIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryLTVUpdatedIterator{contract: _AssetRegistry.contract, event: "LTVUpdated", logs: logs, sub: sub}, nil
}

// WatchLTVUpdated is a free log subscription operation binding the contract event 0xfcf41846e25bb604ecaa822df82486fe3a57d8bfb20248cb82389fa2d859bd46.
//
// Solidity: event LTVUpdated(bytes32 indexed assetId, uint16 oldLTV, uint16 newLTV)
func (_AssetRegistry *AssetRegistryFilterer) WatchLTVUpdated(opts *bind.WatchOpts, sink chan<- *AssetRegistryLTVUpdated, assetId [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "LTVUpdated", assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryLTVUpdated)
				if err := _AssetRegistry.contract.UnpackLog(event, "LTVUpdated", log); err != nil {
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

// ParseLTVUpdated is a log parse operation binding the contract event 0xfcf41846e25bb604ecaa822df82486fe3a57d8bfb20248cb82389fa2d859bd46.
//
// Solidity: event LTVUpdated(bytes32 indexed assetId, uint16 oldLTV, uint16 newLTV)
func (_AssetRegistry *AssetRegistryFilterer) ParseLTVUpdated(log types.Log) (*AssetRegistryLTVUpdated, error) {
	event := new(AssetRegistryLTVUpdated)
	if err := _AssetRegistry.contract.UnpackLog(event, "LTVUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AssetRegistry contract.
type AssetRegistryRoleAdminChangedIterator struct {
	Event *AssetRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AssetRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryRoleAdminChanged)
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
		it.Event = new(AssetRegistryRoleAdminChanged)
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
func (it *AssetRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the AssetRegistry contract.
type AssetRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetRegistry *AssetRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AssetRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryRoleAdminChangedIterator{contract: _AssetRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetRegistry *AssetRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AssetRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryRoleAdminChanged)
				if err := _AssetRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AssetRegistry *AssetRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*AssetRegistryRoleAdminChanged, error) {
	event := new(AssetRegistryRoleAdminChanged)
	if err := _AssetRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AssetRegistry contract.
type AssetRegistryRoleGrantedIterator struct {
	Event *AssetRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *AssetRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryRoleGranted)
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
		it.Event = new(AssetRegistryRoleGranted)
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
func (it *AssetRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryRoleGranted represents a RoleGranted event raised by the AssetRegistry contract.
type AssetRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetRegistry *AssetRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryRoleGrantedIterator{contract: _AssetRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetRegistry *AssetRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AssetRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryRoleGranted)
				if err := _AssetRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AssetRegistry *AssetRegistryFilterer) ParseRoleGranted(log types.Log) (*AssetRegistryRoleGranted, error) {
	event := new(AssetRegistryRoleGranted)
	if err := _AssetRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AssetRegistry contract.
type AssetRegistryRoleRevokedIterator struct {
	Event *AssetRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AssetRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryRoleRevoked)
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
		it.Event = new(AssetRegistryRoleRevoked)
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
func (it *AssetRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryRoleRevoked represents a RoleRevoked event raised by the AssetRegistry contract.
type AssetRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetRegistry *AssetRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryRoleRevokedIterator{contract: _AssetRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetRegistry *AssetRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AssetRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryRoleRevoked)
				if err := _AssetRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AssetRegistry *AssetRegistryFilterer) ParseRoleRevoked(log types.Log) (*AssetRegistryRoleRevoked, error) {
	event := new(AssetRegistryRoleRevoked)
	if err := _AssetRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryStatusChangedIterator is returned from FilterStatusChanged and is used to iterate over the raw logs and unpacked data for StatusChanged events raised by the AssetRegistry contract.
type AssetRegistryStatusChangedIterator struct {
	Event *AssetRegistryStatusChanged // Event containing the contract specifics and raw log

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
func (it *AssetRegistryStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryStatusChanged)
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
		it.Event = new(AssetRegistryStatusChanged)
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
func (it *AssetRegistryStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryStatusChanged represents a StatusChanged event raised by the AssetRegistry contract.
type AssetRegistryStatusChanged struct {
	AssetId   [32]byte
	OldStatus uint8
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStatusChanged is a free log retrieval operation binding the contract event 0x4d7781468081641aba2c04c3349fcf5830b9fedac1b7aaffabc1f1dc6b8883fb.
//
// Solidity: event StatusChanged(bytes32 indexed assetId, uint8 oldStatus, uint8 newStatus)
func (_AssetRegistry *AssetRegistryFilterer) FilterStatusChanged(opts *bind.FilterOpts, assetId [][32]byte) (*AssetRegistryStatusChangedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "StatusChanged", assetIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryStatusChangedIterator{contract: _AssetRegistry.contract, event: "StatusChanged", logs: logs, sub: sub}, nil
}

// WatchStatusChanged is a free log subscription operation binding the contract event 0x4d7781468081641aba2c04c3349fcf5830b9fedac1b7aaffabc1f1dc6b8883fb.
//
// Solidity: event StatusChanged(bytes32 indexed assetId, uint8 oldStatus, uint8 newStatus)
func (_AssetRegistry *AssetRegistryFilterer) WatchStatusChanged(opts *bind.WatchOpts, sink chan<- *AssetRegistryStatusChanged, assetId [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "StatusChanged", assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryStatusChanged)
				if err := _AssetRegistry.contract.UnpackLog(event, "StatusChanged", log); err != nil {
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

// ParseStatusChanged is a log parse operation binding the contract event 0x4d7781468081641aba2c04c3349fcf5830b9fedac1b7aaffabc1f1dc6b8883fb.
//
// Solidity: event StatusChanged(bytes32 indexed assetId, uint8 oldStatus, uint8 newStatus)
func (_AssetRegistry *AssetRegistryFilterer) ParseStatusChanged(log types.Log) (*AssetRegistryStatusChanged, error) {
	event := new(AssetRegistryStatusChanged)
	if err := _AssetRegistry.contract.UnpackLog(event, "StatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AssetRegistry contract.
type AssetRegistryUpgradedIterator struct {
	Event *AssetRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *AssetRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegistryUpgraded)
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
		it.Event = new(AssetRegistryUpgraded)
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
func (it *AssetRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegistryUpgraded represents a Upgraded event raised by the AssetRegistry contract.
type AssetRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AssetRegistry *AssetRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AssetRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AssetRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AssetRegistryUpgradedIterator{contract: _AssetRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AssetRegistry *AssetRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AssetRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AssetRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegistryUpgraded)
				if err := _AssetRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AssetRegistry *AssetRegistryFilterer) ParseUpgraded(log types.Log) (*AssetRegistryUpgraded, error) {
	event := new(AssetRegistryUpgraded)
	if err := _AssetRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
