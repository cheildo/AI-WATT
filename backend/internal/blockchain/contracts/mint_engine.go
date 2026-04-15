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

// MintEngineMetaData contains all meta data concerning the MintEngine contract.
var MintEngineMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InsufficientCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"}],\"name\":\"StablecoinNotAccepted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wattMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wattBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stablecoinReturned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"StablecoinUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTreasury\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"acceptedStablecoins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"}],\"name\":\"collateralBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wattUSD\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"}],\"name\":\"isAcceptedStablecoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wattAmount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stablecoin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"setAcceptedStablecoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wattUSD\",\"outputs\":[{\"internalType\":\"contractIWattUSD\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MintEngineABI is the input ABI used to generate the binding from.
// Deprecated: Use MintEngineMetaData.ABI instead.
var MintEngineABI = MintEngineMetaData.ABI

// MintEngine is an auto generated Go binding around an Ethereum contract.
type MintEngine struct {
	MintEngineCaller     // Read-only binding to the contract
	MintEngineTransactor // Write-only binding to the contract
	MintEngineFilterer   // Log filterer for contract events
}

// MintEngineCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintEngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintEngineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintEngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintEngineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintEngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintEngineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintEngineSession struct {
	Contract     *MintEngine       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintEngineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintEngineCallerSession struct {
	Contract *MintEngineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MintEngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintEngineTransactorSession struct {
	Contract     *MintEngineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MintEngineRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintEngineRaw struct {
	Contract *MintEngine // Generic contract binding to access the raw methods on
}

// MintEngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintEngineCallerRaw struct {
	Contract *MintEngineCaller // Generic read-only contract binding to access the raw methods on
}

// MintEngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintEngineTransactorRaw struct {
	Contract *MintEngineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintEngine creates a new instance of MintEngine, bound to a specific deployed contract.
func NewMintEngine(address common.Address, backend bind.ContractBackend) (*MintEngine, error) {
	contract, err := bindMintEngine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintEngine{MintEngineCaller: MintEngineCaller{contract: contract}, MintEngineTransactor: MintEngineTransactor{contract: contract}, MintEngineFilterer: MintEngineFilterer{contract: contract}}, nil
}

// NewMintEngineCaller creates a new read-only instance of MintEngine, bound to a specific deployed contract.
func NewMintEngineCaller(address common.Address, caller bind.ContractCaller) (*MintEngineCaller, error) {
	contract, err := bindMintEngine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintEngineCaller{contract: contract}, nil
}

// NewMintEngineTransactor creates a new write-only instance of MintEngine, bound to a specific deployed contract.
func NewMintEngineTransactor(address common.Address, transactor bind.ContractTransactor) (*MintEngineTransactor, error) {
	contract, err := bindMintEngine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintEngineTransactor{contract: contract}, nil
}

// NewMintEngineFilterer creates a new log filterer instance of MintEngine, bound to a specific deployed contract.
func NewMintEngineFilterer(address common.Address, filterer bind.ContractFilterer) (*MintEngineFilterer, error) {
	contract, err := bindMintEngine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintEngineFilterer{contract: contract}, nil
}

// bindMintEngine binds a generic wrapper to an already deployed contract.
func bindMintEngine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MintEngineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintEngine *MintEngineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintEngine.Contract.MintEngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintEngine *MintEngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintEngine.Contract.MintEngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintEngine *MintEngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintEngine.Contract.MintEngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintEngine *MintEngineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintEngine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintEngine *MintEngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintEngine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintEngine *MintEngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintEngine.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineSession) ADMINROLE() ([32]byte, error) {
	return _MintEngine.Contract.ADMINROLE(&_MintEngine.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineCallerSession) ADMINROLE() ([32]byte, error) {
	return _MintEngine.Contract.ADMINROLE(&_MintEngine.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MintEngine.Contract.DEFAULTADMINROLE(&_MintEngine.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MintEngine.Contract.DEFAULTADMINROLE(&_MintEngine.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_MintEngine *MintEngineCaller) FEEBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "FEE_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_MintEngine *MintEngineSession) FEEBPS() (*big.Int, error) {
	return _MintEngine.Contract.FEEBPS(&_MintEngine.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_MintEngine *MintEngineCallerSession) FEEBPS() (*big.Int, error) {
	return _MintEngine.Contract.FEEBPS(&_MintEngine.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineSession) PAUSERROLE() ([32]byte, error) {
	return _MintEngine.Contract.PAUSERROLE(&_MintEngine.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineCallerSession) PAUSERROLE() ([32]byte, error) {
	return _MintEngine.Contract.PAUSERROLE(&_MintEngine.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineSession) UPGRADERROLE() ([32]byte, error) {
	return _MintEngine.Contract.UPGRADERROLE(&_MintEngine.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_MintEngine *MintEngineCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _MintEngine.Contract.UPGRADERROLE(&_MintEngine.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MintEngine *MintEngineCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MintEngine *MintEngineSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MintEngine.Contract.UPGRADEINTERFACEVERSION(&_MintEngine.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MintEngine *MintEngineCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MintEngine.Contract.UPGRADEINTERFACEVERSION(&_MintEngine.CallOpts)
}

// AcceptedStablecoins is a free data retrieval call binding the contract method 0x19fee256.
//
// Solidity: function acceptedStablecoins(address ) view returns(bool)
func (_MintEngine *MintEngineCaller) AcceptedStablecoins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "acceptedStablecoins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptedStablecoins is a free data retrieval call binding the contract method 0x19fee256.
//
// Solidity: function acceptedStablecoins(address ) view returns(bool)
func (_MintEngine *MintEngineSession) AcceptedStablecoins(arg0 common.Address) (bool, error) {
	return _MintEngine.Contract.AcceptedStablecoins(&_MintEngine.CallOpts, arg0)
}

// AcceptedStablecoins is a free data retrieval call binding the contract method 0x19fee256.
//
// Solidity: function acceptedStablecoins(address ) view returns(bool)
func (_MintEngine *MintEngineCallerSession) AcceptedStablecoins(arg0 common.Address) (bool, error) {
	return _MintEngine.Contract.AcceptedStablecoins(&_MintEngine.CallOpts, arg0)
}

// CollateralBalance is a free data retrieval call binding the contract method 0xa1bf2840.
//
// Solidity: function collateralBalance(address stablecoin) view returns(uint256)
func (_MintEngine *MintEngineCaller) CollateralBalance(opts *bind.CallOpts, stablecoin common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "collateralBalance", stablecoin)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralBalance is a free data retrieval call binding the contract method 0xa1bf2840.
//
// Solidity: function collateralBalance(address stablecoin) view returns(uint256)
func (_MintEngine *MintEngineSession) CollateralBalance(stablecoin common.Address) (*big.Int, error) {
	return _MintEngine.Contract.CollateralBalance(&_MintEngine.CallOpts, stablecoin)
}

// CollateralBalance is a free data retrieval call binding the contract method 0xa1bf2840.
//
// Solidity: function collateralBalance(address stablecoin) view returns(uint256)
func (_MintEngine *MintEngineCallerSession) CollateralBalance(stablecoin common.Address) (*big.Int, error) {
	return _MintEngine.Contract.CollateralBalance(&_MintEngine.CallOpts, stablecoin)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MintEngine *MintEngineCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MintEngine *MintEngineSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MintEngine.Contract.GetRoleAdmin(&_MintEngine.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MintEngine *MintEngineCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MintEngine.Contract.GetRoleAdmin(&_MintEngine.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MintEngine *MintEngineCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MintEngine *MintEngineSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MintEngine.Contract.HasRole(&_MintEngine.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MintEngine *MintEngineCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MintEngine.Contract.HasRole(&_MintEngine.CallOpts, role, account)
}

// IsAcceptedStablecoin is a free data retrieval call binding the contract method 0xbe684f02.
//
// Solidity: function isAcceptedStablecoin(address stablecoin) view returns(bool)
func (_MintEngine *MintEngineCaller) IsAcceptedStablecoin(opts *bind.CallOpts, stablecoin common.Address) (bool, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "isAcceptedStablecoin", stablecoin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAcceptedStablecoin is a free data retrieval call binding the contract method 0xbe684f02.
//
// Solidity: function isAcceptedStablecoin(address stablecoin) view returns(bool)
func (_MintEngine *MintEngineSession) IsAcceptedStablecoin(stablecoin common.Address) (bool, error) {
	return _MintEngine.Contract.IsAcceptedStablecoin(&_MintEngine.CallOpts, stablecoin)
}

// IsAcceptedStablecoin is a free data retrieval call binding the contract method 0xbe684f02.
//
// Solidity: function isAcceptedStablecoin(address stablecoin) view returns(bool)
func (_MintEngine *MintEngineCallerSession) IsAcceptedStablecoin(stablecoin common.Address) (bool, error) {
	return _MintEngine.Contract.IsAcceptedStablecoin(&_MintEngine.CallOpts, stablecoin)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MintEngine *MintEngineCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MintEngine *MintEngineSession) Paused() (bool, error) {
	return _MintEngine.Contract.Paused(&_MintEngine.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MintEngine *MintEngineCallerSession) Paused() (bool, error) {
	return _MintEngine.Contract.Paused(&_MintEngine.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MintEngine *MintEngineCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MintEngine *MintEngineSession) ProxiableUUID() ([32]byte, error) {
	return _MintEngine.Contract.ProxiableUUID(&_MintEngine.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MintEngine *MintEngineCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MintEngine.Contract.ProxiableUUID(&_MintEngine.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MintEngine *MintEngineCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MintEngine *MintEngineSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MintEngine.Contract.SupportsInterface(&_MintEngine.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MintEngine *MintEngineCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MintEngine.Contract.SupportsInterface(&_MintEngine.CallOpts, interfaceId)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_MintEngine *MintEngineCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_MintEngine *MintEngineSession) Treasury() (common.Address, error) {
	return _MintEngine.Contract.Treasury(&_MintEngine.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_MintEngine *MintEngineCallerSession) Treasury() (common.Address, error) {
	return _MintEngine.Contract.Treasury(&_MintEngine.CallOpts)
}

// WattUSD is a free data retrieval call binding the contract method 0x7aff9566.
//
// Solidity: function wattUSD() view returns(address)
func (_MintEngine *MintEngineCaller) WattUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintEngine.contract.Call(opts, &out, "wattUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WattUSD is a free data retrieval call binding the contract method 0x7aff9566.
//
// Solidity: function wattUSD() view returns(address)
func (_MintEngine *MintEngineSession) WattUSD() (common.Address, error) {
	return _MintEngine.Contract.WattUSD(&_MintEngine.CallOpts)
}

// WattUSD is a free data retrieval call binding the contract method 0x7aff9566.
//
// Solidity: function wattUSD() view returns(address)
func (_MintEngine *MintEngineCallerSession) WattUSD() (common.Address, error) {
	return _MintEngine.Contract.WattUSD(&_MintEngine.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MintEngine *MintEngineTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MintEngine *MintEngineSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.GrantRole(&_MintEngine.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MintEngine *MintEngineTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.GrantRole(&_MintEngine.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address admin, address _wattUSD, address _treasury) returns()
func (_MintEngine *MintEngineTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, _wattUSD common.Address, _treasury common.Address) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "initialize", admin, _wattUSD, _treasury)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address admin, address _wattUSD, address _treasury) returns()
func (_MintEngine *MintEngineSession) Initialize(admin common.Address, _wattUSD common.Address, _treasury common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.Initialize(&_MintEngine.TransactOpts, admin, _wattUSD, _treasury)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address admin, address _wattUSD, address _treasury) returns()
func (_MintEngine *MintEngineTransactorSession) Initialize(admin common.Address, _wattUSD common.Address, _treasury common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.Initialize(&_MintEngine.TransactOpts, admin, _wattUSD, _treasury)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address stablecoin, uint256 amount) returns()
func (_MintEngine *MintEngineTransactor) Mint(opts *bind.TransactOpts, stablecoin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "mint", stablecoin, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address stablecoin, uint256 amount) returns()
func (_MintEngine *MintEngineSession) Mint(stablecoin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MintEngine.Contract.Mint(&_MintEngine.TransactOpts, stablecoin, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address stablecoin, uint256 amount) returns()
func (_MintEngine *MintEngineTransactorSession) Mint(stablecoin common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MintEngine.Contract.Mint(&_MintEngine.TransactOpts, stablecoin, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MintEngine *MintEngineTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MintEngine *MintEngineSession) Pause() (*types.Transaction, error) {
	return _MintEngine.Contract.Pause(&_MintEngine.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MintEngine *MintEngineTransactorSession) Pause() (*types.Transaction, error) {
	return _MintEngine.Contract.Pause(&_MintEngine.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address stablecoin, uint256 wattAmount) returns()
func (_MintEngine *MintEngineTransactor) Redeem(opts *bind.TransactOpts, stablecoin common.Address, wattAmount *big.Int) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "redeem", stablecoin, wattAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address stablecoin, uint256 wattAmount) returns()
func (_MintEngine *MintEngineSession) Redeem(stablecoin common.Address, wattAmount *big.Int) (*types.Transaction, error) {
	return _MintEngine.Contract.Redeem(&_MintEngine.TransactOpts, stablecoin, wattAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address stablecoin, uint256 wattAmount) returns()
func (_MintEngine *MintEngineTransactorSession) Redeem(stablecoin common.Address, wattAmount *big.Int) (*types.Transaction, error) {
	return _MintEngine.Contract.Redeem(&_MintEngine.TransactOpts, stablecoin, wattAmount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MintEngine *MintEngineTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MintEngine *MintEngineSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.RenounceRole(&_MintEngine.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MintEngine *MintEngineTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.RenounceRole(&_MintEngine.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MintEngine *MintEngineTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MintEngine *MintEngineSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.RevokeRole(&_MintEngine.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MintEngine *MintEngineTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.RevokeRole(&_MintEngine.TransactOpts, role, account)
}

// SetAcceptedStablecoin is a paid mutator transaction binding the contract method 0xc16d700f.
//
// Solidity: function setAcceptedStablecoin(address stablecoin, bool accepted) returns()
func (_MintEngine *MintEngineTransactor) SetAcceptedStablecoin(opts *bind.TransactOpts, stablecoin common.Address, accepted bool) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "setAcceptedStablecoin", stablecoin, accepted)
}

// SetAcceptedStablecoin is a paid mutator transaction binding the contract method 0xc16d700f.
//
// Solidity: function setAcceptedStablecoin(address stablecoin, bool accepted) returns()
func (_MintEngine *MintEngineSession) SetAcceptedStablecoin(stablecoin common.Address, accepted bool) (*types.Transaction, error) {
	return _MintEngine.Contract.SetAcceptedStablecoin(&_MintEngine.TransactOpts, stablecoin, accepted)
}

// SetAcceptedStablecoin is a paid mutator transaction binding the contract method 0xc16d700f.
//
// Solidity: function setAcceptedStablecoin(address stablecoin, bool accepted) returns()
func (_MintEngine *MintEngineTransactorSession) SetAcceptedStablecoin(stablecoin common.Address, accepted bool) (*types.Transaction, error) {
	return _MintEngine.Contract.SetAcceptedStablecoin(&_MintEngine.TransactOpts, stablecoin, accepted)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_MintEngine *MintEngineTransactor) SetTreasury(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "setTreasury", newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_MintEngine *MintEngineSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.SetTreasury(&_MintEngine.TransactOpts, newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_MintEngine *MintEngineTransactorSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _MintEngine.Contract.SetTreasury(&_MintEngine.TransactOpts, newTreasury)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MintEngine *MintEngineTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MintEngine *MintEngineSession) Unpause() (*types.Transaction, error) {
	return _MintEngine.Contract.Unpause(&_MintEngine.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MintEngine *MintEngineTransactorSession) Unpause() (*types.Transaction, error) {
	return _MintEngine.Contract.Unpause(&_MintEngine.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MintEngine *MintEngineTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MintEngine.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MintEngine *MintEngineSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MintEngine.Contract.UpgradeToAndCall(&_MintEngine.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MintEngine *MintEngineTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MintEngine.Contract.UpgradeToAndCall(&_MintEngine.TransactOpts, newImplementation, data)
}

// MintEngineInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MintEngine contract.
type MintEngineInitializedIterator struct {
	Event *MintEngineInitialized // Event containing the contract specifics and raw log

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
func (it *MintEngineInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineInitialized)
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
		it.Event = new(MintEngineInitialized)
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
func (it *MintEngineInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineInitialized represents a Initialized event raised by the MintEngine contract.
type MintEngineInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MintEngine *MintEngineFilterer) FilterInitialized(opts *bind.FilterOpts) (*MintEngineInitializedIterator, error) {

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MintEngineInitializedIterator{contract: _MintEngine.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MintEngine *MintEngineFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MintEngineInitialized) (event.Subscription, error) {

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineInitialized)
				if err := _MintEngine.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MintEngine *MintEngineFilterer) ParseInitialized(log types.Log) (*MintEngineInitialized, error) {
	event := new(MintEngineInitialized)
	if err := _MintEngine.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEngineMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the MintEngine contract.
type MintEngineMintedIterator struct {
	Event *MintEngineMinted // Event containing the contract specifics and raw log

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
func (it *MintEngineMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineMinted)
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
		it.Event = new(MintEngineMinted)
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
func (it *MintEngineMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineMinted represents a Minted event raised by the MintEngine contract.
type MintEngineMinted struct {
	Depositor     common.Address
	Stablecoin    common.Address
	DepositAmount *big.Int
	WattMinted    *big.Int
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x68c721f9a38584842eb66945b43c95066397ce5f2576fcc0588471c85d209d0c.
//
// Solidity: event Minted(address indexed depositor, address indexed stablecoin, uint256 depositAmount, uint256 wattMinted, uint256 fee)
func (_MintEngine *MintEngineFilterer) FilterMinted(opts *bind.FilterOpts, depositor []common.Address, stablecoin []common.Address) (*MintEngineMintedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var stablecoinRule []interface{}
	for _, stablecoinItem := range stablecoin {
		stablecoinRule = append(stablecoinRule, stablecoinItem)
	}

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "Minted", depositorRule, stablecoinRule)
	if err != nil {
		return nil, err
	}
	return &MintEngineMintedIterator{contract: _MintEngine.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x68c721f9a38584842eb66945b43c95066397ce5f2576fcc0588471c85d209d0c.
//
// Solidity: event Minted(address indexed depositor, address indexed stablecoin, uint256 depositAmount, uint256 wattMinted, uint256 fee)
func (_MintEngine *MintEngineFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MintEngineMinted, depositor []common.Address, stablecoin []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var stablecoinRule []interface{}
	for _, stablecoinItem := range stablecoin {
		stablecoinRule = append(stablecoinRule, stablecoinItem)
	}

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "Minted", depositorRule, stablecoinRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineMinted)
				if err := _MintEngine.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x68c721f9a38584842eb66945b43c95066397ce5f2576fcc0588471c85d209d0c.
//
// Solidity: event Minted(address indexed depositor, address indexed stablecoin, uint256 depositAmount, uint256 wattMinted, uint256 fee)
func (_MintEngine *MintEngineFilterer) ParseMinted(log types.Log) (*MintEngineMinted, error) {
	event := new(MintEngineMinted)
	if err := _MintEngine.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEnginePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MintEngine contract.
type MintEnginePausedIterator struct {
	Event *MintEnginePaused // Event containing the contract specifics and raw log

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
func (it *MintEnginePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEnginePaused)
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
		it.Event = new(MintEnginePaused)
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
func (it *MintEnginePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEnginePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEnginePaused represents a Paused event raised by the MintEngine contract.
type MintEnginePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MintEngine *MintEngineFilterer) FilterPaused(opts *bind.FilterOpts) (*MintEnginePausedIterator, error) {

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MintEnginePausedIterator{contract: _MintEngine.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MintEngine *MintEngineFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MintEnginePaused) (event.Subscription, error) {

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEnginePaused)
				if err := _MintEngine.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MintEngine *MintEngineFilterer) ParsePaused(log types.Log) (*MintEnginePaused, error) {
	event := new(MintEnginePaused)
	if err := _MintEngine.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEngineRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the MintEngine contract.
type MintEngineRedeemedIterator struct {
	Event *MintEngineRedeemed // Event containing the contract specifics and raw log

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
func (it *MintEngineRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineRedeemed)
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
		it.Event = new(MintEngineRedeemed)
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
func (it *MintEngineRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineRedeemed represents a Redeemed event raised by the MintEngine contract.
type MintEngineRedeemed struct {
	Redeemer           common.Address
	Stablecoin         common.Address
	WattBurned         *big.Int
	StablecoinReturned *big.Int
	Fee                *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x764aeeb2d1ec3f2945d6486e2f7e3fae9ac5fe11aa56b7a9d90c92212e33050c.
//
// Solidity: event Redeemed(address indexed redeemer, address indexed stablecoin, uint256 wattBurned, uint256 stablecoinReturned, uint256 fee)
func (_MintEngine *MintEngineFilterer) FilterRedeemed(opts *bind.FilterOpts, redeemer []common.Address, stablecoin []common.Address) (*MintEngineRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var stablecoinRule []interface{}
	for _, stablecoinItem := range stablecoin {
		stablecoinRule = append(stablecoinRule, stablecoinItem)
	}

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "Redeemed", redeemerRule, stablecoinRule)
	if err != nil {
		return nil, err
	}
	return &MintEngineRedeemedIterator{contract: _MintEngine.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x764aeeb2d1ec3f2945d6486e2f7e3fae9ac5fe11aa56b7a9d90c92212e33050c.
//
// Solidity: event Redeemed(address indexed redeemer, address indexed stablecoin, uint256 wattBurned, uint256 stablecoinReturned, uint256 fee)
func (_MintEngine *MintEngineFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *MintEngineRedeemed, redeemer []common.Address, stablecoin []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var stablecoinRule []interface{}
	for _, stablecoinItem := range stablecoin {
		stablecoinRule = append(stablecoinRule, stablecoinItem)
	}

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "Redeemed", redeemerRule, stablecoinRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineRedeemed)
				if err := _MintEngine.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0x764aeeb2d1ec3f2945d6486e2f7e3fae9ac5fe11aa56b7a9d90c92212e33050c.
//
// Solidity: event Redeemed(address indexed redeemer, address indexed stablecoin, uint256 wattBurned, uint256 stablecoinReturned, uint256 fee)
func (_MintEngine *MintEngineFilterer) ParseRedeemed(log types.Log) (*MintEngineRedeemed, error) {
	event := new(MintEngineRedeemed)
	if err := _MintEngine.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEngineRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MintEngine contract.
type MintEngineRoleAdminChangedIterator struct {
	Event *MintEngineRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MintEngineRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineRoleAdminChanged)
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
		it.Event = new(MintEngineRoleAdminChanged)
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
func (it *MintEngineRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineRoleAdminChanged represents a RoleAdminChanged event raised by the MintEngine contract.
type MintEngineRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MintEngine *MintEngineFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MintEngineRoleAdminChangedIterator, error) {

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

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MintEngineRoleAdminChangedIterator{contract: _MintEngine.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MintEngine *MintEngineFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MintEngineRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineRoleAdminChanged)
				if err := _MintEngine.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_MintEngine *MintEngineFilterer) ParseRoleAdminChanged(log types.Log) (*MintEngineRoleAdminChanged, error) {
	event := new(MintEngineRoleAdminChanged)
	if err := _MintEngine.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEngineRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MintEngine contract.
type MintEngineRoleGrantedIterator struct {
	Event *MintEngineRoleGranted // Event containing the contract specifics and raw log

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
func (it *MintEngineRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineRoleGranted)
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
		it.Event = new(MintEngineRoleGranted)
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
func (it *MintEngineRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineRoleGranted represents a RoleGranted event raised by the MintEngine contract.
type MintEngineRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintEngine *MintEngineFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MintEngineRoleGrantedIterator, error) {

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

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MintEngineRoleGrantedIterator{contract: _MintEngine.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintEngine *MintEngineFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MintEngineRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineRoleGranted)
				if err := _MintEngine.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_MintEngine *MintEngineFilterer) ParseRoleGranted(log types.Log) (*MintEngineRoleGranted, error) {
	event := new(MintEngineRoleGranted)
	if err := _MintEngine.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEngineRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MintEngine contract.
type MintEngineRoleRevokedIterator struct {
	Event *MintEngineRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MintEngineRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineRoleRevoked)
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
		it.Event = new(MintEngineRoleRevoked)
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
func (it *MintEngineRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineRoleRevoked represents a RoleRevoked event raised by the MintEngine contract.
type MintEngineRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintEngine *MintEngineFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MintEngineRoleRevokedIterator, error) {

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

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MintEngineRoleRevokedIterator{contract: _MintEngine.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MintEngine *MintEngineFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MintEngineRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineRoleRevoked)
				if err := _MintEngine.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_MintEngine *MintEngineFilterer) ParseRoleRevoked(log types.Log) (*MintEngineRoleRevoked, error) {
	event := new(MintEngineRoleRevoked)
	if err := _MintEngine.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEngineStablecoinUpdatedIterator is returned from FilterStablecoinUpdated and is used to iterate over the raw logs and unpacked data for StablecoinUpdated events raised by the MintEngine contract.
type MintEngineStablecoinUpdatedIterator struct {
	Event *MintEngineStablecoinUpdated // Event containing the contract specifics and raw log

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
func (it *MintEngineStablecoinUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineStablecoinUpdated)
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
		it.Event = new(MintEngineStablecoinUpdated)
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
func (it *MintEngineStablecoinUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineStablecoinUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineStablecoinUpdated represents a StablecoinUpdated event raised by the MintEngine contract.
type MintEngineStablecoinUpdated struct {
	Stablecoin common.Address
	Accepted   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStablecoinUpdated is a free log retrieval operation binding the contract event 0x855c900f27c6636569173ceb716d3cfc1a9f951524e822fc3b3f597010a057a3.
//
// Solidity: event StablecoinUpdated(address indexed stablecoin, bool accepted)
func (_MintEngine *MintEngineFilterer) FilterStablecoinUpdated(opts *bind.FilterOpts, stablecoin []common.Address) (*MintEngineStablecoinUpdatedIterator, error) {

	var stablecoinRule []interface{}
	for _, stablecoinItem := range stablecoin {
		stablecoinRule = append(stablecoinRule, stablecoinItem)
	}

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "StablecoinUpdated", stablecoinRule)
	if err != nil {
		return nil, err
	}
	return &MintEngineStablecoinUpdatedIterator{contract: _MintEngine.contract, event: "StablecoinUpdated", logs: logs, sub: sub}, nil
}

// WatchStablecoinUpdated is a free log subscription operation binding the contract event 0x855c900f27c6636569173ceb716d3cfc1a9f951524e822fc3b3f597010a057a3.
//
// Solidity: event StablecoinUpdated(address indexed stablecoin, bool accepted)
func (_MintEngine *MintEngineFilterer) WatchStablecoinUpdated(opts *bind.WatchOpts, sink chan<- *MintEngineStablecoinUpdated, stablecoin []common.Address) (event.Subscription, error) {

	var stablecoinRule []interface{}
	for _, stablecoinItem := range stablecoin {
		stablecoinRule = append(stablecoinRule, stablecoinItem)
	}

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "StablecoinUpdated", stablecoinRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineStablecoinUpdated)
				if err := _MintEngine.contract.UnpackLog(event, "StablecoinUpdated", log); err != nil {
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

// ParseStablecoinUpdated is a log parse operation binding the contract event 0x855c900f27c6636569173ceb716d3cfc1a9f951524e822fc3b3f597010a057a3.
//
// Solidity: event StablecoinUpdated(address indexed stablecoin, bool accepted)
func (_MintEngine *MintEngineFilterer) ParseStablecoinUpdated(log types.Log) (*MintEngineStablecoinUpdated, error) {
	event := new(MintEngineStablecoinUpdated)
	if err := _MintEngine.contract.UnpackLog(event, "StablecoinUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEngineTreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the MintEngine contract.
type MintEngineTreasuryUpdatedIterator struct {
	Event *MintEngineTreasuryUpdated // Event containing the contract specifics and raw log

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
func (it *MintEngineTreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineTreasuryUpdated)
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
		it.Event = new(MintEngineTreasuryUpdated)
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
func (it *MintEngineTreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineTreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineTreasuryUpdated represents a TreasuryUpdated event raised by the MintEngine contract.
type MintEngineTreasuryUpdated struct {
	OldTreasury common.Address
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x4ab5be82436d353e61ca18726e984e561f5c1cc7c6d38b29d2553c790434705a.
//
// Solidity: event TreasuryUpdated(address indexed oldTreasury, address indexed newTreasury)
func (_MintEngine *MintEngineFilterer) FilterTreasuryUpdated(opts *bind.FilterOpts, oldTreasury []common.Address, newTreasury []common.Address) (*MintEngineTreasuryUpdatedIterator, error) {

	var oldTreasuryRule []interface{}
	for _, oldTreasuryItem := range oldTreasury {
		oldTreasuryRule = append(oldTreasuryRule, oldTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "TreasuryUpdated", oldTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &MintEngineTreasuryUpdatedIterator{contract: _MintEngine.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x4ab5be82436d353e61ca18726e984e561f5c1cc7c6d38b29d2553c790434705a.
//
// Solidity: event TreasuryUpdated(address indexed oldTreasury, address indexed newTreasury)
func (_MintEngine *MintEngineFilterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *MintEngineTreasuryUpdated, oldTreasury []common.Address, newTreasury []common.Address) (event.Subscription, error) {

	var oldTreasuryRule []interface{}
	for _, oldTreasuryItem := range oldTreasury {
		oldTreasuryRule = append(oldTreasuryRule, oldTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "TreasuryUpdated", oldTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineTreasuryUpdated)
				if err := _MintEngine.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
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

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x4ab5be82436d353e61ca18726e984e561f5c1cc7c6d38b29d2553c790434705a.
//
// Solidity: event TreasuryUpdated(address indexed oldTreasury, address indexed newTreasury)
func (_MintEngine *MintEngineFilterer) ParseTreasuryUpdated(log types.Log) (*MintEngineTreasuryUpdated, error) {
	event := new(MintEngineTreasuryUpdated)
	if err := _MintEngine.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEngineUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MintEngine contract.
type MintEngineUnpausedIterator struct {
	Event *MintEngineUnpaused // Event containing the contract specifics and raw log

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
func (it *MintEngineUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineUnpaused)
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
		it.Event = new(MintEngineUnpaused)
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
func (it *MintEngineUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineUnpaused represents a Unpaused event raised by the MintEngine contract.
type MintEngineUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MintEngine *MintEngineFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MintEngineUnpausedIterator, error) {

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MintEngineUnpausedIterator{contract: _MintEngine.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MintEngine *MintEngineFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MintEngineUnpaused) (event.Subscription, error) {

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineUnpaused)
				if err := _MintEngine.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MintEngine *MintEngineFilterer) ParseUnpaused(log types.Log) (*MintEngineUnpaused, error) {
	event := new(MintEngineUnpaused)
	if err := _MintEngine.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintEngineUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MintEngine contract.
type MintEngineUpgradedIterator struct {
	Event *MintEngineUpgraded // Event containing the contract specifics and raw log

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
func (it *MintEngineUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintEngineUpgraded)
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
		it.Event = new(MintEngineUpgraded)
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
func (it *MintEngineUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintEngineUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintEngineUpgraded represents a Upgraded event raised by the MintEngine contract.
type MintEngineUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MintEngine *MintEngineFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MintEngineUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MintEngine.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MintEngineUpgradedIterator{contract: _MintEngine.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MintEngine *MintEngineFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MintEngineUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MintEngine.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintEngineUpgraded)
				if err := _MintEngine.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_MintEngine *MintEngineFilterer) ParseUpgraded(log types.Log) (*MintEngineUpgraded, error) {
	event := new(MintEngineUpgraded)
	if err := _MintEngine.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
