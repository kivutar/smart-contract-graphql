package main

import (
	"context"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/xid"
)

type watchLogsSubscriber struct {
	stop         <-chan struct{}
	logResolvers chan<- logResolver
}

func (r *resolver) broadcastWatchLogs() {
	subscribers := map[string]*watchLogsSubscriber{}
	unsubscribe := make(chan string)

	// NOTE: subscribing and sending events are at odds.
	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.watchLogsSubscriber:
			subscribers[xid.New().String()] = s
		case logResolver := <-r.logResolvers:
			for id, s := range subscribers {
				go func(id string, s *watchLogsSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.logResolvers <- logResolver:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}

func (r *resolver) WatchLogs(ctx context.Context, args struct {
	Name    string
	Address string
	ABI     string
}) (<-chan logResolver, error) {
	logResolvers := make(chan logResolver)
	r.watchLogsSubscriber <- &watchLogsSubscriber{logResolvers: logResolvers, stop: ctx.Done()}

	// Parse ABI
	parsed, err := abi.JSON(strings.NewReader(args.ABI))
	if err != nil {
		return logResolvers, err
	}

	// Bind deployed smart contract
	contract := bind.NewBoundContract(
		common.HexToAddress(args.Address), parsed, r.conn, r.conn, r.conn)

	// Get logs from contract
	logs, _, err := contract.WatchLogs(
		&bind.WatchOpts{Context: ctx},
		args.Name,
	)
	if err != nil {
		return logResolvers, err
	}

	go func() {
		for {
			r.logResolvers <- logResolver{<-logs, &parsed, args.Name}
		}
	}()

	return logResolvers, nil
}
