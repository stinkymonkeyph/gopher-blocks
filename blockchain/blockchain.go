package blockchain

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/stinkymonkeyph/gopher-blocks/constants"
)

type Blockchain struct {
	TransactionPool  []*Transaction `json:"transaction_pool"`
	Blocks           []*Block       `json:"block_chain"`
	WalletIndex      *WalletIndex
	TransactionIndex *TransactionIndex
}

func (bc *Blockchain) Airdrop(address string) {
	txn := NewTransaction(constants.BLOCKCHAIN_AIRDROP_ADDRESS, address, constants.AIRDROP_AMOUNT, []byte{})
	bc.AddTransactionToTransactionPool(txn)
}

func NewBlockchain(genesisBlock *Block) *Blockchain {
	state, err := ReadFromDb()
	bc := &Blockchain{}

	if err != nil {
		bc = new(Blockchain)
		bc.TransactionPool = []*Transaction{}
		bc.Blocks = append(bc.Blocks, genesisBlock)
		bc.WalletIndex = NewWalletIndex()
		bc.TransactionIndex = NewTransactionIndex()
		err := PutIntoDb(bc)

		if err != nil {
			log.Panicf(err.Error())
		}

	} else {
		log.Println("Found existing blockchain state, persisting state from datastore")
		bc = &state
	}

	return bc
}

func (bc *Blockchain) ToJSON() string {
	bbc, err := json.Marshal(bc)

	if err != nil {
		log.Panic(err.Error())
	}

	return string(bbc)
}

func (bc *Blockchain) AddTransactionToTransactionPool(txn *Transaction) {
	txn.Status = constants.STATUS_PENDING
	bc.TransactionPool = append(bc.TransactionPool, txn)
	err := PutIntoDb(bc)
	if err != nil {
		log.Default().Panicf(err.Error())
	}
}

func (bc *Blockchain) AddBlock(b *Block) {
	m := map[string]bool{}

	nextBlockHeight := len(bc.Blocks)

	for index, txn := range b.Transactions {
		m[txn.TransactioHash] = true
		balance := bc.WalletIndex.CalculateBalance(txn.From)
		log.Printf("\n\nsender balance -> %d \n\n", balance)
		bc.WalletIndex.AddTransaction(txn.From, nextBlockHeight, index, txn)
		bc.WalletIndex.AddTransaction(txn.To, nextBlockHeight, index, txn)
		bc.TransactionIndex.AddTransaction(txn.TransactioHash, index, nextBlockHeight)
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
		log.Panic(err.Error())
	}
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	t := make([]*Transaction, 0)
	for _, txn := range bc.TransactionPool {
		t = append(t, NewTransaction(txn.From, txn.To, txn.Value, txn.Data))
	}
	return t
}

func (bc *Blockchain) ValidProof(nonce int, previousHash string, transactions []*Transaction, difficulty int) bool {
	zeroes := strings.Repeat("0", difficulty)
	guessBlock := &Block{PrevHash: previousHash, Timestamp: 0, Nonce: nonce, Transactions: transactions}
	guessHash := guessBlock.Hash()
	return zeroes == guessHash[2:2+constants.MINING_DIFFICULTY]
}

func (bc *Blockchain) ProofOfWork() (int, []*Transaction) {
	t := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0

	for !bc.ValidProof(nonce, previousHash, t, constants.MINING_DIFFICULTY) {
		nonce += 1
	}

	return nonce, t
}

func (bc *Blockchain) Mining() bool {
	log.Println("Start proof of work")
	nonce, txns := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	block := NewBlock(previousHash, nonce)
	block.Transactions = txns
	bc.AddBlock(block)
	log.Println("Found solution")
	return true
}
