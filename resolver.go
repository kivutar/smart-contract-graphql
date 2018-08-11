package main

import "github.com/ethereum/go-ethereum/ethclient"

type resolver struct {
	helloSaidEvents     chan *helloSaidEvent
	helloSaidSubscriber chan *helloSaidSubscriber
	conn                *ethclient.Client
}

func newResolver(conn *ethclient.Client) *resolver {
	r := &resolver{
		helloSaidEvents:     make(chan *helloSaidEvent),
		helloSaidSubscriber: make(chan *helloSaidSubscriber),
		conn:                conn,
	}

	go r.broadcastHelloSaid()

	return r
}
