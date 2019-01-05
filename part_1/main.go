package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Li")
	bc.AddBlock("Send 2 BTC to Wang")

	for _, block := range bc.blocks {
		fmt.Printf("PrevHash:\t %x\n", block.PrevBlockHash)
		fmt.Printf("Data:\t\t %s\n", block.Data)
		fmt.Printf("Hash:\t\t %x\n", block.Hash)

		fmt.Println()
	}
}
