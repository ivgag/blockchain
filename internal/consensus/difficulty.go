package consensus

import (
	"blockchain/internal/blockchain"
)

const (
	TargetBlockTimeInSeconds      = 600
	DifficultyAdjustmentInterval  = 2016
	ExpectedBlockGenerationTime   = int64(TargetBlockTimeInSeconds * DifficultyAdjustmentInterval)
	MaxDifficultyAdjustmentFactor = int64(4)
)

func CalculateNewDifficulty(chain []blockchain.Block) int {
	if len(chain) < DifficultyAdjustmentInterval {
		return 1
	}

	latestBlock := chain[len(chain)-1]
	firstBlock := chain[len(chain)-DifficultyAdjustmentInterval]

	actualTime := latestBlock.Timestamp.Unix() - firstBlock.Timestamp.Unix()

	if actualTime < ExpectedBlockGenerationTime/MaxDifficultyAdjustmentFactor {
		actualTime = ExpectedBlockGenerationTime / MaxDifficultyAdjustmentFactor
	}
	if actualTime > ExpectedBlockGenerationTime*MaxDifficultyAdjustmentFactor {
		actualTime = ExpectedBlockGenerationTime * MaxDifficultyAdjustmentFactor
	}

	newDifficulty := float64(latestBlock.Difficulty) * float64(ExpectedBlockGenerationTime) / float64(actualTime)
	if newDifficulty < 1 {
		newDifficulty = 1
	}
	return int(newDifficulty)
}
