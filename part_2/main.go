package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to li")
	bc.AddBlock("Send 2 BTC to li")

	for _, block := range bc.blocks {
		fmt.Printf("prevHash:\t%x\n", block.PrevBlockHash)
		fmt.Printf("Data:\t\t%s\n", block.Data)
		fmt.Printf("Hash:\t\t%x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("Pow:\t\t%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
