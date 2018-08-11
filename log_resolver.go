package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type logResolver struct {
	types.Log
	ABI  *abi.ABI
	Name string
}

func (lr logResolver) Address() string {
	return lr.Log.Address.Hex()
}

func (lr logResolver) Topics() *[]string {
	var out []string
	for _, t := range lr.Log.Topics {
		out = append(out, t.Hex())
	}
	return &out
}

func (lr logResolver) Data() string {
	return common.Bytes2Hex(lr.Log.Data)
}

func (lr logResolver) Values() (*[]string, error) {
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

func (lr logResolver) TxHash() string {
	return lr.Log.TxHash.Hex()
}

func (lr logResolver) BlockNumber() int32 {
	return int32(lr.Log.BlockNumber)
}

func (lr logResolver) BlockHash() string {
	return lr.Log.BlockHash.Hex()
}

func (lr logResolver) Index() int32 {
	return int32(lr.Log.Index)
}

func (lr logResolver) TxIndex() int32 {
	return int32(lr.Log.TxIndex)
}

func (lr logResolver) Removed() bool {
	return lr.Log.Removed
}
