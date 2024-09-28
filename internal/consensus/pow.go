package consensus

import (
	"blockchain/internal/blockchain"
	"strings"
)

func ProofOfWork(block *blockchain.Block) (string, int) {
	nonce := 0
	var hash string
	target := strings.Repeat("0", block.Difficulty)

	for {
		hash = block.CalculateHashWithNonce(nonce)
		if strings.HasPrefix(hash, target) {
			break
		}
		nonce++
	}
	return hash, nonce
}
