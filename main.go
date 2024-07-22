package main

import (
	"log"

	"github.com/stinkymonkeyph/gopher-blocks/blockchain"
	"github.com/stinkymonkeyph/gopher-blocks/constants"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ": ")
}

func main() {
	block := blockchain.NewBlock("0x0", 0)
	bc := blockchain.NewBlockchain(block)
	log.Print(bc.ToJSON())

}
