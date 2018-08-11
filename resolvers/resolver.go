package resolvers

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// Resolver is the root resolver of the GraphQL endpoint
type Resolver struct {
	logResolvers        chan LogResolver
	watchLogsSubscriber chan *watchLogsSubscriber
	conn                *ethclient.Client
}

func NewResolver(conn *ethclient.Client) *Resolver {
	r := &Resolver{
		logResolvers:        make(chan LogResolver),
		watchLogsSubscriber: make(chan *watchLogsSubscriber),
		conn:                conn,
	}

	go r.broadcastWatchLogs()

	return r
}
