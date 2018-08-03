package main

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type resolver struct {
	conn *ethclient.Client
}

type logResolver struct {
	types.Log
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

func (r *resolver) Logs(ctx context.Context, args struct {
	Name    string
	Address string
	ABI     string
}) (*[]logResolver, error) {
	contract, err := bindContract(args.ABI, common.HexToAddress(args.Address), r.conn, r.conn, r.conn)
	if err != nil {
		return nil, err
	}

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
		out = append(out, logResolver{iter.log})
	}

	return &out, nil
}

func bindContract(ABI string, address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}
