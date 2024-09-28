package handlers

import (
	"encoding/json"
	"net/http"

	"blockchain/internal/blockchain"
)

func GetChainHandler(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"length": len(bc.Blocks()),
			"blocks": bc.Blocks(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
