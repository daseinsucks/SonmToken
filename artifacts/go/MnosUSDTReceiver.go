// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mnosusdtreceiver

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

// MnosusdtreceiverMetaData contains all meta data concerning the Mnosusdtreceiver contract.
var MnosusdtreceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnAndSendUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintForUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mnosToken\",\"outputs\":[{\"internalType\":\"contractMnosToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mnosToken\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MnosusdtreceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MnosusdtreceiverMetaData.ABI instead.
var MnosusdtreceiverABI = MnosusdtreceiverMetaData.ABI

// Mnosusdtreceiver is an auto generated Go binding around an Ethereum contract.
type Mnosusdtreceiver struct {
	MnosusdtreceiverCaller     // Read-only binding to the contract
	MnosusdtreceiverTransactor // Write-only binding to the contract
	MnosusdtreceiverFilterer   // Log filterer for contract events
}

// MnosusdtreceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MnosusdtreceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MnosusdtreceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MnosusdtreceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MnosusdtreceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MnosusdtreceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MnosusdtreceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MnosusdtreceiverSession struct {
	Contract     *Mnosusdtreceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MnosusdtreceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MnosusdtreceiverCallerSession struct {
	Contract *MnosusdtreceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MnosusdtreceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MnosusdtreceiverTransactorSession struct {
	Contract     *MnosusdtreceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MnosusdtreceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MnosusdtreceiverRaw struct {
	Contract *Mnosusdtreceiver // Generic contract binding to access the raw methods on
}

// MnosusdtreceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MnosusdtreceiverCallerRaw struct {
	Contract *MnosusdtreceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MnosusdtreceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MnosusdtreceiverTransactorRaw struct {
	Contract *MnosusdtreceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMnosusdtreceiver creates a new instance of Mnosusdtreceiver, bound to a specific deployed contract.
func NewMnosusdtreceiver(address common.Address, backend bind.ContractBackend) (*Mnosusdtreceiver, error) {
	contract, err := bindMnosusdtreceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mnosusdtreceiver{MnosusdtreceiverCaller: MnosusdtreceiverCaller{contract: contract}, MnosusdtreceiverTransactor: MnosusdtreceiverTransactor{contract: contract}, MnosusdtreceiverFilterer: MnosusdtreceiverFilterer{contract: contract}}, nil
}

// NewMnosusdtreceiverCaller creates a new read-only instance of Mnosusdtreceiver, bound to a specific deployed contract.
func NewMnosusdtreceiverCaller(address common.Address, caller bind.ContractCaller) (*MnosusdtreceiverCaller, error) {
	contract, err := bindMnosusdtreceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MnosusdtreceiverCaller{contract: contract}, nil
}

// NewMnosusdtreceiverTransactor creates a new write-only instance of Mnosusdtreceiver, bound to a specific deployed contract.
func NewMnosusdtreceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MnosusdtreceiverTransactor, error) {
	contract, err := bindMnosusdtreceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MnosusdtreceiverTransactor{contract: contract}, nil
}

// NewMnosusdtreceiverFilterer creates a new log filterer instance of Mnosusdtreceiver, bound to a specific deployed contract.
func NewMnosusdtreceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MnosusdtreceiverFilterer, error) {
	contract, err := bindMnosusdtreceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MnosusdtreceiverFilterer{contract: contract}, nil
}

// bindMnosusdtreceiver binds a generic wrapper to an already deployed contract.
func bindMnosusdtreceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MnosusdtreceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mnosusdtreceiver *MnosusdtreceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mnosusdtreceiver.Contract.MnosusdtreceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mnosusdtreceiver *MnosusdtreceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.MnosusdtreceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mnosusdtreceiver *MnosusdtreceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.MnosusdtreceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mnosusdtreceiver *MnosusdtreceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mnosusdtreceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.contract.Transact(opts, method, params...)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_Mnosusdtreceiver *MnosusdtreceiverCaller) FeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mnosusdtreceiver.contract.Call(opts, &out, "feePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_Mnosusdtreceiver *MnosusdtreceiverSession) FeePercentage() (*big.Int, error) {
	return _Mnosusdtreceiver.Contract.FeePercentage(&_Mnosusdtreceiver.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_Mnosusdtreceiver *MnosusdtreceiverCallerSession) FeePercentage() (*big.Int, error) {
	return _Mnosusdtreceiver.Contract.FeePercentage(&_Mnosusdtreceiver.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Mnosusdtreceiver *MnosusdtreceiverCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mnosusdtreceiver.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Mnosusdtreceiver *MnosusdtreceiverSession) MintPrice() (*big.Int, error) {
	return _Mnosusdtreceiver.Contract.MintPrice(&_Mnosusdtreceiver.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Mnosusdtreceiver *MnosusdtreceiverCallerSession) MintPrice() (*big.Int, error) {
	return _Mnosusdtreceiver.Contract.MintPrice(&_Mnosusdtreceiver.CallOpts)
}

// MnosToken is a free data retrieval call binding the contract method 0x8c3c2244.
//
// Solidity: function mnosToken() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverCaller) MnosToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mnosusdtreceiver.contract.Call(opts, &out, "mnosToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MnosToken is a free data retrieval call binding the contract method 0x8c3c2244.
//
// Solidity: function mnosToken() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverSession) MnosToken() (common.Address, error) {
	return _Mnosusdtreceiver.Contract.MnosToken(&_Mnosusdtreceiver.CallOpts)
}

// MnosToken is a free data retrieval call binding the contract method 0x8c3c2244.
//
// Solidity: function mnosToken() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverCallerSession) MnosToken() (common.Address, error) {
	return _Mnosusdtreceiver.Contract.MnosToken(&_Mnosusdtreceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mnosusdtreceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverSession) Owner() (common.Address, error) {
	return _Mnosusdtreceiver.Contract.Owner(&_Mnosusdtreceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverCallerSession) Owner() (common.Address, error) {
	return _Mnosusdtreceiver.Contract.Owner(&_Mnosusdtreceiver.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mnosusdtreceiver.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverSession) Treasury() (common.Address, error) {
	return _Mnosusdtreceiver.Contract.Treasury(&_Mnosusdtreceiver.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverCallerSession) Treasury() (common.Address, error) {
	return _Mnosusdtreceiver.Contract.Treasury(&_Mnosusdtreceiver.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mnosusdtreceiver.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverSession) Usdt() (common.Address, error) {
	return _Mnosusdtreceiver.Contract.Usdt(&_Mnosusdtreceiver.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Mnosusdtreceiver *MnosusdtreceiverCallerSession) Usdt() (common.Address, error) {
	return _Mnosusdtreceiver.Contract.Usdt(&_Mnosusdtreceiver.CallOpts)
}

// BurnAndSendUSDT is a paid mutator transaction binding the contract method 0x259470fa.
//
// Solidity: function burnAndSendUSDT(uint256 amount) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactor) BurnAndSendUSDT(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.contract.Transact(opts, "burnAndSendUSDT", amount)
}

// BurnAndSendUSDT is a paid mutator transaction binding the contract method 0x259470fa.
//
// Solidity: function burnAndSendUSDT(uint256 amount) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverSession) BurnAndSendUSDT(amount *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.BurnAndSendUSDT(&_Mnosusdtreceiver.TransactOpts, amount)
}

// BurnAndSendUSDT is a paid mutator transaction binding the contract method 0x259470fa.
//
// Solidity: function burnAndSendUSDT(uint256 amount) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorSession) BurnAndSendUSDT(amount *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.BurnAndSendUSDT(&_Mnosusdtreceiver.TransactOpts, amount)
}

// MintForUSDT is a paid mutator transaction binding the contract method 0xabb47966.
//
// Solidity: function mintForUSDT(uint256 amount) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactor) MintForUSDT(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.contract.Transact(opts, "mintForUSDT", amount)
}

// MintForUSDT is a paid mutator transaction binding the contract method 0xabb47966.
//
// Solidity: function mintForUSDT(uint256 amount) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverSession) MintForUSDT(amount *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.MintForUSDT(&_Mnosusdtreceiver.TransactOpts, amount)
}

// MintForUSDT is a paid mutator transaction binding the contract method 0xabb47966.
//
// Solidity: function mintForUSDT(uint256 amount) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorSession) MintForUSDT(amount *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.MintForUSDT(&_Mnosusdtreceiver.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mnosusdtreceiver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mnosusdtreceiver *MnosusdtreceiverSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.RenounceOwnership(&_Mnosusdtreceiver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.RenounceOwnership(&_Mnosusdtreceiver.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactor) SetFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.contract.Transact(opts, "setFee", _newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverSession) SetFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.SetFee(&_Mnosusdtreceiver.TransactOpts, _newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorSession) SetFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.SetFee(&_Mnosusdtreceiver.TransactOpts, _newFee)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _newPrice) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactor) SetMintPrice(opts *bind.TransactOpts, _newPrice *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.contract.Transact(opts, "setMintPrice", _newPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _newPrice) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverSession) SetMintPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.SetMintPrice(&_Mnosusdtreceiver.TransactOpts, _newPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _newPrice) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorSession) SetMintPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.SetMintPrice(&_Mnosusdtreceiver.TransactOpts, _newPrice)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _mnosToken) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactor) SetToken(opts *bind.TransactOpts, _mnosToken common.Address) (*types.Transaction, error) {
	return _Mnosusdtreceiver.contract.Transact(opts, "setToken", _mnosToken)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _mnosToken) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverSession) SetToken(_mnosToken common.Address) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.SetToken(&_Mnosusdtreceiver.TransactOpts, _mnosToken)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _mnosToken) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorSession) SetToken(_mnosToken common.Address) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.SetToken(&_Mnosusdtreceiver.TransactOpts, _mnosToken)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactor) SetTreasury(opts *bind.TransactOpts, _newTreasury common.Address) (*types.Transaction, error) {
	return _Mnosusdtreceiver.contract.Transact(opts, "setTreasury", _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverSession) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.SetTreasury(&_Mnosusdtreceiver.TransactOpts, _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorSession) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.SetTreasury(&_Mnosusdtreceiver.TransactOpts, _newTreasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mnosusdtreceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.TransferOwnership(&_Mnosusdtreceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mnosusdtreceiver *MnosusdtreceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mnosusdtreceiver.Contract.TransferOwnership(&_Mnosusdtreceiver.TransactOpts, newOwner)
}

// MnosusdtreceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mnosusdtreceiver contract.
type MnosusdtreceiverOwnershipTransferredIterator struct {
	Event *MnosusdtreceiverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MnosusdtreceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MnosusdtreceiverOwnershipTransferred)
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
		it.Event = new(MnosusdtreceiverOwnershipTransferred)
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
func (it *MnosusdtreceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MnosusdtreceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MnosusdtreceiverOwnershipTransferred represents a OwnershipTransferred event raised by the Mnosusdtreceiver contract.
type MnosusdtreceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mnosusdtreceiver *MnosusdtreceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MnosusdtreceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mnosusdtreceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MnosusdtreceiverOwnershipTransferredIterator{contract: _Mnosusdtreceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mnosusdtreceiver *MnosusdtreceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MnosusdtreceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mnosusdtreceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MnosusdtreceiverOwnershipTransferred)
				if err := _Mnosusdtreceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mnosusdtreceiver *MnosusdtreceiverFilterer) ParseOwnershipTransferred(log types.Log) (*MnosusdtreceiverOwnershipTransferred, error) {
	event := new(MnosusdtreceiverOwnershipTransferred)
	if err := _Mnosusdtreceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
