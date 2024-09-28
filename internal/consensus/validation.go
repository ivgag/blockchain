package consensus

import (
	"blockchain/internal/blockchain"
	"errors"
	"strings"
)

func ValidateBlockChain(blockchain blockchain.Blockchain) error {
	// skip genesis block
	for i := 1; i < len(blockchain.Blocks); i++ {
		block := blockchain.Blocks[i]
		previousBlock := blockchain.Blocks[i-1]

		if err := validateBlock(block); err != nil {
			return err
		}

		if block.Index != previousBlock.Index+1 {
			return errors.New("invalid index")
		}

		if block.PrevHash != previousBlock.Hash {
			return errors.New("invalid previous hash")
		}
	}

	return nil
}

func validateBlock(block blockchain.Block) error {
	if block.Hash != block.CalculateHash() {
		return errors.New("invalid hash")
	}

	target := strings.Repeat("0", block.Difficulty)
	if !strings.HasPrefix(block.Hash, target) || block.Hash != block.CalculateHash() {
		return errors.New("invalid proof of work")
	}

	return nil
}
