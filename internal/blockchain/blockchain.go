package blockchain

import (
	"errors"
	"sync"
)

type Blockchain struct {
	blocks     []Block
	Difficulty int
	mutex      sync.Mutex
}

func NewBlockchain() Blockchain {
	return Blockchain{
		blocks:     []Block{GenesisBlock(2)},
		Difficulty: 2,
		mutex:      sync.Mutex{},
	}
}

func (bc *Blockchain) Blocks() []Block {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	return bc.blocks
}

func (bc *Blockchain) AddBlock(block Block) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	bc.blocks = append(bc.blocks, block)
}

func (bc *Blockchain) FirstBlock() Block {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	return bc.blocks[0]
}

func (bc *Blockchain) LastBlock() Block {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) GetBlock(index int) (Block, error) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	if index < 0 || index >= len(bc.blocks) {
		return Block{}, errors.New("invalid index")
	}
	return bc.blocks[index], nil
}
