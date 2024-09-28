package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"blockchain/internal/blockchain"

	"github.com/gorilla/mux"
)

func GetBlockHandler(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		indexStr, ok := vars["index"]
		if !ok {
			http.Error(w, "Index is required", http.StatusBadRequest)
			return
		}

		index, err := strconv.Atoi(indexStr)
		if err != nil {
			http.Error(w, "Invalid index format", http.StatusBadRequest)
			return
		}

		block, err := bc.GetBlock(index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(block)
	}
}
