package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/kivutar/smart-contract-graphql/graphqlwshandler"
)

func main() {
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

		type HelloSaidEvent {
			id: String!
			msg: String!
		}

		type Query {
				logs(name: String!, address: String!, abi: String!): [Log!]
		}

		type Mutation {
			sayHello(msg: String!): HelloSaidEvent!
		}

		type Subscription {
			helloSaid(): HelloSaidEvent!
		}

		schema {
			subscription: Subscription
			mutation: Mutation
			query: Query
		}`

	s := graphql.MustParseSchema(schema, newResolver(conn))

	http.Handle("/graphql", graphqlwshandler.NewHandler(s, &relay.Handler{Schema: s}))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	log.Println("listening on", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
