package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
			subscribers[randomID()] = s
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
			select {
			case log := <-logs:
				fmt.Println(log)
				r.logResolvers <- logResolver{log, &parsed, args.Name}
			}
		}
	}()

	return logResolvers, nil
}

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
