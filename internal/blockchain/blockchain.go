package blockchain

type Blockchain struct {
	Blocks     []Block
	Difficulty int
}

func NewBlockchain() Blockchain {
	return Blockchain{
		Blocks:     []Block{GenesisBlock(2)},
		Difficulty: 2,
	}
}

func (bc *Blockchain) AddBlock(block Block) {
	bc.Blocks = append(bc.Blocks, block)
}

func (bc *Blockchain) FirstBlock() Block {
	return bc.Blocks[0]
}

func (bc *Blockchain) LastBlock() Block {
	return bc.Blocks[len(bc.Blocks)-1]
}
