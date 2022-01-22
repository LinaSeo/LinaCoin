package main

import (
	"fmt"

	"github.com/LinaSeo/LinaCoin/blockchain"
)

func main() {
	mainChain := blockchain.GetBlockchain()
	mainChain.AddBlock("Second")
	mainChain.AddBlock("Third")
	mainChain.AddBlock("Fourth")

	for _, block := range mainChain.AllBlocks() {
		fmt.Println("Data: %s\n", block.Data)
		fmt.Println("Hash: %s\n", block.Hash)
		fmt.Println("prevHash: %s\n", block.PrevHash)
	}
}
