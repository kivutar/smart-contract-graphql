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

func (lr LogResolver) Address() string {
	return lr.Log.Address.Hex()
}

func (lr LogResolver) Topics() *[]string {
	var out []string
	for _, t := range lr.Log.Topics {
		out = append(out, t.Hex())
	}
	return &out
}

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

func (lr LogResolver) TxHash() string {
	return lr.Log.TxHash.Hex()
}

func (lr LogResolver) BlockNumber() int32 {
	return int32(lr.Log.BlockNumber)
}

func (lr LogResolver) BlockHash() string {
	return lr.Log.BlockHash.Hex()
}

func (lr LogResolver) Index() int32 {
	return int32(lr.Log.Index)
}

func (lr LogResolver) TxIndex() int32 {
	return int32(lr.Log.TxIndex)
}

func (lr LogResolver) Removed() bool {
	return lr.Log.Removed
}
