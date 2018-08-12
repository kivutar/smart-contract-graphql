package resolvers

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// LogResolver is a resolver for the Log GraphQL type.
// It defines methods to resolve every field of Log.
type LogResolver struct {
	types.Log
	ABI  *abi.ABI
	Name string
}

// Address returns the hexadecimal representation of the contract that generated the event log
func (lr LogResolver) Address() string {
	return lr.Log.Address.Hex()
}

// Topics returns a list of topics (indexed fields) for this event log
func (lr LogResolver) Topics() *[]string {
	var out []string
	for _, t := range lr.Log.Topics {
		out = append(out, t.Hex())
	}
	return &out
}

// Data is the raw data of this event log. For a decoded version, see Values.
func (lr LogResolver) Data() string {
	return common.Bytes2Hex(lr.Log.Data)
}

// Values represent the decoded version of Data
func (lr LogResolver) Values() (*[]string, error) {
	values, err := lr.ABI.Events[lr.Name].Inputs.UnpackValues(lr.Log.Data)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, v := range values {
		out = append(out, fmt.Sprint(v))
	}
	return &out, nil
}

// TxHash is the hash of the transaction that generated this event log
func (lr LogResolver) TxHash() string {
	return lr.Log.TxHash.Hex()
}

// BlockNumber is the height of the block in which the transaction was included
func (lr LogResolver) BlockNumber() int32 {
	return int32(lr.Log.BlockNumber)
}

// BlockHash is the hexadecimal representation of the hash of the block in which the transaction was included
func (lr LogResolver) BlockHash() string {
	return lr.Log.BlockHash.Hex()
}

// Index is the index of the event log in the transaction receipt
func (lr LogResolver) Index() int32 {
	return int32(lr.Log.Index)
}

// TxIndex is the index of the transaction in the block
func (lr LogResolver) TxIndex() int32 {
	return int32(lr.Log.TxIndex)
}

// Removed is true if this log was reverted due to a chain reorganisation.
// You must pay attention to this if you receive logs through a filter query.
func (lr LogResolver) Removed() bool {
	return lr.Log.Removed
}
