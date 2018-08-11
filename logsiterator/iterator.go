package logsiterator

import (
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

// LogsIterator is a generic helper to iterate over smart contract event logs
type LogsIterator struct {
	Log  types.Log
	Logs chan types.Log
	Sub  ethereum.Subscription
	Done bool
	Fail error
}

// Next consumes the next entry from the logs channel
func (it *LogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.Fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.Done {
		select {
		case log := <-it.Logs:
			it.Log = log
			return true
		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.Logs:
		it.Log = log
		return true
	case err := <-it.Sub.Err():
		it.Done = true
		it.Fail = err
		return it.Next()
	}
}
