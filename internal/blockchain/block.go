package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index      int
	Timestamp  time.Time
	Data       string
	PrevHash   string
	Hash       string
	Nonce      int
	Difficulty int
}

func GenesisBlock(difficulty int) Block {
	return Block{
		Index:      0,
		Timestamp:  time.Now(),
		Data:       "Genesis Block",
		PrevHash:   "",
		Hash:       "",
		Nonce:      0,
		Difficulty: difficulty,
	}
}

func (b *Block) CalculateHash() string {
	return b.CalculateHashWithNonce(b.Nonce)
}

func (b *Block) CalculateHashWithNonce(nonce int) string {
	record := strconv.Itoa(b.Index) + b.Timestamp.String() + b.Data + b.PrevHash + strconv.Itoa(nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
