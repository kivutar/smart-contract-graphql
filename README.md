# Smart Contract GraphQL [![Build Status](https://travis-ci.org/kivutar/smart-contract-graphql.svg?branch=master)](https://travis-ci.org/kivutar/smart-contract-graphql) [![GoDoc](https://godoc.org/github.com/kivutar/smart-contract-graphql?status.svg)](https://godoc.org/github.com/kivutar/smart-contract-graphql)

This is a stateless smart contract event watcher exposing a GraphQL API. You can use it to monitor events from an already deployed smart contract using a GraphQL client.

Please note that you can perform the same tasks client side using web3 filters API instead of this project.

## Getting the source code

You need [dep](https://github.com/golang/dep)

    go get github.com/kivutar/smart-contract-graphql
    cd $GOPATH/src/github.com/kivutar/smart-contract-graphql
    dep ensure

## Building and running locally

    go build && PORT=3000 RPC_ENDPOINT="wss://rinkeby.infura.io/ws" ./smart-contract-graphql

## Example query

```
query {
  filterLogs (name: "Message", address: "0x5e626b58388f8b083d6d399f385c901675636b6e", abi: "[{ \"constant\": false, \"inputs\": [{ \"name\": \"_hashContent\", \"type\": \"string\" }, { \"name\": \"_hashImage\", \"type\": \"string\" }], \"name\": \"post\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [{ \"name\": \"\", \"type\": \"uint256\" }], \"name\": \"hashes\", \"outputs\": [{ \"name\": \"sender\", \"type\": \"address\" }, { \"name\": \"content\", \"type\": \"string\" }, { \"name\": \"image\", \"type\": \"string\" }, { \"name\": \"timestamp\", \"type\": \"uint256\" }], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [], \"name\": \"lastHashId\", \"outputs\": [{ \"name\": \"\", \"type\": \"uint256\" }], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [{ \"name\": \"_hashAvatar\", \"type\": \"string\" }], \"name\": \"set_avatar\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [{ \"name\": \"\", \"type\": \"address\" }], \"name\": \"avatars\", \"outputs\": [{ \"name\": \"\", \"type\": \"string\" }], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"constructor\" }, { \"anonymous\": false, \"inputs\": [{ \"indexed\": true, \"name\": \"sender\", \"type\": \"address\" }, { \"indexed\": false, \"name\": \"id\", \"type\": \"uint256\" }, { \"indexed\": false, \"name\": \"cid\", \"type\": \"string\" }, { \"indexed\": false, \"name\": \"img\", \"type\": \"string\" }, { \"indexed\": false, \"name\": \"timestamp\", \"type\": \"uint256\" }], \"name\": \"Message\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [{ \"indexed\": true, \"name\": \"sender\", \"type\": \"address\" }, { \"indexed\": false, \"name\": \"cid\", \"type\": \"string\" }], \"name\": \"Avatar\", \"type\": \"event\" }]") {
    Topics
    Data
    TxHash
    BlockNumber
    BlockHash
    Index
    TxIndex
    Address
    Removed
    Values
  }
}
```

## Example subscription

```
subscription {
  watchLogs (name: "Message", address: "0x5e626b58388f8b083d6d399f385c901675636b6e", abi: "[{ \"constant\": false, \"inputs\": [{ \"name\": \"_hashContent\", \"type\": \"string\" }, { \"name\": \"_hashImage\", \"type\": \"string\" }], \"name\": \"post\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [{ \"name\": \"\", \"type\": \"uint256\" }], \"name\": \"hashes\", \"outputs\": [{ \"name\": \"sender\", \"type\": \"address\" }, { \"name\": \"content\", \"type\": \"string\" }, { \"name\": \"image\", \"type\": \"string\" }, { \"name\": \"timestamp\", \"type\": \"uint256\" }], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [], \"name\": \"lastHashId\", \"outputs\": [{ \"name\": \"\", \"type\": \"uint256\" }], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [{ \"name\": \"_hashAvatar\", \"type\": \"string\" }], \"name\": \"set_avatar\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [{ \"name\": \"\", \"type\": \"address\" }], \"name\": \"avatars\", \"outputs\": [{ \"name\": \"\", \"type\": \"string\" }], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"constructor\" }, { \"anonymous\": false, \"inputs\": [{ \"indexed\": true, \"name\": \"sender\", \"type\": \"address\" }, { \"indexed\": false, \"name\": \"id\", \"type\": \"uint256\" }, { \"indexed\": false, \"name\": \"cid\", \"type\": \"string\" }, { \"indexed\": false, \"name\": \"img\", \"type\": \"string\" }, { \"indexed\": false, \"name\": \"timestamp\", \"type\": \"uint256\" }], \"name\": \"Message\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [{ \"indexed\": true, \"name\": \"sender\", \"type\": \"address\" }, { \"indexed\": false, \"name\": \"cid\", \"type\": \"string\" }], \"name\": \"Avatar\", \"type\": \"event\" }]") {
    Topics
    Data
    TxHash
    BlockNumber
    BlockHash
    Index
    TxIndex
    Address
    Removed
    Values
  }
}
```

Go to http://ipfs.io/ipfs/Qma2o1KZ8z75cbSZhVhEJPHW3L4hdM1jM6PWRuoS14YJqW/ and post a message using Metamask on Rinkeby. Wait a minute, you should see your message event log in graphiql right panel.

## Online demo

You can try it yourself on [https://smart-contract-graphql.herokuapp.com/](https://smart-contract-graphql.herokuapp.com/). The contract needs to be deployed on Rinkeby Testnet.

## TODO

 - Event Data decoding: Values should return a map with the right argument names and types
 - Passing `abi` could be avoided by calling `ethclient.Client.FilterLogs` instead of `BoundContract.FilterLogs`
 - Expose FilterOpts params to set the blockNumber range
 - Expose query params to build complex queries
 - Refactoring + tests + scaling tests
