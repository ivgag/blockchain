// cmd/myblockchain/main.go
package main

import (
	"blockchain/internal/api"
	"blockchain/internal/blockchain"
	"blockchain/internal/consensus"
	"log"
	"net/http"
)

func main() {
	bc := blockchain.NewBlockchain()

	cons := consensus.NewPoWConsensus(&bc)

	router := api.NewRouter(&bc, cons)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
