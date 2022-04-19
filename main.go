package main

import "fmt"

func main() {
	bc := NewBlockchain()
	bc.AddBlock("send 1 btc to vincent")
	bc.AddBlock("sent more vincent")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x \n", block.PrevBlockHash)
		fmt.Printf("Data: %s \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)
		fmt.Println()
	}
}
