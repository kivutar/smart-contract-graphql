package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type rootResolver struct {
	conn *ethclient.Client
}

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

func (qr *rootResolver) Logs(ctx context.Context, args struct {
	Name    string
	Address string
	ABI     string
}) (*[]logResolver, error) {
	// Parse ABI
	parsed, err := abi.JSON(strings.NewReader(args.ABI))
	if err != nil {
		return nil, err
	}

	// Bind deployed smart contract
	contract := bind.NewBoundContract(
		common.HexToAddress(args.Address), parsed, qr.conn, qr.conn, qr.conn)

	// Get logs from contract
	logs, sub, err := contract.FilterLogs(
		&bind.FilterOpts{Context: ctx},
		args.Name,
	)
	if err != nil {
		return nil, err
	}

	// Iterate over logs to build the output
	iter := logsIterator{
		logs: logs,
		sub:  sub,
	}
	var out []logResolver
	for iter.Next() {
		out = append(out, logResolver{iter.log, &parsed, args.Name})
	}

	return &out, nil
}
