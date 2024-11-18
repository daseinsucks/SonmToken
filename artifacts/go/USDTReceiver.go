// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdtreceiver

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
)

// UsdtreceiverMetaData contains all meta data concerning the Usdtreceiver contract.
var UsdtreceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"feePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mnosToken\",\"outputs\":[{\"internalType\":\"contractMnosToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintForUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnAndSendUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mnosToken\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMinMintAmount\",\"type\":\"uint256\"}],\"name\":\"setMinMintAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinMintAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UsdtreceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use UsdtreceiverMetaData.ABI instead.
var UsdtreceiverABI = UsdtreceiverMetaData.ABI

// Usdtreceiver is an auto generated Go binding around an Ethereum contract.
type Usdtreceiver struct {
	UsdtreceiverCaller     // Read-only binding to the contract
	UsdtreceiverTransactor // Write-only binding to the contract
	UsdtreceiverFilterer   // Log filterer for contract events
}

// UsdtreceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdtreceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtreceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdtreceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtreceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdtreceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtreceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdtreceiverSession struct {
	Contract     *Usdtreceiver     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdtreceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdtreceiverCallerSession struct {
	Contract *UsdtreceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UsdtreceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdtreceiverTransactorSession struct {
	Contract     *UsdtreceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UsdtreceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdtreceiverRaw struct {
	Contract *Usdtreceiver // Generic contract binding to access the raw methods on
}

// UsdtreceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdtreceiverCallerRaw struct {
	Contract *UsdtreceiverCaller // Generic read-only contract binding to access the raw methods on
}

// UsdtreceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdtreceiverTransactorRaw struct {
	Contract *UsdtreceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdtreceiver creates a new instance of Usdtreceiver, bound to a specific deployed contract.
func NewUsdtreceiver(address common.Address, backend bind.ContractBackend) (*Usdtreceiver, error) {
	contract, err := bindUsdtreceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdtreceiver{UsdtreceiverCaller: UsdtreceiverCaller{contract: contract}, UsdtreceiverTransactor: UsdtreceiverTransactor{contract: contract}, UsdtreceiverFilterer: UsdtreceiverFilterer{contract: contract}}, nil
}

// NewUsdtreceiverCaller creates a new read-only instance of Usdtreceiver, bound to a specific deployed contract.
func NewUsdtreceiverCaller(address common.Address, caller bind.ContractCaller) (*UsdtreceiverCaller, error) {
	contract, err := bindUsdtreceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtreceiverCaller{contract: contract}, nil
}

// NewUsdtreceiverTransactor creates a new write-only instance of Usdtreceiver, bound to a specific deployed contract.
func NewUsdtreceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdtreceiverTransactor, error) {
	contract, err := bindUsdtreceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtreceiverTransactor{contract: contract}, nil
}

// NewUsdtreceiverFilterer creates a new log filterer instance of Usdtreceiver, bound to a specific deployed contract.
func NewUsdtreceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdtreceiverFilterer, error) {
	contract, err := bindUsdtreceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdtreceiverFilterer{contract: contract}, nil
}

