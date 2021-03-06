package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/kivutar/smart-contract-graphql/handlers"
	"github.com/kivutar/smart-contract-graphql/resolvers"
)

func main() {
	port := os.Getenv("PORT")
	conn, err := ethclient.Dial(os.Getenv("RPC_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	schema := `
		type Log {
			Address: String!
			Topics: [String!]
			Data: String!
			Values: [String!]
			TxHash: String!
			BlockNumber: Int!
			BlockHash: String!
			Index: Int!
			TxIndex: Int!
			Removed: Boolean!
		}

		type Query {
			filterLogs(name: String!, address: String!, abi: String!): [Log!]
		}

		type Subscription {
			watchLogs(name: String!, address: String!, abi: String!): Log!
		}

		schema {
			subscription: Subscription
			query: Query
		}`

	s := graphql.MustParseSchema(schema, resolvers.NewResolver(conn))

	http.Handle("/graphql", handlers.NewGraphQLWSHandler(s, &relay.Handler{Schema: s}))

	http.HandleFunc("/", handlers.GraphiQL)

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
