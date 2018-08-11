package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type resolver struct {
	logResolvers        chan logResolver
	watchLogsSubscriber chan *watchLogsSubscriber
	conn                *ethclient.Client
}

func newResolver(conn *ethclient.Client) *resolver {
	r := &resolver{
		logResolvers:        make(chan logResolver),
		watchLogsSubscriber: make(chan *watchLogsSubscriber),
		conn:                conn,
	}

	go r.broadcastWatchLogs()

	return r
}
