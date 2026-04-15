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

// IHealthAttestationAttestation is an auto generated low-level Go binding around an user-defined struct.
type IHealthAttestationAttestation struct {
	AssetId    [32]byte
	HealthHash [32]byte
	Score      uint8
	Timestamp  *big.Int
}

// HealthAttestationMetaData contains all meta data concerning the HealthAttestation contract.
var HealthAttestationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nextAllowed\",\"type\":\"uint256\"}],\"name\":\"AttestationTooSoon\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"score\",\"type\":\"uint8\"}],\"name\":\"InvalidScore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"healthHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"score\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AttestationSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COOLDOWN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFLOW_SIGNER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getAttestationHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"healthHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"score\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIHealthAttestation.Attestation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"getLatestAttestation\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"healthHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"score\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIHealthAttestation.Attestation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"hasAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"healthHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"score\",\"type\":\"uint8\"}],\"name\":\"submitAttestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// HealthAttestationABI is the input ABI used to generate the binding from.
// Deprecated: Use HealthAttestationMetaData.ABI instead.
var HealthAttestationABI = HealthAttestationMetaData.ABI

// HealthAttestation is an auto generated Go binding around an Ethereum contract.
type HealthAttestation struct {
	HealthAttestationCaller     // Read-only binding to the contract
	HealthAttestationTransactor // Write-only binding to the contract
	HealthAttestationFilterer   // Log filterer for contract events
}

