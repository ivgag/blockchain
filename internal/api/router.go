// internal/api/router.go
package api

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/consensus"

	"blockchain/internal/api/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(bc *blockchain.Blockchain, cons *consensus.PoWConsensus) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/chain", handlers.GetChainHandler(bc)).Methods("GET")
	router.HandleFunc("/blocks/{index}", handlers.GetBlockHandler(bc)).Methods("GET")

	return router
}
