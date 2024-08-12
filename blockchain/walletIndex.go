package blockchain

import (
	"log"
)

type WalletTransactionIndex struct {
	BlockIndex       int          `json:"block_index"`
	TransactionIndex int          `json:"transaction_index"`
	Transaction      *Transaction `json:"transactions"`
}

type WalletIndex struct {
	Transactions map[string][]*WalletTransactionIndex `json:"transactions"`
}

func NewWalletIndex() *WalletIndex {
	return &WalletIndex{
		Transactions: make(map[string][]*WalletTransactionIndex, 0),
	}
}

func (w *WalletIndex) AddTransaction(address string, blockIndex int, txIndex int, transaction *Transaction) {
	if w.Transactions == nil {
		w.Transactions = make(map[string][]*WalletTransactionIndex)
	}

	wti := &WalletTransactionIndex{
		BlockIndex:       blockIndex,
		TransactionIndex: txIndex,
		Transaction:      transaction,
	}

	log.Printf("Adding transaction: Address=%s, BlockIndex=%d, TxIndex=%d", address, blockIndex, txIndex)
	w.Transactions[address] = append(w.Transactions[address], wti)
	log.Printf("Transactions for %s: %+v", address, w.Transactions[address])
}

func (w *WalletIndex) GetWalletTransactions(address string) []*WalletTransactionIndex {
	if w.Transactions == nil {
		w.Transactions = make(map[string][]*WalletTransactionIndex, 0)
	}

	return w.Transactions[address]
}
