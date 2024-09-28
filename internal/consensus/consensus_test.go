package consensus

import (
	"blockchain/internal/blockchain"
	"testing"
	"time"
)

func TestCalculateNewDifficulty(t *testing.T) {
	var chain []blockchain.Block
	baseTime := time.Now()
	for i := 0; i < DifficultyAdjustmentInterval; i++ {
		block := blockchain.Block{
			Index:      i,
			Timestamp:  baseTime.Add(time.Duration(i*TargetBlockTimeInSeconds) * time.Second),
			Data:       "Test Block",
			PrevHash:   "prevhash",
			Hash:       "hash",
			Nonce:      0,
			Difficulty: 2,
		}
		chain = append(chain, block)
	}

	newDifficulty := CalculateNewDifficulty(chain)
	expectedDifficulty := 2
	if newDifficulty != expectedDifficulty {
		t.Errorf("expected difficulty %d, got %d", expectedDifficulty, newDifficulty)
	}
}

func TestAddNewBlock(t *testing.T) {
	var bc = blockchain.NewBlockchain()
	var pow = NewPoWConsensus(&bc)

	var block = pow.CreateNewBlock("Test Block")

	if result := pow.ValidateBlock(block); result != nil {
		t.Errorf("failed to validate block: " + result.Error())
	}

	if result := pow.AddNewBlock(block); result != nil {
		t.Errorf("failed to add block to blockchain: " + result.Error())
	}

	if result := ValidateBlockChain(bc); result != nil {
		t.Errorf("failed to validate blockchain: " + result.Error())
	}
}