// bindUsdtreceiver binds a generic wrapper to an already deployed contract.
func bindUsdtreceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdtreceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdtreceiver *UsdtreceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdtreceiver.Contract.UsdtreceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdtreceiver *UsdtreceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.UsdtreceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdtreceiver *UsdtreceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.UsdtreceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdtreceiver *UsdtreceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdtreceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdtreceiver *UsdtreceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdtreceiver *UsdtreceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.contract.Transact(opts, method, params...)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_Usdtreceiver *UsdtreceiverCaller) FeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdtreceiver.contract.Call(opts, &out, "feePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_Usdtreceiver *UsdtreceiverSession) FeePercentage() (*big.Int, error) {
	return _Usdtreceiver.Contract.FeePercentage(&_Usdtreceiver.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_Usdtreceiver *UsdtreceiverCallerSession) FeePercentage() (*big.Int, error) {
	return _Usdtreceiver.Contract.FeePercentage(&_Usdtreceiver.CallOpts)
}

// GetMinMintAmount is a free data retrieval call binding the contract method 0x50ea3900.
//
// Solidity: function getMinMintAmount() view returns(uint256)
func (_Usdtreceiver *UsdtreceiverCaller) GetMinMintAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdtreceiver.contract.Call(opts, &out, "getMinMintAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinMintAmount is a free data retrieval call binding the contract method 0x50ea3900.
//
// Solidity: function getMinMintAmount() view returns(uint256)
func (_Usdtreceiver *UsdtreceiverSession) GetMinMintAmount() (*big.Int, error) {
	return _Usdtreceiver.Contract.GetMinMintAmount(&_Usdtreceiver.CallOpts)
}

// GetMinMintAmount is a free data retrieval call binding the contract method 0x50ea3900.
//
// Solidity: function getMinMintAmount() view returns(uint256)
func (_Usdtreceiver *UsdtreceiverCallerSession) GetMinMintAmount() (*big.Int, error) {
	return _Usdtreceiver.Contract.GetMinMintAmount(&_Usdtreceiver.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Usdtreceiver *UsdtreceiverCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdtreceiver.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Usdtreceiver *UsdtreceiverSession) MintPrice() (*big.Int, error) {
	return _Usdtreceiver.Contract.MintPrice(&_Usdtreceiver.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Usdtreceiver *UsdtreceiverCallerSession) MintPrice() (*big.Int, error) {
	return _Usdtreceiver.Contract.MintPrice(&_Usdtreceiver.CallOpts)
}

// MnosToken is a free data retrieval call binding the contract method 0x8c3c2244.
//
// Solidity: function mnosToken() view returns(address)
func (_Usdtreceiver *UsdtreceiverCaller) MnosToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdtreceiver.contract.Call(opts, &out, "mnosToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MnosToken is a free data retrieval call binding the contract method 0x8c3c2244.
//
// Solidity: function mnosToken() view returns(address)
func (_Usdtreceiver *UsdtreceiverSession) MnosToken() (common.Address, error) {
	return _Usdtreceiver.Contract.MnosToken(&_Usdtreceiver.CallOpts)
}

// MnosToken is a free data retrieval call binding the contract method 0x8c3c2244.
//
// Solidity: function mnosToken() view returns(address)
func (_Usdtreceiver *UsdtreceiverCallerSession) MnosToken() (common.Address, error) {
	return _Usdtreceiver.Contract.MnosToken(&_Usdtreceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdtreceiver *UsdtreceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdtreceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdtreceiver *UsdtreceiverSession) Owner() (common.Address, error) {
	return _Usdtreceiver.Contract.Owner(&_Usdtreceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Usdtreceiver *UsdtreceiverCallerSession) Owner() (common.Address, error) {
	return _Usdtreceiver.Contract.Owner(&_Usdtreceiver.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Usdtreceiver *UsdtreceiverCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdtreceiver.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Usdtreceiver *UsdtreceiverSession) Treasury() (common.Address, error) {
	return _Usdtreceiver.Contract.Treasury(&_Usdtreceiver.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Usdtreceiver *UsdtreceiverCallerSession) Treasury() (common.Address, error) {
	return _Usdtreceiver.Contract.Treasury(&_Usdtreceiver.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Usdtreceiver *UsdtreceiverCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Usdtreceiver.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Usdtreceiver *UsdtreceiverSession) Usdt() (common.Address, error) {
	return _Usdtreceiver.Contract.Usdt(&_Usdtreceiver.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Usdtreceiver *UsdtreceiverCallerSession) Usdt() (common.Address, error) {
	return _Usdtreceiver.Contract.Usdt(&_Usdtreceiver.CallOpts)
}

// BurnAndSendUSDT is a paid mutator transaction binding the contract method 0x259470fa.
//
// Solidity: function burnAndSendUSDT(uint256 amount) returns()
func (_Usdtreceiver *UsdtreceiverTransactor) BurnAndSendUSDT(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.contract.Transact(opts, "burnAndSendUSDT", amount)
}

// BurnAndSendUSDT is a paid mutator transaction binding the contract method 0x259470fa.
//
// Solidity: function burnAndSendUSDT(uint256 amount) returns()
func (_Usdtreceiver *UsdtreceiverSession) BurnAndSendUSDT(amount *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.BurnAndSendUSDT(&_Usdtreceiver.TransactOpts, amount)
}

// BurnAndSendUSDT is a paid mutator transaction binding the contract method 0x259470fa.
//
// Solidity: function burnAndSendUSDT(uint256 amount) returns()
func (_Usdtreceiver *UsdtreceiverTransactorSession) BurnAndSendUSDT(amount *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.BurnAndSendUSDT(&_Usdtreceiver.TransactOpts, amount)
}

// MintForUSDT is a paid mutator transaction binding the contract method 0xabb47966.
//
// Solidity: function mintForUSDT(uint256 amount) returns()
func (_Usdtreceiver *UsdtreceiverTransactor) MintForUSDT(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.contract.Transact(opts, "mintForUSDT", amount)
}

// MintForUSDT is a paid mutator transaction binding the contract method 0xabb47966.
//
// Solidity: function mintForUSDT(uint256 amount) returns()
func (_Usdtreceiver *UsdtreceiverSession) MintForUSDT(amount *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.MintForUSDT(&_Usdtreceiver.TransactOpts, amount)
}

// MintForUSDT is a paid mutator transaction binding the contract method 0xabb47966.
//
// Solidity: function mintForUSDT(uint256 amount) returns()
func (_Usdtreceiver *UsdtreceiverTransactorSession) MintForUSDT(amount *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.MintForUSDT(&_Usdtreceiver.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Usdtreceiver *UsdtreceiverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdtreceiver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Usdtreceiver *UsdtreceiverSession) RenounceOwnership() (*types.Transaction, error) {
	return _Usdtreceiver.Contract.RenounceOwnership(&_Usdtreceiver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Usdtreceiver *UsdtreceiverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Usdtreceiver.Contract.RenounceOwnership(&_Usdtreceiver.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_Usdtreceiver *UsdtreceiverTransactor) SetFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.contract.Transact(opts, "setFee", _newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_Usdtreceiver *UsdtreceiverSession) SetFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetFee(&_Usdtreceiver.TransactOpts, _newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_Usdtreceiver *UsdtreceiverTransactorSession) SetFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetFee(&_Usdtreceiver.TransactOpts, _newFee)
}

// SetMinMintAmount is a paid mutator transaction binding the contract method 0x1d85d796.
//
// Solidity: function setMinMintAmount(uint256 _newMinMintAmount) returns()
func (_Usdtreceiver *UsdtreceiverTransactor) SetMinMintAmount(opts *bind.TransactOpts, _newMinMintAmount *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.contract.Transact(opts, "setMinMintAmount", _newMinMintAmount)
}

// SetMinMintAmount is a paid mutator transaction binding the contract method 0x1d85d796.
//
// Solidity: function setMinMintAmount(uint256 _newMinMintAmount) returns()
func (_Usdtreceiver *UsdtreceiverSession) SetMinMintAmount(_newMinMintAmount *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetMinMintAmount(&_Usdtreceiver.TransactOpts, _newMinMintAmount)
}

// SetMinMintAmount is a paid mutator transaction binding the contract method 0x1d85d796.
//
// Solidity: function setMinMintAmount(uint256 _newMinMintAmount) returns()
func (_Usdtreceiver *UsdtreceiverTransactorSession) SetMinMintAmount(_newMinMintAmount *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetMinMintAmount(&_Usdtreceiver.TransactOpts, _newMinMintAmount)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _newPrice) returns()
func (_Usdtreceiver *UsdtreceiverTransactor) SetMintPrice(opts *bind.TransactOpts, _newPrice *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.contract.Transact(opts, "setMintPrice", _newPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _newPrice) returns()
func (_Usdtreceiver *UsdtreceiverSession) SetMintPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetMintPrice(&_Usdtreceiver.TransactOpts, _newPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _newPrice) returns()
func (_Usdtreceiver *UsdtreceiverTransactorSession) SetMintPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetMintPrice(&_Usdtreceiver.TransactOpts, _newPrice)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _mnosToken) returns()
func (_Usdtreceiver *UsdtreceiverTransactor) SetToken(opts *bind.TransactOpts, _mnosToken common.Address) (*types.Transaction, error) {
	return _Usdtreceiver.contract.Transact(opts, "setToken", _mnosToken)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _mnosToken) returns()
func (_Usdtreceiver *UsdtreceiverSession) SetToken(_mnosToken common.Address) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetToken(&_Usdtreceiver.TransactOpts, _mnosToken)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _mnosToken) returns()
func (_Usdtreceiver *UsdtreceiverTransactorSession) SetToken(_mnosToken common.Address) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetToken(&_Usdtreceiver.TransactOpts, _mnosToken)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_Usdtreceiver *UsdtreceiverTransactor) SetTreasury(opts *bind.TransactOpts, _newTreasury common.Address) (*types.Transaction, error) {
	return _Usdtreceiver.contract.Transact(opts, "setTreasury", _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_Usdtreceiver *UsdtreceiverSession) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetTreasury(&_Usdtreceiver.TransactOpts, _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_Usdtreceiver *UsdtreceiverTransactorSession) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.SetTreasury(&_Usdtreceiver.TransactOpts, _newTreasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdtreceiver *UsdtreceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Usdtreceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdtreceiver *UsdtreceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.TransferOwnership(&_Usdtreceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Usdtreceiver *UsdtreceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Usdtreceiver.Contract.TransferOwnership(&_Usdtreceiver.TransactOpts, newOwner)
}

// UsdtreceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Usdtreceiver contract.
type UsdtreceiverOwnershipTransferredIterator struct {
	Event *UsdtreceiverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UsdtreceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtreceiverOwnershipTransferred)
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
		it.Event = new(UsdtreceiverOwnershipTransferred)
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
func (it *UsdtreceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtreceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtreceiverOwnershipTransferred represents a OwnershipTransferred event raised by the Usdtreceiver contract.
type UsdtreceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Usdtreceiver *UsdtreceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UsdtreceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Usdtreceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UsdtreceiverOwnershipTransferredIterator{contract: _Usdtreceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Usdtreceiver *UsdtreceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UsdtreceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Usdtreceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtreceiverOwnershipTransferred)
				if err := _Usdtreceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Usdtreceiver *UsdtreceiverFilterer) ParseOwnershipTransferred(log types.Log) (*UsdtreceiverOwnershipTransferred, error) {
	event := new(UsdtreceiverOwnershipTransferred)
	if err := _Usdtreceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
