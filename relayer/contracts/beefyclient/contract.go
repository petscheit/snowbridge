// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package beefyclient

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

// BeefyClientCommitment is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientCommitment struct {
	BlockNumber    uint32
	ValidatorSetID uint64
	Payload        BeefyClientPayload
}

// BeefyClientMMRLeaf is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientMMRLeaf struct {
	Version              uint8
	ParentNumber         uint32
	ParentHash           [32]byte
	NextAuthoritySetID   uint64
	NextAuthoritySetLen  uint32
	NextAuthoritySetRoot [32]byte
	ParachainHeadsRoot   [32]byte
}

// BeefyClientPayload is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientPayload struct {
	MmrRootHash [32]byte
	Prefix      []byte
	Suffix      []byte
}

// BeefyClientValidatorProof is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientValidatorProof struct {
	V       uint8
	R       [32]byte
	S       [32]byte
	Index   *big.Int
	Account common.Address
	Proof   [][32]byte
}

// BeefyClientValidatorSet is an auto generated low-level Go binding around an user-defined struct.
type BeefyClientValidatorSet struct {
	Id     *big.Int
	Length *big.Int
	Root   [32]byte
}

// BeefyClientMetaData contains all meta data concerning the BeefyClient contract.
var BeefyClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_randaoCommitDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_randaoCommitExpiration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidBitfield\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMMRLeaf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMMRLeafProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTask\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidatorProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughClaims\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrevRandaoAlreadyCaptured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrevRandaoNotCaptured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TaskExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WaitPeriodNotOver\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mmrRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"name\":\"NewMMRRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentHash\",\"type\":\"bytes32\"}],\"name\":\"commitPrevRandao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"bitfield\",\"type\":\"uint256[]\"}],\"name\":\"createFinalBitfield\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"bitsToSet\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"createInitialBitfield\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"length\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_initialBeefyBlock\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"length\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structBeefyClient.ValidatorSet\",\"name\":\"_initialValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"length\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structBeefyClient.ValidatorSet\",\"name\":\"_nextValidatorSet\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBeefyBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestMMRRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidatorSet\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"length\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randaoCommitDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randaoCommitExpiration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"validatorSetID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"mmrRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structBeefyClient.Payload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structBeefyClient.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"bitfield\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structBeefyClient.ValidatorProof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"}],\"name\":\"submitFinal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"validatorSetID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"mmrRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structBeefyClient.Payload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structBeefyClient.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"bitfield\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structBeefyClient.ValidatorProof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"parentNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nextAuthoritySetID\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nextAuthoritySetLen\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAuthoritySetRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"parachainHeadsRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeefyClient.MMRLeaf\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"leafProofOrder\",\"type\":\"uint256\"}],\"name\":\"submitFinalWithHandover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"bitfield\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structBeefyClient.ValidatorProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"submitInitial\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"bitfield\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structBeefyClient.ValidatorProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"submitInitialWithHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"validatorSetLen\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"prevRandao\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"bitfieldHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofOrder\",\"type\":\"uint256\"}],\"name\":\"verifyMMRLeafProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BeefyClientABI is the input ABI used to generate the binding from.
// Deprecated: Use BeefyClientMetaData.ABI instead.
var BeefyClientABI = BeefyClientMetaData.ABI

// BeefyClient is an auto generated Go binding around an Ethereum contract.
type BeefyClient struct {
	BeefyClientCaller     // Read-only binding to the contract
	BeefyClientTransactor // Write-only binding to the contract
	BeefyClientFilterer   // Log filterer for contract events
}

// BeefyClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeefyClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeefyClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeefyClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeefyClientSession struct {
	Contract     *BeefyClient      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeefyClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeefyClientCallerSession struct {
	Contract *BeefyClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BeefyClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeefyClientTransactorSession struct {
	Contract     *BeefyClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BeefyClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeefyClientRaw struct {
	Contract *BeefyClient // Generic contract binding to access the raw methods on
}

// BeefyClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeefyClientCallerRaw struct {
	Contract *BeefyClientCaller // Generic read-only contract binding to access the raw methods on
}

// BeefyClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeefyClientTransactorRaw struct {
	Contract *BeefyClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeefyClient creates a new instance of BeefyClient, bound to a specific deployed contract.
func NewBeefyClient(address common.Address, backend bind.ContractBackend) (*BeefyClient, error) {
	contract, err := bindBeefyClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeefyClient{BeefyClientCaller: BeefyClientCaller{contract: contract}, BeefyClientTransactor: BeefyClientTransactor{contract: contract}, BeefyClientFilterer: BeefyClientFilterer{contract: contract}}, nil
}

// NewBeefyClientCaller creates a new read-only instance of BeefyClient, bound to a specific deployed contract.
func NewBeefyClientCaller(address common.Address, caller bind.ContractCaller) (*BeefyClientCaller, error) {
	contract, err := bindBeefyClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeefyClientCaller{contract: contract}, nil
}

// NewBeefyClientTransactor creates a new write-only instance of BeefyClient, bound to a specific deployed contract.
func NewBeefyClientTransactor(address common.Address, transactor bind.ContractTransactor) (*BeefyClientTransactor, error) {
	contract, err := bindBeefyClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeefyClientTransactor{contract: contract}, nil
}

// NewBeefyClientFilterer creates a new log filterer instance of BeefyClient, bound to a specific deployed contract.
func NewBeefyClientFilterer(address common.Address, filterer bind.ContractFilterer) (*BeefyClientFilterer, error) {
	contract, err := bindBeefyClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeefyClientFilterer{contract: contract}, nil
}

// bindBeefyClient binds a generic wrapper to an already deployed contract.
func bindBeefyClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BeefyClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeefyClient *BeefyClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeefyClient.Contract.BeefyClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeefyClient *BeefyClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyClient.Contract.BeefyClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeefyClient *BeefyClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeefyClient.Contract.BeefyClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeefyClient *BeefyClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeefyClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeefyClient *BeefyClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeefyClient *BeefyClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeefyClient.Contract.contract.Transact(opts, method, params...)
}

// CreateFinalBitfield is a free data retrieval call binding the contract method 0x8ab81d13.
//
// Solidity: function createFinalBitfield(bytes32 commitmentHash, uint256[] bitfield) view returns(uint256[])
func (_BeefyClient *BeefyClientCaller) CreateFinalBitfield(opts *bind.CallOpts, commitmentHash [32]byte, bitfield []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "createFinalBitfield", commitmentHash, bitfield)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CreateFinalBitfield is a free data retrieval call binding the contract method 0x8ab81d13.
//
// Solidity: function createFinalBitfield(bytes32 commitmentHash, uint256[] bitfield) view returns(uint256[])
func (_BeefyClient *BeefyClientSession) CreateFinalBitfield(commitmentHash [32]byte, bitfield []*big.Int) ([]*big.Int, error) {
	return _BeefyClient.Contract.CreateFinalBitfield(&_BeefyClient.CallOpts, commitmentHash, bitfield)
}

// CreateFinalBitfield is a free data retrieval call binding the contract method 0x8ab81d13.
//
// Solidity: function createFinalBitfield(bytes32 commitmentHash, uint256[] bitfield) view returns(uint256[])
func (_BeefyClient *BeefyClientCallerSession) CreateFinalBitfield(commitmentHash [32]byte, bitfield []*big.Int) ([]*big.Int, error) {
	return _BeefyClient.Contract.CreateFinalBitfield(&_BeefyClient.CallOpts, commitmentHash, bitfield)
}

// CreateInitialBitfield is a free data retrieval call binding the contract method 0x5da57fe9.
//
// Solidity: function createInitialBitfield(uint256[] bitsToSet, uint256 length) pure returns(uint256[])
func (_BeefyClient *BeefyClientCaller) CreateInitialBitfield(opts *bind.CallOpts, bitsToSet []*big.Int, length *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "createInitialBitfield", bitsToSet, length)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CreateInitialBitfield is a free data retrieval call binding the contract method 0x5da57fe9.
//
// Solidity: function createInitialBitfield(uint256[] bitsToSet, uint256 length) pure returns(uint256[])
func (_BeefyClient *BeefyClientSession) CreateInitialBitfield(bitsToSet []*big.Int, length *big.Int) ([]*big.Int, error) {
	return _BeefyClient.Contract.CreateInitialBitfield(&_BeefyClient.CallOpts, bitsToSet, length)
}

