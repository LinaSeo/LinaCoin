package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) addBlock(msg string) {
	prevHash := ""
	if len(b.blocks) > 0 {
		prevHash = b.blocks[len(b.blocks)-1].hash
	}
	newBlock := block{msg, "", prevHash}

	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	newBlock.hash = hexHash

	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Println(block.data)
		fmt.Println(block.prevHash)
		fmt.Println(block.hash)
	}
}

func main() {

}
