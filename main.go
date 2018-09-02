package main

import (
	"github.com/kodburn/blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.CloseDBConn()

	cli := blockchain.NewCLI(bc)
	cli.Run()
}