// CreateInitialBitfield is a free data retrieval call binding the contract method 0x5da57fe9.
//
// Solidity: function createInitialBitfield(uint256[] bitsToSet, uint256 length) pure returns(uint256[])
func (_BeefyClient *BeefyClientCallerSession) CreateInitialBitfield(bitsToSet []*big.Int, length *big.Int) ([]*big.Int, error) {
	return _BeefyClient.Contract.CreateInitialBitfield(&_BeefyClient.CallOpts, bitsToSet, length)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x2cdea717.
//
// Solidity: function currentValidatorSet() view returns(uint128 id, uint128 length, bytes32 root)
func (_BeefyClient *BeefyClientCaller) CurrentValidatorSet(opts *bind.CallOpts) (struct {
	Id     *big.Int
	Length *big.Int
	Root   [32]byte
}, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "currentValidatorSet")

	outstruct := new(struct {
		Id     *big.Int
		Length *big.Int
		Root   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Root = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x2cdea717.
//
// Solidity: function currentValidatorSet() view returns(uint128 id, uint128 length, bytes32 root)
func (_BeefyClient *BeefyClientSession) CurrentValidatorSet() (struct {
	Id     *big.Int
	Length *big.Int
	Root   [32]byte
}, error) {
	return _BeefyClient.Contract.CurrentValidatorSet(&_BeefyClient.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x2cdea717.
//
// Solidity: function currentValidatorSet() view returns(uint128 id, uint128 length, bytes32 root)
func (_BeefyClient *BeefyClientCallerSession) CurrentValidatorSet() (struct {
	Id     *big.Int
	Length *big.Int
	Root   [32]byte
}, error) {
	return _BeefyClient.Contract.CurrentValidatorSet(&_BeefyClient.CallOpts)
}

// LatestBeefyBlock is a free data retrieval call binding the contract method 0x66ae69a0.
//
// Solidity: function latestBeefyBlock() view returns(uint64)
func (_BeefyClient *BeefyClientCaller) LatestBeefyBlock(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "latestBeefyBlock")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestBeefyBlock is a free data retrieval call binding the contract method 0x66ae69a0.
//
// Solidity: function latestBeefyBlock() view returns(uint64)
func (_BeefyClient *BeefyClientSession) LatestBeefyBlock() (uint64, error) {
	return _BeefyClient.Contract.LatestBeefyBlock(&_BeefyClient.CallOpts)
}

// LatestBeefyBlock is a free data retrieval call binding the contract method 0x66ae69a0.
//
// Solidity: function latestBeefyBlock() view returns(uint64)
func (_BeefyClient *BeefyClientCallerSession) LatestBeefyBlock() (uint64, error) {
	return _BeefyClient.Contract.LatestBeefyBlock(&_BeefyClient.CallOpts)
}

// LatestMMRRoot is a free data retrieval call binding the contract method 0x41c9634e.
//
// Solidity: function latestMMRRoot() view returns(bytes32)
func (_BeefyClient *BeefyClientCaller) LatestMMRRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "latestMMRRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestMMRRoot is a free data retrieval call binding the contract method 0x41c9634e.
//
// Solidity: function latestMMRRoot() view returns(bytes32)
func (_BeefyClient *BeefyClientSession) LatestMMRRoot() ([32]byte, error) {
	return _BeefyClient.Contract.LatestMMRRoot(&_BeefyClient.CallOpts)
}

// LatestMMRRoot is a free data retrieval call binding the contract method 0x41c9634e.
//
// Solidity: function latestMMRRoot() view returns(bytes32)
func (_BeefyClient *BeefyClientCallerSession) LatestMMRRoot() ([32]byte, error) {
	return _BeefyClient.Contract.LatestMMRRoot(&_BeefyClient.CallOpts)
}

// NextValidatorSet is a free data retrieval call binding the contract method 0x36667513.
//
// Solidity: function nextValidatorSet() view returns(uint128 id, uint128 length, bytes32 root)
func (_BeefyClient *BeefyClientCaller) NextValidatorSet(opts *bind.CallOpts) (struct {
	Id     *big.Int
	Length *big.Int
	Root   [32]byte
}, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "nextValidatorSet")

	outstruct := new(struct {
		Id     *big.Int
		Length *big.Int
		Root   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Root = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// NextValidatorSet is a free data retrieval call binding the contract method 0x36667513.
//
// Solidity: function nextValidatorSet() view returns(uint128 id, uint128 length, bytes32 root)
func (_BeefyClient *BeefyClientSession) NextValidatorSet() (struct {
	Id     *big.Int
	Length *big.Int
	Root   [32]byte
}, error) {
	return _BeefyClient.Contract.NextValidatorSet(&_BeefyClient.CallOpts)
}

// NextValidatorSet is a free data retrieval call binding the contract method 0x36667513.
//
// Solidity: function nextValidatorSet() view returns(uint128 id, uint128 length, bytes32 root)
func (_BeefyClient *BeefyClientCallerSession) NextValidatorSet() (struct {
	Id     *big.Int
	Length *big.Int
	Root   [32]byte
}, error) {
	return _BeefyClient.Contract.NextValidatorSet(&_BeefyClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeefyClient *BeefyClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeefyClient *BeefyClientSession) Owner() (common.Address, error) {
	return _BeefyClient.Contract.Owner(&_BeefyClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeefyClient *BeefyClientCallerSession) Owner() (common.Address, error) {
	return _BeefyClient.Contract.Owner(&_BeefyClient.CallOpts)
}

// RandaoCommitDelay is a free data retrieval call binding the contract method 0x591d99ee.
//
// Solidity: function randaoCommitDelay() view returns(uint256)
func (_BeefyClient *BeefyClientCaller) RandaoCommitDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "randaoCommitDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandaoCommitDelay is a free data retrieval call binding the contract method 0x591d99ee.
//
// Solidity: function randaoCommitDelay() view returns(uint256)
func (_BeefyClient *BeefyClientSession) RandaoCommitDelay() (*big.Int, error) {
	return _BeefyClient.Contract.RandaoCommitDelay(&_BeefyClient.CallOpts)
}

// RandaoCommitDelay is a free data retrieval call binding the contract method 0x591d99ee.
//
// Solidity: function randaoCommitDelay() view returns(uint256)
func (_BeefyClient *BeefyClientCallerSession) RandaoCommitDelay() (*big.Int, error) {
	return _BeefyClient.Contract.RandaoCommitDelay(&_BeefyClient.CallOpts)
}

// RandaoCommitExpiration is a free data retrieval call binding the contract method 0xad209a9b.
//
// Solidity: function randaoCommitExpiration() view returns(uint256)
func (_BeefyClient *BeefyClientCaller) RandaoCommitExpiration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "randaoCommitExpiration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandaoCommitExpiration is a free data retrieval call binding the contract method 0xad209a9b.
//
// Solidity: function randaoCommitExpiration() view returns(uint256)
func (_BeefyClient *BeefyClientSession) RandaoCommitExpiration() (*big.Int, error) {
	return _BeefyClient.Contract.RandaoCommitExpiration(&_BeefyClient.CallOpts)
}

// RandaoCommitExpiration is a free data retrieval call binding the contract method 0xad209a9b.
//
// Solidity: function randaoCommitExpiration() view returns(uint256)
func (_BeefyClient *BeefyClientCallerSession) RandaoCommitExpiration() (*big.Int, error) {
	return _BeefyClient.Contract.RandaoCommitExpiration(&_BeefyClient.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0xe579f500.
//
// Solidity: function tasks(bytes32 ) view returns(address account, uint64 blockNumber, uint32 validatorSetLen, uint256 prevRandao, bytes32 bitfieldHash)
func (_BeefyClient *BeefyClientCaller) Tasks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Account         common.Address
	BlockNumber     uint64
	ValidatorSetLen uint32
	PrevRandao      *big.Int
	BitfieldHash    [32]byte
}, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Account         common.Address
		BlockNumber     uint64
		ValidatorSetLen uint32
		PrevRandao      *big.Int
		BitfieldHash    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ValidatorSetLen = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.PrevRandao = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BitfieldHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0xe579f500.
//
// Solidity: function tasks(bytes32 ) view returns(address account, uint64 blockNumber, uint32 validatorSetLen, uint256 prevRandao, bytes32 bitfieldHash)
func (_BeefyClient *BeefyClientSession) Tasks(arg0 [32]byte) (struct {
	Account         common.Address
	BlockNumber     uint64
	ValidatorSetLen uint32
	PrevRandao      *big.Int
	BitfieldHash    [32]byte
}, error) {
	return _BeefyClient.Contract.Tasks(&_BeefyClient.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0xe579f500.
//
// Solidity: function tasks(bytes32 ) view returns(address account, uint64 blockNumber, uint32 validatorSetLen, uint256 prevRandao, bytes32 bitfieldHash)
func (_BeefyClient *BeefyClientCallerSession) Tasks(arg0 [32]byte) (struct {
	Account         common.Address
	BlockNumber     uint64
	ValidatorSetLen uint32
	PrevRandao      *big.Int
	BitfieldHash    [32]byte
}, error) {
	return _BeefyClient.Contract.Tasks(&_BeefyClient.CallOpts, arg0)
}

// VerifyMMRLeafProof is a free data retrieval call binding the contract method 0xa401662b.
//
// Solidity: function verifyMMRLeafProof(bytes32 leafHash, bytes32[] proof, uint256 proofOrder) view returns(bool)
func (_BeefyClient *BeefyClientCaller) VerifyMMRLeafProof(opts *bind.CallOpts, leafHash [32]byte, proof [][32]byte, proofOrder *big.Int) (bool, error) {
	var out []interface{}
	err := _BeefyClient.contract.Call(opts, &out, "verifyMMRLeafProof", leafHash, proof, proofOrder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMMRLeafProof is a free data retrieval call binding the contract method 0xa401662b.
//
// Solidity: function verifyMMRLeafProof(bytes32 leafHash, bytes32[] proof, uint256 proofOrder) view returns(bool)
func (_BeefyClient *BeefyClientSession) VerifyMMRLeafProof(leafHash [32]byte, proof [][32]byte, proofOrder *big.Int) (bool, error) {
	return _BeefyClient.Contract.VerifyMMRLeafProof(&_BeefyClient.CallOpts, leafHash, proof, proofOrder)
}

// VerifyMMRLeafProof is a free data retrieval call binding the contract method 0xa401662b.
//
// Solidity: function verifyMMRLeafProof(bytes32 leafHash, bytes32[] proof, uint256 proofOrder) view returns(bool)
func (_BeefyClient *BeefyClientCallerSession) VerifyMMRLeafProof(leafHash [32]byte, proof [][32]byte, proofOrder *big.Int) (bool, error) {
	return _BeefyClient.Contract.VerifyMMRLeafProof(&_BeefyClient.CallOpts, leafHash, proof, proofOrder)
}

// CommitPrevRandao is a paid mutator transaction binding the contract method 0xa77cf3d2.
//
// Solidity: function commitPrevRandao(bytes32 commitmentHash) returns()
func (_BeefyClient *BeefyClientTransactor) CommitPrevRandao(opts *bind.TransactOpts, commitmentHash [32]byte) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "commitPrevRandao", commitmentHash)
}

// CommitPrevRandao is a paid mutator transaction binding the contract method 0xa77cf3d2.
//
// Solidity: function commitPrevRandao(bytes32 commitmentHash) returns()
func (_BeefyClient *BeefyClientSession) CommitPrevRandao(commitmentHash [32]byte) (*types.Transaction, error) {
	return _BeefyClient.Contract.CommitPrevRandao(&_BeefyClient.TransactOpts, commitmentHash)
}

// CommitPrevRandao is a paid mutator transaction binding the contract method 0xa77cf3d2.
//
// Solidity: function commitPrevRandao(bytes32 commitmentHash) returns()
func (_BeefyClient *BeefyClientTransactorSession) CommitPrevRandao(commitmentHash [32]byte) (*types.Transaction, error) {
	return _BeefyClient.Contract.CommitPrevRandao(&_BeefyClient.TransactOpts, commitmentHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xe104815d.
//
// Solidity: function initialize(uint64 _initialBeefyBlock, (uint128,uint128,bytes32) _initialValidatorSet, (uint128,uint128,bytes32) _nextValidatorSet) returns()
func (_BeefyClient *BeefyClientTransactor) Initialize(opts *bind.TransactOpts, _initialBeefyBlock uint64, _initialValidatorSet BeefyClientValidatorSet, _nextValidatorSet BeefyClientValidatorSet) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "initialize", _initialBeefyBlock, _initialValidatorSet, _nextValidatorSet)
}

// Initialize is a paid mutator transaction binding the contract method 0xe104815d.
//
// Solidity: function initialize(uint64 _initialBeefyBlock, (uint128,uint128,bytes32) _initialValidatorSet, (uint128,uint128,bytes32) _nextValidatorSet) returns()
func (_BeefyClient *BeefyClientSession) Initialize(_initialBeefyBlock uint64, _initialValidatorSet BeefyClientValidatorSet, _nextValidatorSet BeefyClientValidatorSet) (*types.Transaction, error) {
	return _BeefyClient.Contract.Initialize(&_BeefyClient.TransactOpts, _initialBeefyBlock, _initialValidatorSet, _nextValidatorSet)
}

// Initialize is a paid mutator transaction binding the contract method 0xe104815d.
//
// Solidity: function initialize(uint64 _initialBeefyBlock, (uint128,uint128,bytes32) _initialValidatorSet, (uint128,uint128,bytes32) _nextValidatorSet) returns()
func (_BeefyClient *BeefyClientTransactorSession) Initialize(_initialBeefyBlock uint64, _initialValidatorSet BeefyClientValidatorSet, _nextValidatorSet BeefyClientValidatorSet) (*types.Transaction, error) {
	return _BeefyClient.Contract.Initialize(&_BeefyClient.TransactOpts, _initialBeefyBlock, _initialValidatorSet, _nextValidatorSet)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeefyClient *BeefyClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeefyClient *BeefyClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _BeefyClient.Contract.RenounceOwnership(&_BeefyClient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeefyClient *BeefyClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BeefyClient.Contract.RenounceOwnership(&_BeefyClient.TransactOpts)
}

// SubmitFinal is a paid mutator transaction binding the contract method 0xc46af85d.
//
// Solidity: function submitFinal((uint32,uint64,(bytes32,bytes,bytes)) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[])[] proofs) returns()
func (_BeefyClient *BeefyClientTransactor) SubmitFinal(opts *bind.TransactOpts, commitment BeefyClientCommitment, bitfield []*big.Int, proofs []BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "submitFinal", commitment, bitfield, proofs)
}

// SubmitFinal is a paid mutator transaction binding the contract method 0xc46af85d.
//
// Solidity: function submitFinal((uint32,uint64,(bytes32,bytes,bytes)) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[])[] proofs) returns()
func (_BeefyClient *BeefyClientSession) SubmitFinal(commitment BeefyClientCommitment, bitfield []*big.Int, proofs []BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitFinal(&_BeefyClient.TransactOpts, commitment, bitfield, proofs)
}

// SubmitFinal is a paid mutator transaction binding the contract method 0xc46af85d.
//
// Solidity: function submitFinal((uint32,uint64,(bytes32,bytes,bytes)) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[])[] proofs) returns()
func (_BeefyClient *BeefyClientTransactorSession) SubmitFinal(commitment BeefyClientCommitment, bitfield []*big.Int, proofs []BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitFinal(&_BeefyClient.TransactOpts, commitment, bitfield, proofs)
}

// SubmitFinalWithHandover is a paid mutator transaction binding the contract method 0x8114d5e8.
//
// Solidity: function submitFinalWithHandover((uint32,uint64,(bytes32,bytes,bytes)) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[])[] proofs, (uint8,uint32,bytes32,uint64,uint32,bytes32,bytes32) leaf, bytes32[] leafProof, uint256 leafProofOrder) returns()
func (_BeefyClient *BeefyClientTransactor) SubmitFinalWithHandover(opts *bind.TransactOpts, commitment BeefyClientCommitment, bitfield []*big.Int, proofs []BeefyClientValidatorProof, leaf BeefyClientMMRLeaf, leafProof [][32]byte, leafProofOrder *big.Int) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "submitFinalWithHandover", commitment, bitfield, proofs, leaf, leafProof, leafProofOrder)
}

// SubmitFinalWithHandover is a paid mutator transaction binding the contract method 0x8114d5e8.
//
// Solidity: function submitFinalWithHandover((uint32,uint64,(bytes32,bytes,bytes)) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[])[] proofs, (uint8,uint32,bytes32,uint64,uint32,bytes32,bytes32) leaf, bytes32[] leafProof, uint256 leafProofOrder) returns()
func (_BeefyClient *BeefyClientSession) SubmitFinalWithHandover(commitment BeefyClientCommitment, bitfield []*big.Int, proofs []BeefyClientValidatorProof, leaf BeefyClientMMRLeaf, leafProof [][32]byte, leafProofOrder *big.Int) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitFinalWithHandover(&_BeefyClient.TransactOpts, commitment, bitfield, proofs, leaf, leafProof, leafProofOrder)
}

// SubmitFinalWithHandover is a paid mutator transaction binding the contract method 0x8114d5e8.
//
// Solidity: function submitFinalWithHandover((uint32,uint64,(bytes32,bytes,bytes)) commitment, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[])[] proofs, (uint8,uint32,bytes32,uint64,uint32,bytes32,bytes32) leaf, bytes32[] leafProof, uint256 leafProofOrder) returns()
func (_BeefyClient *BeefyClientTransactorSession) SubmitFinalWithHandover(commitment BeefyClientCommitment, bitfield []*big.Int, proofs []BeefyClientValidatorProof, leaf BeefyClientMMRLeaf, leafProof [][32]byte, leafProofOrder *big.Int) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitFinalWithHandover(&_BeefyClient.TransactOpts, commitment, bitfield, proofs, leaf, leafProof, leafProofOrder)
}

// SubmitInitial is a paid mutator transaction binding the contract method 0x06a9eff7.
//
// Solidity: function submitInitial(bytes32 commitmentHash, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[]) proof) payable returns()
func (_BeefyClient *BeefyClientTransactor) SubmitInitial(opts *bind.TransactOpts, commitmentHash [32]byte, bitfield []*big.Int, proof BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "submitInitial", commitmentHash, bitfield, proof)
}

// SubmitInitial is a paid mutator transaction binding the contract method 0x06a9eff7.
//
// Solidity: function submitInitial(bytes32 commitmentHash, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[]) proof) payable returns()
func (_BeefyClient *BeefyClientSession) SubmitInitial(commitmentHash [32]byte, bitfield []*big.Int, proof BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitInitial(&_BeefyClient.TransactOpts, commitmentHash, bitfield, proof)
}

// SubmitInitial is a paid mutator transaction binding the contract method 0x06a9eff7.
//
// Solidity: function submitInitial(bytes32 commitmentHash, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[]) proof) payable returns()
func (_BeefyClient *BeefyClientTransactorSession) SubmitInitial(commitmentHash [32]byte, bitfield []*big.Int, proof BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitInitial(&_BeefyClient.TransactOpts, commitmentHash, bitfield, proof)
}

// SubmitInitialWithHandover is a paid mutator transaction binding the contract method 0x61de6237.
//
// Solidity: function submitInitialWithHandover(bytes32 commitmentHash, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[]) proof) payable returns()
func (_BeefyClient *BeefyClientTransactor) SubmitInitialWithHandover(opts *bind.TransactOpts, commitmentHash [32]byte, bitfield []*big.Int, proof BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "submitInitialWithHandover", commitmentHash, bitfield, proof)
}

// SubmitInitialWithHandover is a paid mutator transaction binding the contract method 0x61de6237.
//
// Solidity: function submitInitialWithHandover(bytes32 commitmentHash, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[]) proof) payable returns()
func (_BeefyClient *BeefyClientSession) SubmitInitialWithHandover(commitmentHash [32]byte, bitfield []*big.Int, proof BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitInitialWithHandover(&_BeefyClient.TransactOpts, commitmentHash, bitfield, proof)
}

// SubmitInitialWithHandover is a paid mutator transaction binding the contract method 0x61de6237.
//
// Solidity: function submitInitialWithHandover(bytes32 commitmentHash, uint256[] bitfield, (uint8,bytes32,bytes32,uint256,address,bytes32[]) proof) payable returns()
func (_BeefyClient *BeefyClientTransactorSession) SubmitInitialWithHandover(commitmentHash [32]byte, bitfield []*big.Int, proof BeefyClientValidatorProof) (*types.Transaction, error) {
	return _BeefyClient.Contract.SubmitInitialWithHandover(&_BeefyClient.TransactOpts, commitmentHash, bitfield, proof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeefyClient *BeefyClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BeefyClient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeefyClient *BeefyClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BeefyClient.Contract.TransferOwnership(&_BeefyClient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeefyClient *BeefyClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BeefyClient.Contract.TransferOwnership(&_BeefyClient.TransactOpts, newOwner)
}

// BeefyClientNewMMRRootIterator is returned from FilterNewMMRRoot and is used to iterate over the raw logs and unpacked data for NewMMRRoot events raised by the BeefyClient contract.
type BeefyClientNewMMRRootIterator struct {
	Event *BeefyClientNewMMRRoot // Event containing the contract specifics and raw log

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
func (it *BeefyClientNewMMRRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeefyClientNewMMRRoot)
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
		it.Event = new(BeefyClientNewMMRRoot)
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
func (it *BeefyClientNewMMRRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeefyClientNewMMRRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeefyClientNewMMRRoot represents a NewMMRRoot event raised by the BeefyClient contract.
type BeefyClientNewMMRRoot struct {
	MmrRoot     [32]byte
	BlockNumber uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewMMRRoot is a free log retrieval operation binding the contract event 0xd95fe1258d152dc91c81b09380498adc76ed36a6079bcb2ed31eff622ae2d0f1.
//
// Solidity: event NewMMRRoot(bytes32 mmrRoot, uint64 blockNumber)
func (_BeefyClient *BeefyClientFilterer) FilterNewMMRRoot(opts *bind.FilterOpts) (*BeefyClientNewMMRRootIterator, error) {

	logs, sub, err := _BeefyClient.contract.FilterLogs(opts, "NewMMRRoot")
	if err != nil {
		return nil, err
	}
	return &BeefyClientNewMMRRootIterator{contract: _BeefyClient.contract, event: "NewMMRRoot", logs: logs, sub: sub}, nil
}

// WatchNewMMRRoot is a free log subscription operation binding the contract event 0xd95fe1258d152dc91c81b09380498adc76ed36a6079bcb2ed31eff622ae2d0f1.
//
// Solidity: event NewMMRRoot(bytes32 mmrRoot, uint64 blockNumber)
func (_BeefyClient *BeefyClientFilterer) WatchNewMMRRoot(opts *bind.WatchOpts, sink chan<- *BeefyClientNewMMRRoot) (event.Subscription, error) {

	logs, sub, err := _BeefyClient.contract.WatchLogs(opts, "NewMMRRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeefyClientNewMMRRoot)
				if err := _BeefyClient.contract.UnpackLog(event, "NewMMRRoot", log); err != nil {
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

// ParseNewMMRRoot is a log parse operation binding the contract event 0xd95fe1258d152dc91c81b09380498adc76ed36a6079bcb2ed31eff622ae2d0f1.
//
// Solidity: event NewMMRRoot(bytes32 mmrRoot, uint64 blockNumber)
func (_BeefyClient *BeefyClientFilterer) ParseNewMMRRoot(log types.Log) (*BeefyClientNewMMRRoot, error) {
	event := new(BeefyClientNewMMRRoot)
	if err := _BeefyClient.contract.UnpackLog(event, "NewMMRRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeefyClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BeefyClient contract.
type BeefyClientOwnershipTransferredIterator struct {
	Event *BeefyClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BeefyClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeefyClientOwnershipTransferred)
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
		it.Event = new(BeefyClientOwnershipTransferred)
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
func (it *BeefyClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeefyClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeefyClientOwnershipTransferred represents a OwnershipTransferred event raised by the BeefyClient contract.
type BeefyClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BeefyClient *BeefyClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BeefyClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeefyClient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BeefyClientOwnershipTransferredIterator{contract: _BeefyClient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BeefyClient *BeefyClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BeefyClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeefyClient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeefyClientOwnershipTransferred)
				if err := _BeefyClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BeefyClient *BeefyClientFilterer) ParseOwnershipTransferred(log types.Log) (*BeefyClientOwnershipTransferred, error) {
	event := new(BeefyClientOwnershipTransferred)
	if err := _BeefyClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
