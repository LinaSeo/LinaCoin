package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	// store pointers not all data of blocks
	blocks []*block
}

// Singleton
var b *blockchain

// Run only once while parallel env
var once sync.Once

func getLastHash() string {
	height := len(b.AllBlocks()) - 1
	if height == 0 {
		return ""
	}
	return b.blocks[height].Hash
}

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	hexHash := fmt.Sprintf("%x", hash)
	b.Hash = hexHash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()

	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

// Blockchain reset before creating object
// Initialization
func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}
