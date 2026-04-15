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

// ILendingPoolLoan is an auto generated low-level Go binding around an user-defined struct.
type ILendingPoolLoan struct {
	LoanId       [32]byte
	AssetId      [32]byte
	Borrower     common.Address
	Curator      common.Address
	Principal    *big.Int
	Outstanding  *big.Int
	InterestRate *big.Int
	Status       uint8
	EngineType   uint8
	OriginatedAt *big.Int
	MaturityAt   *big.Int
}

// LendingPoolMetaData contains all meta data concerning the LendingPool contract.
var LendingPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"AssetAlreadyEncumbered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"AssetNotActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastTimestamp\",\"type\":\"uint256\"}],\"name\":\"AttestationStale\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outstanding\",\"type\":\"uint256\"}],\"name\":\"ExceedsOutstanding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"score\",\"type\":\"uint8\"}],\"name\":\"HealthScoreTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"},{\"internalType\":\"enumILendingPool.LoanStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"LoanNotActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"}],\"name\":\"LoanNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"LoanDefaulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"LoanLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"curator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturityAt\",\"type\":\"uint256\"}],\"name\":\"LoanOriginated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"LoanSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principalPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outstandingAfter\",\"type\":\"uint256\"}],\"name\":\"RepaymentReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ATTESTATION_MAX_AGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CURATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIQUIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_HEALTH_SCORE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROTOCOL_FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetRegistry\",\"outputs\":[{\"internalType\":\"contractIAssetRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"}],\"name\":\"flagDefaulted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"}],\"name\":\"fullRepay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getBorrowerLoans\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"}],\"name\":\"getLoan\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"curator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outstanding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"enumILendingPool.LoanStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"engineType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"originatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityAt\",\"type\":\"uint256\"}],\"internalType\":\"structILendingPool.Loan\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"healthAttestation\",\"outputs\":[{\"internalType\":\"contractIHealthAttestation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetRegistry_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"healthAttestation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wattUSD_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sWattUSD_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"termDays\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"engineType\",\"type\":\"uint8\"}],\"name\":\"originateLoan\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sWattUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wattUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LendingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LendingPoolMetaData.ABI instead.
var LendingPoolABI = LendingPoolMetaData.ABI

// LendingPool is an auto generated Go binding around an Ethereum contract.
type LendingPool struct {
	LendingPoolCaller     // Read-only binding to the contract
	LendingPoolTransactor // Write-only binding to the contract
	LendingPoolFilterer   // Log filterer for contract events
}

// LendingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingPoolSession struct {
	Contract     *LendingPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingPoolCallerSession struct {
	Contract *LendingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LendingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingPoolTransactorSession struct {
	Contract     *LendingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LendingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingPoolRaw struct {
	Contract *LendingPool // Generic contract binding to access the raw methods on
}

// LendingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingPoolCallerRaw struct {
	Contract *LendingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// LendingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingPoolTransactorRaw struct {
	Contract *LendingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingPool creates a new instance of LendingPool, bound to a specific deployed contract.
func NewLendingPool(address common.Address, backend bind.ContractBackend) (*LendingPool, error) {
	contract, err := bindLendingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingPool{LendingPoolCaller: LendingPoolCaller{contract: contract}, LendingPoolTransactor: LendingPoolTransactor{contract: contract}, LendingPoolFilterer: LendingPoolFilterer{contract: contract}}, nil
}

// NewLendingPoolCaller creates a new read-only instance of LendingPool, bound to a specific deployed contract.
func NewLendingPoolCaller(address common.Address, caller bind.ContractCaller) (*LendingPoolCaller, error) {
	contract, err := bindLendingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolCaller{contract: contract}, nil
}

// NewLendingPoolTransactor creates a new write-only instance of LendingPool, bound to a specific deployed contract.
func NewLendingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingPoolTransactor, error) {
	contract, err := bindLendingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolTransactor{contract: contract}, nil
}

// NewLendingPoolFilterer creates a new log filterer instance of LendingPool, bound to a specific deployed contract.
func NewLendingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingPoolFilterer, error) {
	contract, err := bindLendingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingPoolFilterer{contract: contract}, nil
}

// bindLendingPool binds a generic wrapper to an already deployed contract.
func bindLendingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LendingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPool *LendingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPool.Contract.LendingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPool *LendingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPool.Contract.LendingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPool *LendingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPool.Contract.LendingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPool *LendingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPool *LendingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPool *LendingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPool.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolSession) ADMINROLE() ([32]byte, error) {
	return _LendingPool.Contract.ADMINROLE(&_LendingPool.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCallerSession) ADMINROLE() ([32]byte, error) {
	return _LendingPool.Contract.ADMINROLE(&_LendingPool.CallOpts)
}

// ATTESTATIONMAXAGE is a free data retrieval call binding the contract method 0x9aec990e.
//
// Solidity: function ATTESTATION_MAX_AGE() view returns(uint256)
func (_LendingPool *LendingPoolCaller) ATTESTATIONMAXAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "ATTESTATION_MAX_AGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ATTESTATIONMAXAGE is a free data retrieval call binding the contract method 0x9aec990e.
//
// Solidity: function ATTESTATION_MAX_AGE() view returns(uint256)
func (_LendingPool *LendingPoolSession) ATTESTATIONMAXAGE() (*big.Int, error) {
	return _LendingPool.Contract.ATTESTATIONMAXAGE(&_LendingPool.CallOpts)
}

// ATTESTATIONMAXAGE is a free data retrieval call binding the contract method 0x9aec990e.
//
// Solidity: function ATTESTATION_MAX_AGE() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) ATTESTATIONMAXAGE() (*big.Int, error) {
	return _LendingPool.Contract.ATTESTATIONMAXAGE(&_LendingPool.CallOpts)
}