// HealthAttestationCaller is an auto generated read-only Go binding around an Ethereum contract.
type HealthAttestationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthAttestationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HealthAttestationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthAttestationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HealthAttestationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthAttestationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HealthAttestationSession struct {
	Contract     *HealthAttestation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HealthAttestationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HealthAttestationCallerSession struct {
	Contract *HealthAttestationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// HealthAttestationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HealthAttestationTransactorSession struct {
	Contract     *HealthAttestationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// HealthAttestationRaw is an auto generated low-level Go binding around an Ethereum contract.
type HealthAttestationRaw struct {
	Contract *HealthAttestation // Generic contract binding to access the raw methods on
}

// HealthAttestationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HealthAttestationCallerRaw struct {
	Contract *HealthAttestationCaller // Generic read-only contract binding to access the raw methods on
}

// HealthAttestationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HealthAttestationTransactorRaw struct {
	Contract *HealthAttestationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHealthAttestation creates a new instance of HealthAttestation, bound to a specific deployed contract.
func NewHealthAttestation(address common.Address, backend bind.ContractBackend) (*HealthAttestation, error) {
	contract, err := bindHealthAttestation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HealthAttestation{HealthAttestationCaller: HealthAttestationCaller{contract: contract}, HealthAttestationTransactor: HealthAttestationTransactor{contract: contract}, HealthAttestationFilterer: HealthAttestationFilterer{contract: contract}}, nil
}

// NewHealthAttestationCaller creates a new read-only instance of HealthAttestation, bound to a specific deployed contract.
func NewHealthAttestationCaller(address common.Address, caller bind.ContractCaller) (*HealthAttestationCaller, error) {
	contract, err := bindHealthAttestation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HealthAttestationCaller{contract: contract}, nil
}

// NewHealthAttestationTransactor creates a new write-only instance of HealthAttestation, bound to a specific deployed contract.
func NewHealthAttestationTransactor(address common.Address, transactor bind.ContractTransactor) (*HealthAttestationTransactor, error) {
	contract, err := bindHealthAttestation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HealthAttestationTransactor{contract: contract}, nil
}

// NewHealthAttestationFilterer creates a new log filterer instance of HealthAttestation, bound to a specific deployed contract.
func NewHealthAttestationFilterer(address common.Address, filterer bind.ContractFilterer) (*HealthAttestationFilterer, error) {
	contract, err := bindHealthAttestation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HealthAttestationFilterer{contract: contract}, nil
}

// bindHealthAttestation binds a generic wrapper to an already deployed contract.
func bindHealthAttestation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HealthAttestationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HealthAttestation *HealthAttestationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HealthAttestation.Contract.HealthAttestationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HealthAttestation *HealthAttestationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HealthAttestation.Contract.HealthAttestationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HealthAttestation *HealthAttestationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HealthAttestation.Contract.HealthAttestationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HealthAttestation *HealthAttestationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HealthAttestation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HealthAttestation *HealthAttestationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HealthAttestation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HealthAttestation *HealthAttestationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HealthAttestation.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_HealthAttestation *HealthAttestationSession) ADMINROLE() ([32]byte, error) {
	return _HealthAttestation.Contract.ADMINROLE(&_HealthAttestation.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCallerSession) ADMINROLE() ([32]byte, error) {
	return _HealthAttestation.Contract.ADMINROLE(&_HealthAttestation.CallOpts)
}

// COOLDOWN is a free data retrieval call binding the contract method 0xa2724a4d.
//
// Solidity: function COOLDOWN() view returns(uint256)
func (_HealthAttestation *HealthAttestationCaller) COOLDOWN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "COOLDOWN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COOLDOWN is a free data retrieval call binding the contract method 0xa2724a4d.
//
// Solidity: function COOLDOWN() view returns(uint256)
func (_HealthAttestation *HealthAttestationSession) COOLDOWN() (*big.Int, error) {
	return _HealthAttestation.Contract.COOLDOWN(&_HealthAttestation.CallOpts)
}

// COOLDOWN is a free data retrieval call binding the contract method 0xa2724a4d.
//
// Solidity: function COOLDOWN() view returns(uint256)
func (_HealthAttestation *HealthAttestationCallerSession) COOLDOWN() (*big.Int, error) {
	return _HealthAttestation.Contract.COOLDOWN(&_HealthAttestation.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HealthAttestation *HealthAttestationSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HealthAttestation.Contract.DEFAULTADMINROLE(&_HealthAttestation.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HealthAttestation.Contract.DEFAULTADMINROLE(&_HealthAttestation.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_HealthAttestation *HealthAttestationSession) UPGRADERROLE() ([32]byte, error) {
	return _HealthAttestation.Contract.UPGRADERROLE(&_HealthAttestation.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _HealthAttestation.Contract.UPGRADERROLE(&_HealthAttestation.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_HealthAttestation *HealthAttestationCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_HealthAttestation *HealthAttestationSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _HealthAttestation.Contract.UPGRADEINTERFACEVERSION(&_HealthAttestation.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_HealthAttestation *HealthAttestationCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _HealthAttestation.Contract.UPGRADEINTERFACEVERSION(&_HealthAttestation.CallOpts)
}

// VERIFLOWSIGNER is a free data retrieval call binding the contract method 0x2aab8176.
//
// Solidity: function VERIFLOW_SIGNER() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCaller) VERIFLOWSIGNER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "VERIFLOW_SIGNER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFLOWSIGNER is a free data retrieval call binding the contract method 0x2aab8176.
//
// Solidity: function VERIFLOW_SIGNER() view returns(bytes32)
func (_HealthAttestation *HealthAttestationSession) VERIFLOWSIGNER() ([32]byte, error) {
	return _HealthAttestation.Contract.VERIFLOWSIGNER(&_HealthAttestation.CallOpts)
}

// VERIFLOWSIGNER is a free data retrieval call binding the contract method 0x2aab8176.
//
// Solidity: function VERIFLOW_SIGNER() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCallerSession) VERIFLOWSIGNER() ([32]byte, error) {
	return _HealthAttestation.Contract.VERIFLOWSIGNER(&_HealthAttestation.CallOpts)
}

// GetAttestationHistory is a free data retrieval call binding the contract method 0x7bac1d2b.
//
// Solidity: function getAttestationHistory(bytes32 assetId, uint256 limit) view returns((bytes32,bytes32,uint8,uint256)[])
func (_HealthAttestation *HealthAttestationCaller) GetAttestationHistory(opts *bind.CallOpts, assetId [32]byte, limit *big.Int) ([]IHealthAttestationAttestation, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "getAttestationHistory", assetId, limit)

	if err != nil {
		return *new([]IHealthAttestationAttestation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IHealthAttestationAttestation)).(*[]IHealthAttestationAttestation)

	return out0, err

}

// GetAttestationHistory is a free data retrieval call binding the contract method 0x7bac1d2b.
//
// Solidity: function getAttestationHistory(bytes32 assetId, uint256 limit) view returns((bytes32,bytes32,uint8,uint256)[])
func (_HealthAttestation *HealthAttestationSession) GetAttestationHistory(assetId [32]byte, limit *big.Int) ([]IHealthAttestationAttestation, error) {
	return _HealthAttestation.Contract.GetAttestationHistory(&_HealthAttestation.CallOpts, assetId, limit)
}

// GetAttestationHistory is a free data retrieval call binding the contract method 0x7bac1d2b.
//
// Solidity: function getAttestationHistory(bytes32 assetId, uint256 limit) view returns((bytes32,bytes32,uint8,uint256)[])
func (_HealthAttestation *HealthAttestationCallerSession) GetAttestationHistory(assetId [32]byte, limit *big.Int) ([]IHealthAttestationAttestation, error) {
	return _HealthAttestation.Contract.GetAttestationHistory(&_HealthAttestation.CallOpts, assetId, limit)
}

// GetLatestAttestation is a free data retrieval call binding the contract method 0x5edf0314.
//
// Solidity: function getLatestAttestation(bytes32 assetId) view returns((bytes32,bytes32,uint8,uint256))
func (_HealthAttestation *HealthAttestationCaller) GetLatestAttestation(opts *bind.CallOpts, assetId [32]byte) (IHealthAttestationAttestation, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "getLatestAttestation", assetId)

	if err != nil {
		return *new(IHealthAttestationAttestation), err
	}

	out0 := *abi.ConvertType(out[0], new(IHealthAttestationAttestation)).(*IHealthAttestationAttestation)

	return out0, err

}

// GetLatestAttestation is a free data retrieval call binding the contract method 0x5edf0314.
//
// Solidity: function getLatestAttestation(bytes32 assetId) view returns((bytes32,bytes32,uint8,uint256))
func (_HealthAttestation *HealthAttestationSession) GetLatestAttestation(assetId [32]byte) (IHealthAttestationAttestation, error) {
	return _HealthAttestation.Contract.GetLatestAttestation(&_HealthAttestation.CallOpts, assetId)
}

// GetLatestAttestation is a free data retrieval call binding the contract method 0x5edf0314.
//
// Solidity: function getLatestAttestation(bytes32 assetId) view returns((bytes32,bytes32,uint8,uint256))
func (_HealthAttestation *HealthAttestationCallerSession) GetLatestAttestation(assetId [32]byte) (IHealthAttestationAttestation, error) {
	return _HealthAttestation.Contract.GetLatestAttestation(&_HealthAttestation.CallOpts, assetId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HealthAttestation *HealthAttestationCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HealthAttestation *HealthAttestationSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HealthAttestation.Contract.GetRoleAdmin(&_HealthAttestation.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HealthAttestation *HealthAttestationCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HealthAttestation.Contract.GetRoleAdmin(&_HealthAttestation.CallOpts, role)
}

// HasAttestation is a free data retrieval call binding the contract method 0x35deba33.
//
// Solidity: function hasAttestation(bytes32 assetId) view returns(bool)
func (_HealthAttestation *HealthAttestationCaller) HasAttestation(opts *bind.CallOpts, assetId [32]byte) (bool, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "hasAttestation", assetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAttestation is a free data retrieval call binding the contract method 0x35deba33.
//
// Solidity: function hasAttestation(bytes32 assetId) view returns(bool)
func (_HealthAttestation *HealthAttestationSession) HasAttestation(assetId [32]byte) (bool, error) {
	return _HealthAttestation.Contract.HasAttestation(&_HealthAttestation.CallOpts, assetId)
}

// HasAttestation is a free data retrieval call binding the contract method 0x35deba33.
//
// Solidity: function hasAttestation(bytes32 assetId) view returns(bool)
func (_HealthAttestation *HealthAttestationCallerSession) HasAttestation(assetId [32]byte) (bool, error) {
	return _HealthAttestation.Contract.HasAttestation(&_HealthAttestation.CallOpts, assetId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HealthAttestation *HealthAttestationCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HealthAttestation *HealthAttestationSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HealthAttestation.Contract.HasRole(&_HealthAttestation.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HealthAttestation *HealthAttestationCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HealthAttestation.Contract.HasRole(&_HealthAttestation.CallOpts, role, account)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_HealthAttestation *HealthAttestationSession) ProxiableUUID() ([32]byte, error) {
	return _HealthAttestation.Contract.ProxiableUUID(&_HealthAttestation.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_HealthAttestation *HealthAttestationCallerSession) ProxiableUUID() ([32]byte, error) {
	return _HealthAttestation.Contract.ProxiableUUID(&_HealthAttestation.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HealthAttestation *HealthAttestationCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HealthAttestation.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HealthAttestation *HealthAttestationSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HealthAttestation.Contract.SupportsInterface(&_HealthAttestation.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HealthAttestation *HealthAttestationCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HealthAttestation.Contract.SupportsInterface(&_HealthAttestation.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HealthAttestation *HealthAttestationTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HealthAttestation.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HealthAttestation *HealthAttestationSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HealthAttestation.Contract.GrantRole(&_HealthAttestation.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HealthAttestation *HealthAttestationTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HealthAttestation.Contract.GrantRole(&_HealthAttestation.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_HealthAttestation *HealthAttestationTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _HealthAttestation.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_HealthAttestation *HealthAttestationSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _HealthAttestation.Contract.Initialize(&_HealthAttestation.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_HealthAttestation *HealthAttestationTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _HealthAttestation.Contract.Initialize(&_HealthAttestation.TransactOpts, admin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_HealthAttestation *HealthAttestationTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _HealthAttestation.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_HealthAttestation *HealthAttestationSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _HealthAttestation.Contract.RenounceRole(&_HealthAttestation.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_HealthAttestation *HealthAttestationTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _HealthAttestation.Contract.RenounceRole(&_HealthAttestation.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HealthAttestation *HealthAttestationTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HealthAttestation.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HealthAttestation *HealthAttestationSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HealthAttestation.Contract.RevokeRole(&_HealthAttestation.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HealthAttestation *HealthAttestationTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HealthAttestation.Contract.RevokeRole(&_HealthAttestation.TransactOpts, role, account)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xbcf0b688.
//
// Solidity: function submitAttestation(bytes32 assetId, bytes32 healthHash, uint8 score) returns()
func (_HealthAttestation *HealthAttestationTransactor) SubmitAttestation(opts *bind.TransactOpts, assetId [32]byte, healthHash [32]byte, score uint8) (*types.Transaction, error) {
	return _HealthAttestation.contract.Transact(opts, "submitAttestation", assetId, healthHash, score)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xbcf0b688.
//
// Solidity: function submitAttestation(bytes32 assetId, bytes32 healthHash, uint8 score) returns()
func (_HealthAttestation *HealthAttestationSession) SubmitAttestation(assetId [32]byte, healthHash [32]byte, score uint8) (*types.Transaction, error) {
	return _HealthAttestation.Contract.SubmitAttestation(&_HealthAttestation.TransactOpts, assetId, healthHash, score)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xbcf0b688.
//
// Solidity: function submitAttestation(bytes32 assetId, bytes32 healthHash, uint8 score) returns()
func (_HealthAttestation *HealthAttestationTransactorSession) SubmitAttestation(assetId [32]byte, healthHash [32]byte, score uint8) (*types.Transaction, error) {
	return _HealthAttestation.Contract.SubmitAttestation(&_HealthAttestation.TransactOpts, assetId, healthHash, score)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_HealthAttestation *HealthAttestationTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _HealthAttestation.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_HealthAttestation *HealthAttestationSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _HealthAttestation.Contract.UpgradeToAndCall(&_HealthAttestation.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_HealthAttestation *HealthAttestationTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _HealthAttestation.Contract.UpgradeToAndCall(&_HealthAttestation.TransactOpts, newImplementation, data)
}

// HealthAttestationAttestationSubmittedIterator is returned from FilterAttestationSubmitted and is used to iterate over the raw logs and unpacked data for AttestationSubmitted events raised by the HealthAttestation contract.
type HealthAttestationAttestationSubmittedIterator struct {
	Event *HealthAttestationAttestationSubmitted // Event containing the contract specifics and raw log

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
func (it *HealthAttestationAttestationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealthAttestationAttestationSubmitted)
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
		it.Event = new(HealthAttestationAttestationSubmitted)
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
func (it *HealthAttestationAttestationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealthAttestationAttestationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealthAttestationAttestationSubmitted represents a AttestationSubmitted event raised by the HealthAttestation contract.
type HealthAttestationAttestationSubmitted struct {
	AssetId    [32]byte
	HealthHash [32]byte
	Score      uint8
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAttestationSubmitted is a free log retrieval operation binding the contract event 0x6a11e133dc9fe8dc33fdf31bf5fd30a1c1a0084616c52af05cf418af5fa63be3.
//
// Solidity: event AttestationSubmitted(bytes32 indexed assetId, bytes32 healthHash, uint8 score, uint256 timestamp)
func (_HealthAttestation *HealthAttestationFilterer) FilterAttestationSubmitted(opts *bind.FilterOpts, assetId [][32]byte) (*HealthAttestationAttestationSubmittedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _HealthAttestation.contract.FilterLogs(opts, "AttestationSubmitted", assetIdRule)
	if err != nil {
		return nil, err
	}
	return &HealthAttestationAttestationSubmittedIterator{contract: _HealthAttestation.contract, event: "AttestationSubmitted", logs: logs, sub: sub}, nil
}

// WatchAttestationSubmitted is a free log subscription operation binding the contract event 0x6a11e133dc9fe8dc33fdf31bf5fd30a1c1a0084616c52af05cf418af5fa63be3.
//
// Solidity: event AttestationSubmitted(bytes32 indexed assetId, bytes32 healthHash, uint8 score, uint256 timestamp)
func (_HealthAttestation *HealthAttestationFilterer) WatchAttestationSubmitted(opts *bind.WatchOpts, sink chan<- *HealthAttestationAttestationSubmitted, assetId [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _HealthAttestation.contract.WatchLogs(opts, "AttestationSubmitted", assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealthAttestationAttestationSubmitted)
				if err := _HealthAttestation.contract.UnpackLog(event, "AttestationSubmitted", log); err != nil {
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

// ParseAttestationSubmitted is a log parse operation binding the contract event 0x6a11e133dc9fe8dc33fdf31bf5fd30a1c1a0084616c52af05cf418af5fa63be3.
//
// Solidity: event AttestationSubmitted(bytes32 indexed assetId, bytes32 healthHash, uint8 score, uint256 timestamp)
func (_HealthAttestation *HealthAttestationFilterer) ParseAttestationSubmitted(log types.Log) (*HealthAttestationAttestationSubmitted, error) {
	event := new(HealthAttestationAttestationSubmitted)
	if err := _HealthAttestation.contract.UnpackLog(event, "AttestationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealthAttestationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HealthAttestation contract.
type HealthAttestationInitializedIterator struct {
	Event *HealthAttestationInitialized // Event containing the contract specifics and raw log

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
func (it *HealthAttestationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealthAttestationInitialized)
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
		it.Event = new(HealthAttestationInitialized)
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
func (it *HealthAttestationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealthAttestationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealthAttestationInitialized represents a Initialized event raised by the HealthAttestation contract.
type HealthAttestationInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_HealthAttestation *HealthAttestationFilterer) FilterInitialized(opts *bind.FilterOpts) (*HealthAttestationInitializedIterator, error) {

	logs, sub, err := _HealthAttestation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HealthAttestationInitializedIterator{contract: _HealthAttestation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_HealthAttestation *HealthAttestationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HealthAttestationInitialized) (event.Subscription, error) {

	logs, sub, err := _HealthAttestation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealthAttestationInitialized)
				if err := _HealthAttestation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_HealthAttestation *HealthAttestationFilterer) ParseInitialized(log types.Log) (*HealthAttestationInitialized, error) {
	event := new(HealthAttestationInitialized)
	if err := _HealthAttestation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealthAttestationRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HealthAttestation contract.
type HealthAttestationRoleAdminChangedIterator struct {
	Event *HealthAttestationRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *HealthAttestationRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealthAttestationRoleAdminChanged)
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
		it.Event = new(HealthAttestationRoleAdminChanged)
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
func (it *HealthAttestationRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealthAttestationRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealthAttestationRoleAdminChanged represents a RoleAdminChanged event raised by the HealthAttestation contract.
type HealthAttestationRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HealthAttestation *HealthAttestationFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HealthAttestationRoleAdminChangedIterator, error) {

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

	logs, sub, err := _HealthAttestation.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HealthAttestationRoleAdminChangedIterator{contract: _HealthAttestation.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HealthAttestation *HealthAttestationFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HealthAttestationRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HealthAttestation.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealthAttestationRoleAdminChanged)
				if err := _HealthAttestation.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_HealthAttestation *HealthAttestationFilterer) ParseRoleAdminChanged(log types.Log) (*HealthAttestationRoleAdminChanged, error) {
	event := new(HealthAttestationRoleAdminChanged)
	if err := _HealthAttestation.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealthAttestationRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HealthAttestation contract.
type HealthAttestationRoleGrantedIterator struct {
	Event *HealthAttestationRoleGranted // Event containing the contract specifics and raw log

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
func (it *HealthAttestationRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealthAttestationRoleGranted)
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
		it.Event = new(HealthAttestationRoleGranted)
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
func (it *HealthAttestationRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealthAttestationRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealthAttestationRoleGranted represents a RoleGranted event raised by the HealthAttestation contract.
type HealthAttestationRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HealthAttestation *HealthAttestationFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HealthAttestationRoleGrantedIterator, error) {

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

	logs, sub, err := _HealthAttestation.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HealthAttestationRoleGrantedIterator{contract: _HealthAttestation.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HealthAttestation *HealthAttestationFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HealthAttestationRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HealthAttestation.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealthAttestationRoleGranted)
				if err := _HealthAttestation.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_HealthAttestation *HealthAttestationFilterer) ParseRoleGranted(log types.Log) (*HealthAttestationRoleGranted, error) {
	event := new(HealthAttestationRoleGranted)
	if err := _HealthAttestation.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealthAttestationRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HealthAttestation contract.
type HealthAttestationRoleRevokedIterator struct {
	Event *HealthAttestationRoleRevoked // Event containing the contract specifics and raw log

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
func (it *HealthAttestationRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealthAttestationRoleRevoked)
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
		it.Event = new(HealthAttestationRoleRevoked)
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
func (it *HealthAttestationRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealthAttestationRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealthAttestationRoleRevoked represents a RoleRevoked event raised by the HealthAttestation contract.
type HealthAttestationRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HealthAttestation *HealthAttestationFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HealthAttestationRoleRevokedIterator, error) {

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

	logs, sub, err := _HealthAttestation.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HealthAttestationRoleRevokedIterator{contract: _HealthAttestation.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HealthAttestation *HealthAttestationFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HealthAttestationRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HealthAttestation.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealthAttestationRoleRevoked)
				if err := _HealthAttestation.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_HealthAttestation *HealthAttestationFilterer) ParseRoleRevoked(log types.Log) (*HealthAttestationRoleRevoked, error) {
	event := new(HealthAttestationRoleRevoked)
	if err := _HealthAttestation.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealthAttestationUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the HealthAttestation contract.
type HealthAttestationUpgradedIterator struct {
	Event *HealthAttestationUpgraded // Event containing the contract specifics and raw log

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
func (it *HealthAttestationUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealthAttestationUpgraded)
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
		it.Event = new(HealthAttestationUpgraded)
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
func (it *HealthAttestationUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealthAttestationUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealthAttestationUpgraded represents a Upgraded event raised by the HealthAttestation contract.
type HealthAttestationUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_HealthAttestation *HealthAttestationFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*HealthAttestationUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _HealthAttestation.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &HealthAttestationUpgradedIterator{contract: _HealthAttestation.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_HealthAttestation *HealthAttestationFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *HealthAttestationUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _HealthAttestation.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealthAttestationUpgraded)
				if err := _HealthAttestation.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_HealthAttestation *HealthAttestationFilterer) ParseUpgraded(log types.Log) (*HealthAttestationUpgraded, error) {
	event := new(HealthAttestationUpgraded)
	if err := _HealthAttestation.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
