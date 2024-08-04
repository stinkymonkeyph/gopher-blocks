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
	bc.Airdrop("0x1")
	bc.Mining()
	log.Print(bc.ToJSON())

	transaction1 := blockchain.NewTransaction("0x1", "0x2", 12, []byte{})
	bc.AddTransactionToTransactionPool(transaction1)
	bc.Mining()
	log.Print(bc.ToJSON())
	senderBalance := bc.WalletIndex.CalculateBalance("0x1")
	log.Printf("\n\n\nSender Balance: %d \n", senderBalance)
}
