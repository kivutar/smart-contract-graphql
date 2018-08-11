package main

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (r *resolver) FilterLogs(ctx context.Context, args struct {
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
		common.HexToAddress(args.Address), parsed, r.conn, r.conn, r.conn)

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
