package main

import (
	"fmt"

	"github.com/kodburn/blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	bi := bc.Iterator()

	for {
		block, last := bi.Next()
		if last {
			break
		} else {
			fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			pow := blockchain.NewProofOfWork(block)
			fmt.Printf("PoW: %t\n", pow.Validate())
			fmt.Println()
		}
	}
}
