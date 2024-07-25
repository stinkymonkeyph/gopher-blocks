package blockchain

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/stinkymonkeyph/gopher-blocks/constants"
)

type Blockchain struct {
	TransactionPool []*Transaction `json:"transaction_pool"`
	Blocks          []*Block       `json:"block_chain"`
}

func NewBlockchain(genesisBlock *Block) *Blockchain {
	bc := new(Blockchain)
	bc.TransactionPool = []*Transaction{}
	bc.Blocks = append(bc.Blocks, genesisBlock)

	return bc
}

func (bc *Blockchain) ToJSON() string {
	bbc, err := json.Marshal(bc)

	if err != nil {
		panic("Something went wrong while serializing blockchain object")
	}

	return string(bbc)
}

func (bc *Blockchain) AddTransactionToTransactionPool(txn *Transaction) {
	txn.Status = constants.STATUS_PENDING
	bc.TransactionPool = append(bc.TransactionPool, txn)
}

func (bc *Blockchain) AddBlock(b *Block) {
	m := map[string]bool{}
	for _, txn := range b.Transactions {
		m[txn.TransactioHash] = true
	}

	for idx, txn := range bc.TransactionPool {
		_, ok := m[txn.TransactioHash]

		if ok {
			bc.TransactionPool = append(bc.TransactionPool[:idx], bc.TransactionPool[idx+1:]...)
		}
	}

	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) ProofOfWorkMining(minerAddress string) {
	nonce := 0
	for {
		prevHash := bc.Blocks[len(bc.Blocks)-1].Hash()
		guessBlock := NewBlock(prevHash, nonce)

		for _, txn := range bc.TransactionPool {
			tx := NewTransaction(txn.From, txn.To, txn.Value, txn.Data)
			guessBlock.AddTransactionToTheBlock(tx)
		}

		zeroes := strings.Repeat("0", constants.MINING_DIFFICULTY)
		guessHash := guessBlock.Hash()

		if zeroes == guessHash[2:2+constants.MINING_DIFFICULTY] {
			rewardTxn := NewTransaction(constants.BLOCKCHAIN_REWARD_ADDRESS, minerAddress, constants.MINING_REWARD, []byte{})
			guessBlock.Transactions = append(guessBlock.Transactions, rewardTxn)
			bc.AddBlock(guessBlock)
			log.Printf(bc.ToJSON(), "\n\n")
			nonce = 0
			continue
		}
	}
}