// CURATORROLE is a free data retrieval call binding the contract method 0x4a2cfb54.
//
// Solidity: function CURATOR_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCaller) CURATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "CURATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CURATORROLE is a free data retrieval call binding the contract method 0x4a2cfb54.
//
// Solidity: function CURATOR_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolSession) CURATORROLE() ([32]byte, error) {
	return _LendingPool.Contract.CURATORROLE(&_LendingPool.CallOpts)
}

// CURATORROLE is a free data retrieval call binding the contract method 0x4a2cfb54.
//
// Solidity: function CURATOR_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCallerSession) CURATORROLE() ([32]byte, error) {
	return _LendingPool.Contract.CURATORROLE(&_LendingPool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LendingPool.Contract.DEFAULTADMINROLE(&_LendingPool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LendingPool.Contract.DEFAULTADMINROLE(&_LendingPool.CallOpts)
}

// LIQUIDATORROLE is a free data retrieval call binding the contract method 0x16d8887a.
//
// Solidity: function LIQUIDATOR_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCaller) LIQUIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "LIQUIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LIQUIDATORROLE is a free data retrieval call binding the contract method 0x16d8887a.
//
// Solidity: function LIQUIDATOR_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolSession) LIQUIDATORROLE() ([32]byte, error) {
	return _LendingPool.Contract.LIQUIDATORROLE(&_LendingPool.CallOpts)
}

// LIQUIDATORROLE is a free data retrieval call binding the contract method 0x16d8887a.
//
// Solidity: function LIQUIDATOR_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCallerSession) LIQUIDATORROLE() ([32]byte, error) {
	return _LendingPool.Contract.LIQUIDATORROLE(&_LendingPool.CallOpts)
}

// MINHEALTHSCORE is a free data retrieval call binding the contract method 0xd1ceda4f.
//
// Solidity: function MIN_HEALTH_SCORE() view returns(uint8)
func (_LendingPool *LendingPoolCaller) MINHEALTHSCORE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "MIN_HEALTH_SCORE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINHEALTHSCORE is a free data retrieval call binding the contract method 0xd1ceda4f.
//
// Solidity: function MIN_HEALTH_SCORE() view returns(uint8)
func (_LendingPool *LendingPoolSession) MINHEALTHSCORE() (uint8, error) {
	return _LendingPool.Contract.MINHEALTHSCORE(&_LendingPool.CallOpts)
}

