package resolvers

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kivutar/smart-contract-graphql/logsiterator"
)

// FilterLogs is a GraphQL query that returns logs by their name
func (r *Resolver) FilterLogs(ctx context.Context, args struct {
	Name    string
	Address string
	ABI     string
}) (*[]LogResolver, error) {
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
	iter := logsiterator.LogsIterator{
		Logs: logs,
		Sub:  sub,
	}
	var out []LogResolver
	for iter.Next() {
		out = append(out, LogResolver{iter.Log, &parsed, args.Name})
	}

	return &out, nil
}
