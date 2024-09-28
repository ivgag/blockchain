package consensus

import (
	"blockchain/internal/blockchain"
	"errors"
	"log"
	"sync"
	"time"
)

type Consensus interface {
	ValidateBlock(block blockchain.Block) error
	SelectChain(chains [][]blockchain.Block) []blockchain.Block
	AdjustDifficulty(chain []blockchain.Block) int
	AddNewBlock(block blockchain.Block) error
}

type PoWConsensus struct {
	blockchain *blockchain.Blockchain
	mutex      sync.Mutex
}

func NewPoWConsensus(bc *blockchain.Blockchain) *PoWConsensus {
	return &PoWConsensus{
		blockchain: bc,
	}
}

func (c *PoWConsensus) CreateNewBlock(data string) blockchain.Block {
	lastBlock := c.blockchain.LastBlock()
	newBlock := blockchain.Block{
		Index:      lastBlock.Index + 1,
		Timestamp:  time.Now(),
		Data:       data,
		PrevHash:   lastBlock.Hash,
		Difficulty: lastBlock.Difficulty,
	}

	newBlock.Hash, newBlock.Nonce = ProofOfWork(&newBlock)
	return newBlock
}

func (c *PoWConsensus) AdjustDifficulty() int {
	return CalculateNewDifficulty(c.blockchain.Blocks)
}

func (c *PoWConsensus) AddNewBlock(block blockchain.Block) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if err := c.ValidateBlock(block); err != nil {
		return err
	}

	lastBlock := c.blockchain.LastBlock()
	if block.PrevHash != lastBlock.Hash {
		return errors.New("invalid previous hash")
	}

	c.blockchain.AddBlock(block)
	log.Printf("block %d added to the chain\n", block.Index)

	c.AdjustDifficulty()

	return nil
}

func (c *PoWConsensus) ValidateBlock(block blockchain.Block) error {
	return validateBlock(block)
}