// MINHEALTHSCORE is a free data retrieval call binding the contract method 0xd1ceda4f.
//
// Solidity: function MIN_HEALTH_SCORE() view returns(uint8)
func (_LendingPool *LendingPoolCallerSession) MINHEALTHSCORE() (uint8, error) {
	return _LendingPool.Contract.MINHEALTHSCORE(&_LendingPool.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolSession) PAUSERROLE() ([32]byte, error) {
	return _LendingPool.Contract.PAUSERROLE(&_LendingPool.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCallerSession) PAUSERROLE() ([32]byte, error) {
	return _LendingPool.Contract.PAUSERROLE(&_LendingPool.CallOpts)
}

// PROTOCOLFEEBPS is a free data retrieval call binding the contract method 0xbe378228.
//
// Solidity: function PROTOCOL_FEE_BPS() view returns(uint256)
func (_LendingPool *LendingPoolCaller) PROTOCOLFEEBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "PROTOCOL_FEE_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROTOCOLFEEBPS is a free data retrieval call binding the contract method 0xbe378228.
//
// Solidity: function PROTOCOL_FEE_BPS() view returns(uint256)
func (_LendingPool *LendingPoolSession) PROTOCOLFEEBPS() (*big.Int, error) {
	return _LendingPool.Contract.PROTOCOLFEEBPS(&_LendingPool.CallOpts)
}

// PROTOCOLFEEBPS is a free data retrieval call binding the contract method 0xbe378228.
//
// Solidity: function PROTOCOL_FEE_BPS() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) PROTOCOLFEEBPS() (*big.Int, error) {
	return _LendingPool.Contract.PROTOCOLFEEBPS(&_LendingPool.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolSession) UPGRADERROLE() ([32]byte, error) {
	return _LendingPool.Contract.UPGRADERROLE(&_LendingPool.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_LendingPool *LendingPoolCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _LendingPool.Contract.UPGRADERROLE(&_LendingPool.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LendingPool *LendingPoolCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LendingPool *LendingPoolSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LendingPool.Contract.UPGRADEINTERFACEVERSION(&_LendingPool.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LendingPool *LendingPoolCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LendingPool.Contract.UPGRADEINTERFACEVERSION(&_LendingPool.CallOpts)
}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_LendingPool *LendingPoolCaller) AssetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "assetRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_LendingPool *LendingPoolSession) AssetRegistry() (common.Address, error) {
	return _LendingPool.Contract.AssetRegistry(&_LendingPool.CallOpts)
}

// AssetRegistry is a free data retrieval call binding the contract method 0x979d7e86.
//
// Solidity: function assetRegistry() view returns(address)
func (_LendingPool *LendingPoolCallerSession) AssetRegistry() (common.Address, error) {
	return _LendingPool.Contract.AssetRegistry(&_LendingPool.CallOpts)
}

// GetBorrowerLoans is a free data retrieval call binding the contract method 0xeee8b7ff.
//
// Solidity: function getBorrowerLoans(address borrower) view returns(bytes32[])
func (_LendingPool *LendingPoolCaller) GetBorrowerLoans(opts *bind.CallOpts, borrower common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "getBorrowerLoans", borrower)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBorrowerLoans is a free data retrieval call binding the contract method 0xeee8b7ff.
//
// Solidity: function getBorrowerLoans(address borrower) view returns(bytes32[])
func (_LendingPool *LendingPoolSession) GetBorrowerLoans(borrower common.Address) ([][32]byte, error) {
	return _LendingPool.Contract.GetBorrowerLoans(&_LendingPool.CallOpts, borrower)
}

// GetBorrowerLoans is a free data retrieval call binding the contract method 0xeee8b7ff.
//
// Solidity: function getBorrowerLoans(address borrower) view returns(bytes32[])
func (_LendingPool *LendingPoolCallerSession) GetBorrowerLoans(borrower common.Address) ([][32]byte, error) {
	return _LendingPool.Contract.GetBorrowerLoans(&_LendingPool.CallOpts, borrower)
}

// GetLoan is a free data retrieval call binding the contract method 0x8932f5f7.
//
// Solidity: function getLoan(bytes32 loanId) view returns((bytes32,bytes32,address,address,uint256,uint256,uint256,uint8,uint8,uint256,uint256))
func (_LendingPool *LendingPoolCaller) GetLoan(opts *bind.CallOpts, loanId [32]byte) (ILendingPoolLoan, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "getLoan", loanId)

	if err != nil {
		return *new(ILendingPoolLoan), err
	}

	out0 := *abi.ConvertType(out[0], new(ILendingPoolLoan)).(*ILendingPoolLoan)

	return out0, err

}

// GetLoan is a free data retrieval call binding the contract method 0x8932f5f7.
//
// Solidity: function getLoan(bytes32 loanId) view returns((bytes32,bytes32,address,address,uint256,uint256,uint256,uint8,uint8,uint256,uint256))
func (_LendingPool *LendingPoolSession) GetLoan(loanId [32]byte) (ILendingPoolLoan, error) {
	return _LendingPool.Contract.GetLoan(&_LendingPool.CallOpts, loanId)
}

// GetLoan is a free data retrieval call binding the contract method 0x8932f5f7.
//
// Solidity: function getLoan(bytes32 loanId) view returns((bytes32,bytes32,address,address,uint256,uint256,uint256,uint8,uint8,uint256,uint256))
func (_LendingPool *LendingPoolCallerSession) GetLoan(loanId [32]byte) (ILendingPoolLoan, error) {
	return _LendingPool.Contract.GetLoan(&_LendingPool.CallOpts, loanId)
}

// GetProtocolFees is a free data retrieval call binding the contract method 0xd8dfcea0.
//
// Solidity: function getProtocolFees() view returns(uint256)
func (_LendingPool *LendingPoolCaller) GetProtocolFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "getProtocolFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFees is a free data retrieval call binding the contract method 0xd8dfcea0.
//
// Solidity: function getProtocolFees() view returns(uint256)
func (_LendingPool *LendingPoolSession) GetProtocolFees() (*big.Int, error) {
	return _LendingPool.Contract.GetProtocolFees(&_LendingPool.CallOpts)
}

// GetProtocolFees is a free data retrieval call binding the contract method 0xd8dfcea0.
//
// Solidity: function getProtocolFees() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) GetProtocolFees() (*big.Int, error) {
	return _LendingPool.Contract.GetProtocolFees(&_LendingPool.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LendingPool *LendingPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LendingPool *LendingPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LendingPool.Contract.GetRoleAdmin(&_LendingPool.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LendingPool *LendingPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LendingPool.Contract.GetRoleAdmin(&_LendingPool.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LendingPool *LendingPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LendingPool *LendingPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LendingPool.Contract.HasRole(&_LendingPool.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LendingPool *LendingPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LendingPool.Contract.HasRole(&_LendingPool.CallOpts, role, account)
}

// HealthAttestation is a free data retrieval call binding the contract method 0x1997b0a1.
//
// Solidity: function healthAttestation() view returns(address)
func (_LendingPool *LendingPoolCaller) HealthAttestation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "healthAttestation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HealthAttestation is a free data retrieval call binding the contract method 0x1997b0a1.
//
// Solidity: function healthAttestation() view returns(address)
func (_LendingPool *LendingPoolSession) HealthAttestation() (common.Address, error) {
	return _LendingPool.Contract.HealthAttestation(&_LendingPool.CallOpts)
}

// HealthAttestation is a free data retrieval call binding the contract method 0x1997b0a1.
//
// Solidity: function healthAttestation() view returns(address)
func (_LendingPool *LendingPoolCallerSession) HealthAttestation() (common.Address, error) {
	return _LendingPool.Contract.HealthAttestation(&_LendingPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LendingPool *LendingPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LendingPool *LendingPoolSession) Paused() (bool, error) {
	return _LendingPool.Contract.Paused(&_LendingPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LendingPool *LendingPoolCallerSession) Paused() (bool, error) {
	return _LendingPool.Contract.Paused(&_LendingPool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LendingPool *LendingPoolCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LendingPool *LendingPoolSession) ProxiableUUID() ([32]byte, error) {
	return _LendingPool.Contract.ProxiableUUID(&_LendingPool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LendingPool *LendingPoolCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LendingPool.Contract.ProxiableUUID(&_LendingPool.CallOpts)
}

// SWattUSD is a free data retrieval call binding the contract method 0xe02fd053.
//
// Solidity: function sWattUSD() view returns(address)
func (_LendingPool *LendingPoolCaller) SWattUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "sWattUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SWattUSD is a free data retrieval call binding the contract method 0xe02fd053.
//
// Solidity: function sWattUSD() view returns(address)
func (_LendingPool *LendingPoolSession) SWattUSD() (common.Address, error) {
	return _LendingPool.Contract.SWattUSD(&_LendingPool.CallOpts)
}

// SWattUSD is a free data retrieval call binding the contract method 0xe02fd053.
//
// Solidity: function sWattUSD() view returns(address)
func (_LendingPool *LendingPoolCallerSession) SWattUSD() (common.Address, error) {
	return _LendingPool.Contract.SWattUSD(&_LendingPool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LendingPool *LendingPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LendingPool *LendingPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LendingPool.Contract.SupportsInterface(&_LendingPool.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LendingPool *LendingPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LendingPool.Contract.SupportsInterface(&_LendingPool.CallOpts, interfaceId)
}

// WattUSD is a free data retrieval call binding the contract method 0x7aff9566.
//
// Solidity: function wattUSD() view returns(address)
func (_LendingPool *LendingPoolCaller) WattUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "wattUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WattUSD is a free data retrieval call binding the contract method 0x7aff9566.
//
// Solidity: function wattUSD() view returns(address)
func (_LendingPool *LendingPoolSession) WattUSD() (common.Address, error) {
	return _LendingPool.Contract.WattUSD(&_LendingPool.CallOpts)
}

// WattUSD is a free data retrieval call binding the contract method 0x7aff9566.
//
// Solidity: function wattUSD() view returns(address)
func (_LendingPool *LendingPoolCallerSession) WattUSD() (common.Address, error) {
	return _LendingPool.Contract.WattUSD(&_LendingPool.CallOpts)
}

// FlagDefaulted is a paid mutator transaction binding the contract method 0xb6742d97.
//
// Solidity: function flagDefaulted(bytes32 loanId) returns()
func (_LendingPool *LendingPoolTransactor) FlagDefaulted(opts *bind.TransactOpts, loanId [32]byte) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "flagDefaulted", loanId)
}

// FlagDefaulted is a paid mutator transaction binding the contract method 0xb6742d97.
//
// Solidity: function flagDefaulted(bytes32 loanId) returns()
func (_LendingPool *LendingPoolSession) FlagDefaulted(loanId [32]byte) (*types.Transaction, error) {
	return _LendingPool.Contract.FlagDefaulted(&_LendingPool.TransactOpts, loanId)
}

// FlagDefaulted is a paid mutator transaction binding the contract method 0xb6742d97.
//
// Solidity: function flagDefaulted(bytes32 loanId) returns()
func (_LendingPool *LendingPoolTransactorSession) FlagDefaulted(loanId [32]byte) (*types.Transaction, error) {
	return _LendingPool.Contract.FlagDefaulted(&_LendingPool.TransactOpts, loanId)
}

// FullRepay is a paid mutator transaction binding the contract method 0x92f3ac8a.
//
// Solidity: function fullRepay(bytes32 loanId) returns()
func (_LendingPool *LendingPoolTransactor) FullRepay(opts *bind.TransactOpts, loanId [32]byte) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "fullRepay", loanId)
}

// FullRepay is a paid mutator transaction binding the contract method 0x92f3ac8a.
//
// Solidity: function fullRepay(bytes32 loanId) returns()
func (_LendingPool *LendingPoolSession) FullRepay(loanId [32]byte) (*types.Transaction, error) {
	return _LendingPool.Contract.FullRepay(&_LendingPool.TransactOpts, loanId)
}

// FullRepay is a paid mutator transaction binding the contract method 0x92f3ac8a.
//
// Solidity: function fullRepay(bytes32 loanId) returns()
func (_LendingPool *LendingPoolTransactorSession) FullRepay(loanId [32]byte) (*types.Transaction, error) {
	return _LendingPool.Contract.FullRepay(&_LendingPool.TransactOpts, loanId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LendingPool *LendingPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LendingPool *LendingPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.GrantRole(&_LendingPool.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LendingPool *LendingPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.GrantRole(&_LendingPool.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address admin, address assetRegistry_, address healthAttestation_, address wattUSD_, address sWattUSD_) returns()
func (_LendingPool *LendingPoolTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, assetRegistry_ common.Address, healthAttestation_ common.Address, wattUSD_ common.Address, sWattUSD_ common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "initialize", admin, assetRegistry_, healthAttestation_, wattUSD_, sWattUSD_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address admin, address assetRegistry_, address healthAttestation_, address wattUSD_, address sWattUSD_) returns()
func (_LendingPool *LendingPoolSession) Initialize(admin common.Address, assetRegistry_ common.Address, healthAttestation_ common.Address, wattUSD_ common.Address, sWattUSD_ common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.Initialize(&_LendingPool.TransactOpts, admin, assetRegistry_, healthAttestation_, wattUSD_, sWattUSD_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address admin, address assetRegistry_, address healthAttestation_, address wattUSD_, address sWattUSD_) returns()
func (_LendingPool *LendingPoolTransactorSession) Initialize(admin common.Address, assetRegistry_ common.Address, healthAttestation_ common.Address, wattUSD_ common.Address, sWattUSD_ common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.Initialize(&_LendingPool.TransactOpts, admin, assetRegistry_, healthAttestation_, wattUSD_, sWattUSD_)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0a71096e.
//
// Solidity: function liquidate(bytes32 loanId) returns()
func (_LendingPool *LendingPoolTransactor) Liquidate(opts *bind.TransactOpts, loanId [32]byte) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "liquidate", loanId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0a71096e.
//
// Solidity: function liquidate(bytes32 loanId) returns()
func (_LendingPool *LendingPoolSession) Liquidate(loanId [32]byte) (*types.Transaction, error) {
	return _LendingPool.Contract.Liquidate(&_LendingPool.TransactOpts, loanId)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0a71096e.
//
// Solidity: function liquidate(bytes32 loanId) returns()
func (_LendingPool *LendingPoolTransactorSession) Liquidate(loanId [32]byte) (*types.Transaction, error) {
	return _LendingPool.Contract.Liquidate(&_LendingPool.TransactOpts, loanId)
}

// OriginateLoan is a paid mutator transaction binding the contract method 0x9a334a61.
//
// Solidity: function originateLoan(bytes32 assetId, address borrower, uint256 principal, uint256 interestRate, uint256 termDays, uint8 engineType) returns(bytes32 loanId)
func (_LendingPool *LendingPoolTransactor) OriginateLoan(opts *bind.TransactOpts, assetId [32]byte, borrower common.Address, principal *big.Int, interestRate *big.Int, termDays *big.Int, engineType uint8) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "originateLoan", assetId, borrower, principal, interestRate, termDays, engineType)
}

// OriginateLoan is a paid mutator transaction binding the contract method 0x9a334a61.
//
// Solidity: function originateLoan(bytes32 assetId, address borrower, uint256 principal, uint256 interestRate, uint256 termDays, uint8 engineType) returns(bytes32 loanId)
func (_LendingPool *LendingPoolSession) OriginateLoan(assetId [32]byte, borrower common.Address, principal *big.Int, interestRate *big.Int, termDays *big.Int, engineType uint8) (*types.Transaction, error) {
	return _LendingPool.Contract.OriginateLoan(&_LendingPool.TransactOpts, assetId, borrower, principal, interestRate, termDays, engineType)
}

// OriginateLoan is a paid mutator transaction binding the contract method 0x9a334a61.
//
// Solidity: function originateLoan(bytes32 assetId, address borrower, uint256 principal, uint256 interestRate, uint256 termDays, uint8 engineType) returns(bytes32 loanId)
func (_LendingPool *LendingPoolTransactorSession) OriginateLoan(assetId [32]byte, borrower common.Address, principal *big.Int, interestRate *big.Int, termDays *big.Int, engineType uint8) (*types.Transaction, error) {
	return _LendingPool.Contract.OriginateLoan(&_LendingPool.TransactOpts, assetId, borrower, principal, interestRate, termDays, engineType)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LendingPool *LendingPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LendingPool *LendingPoolSession) Pause() (*types.Transaction, error) {
	return _LendingPool.Contract.Pause(&_LendingPool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LendingPool *LendingPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _LendingPool.Contract.Pause(&_LendingPool.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LendingPool *LendingPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LendingPool *LendingPoolSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.RenounceRole(&_LendingPool.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LendingPool *LendingPoolTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.RenounceRole(&_LendingPool.TransactOpts, role, callerConfirmation)
}

// Repay is a paid mutator transaction binding the contract method 0x00f989ad.
//
// Solidity: function repay(bytes32 loanId, uint256 amount) returns()
func (_LendingPool *LendingPoolTransactor) Repay(opts *bind.TransactOpts, loanId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "repay", loanId, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x00f989ad.
//
// Solidity: function repay(bytes32 loanId, uint256 amount) returns()
func (_LendingPool *LendingPoolSession) Repay(loanId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.Repay(&_LendingPool.TransactOpts, loanId, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x00f989ad.
//
// Solidity: function repay(bytes32 loanId, uint256 amount) returns()
func (_LendingPool *LendingPoolTransactorSession) Repay(loanId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.Repay(&_LendingPool.TransactOpts, loanId, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LendingPool *LendingPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LendingPool *LendingPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.RevokeRole(&_LendingPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LendingPool *LendingPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.RevokeRole(&_LendingPool.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LendingPool *LendingPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LendingPool *LendingPoolSession) Unpause() (*types.Transaction, error) {
	return _LendingPool.Contract.Unpause(&_LendingPool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LendingPool *LendingPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _LendingPool.Contract.Unpause(&_LendingPool.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LendingPool *LendingPoolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LendingPool *LendingPoolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LendingPool.Contract.UpgradeToAndCall(&_LendingPool.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LendingPool *LendingPoolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LendingPool.Contract.UpgradeToAndCall(&_LendingPool.TransactOpts, newImplementation, data)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address to) returns()
func (_LendingPool *LendingPoolTransactor) WithdrawFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "withdrawFees", to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address to) returns()
func (_LendingPool *LendingPoolSession) WithdrawFees(to common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.WithdrawFees(&_LendingPool.TransactOpts, to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address to) returns()
func (_LendingPool *LendingPoolTransactorSession) WithdrawFees(to common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.WithdrawFees(&_LendingPool.TransactOpts, to)
}

// LendingPoolFeesWithdrawnIterator is returned from FilterFeesWithdrawn and is used to iterate over the raw logs and unpacked data for FeesWithdrawn events raised by the LendingPool contract.
type LendingPoolFeesWithdrawnIterator struct {
	Event *LendingPoolFeesWithdrawn // Event containing the contract specifics and raw log

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
func (it *LendingPoolFeesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolFeesWithdrawn)
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
		it.Event = new(LendingPoolFeesWithdrawn)
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
func (it *LendingPoolFeesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolFeesWithdrawn represents a FeesWithdrawn event raised by the LendingPool contract.
type LendingPoolFeesWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeesWithdrawn is a free log retrieval operation binding the contract event 0xc0819c13be868895eb93e40eaceb96de976442fa1d404e5c55f14bb65a8c489a.
//
// Solidity: event FeesWithdrawn(address indexed to, uint256 amount)
func (_LendingPool *LendingPoolFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts, to []common.Address) (*LendingPoolFeesWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "FeesWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolFeesWithdrawnIterator{contract: _LendingPool.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeesWithdrawn is a free log subscription operation binding the contract event 0xc0819c13be868895eb93e40eaceb96de976442fa1d404e5c55f14bb65a8c489a.
//
// Solidity: event FeesWithdrawn(address indexed to, uint256 amount)
func (_LendingPool *LendingPoolFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *LendingPoolFeesWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "FeesWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolFeesWithdrawn)
				if err := _LendingPool.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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
func (_LendingPool *LendingPoolFilterer) ParseFeesWithdrawn(log types.Log) (*LendingPoolFeesWithdrawn, error) {
	event := new(LendingPoolFeesWithdrawn)
	if err := _LendingPool.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LendingPool contract.
type LendingPoolInitializedIterator struct {
	Event *LendingPoolInitialized // Event containing the contract specifics and raw log

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
func (it *LendingPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolInitialized)
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
		it.Event = new(LendingPoolInitialized)
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
func (it *LendingPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolInitialized represents a Initialized event raised by the LendingPool contract.
type LendingPoolInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LendingPool *LendingPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*LendingPoolInitializedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LendingPoolInitializedIterator{contract: _LendingPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LendingPool *LendingPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LendingPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolInitialized)
				if err := _LendingPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LendingPool *LendingPoolFilterer) ParseInitialized(log types.Log) (*LendingPoolInitialized, error) {
	event := new(LendingPoolInitialized)
	if err := _LendingPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolLoanDefaultedIterator is returned from FilterLoanDefaulted and is used to iterate over the raw logs and unpacked data for LoanDefaulted events raised by the LendingPool contract.
type LendingPoolLoanDefaultedIterator struct {
	Event *LendingPoolLoanDefaulted // Event containing the contract specifics and raw log

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
func (it *LendingPoolLoanDefaultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolLoanDefaulted)
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
		it.Event = new(LendingPoolLoanDefaulted)
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
func (it *LendingPoolLoanDefaultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolLoanDefaultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolLoanDefaulted represents a LoanDefaulted event raised by the LendingPool contract.
type LendingPoolLoanDefaulted struct {
	LoanId   [32]byte
	Borrower common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLoanDefaulted is a free log retrieval operation binding the contract event 0x36cb5ae03feffae0ec99449f0264e31b98f44c687c34916193267f2073ed23c5.
//
// Solidity: event LoanDefaulted(bytes32 indexed loanId, address indexed borrower)
func (_LendingPool *LendingPoolFilterer) FilterLoanDefaulted(opts *bind.FilterOpts, loanId [][32]byte, borrower []common.Address) (*LendingPoolLoanDefaultedIterator, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "LoanDefaulted", loanIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLoanDefaultedIterator{contract: _LendingPool.contract, event: "LoanDefaulted", logs: logs, sub: sub}, nil
}

// WatchLoanDefaulted is a free log subscription operation binding the contract event 0x36cb5ae03feffae0ec99449f0264e31b98f44c687c34916193267f2073ed23c5.
//
// Solidity: event LoanDefaulted(bytes32 indexed loanId, address indexed borrower)
func (_LendingPool *LendingPoolFilterer) WatchLoanDefaulted(opts *bind.WatchOpts, sink chan<- *LendingPoolLoanDefaulted, loanId [][32]byte, borrower []common.Address) (event.Subscription, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "LoanDefaulted", loanIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolLoanDefaulted)
				if err := _LendingPool.contract.UnpackLog(event, "LoanDefaulted", log); err != nil {
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

// ParseLoanDefaulted is a log parse operation binding the contract event 0x36cb5ae03feffae0ec99449f0264e31b98f44c687c34916193267f2073ed23c5.
//
// Solidity: event LoanDefaulted(bytes32 indexed loanId, address indexed borrower)
func (_LendingPool *LendingPoolFilterer) ParseLoanDefaulted(log types.Log) (*LendingPoolLoanDefaulted, error) {
	event := new(LendingPoolLoanDefaulted)
	if err := _LendingPool.contract.UnpackLog(event, "LoanDefaulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolLoanLiquidatedIterator is returned from FilterLoanLiquidated and is used to iterate over the raw logs and unpacked data for LoanLiquidated events raised by the LendingPool contract.
type LendingPoolLoanLiquidatedIterator struct {
	Event *LendingPoolLoanLiquidated // Event containing the contract specifics and raw log

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
func (it *LendingPoolLoanLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolLoanLiquidated)
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
		it.Event = new(LendingPoolLoanLiquidated)
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
func (it *LendingPoolLoanLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolLoanLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolLoanLiquidated represents a LoanLiquidated event raised by the LendingPool contract.
type LendingPoolLoanLiquidated struct {
	LoanId     [32]byte
	Borrower   common.Address
	Liquidator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLoanLiquidated is a free log retrieval operation binding the contract event 0xa23ea239f34c12f95374b6b6e12d8f65517b6e98aae01e5cfe3be6905b795b9d.
//
// Solidity: event LoanLiquidated(bytes32 indexed loanId, address indexed borrower, address liquidator)
func (_LendingPool *LendingPoolFilterer) FilterLoanLiquidated(opts *bind.FilterOpts, loanId [][32]byte, borrower []common.Address) (*LendingPoolLoanLiquidatedIterator, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "LoanLiquidated", loanIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLoanLiquidatedIterator{contract: _LendingPool.contract, event: "LoanLiquidated", logs: logs, sub: sub}, nil
}

// WatchLoanLiquidated is a free log subscription operation binding the contract event 0xa23ea239f34c12f95374b6b6e12d8f65517b6e98aae01e5cfe3be6905b795b9d.
//
// Solidity: event LoanLiquidated(bytes32 indexed loanId, address indexed borrower, address liquidator)
func (_LendingPool *LendingPoolFilterer) WatchLoanLiquidated(opts *bind.WatchOpts, sink chan<- *LendingPoolLoanLiquidated, loanId [][32]byte, borrower []common.Address) (event.Subscription, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "LoanLiquidated", loanIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolLoanLiquidated)
				if err := _LendingPool.contract.UnpackLog(event, "LoanLiquidated", log); err != nil {
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

// ParseLoanLiquidated is a log parse operation binding the contract event 0xa23ea239f34c12f95374b6b6e12d8f65517b6e98aae01e5cfe3be6905b795b9d.
//
// Solidity: event LoanLiquidated(bytes32 indexed loanId, address indexed borrower, address liquidator)
func (_LendingPool *LendingPoolFilterer) ParseLoanLiquidated(log types.Log) (*LendingPoolLoanLiquidated, error) {
	event := new(LendingPoolLoanLiquidated)
	if err := _LendingPool.contract.UnpackLog(event, "LoanLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolLoanOriginatedIterator is returned from FilterLoanOriginated and is used to iterate over the raw logs and unpacked data for LoanOriginated events raised by the LendingPool contract.
type LendingPoolLoanOriginatedIterator struct {
	Event *LendingPoolLoanOriginated // Event containing the contract specifics and raw log

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
func (it *LendingPoolLoanOriginatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolLoanOriginated)
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
		it.Event = new(LendingPoolLoanOriginated)
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
func (it *LendingPoolLoanOriginatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolLoanOriginatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolLoanOriginated represents a LoanOriginated event raised by the LendingPool contract.
type LendingPoolLoanOriginated struct {
	LoanId       [32]byte
	AssetId      [32]byte
	Borrower     common.Address
	Curator      common.Address
	Principal    *big.Int
	InterestRate *big.Int
	MaturityAt   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLoanOriginated is a free log retrieval operation binding the contract event 0x36a923fe25811f9e4594ef7316c44431284532e670e96937067354b46e45a95c.
//
// Solidity: event LoanOriginated(bytes32 indexed loanId, bytes32 indexed assetId, address indexed borrower, address curator, uint256 principal, uint256 interestRate, uint256 maturityAt)
func (_LendingPool *LendingPoolFilterer) FilterLoanOriginated(opts *bind.FilterOpts, loanId [][32]byte, assetId [][32]byte, borrower []common.Address) (*LendingPoolLoanOriginatedIterator, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "LoanOriginated", loanIdRule, assetIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLoanOriginatedIterator{contract: _LendingPool.contract, event: "LoanOriginated", logs: logs, sub: sub}, nil
}

// WatchLoanOriginated is a free log subscription operation binding the contract event 0x36a923fe25811f9e4594ef7316c44431284532e670e96937067354b46e45a95c.
//
// Solidity: event LoanOriginated(bytes32 indexed loanId, bytes32 indexed assetId, address indexed borrower, address curator, uint256 principal, uint256 interestRate, uint256 maturityAt)
func (_LendingPool *LendingPoolFilterer) WatchLoanOriginated(opts *bind.WatchOpts, sink chan<- *LendingPoolLoanOriginated, loanId [][32]byte, assetId [][32]byte, borrower []common.Address) (event.Subscription, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "LoanOriginated", loanIdRule, assetIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolLoanOriginated)
				if err := _LendingPool.contract.UnpackLog(event, "LoanOriginated", log); err != nil {
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

// ParseLoanOriginated is a log parse operation binding the contract event 0x36a923fe25811f9e4594ef7316c44431284532e670e96937067354b46e45a95c.
//
// Solidity: event LoanOriginated(bytes32 indexed loanId, bytes32 indexed assetId, address indexed borrower, address curator, uint256 principal, uint256 interestRate, uint256 maturityAt)
func (_LendingPool *LendingPoolFilterer) ParseLoanOriginated(log types.Log) (*LendingPoolLoanOriginated, error) {
	event := new(LendingPoolLoanOriginated)
	if err := _LendingPool.contract.UnpackLog(event, "LoanOriginated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolLoanSettledIterator is returned from FilterLoanSettled and is used to iterate over the raw logs and unpacked data for LoanSettled events raised by the LendingPool contract.
type LendingPoolLoanSettledIterator struct {
	Event *LendingPoolLoanSettled // Event containing the contract specifics and raw log

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
func (it *LendingPoolLoanSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolLoanSettled)
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
		it.Event = new(LendingPoolLoanSettled)
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
func (it *LendingPoolLoanSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolLoanSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolLoanSettled represents a LoanSettled event raised by the LendingPool contract.
type LendingPoolLoanSettled struct {
	LoanId   [32]byte
	Borrower common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLoanSettled is a free log retrieval operation binding the contract event 0x5046121d1aca8511349b51b33b8251c8df15585547b26de5a72b6b7f53ad6653.
//
// Solidity: event LoanSettled(bytes32 indexed loanId, address indexed borrower)
func (_LendingPool *LendingPoolFilterer) FilterLoanSettled(opts *bind.FilterOpts, loanId [][32]byte, borrower []common.Address) (*LendingPoolLoanSettledIterator, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "LoanSettled", loanIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLoanSettledIterator{contract: _LendingPool.contract, event: "LoanSettled", logs: logs, sub: sub}, nil
}

// WatchLoanSettled is a free log subscription operation binding the contract event 0x5046121d1aca8511349b51b33b8251c8df15585547b26de5a72b6b7f53ad6653.
//
// Solidity: event LoanSettled(bytes32 indexed loanId, address indexed borrower)
func (_LendingPool *LendingPoolFilterer) WatchLoanSettled(opts *bind.WatchOpts, sink chan<- *LendingPoolLoanSettled, loanId [][32]byte, borrower []common.Address) (event.Subscription, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "LoanSettled", loanIdRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolLoanSettled)
				if err := _LendingPool.contract.UnpackLog(event, "LoanSettled", log); err != nil {
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

// ParseLoanSettled is a log parse operation binding the contract event 0x5046121d1aca8511349b51b33b8251c8df15585547b26de5a72b6b7f53ad6653.
//
// Solidity: event LoanSettled(bytes32 indexed loanId, address indexed borrower)
func (_LendingPool *LendingPoolFilterer) ParseLoanSettled(log types.Log) (*LendingPoolLoanSettled, error) {
	event := new(LendingPoolLoanSettled)
	if err := _LendingPool.contract.UnpackLog(event, "LoanSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LendingPool contract.
type LendingPoolPausedIterator struct {
	Event *LendingPoolPaused // Event containing the contract specifics and raw log

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
func (it *LendingPoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolPaused)
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
		it.Event = new(LendingPoolPaused)
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
func (it *LendingPoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolPaused represents a Paused event raised by the LendingPool contract.
type LendingPoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LendingPool *LendingPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*LendingPoolPausedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LendingPoolPausedIterator{contract: _LendingPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LendingPool *LendingPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LendingPoolPaused) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolPaused)
				if err := _LendingPool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LendingPool *LendingPoolFilterer) ParsePaused(log types.Log) (*LendingPoolPaused, error) {
	event := new(LendingPoolPaused)
	if err := _LendingPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolRepaymentReceivedIterator is returned from FilterRepaymentReceived and is used to iterate over the raw logs and unpacked data for RepaymentReceived events raised by the LendingPool contract.
type LendingPoolRepaymentReceivedIterator struct {
	Event *LendingPoolRepaymentReceived // Event containing the contract specifics and raw log

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
func (it *LendingPoolRepaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolRepaymentReceived)
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
		it.Event = new(LendingPoolRepaymentReceived)
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
func (it *LendingPoolRepaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolRepaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolRepaymentReceived represents a RepaymentReceived event raised by the LendingPool contract.
type LendingPoolRepaymentReceived struct {
	LoanId           [32]byte
	Payer            common.Address
	Amount           *big.Int
	PrincipalPaid    *big.Int
	InterestPaid     *big.Int
	OutstandingAfter *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRepaymentReceived is a free log retrieval operation binding the contract event 0xd90ea40b1e020448a0c99a0395517f850db856f7b23f7ae22d94deb2d183a027.
//
// Solidity: event RepaymentReceived(bytes32 indexed loanId, address indexed payer, uint256 amount, uint256 principalPaid, uint256 interestPaid, uint256 outstandingAfter)
func (_LendingPool *LendingPoolFilterer) FilterRepaymentReceived(opts *bind.FilterOpts, loanId [][32]byte, payer []common.Address) (*LendingPoolRepaymentReceivedIterator, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "RepaymentReceived", loanIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolRepaymentReceivedIterator{contract: _LendingPool.contract, event: "RepaymentReceived", logs: logs, sub: sub}, nil
}

// WatchRepaymentReceived is a free log subscription operation binding the contract event 0xd90ea40b1e020448a0c99a0395517f850db856f7b23f7ae22d94deb2d183a027.
//
// Solidity: event RepaymentReceived(bytes32 indexed loanId, address indexed payer, uint256 amount, uint256 principalPaid, uint256 interestPaid, uint256 outstandingAfter)
func (_LendingPool *LendingPoolFilterer) WatchRepaymentReceived(opts *bind.WatchOpts, sink chan<- *LendingPoolRepaymentReceived, loanId [][32]byte, payer []common.Address) (event.Subscription, error) {

	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "RepaymentReceived", loanIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolRepaymentReceived)
				if err := _LendingPool.contract.UnpackLog(event, "RepaymentReceived", log); err != nil {
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

// ParseRepaymentReceived is a log parse operation binding the contract event 0xd90ea40b1e020448a0c99a0395517f850db856f7b23f7ae22d94deb2d183a027.
//
// Solidity: event RepaymentReceived(bytes32 indexed loanId, address indexed payer, uint256 amount, uint256 principalPaid, uint256 interestPaid, uint256 outstandingAfter)
func (_LendingPool *LendingPoolFilterer) ParseRepaymentReceived(log types.Log) (*LendingPoolRepaymentReceived, error) {
	event := new(LendingPoolRepaymentReceived)
	if err := _LendingPool.contract.UnpackLog(event, "RepaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LendingPool contract.
type LendingPoolRoleAdminChangedIterator struct {
	Event *LendingPoolRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LendingPoolRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolRoleAdminChanged)
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
		it.Event = new(LendingPoolRoleAdminChanged)
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
func (it *LendingPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolRoleAdminChanged represents a RoleAdminChanged event raised by the LendingPool contract.
type LendingPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LendingPool *LendingPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LendingPoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolRoleAdminChangedIterator{contract: _LendingPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LendingPool *LendingPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LendingPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolRoleAdminChanged)
				if err := _LendingPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LendingPool *LendingPoolFilterer) ParseRoleAdminChanged(log types.Log) (*LendingPoolRoleAdminChanged, error) {
	event := new(LendingPoolRoleAdminChanged)
	if err := _LendingPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LendingPool contract.
type LendingPoolRoleGrantedIterator struct {
	Event *LendingPoolRoleGranted // Event containing the contract specifics and raw log

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
func (it *LendingPoolRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolRoleGranted)
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
		it.Event = new(LendingPoolRoleGranted)
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
func (it *LendingPoolRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolRoleGranted represents a RoleGranted event raised by the LendingPool contract.
type LendingPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LendingPool *LendingPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LendingPoolRoleGrantedIterator, error) {

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

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolRoleGrantedIterator{contract: _LendingPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LendingPool *LendingPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LendingPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolRoleGranted)
				if err := _LendingPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LendingPool *LendingPoolFilterer) ParseRoleGranted(log types.Log) (*LendingPoolRoleGranted, error) {
	event := new(LendingPoolRoleGranted)
	if err := _LendingPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LendingPool contract.
type LendingPoolRoleRevokedIterator struct {
	Event *LendingPoolRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LendingPoolRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolRoleRevoked)
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
		it.Event = new(LendingPoolRoleRevoked)
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
func (it *LendingPoolRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolRoleRevoked represents a RoleRevoked event raised by the LendingPool contract.
type LendingPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LendingPool *LendingPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LendingPoolRoleRevokedIterator, error) {

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

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolRoleRevokedIterator{contract: _LendingPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LendingPool *LendingPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LendingPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolRoleRevoked)
				if err := _LendingPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LendingPool *LendingPoolFilterer) ParseRoleRevoked(log types.Log) (*LendingPoolRoleRevoked, error) {
	event := new(LendingPoolRoleRevoked)
	if err := _LendingPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LendingPool contract.
type LendingPoolUnpausedIterator struct {
	Event *LendingPoolUnpaused // Event containing the contract specifics and raw log

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
func (it *LendingPoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolUnpaused)
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
		it.Event = new(LendingPoolUnpaused)
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
func (it *LendingPoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolUnpaused represents a Unpaused event raised by the LendingPool contract.
type LendingPoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LendingPool *LendingPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LendingPoolUnpausedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LendingPoolUnpausedIterator{contract: _LendingPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LendingPool *LendingPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LendingPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolUnpaused)
				if err := _LendingPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LendingPool *LendingPoolFilterer) ParseUnpaused(log types.Log) (*LendingPoolUnpaused, error) {
	event := new(LendingPoolUnpaused)
	if err := _LendingPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LendingPool contract.
type LendingPoolUpgradedIterator struct {
	Event *LendingPoolUpgraded // Event containing the contract specifics and raw log

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
func (it *LendingPoolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolUpgraded)
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
		it.Event = new(LendingPoolUpgraded)
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
func (it *LendingPoolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolUpgraded represents a Upgraded event raised by the LendingPool contract.
type LendingPoolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LendingPool *LendingPoolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LendingPoolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolUpgradedIterator{contract: _LendingPool.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LendingPool *LendingPoolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LendingPoolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolUpgraded)
				if err := _LendingPool.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_LendingPool *LendingPoolFilterer) ParseUpgraded(log types.Log) (*LendingPoolUpgraded, error) {
	event := new(LendingPoolUpgraded)
	if err := _LendingPool.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
