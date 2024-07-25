package main

import (
	"log"
	"sync"
	"time"

	"github.com/stinkymonkeyph/gopher-blocks/blockchain"
	"github.com/stinkymonkeyph/gopher-blocks/constants"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ": ")
}

func main() {

	var wg sync.WaitGroup

	block := blockchain.NewBlock("0x0", 0)
	transaction1 := blockchain.NewTransaction("0x1", "0x2", 12, []byte{})
	bc := blockchain.NewBlockchain(block)
	log.Print(bc.ToJSON())

	wg.Add(1)
	go bc.ProofOfWorkMining("alice")
	time.Sleep(2000)
	bc.AddTransactionToTransactionPool(transaction1)
	wg.Wait()
}
