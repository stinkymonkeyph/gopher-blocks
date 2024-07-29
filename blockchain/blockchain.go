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
	state, err := ReadFromDb()
	bc := &Blockchain{}

	if err != nil {
		bc = new(Blockchain)
	} else {
		log.Println("Found existing blockchain state, persisting state from datastore")
		bc = &state
	}

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

	err := PutIntoDb(bc)

	if err != nil {
		log.Panicf("Something went wrong while saving state to database, halting entire process %s", err)
	}
}

func (bc *Blockchain) ProofOfWorkMining(minerAddress string) {
	nonce := 0
	log.Println("Starting Proof of Work")
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
			log.Println("Found solution")
			log.Printf("Mining Difficulty is %d", constants.MINING_DIFFICULTY)
			log.Printf("Hash Solution: %s", guessHash)
			rewardTxn := NewTransaction(constants.BLOCKCHAIN_REWARD_ADDRESS, minerAddress, constants.MINING_REWARD, []byte{})
			rewardTxn.Status = constants.STATUS_SUCCESS
			guessBlock.Transactions = append(guessBlock.Transactions, rewardTxn)
			bc.AddBlock(guessBlock)
			log.Printf("%s \n\n", bc.ToJSON())
			nonce = 0
			continue
		}
	}
}
