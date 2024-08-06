package blockchain

import (
	"log"

	"github.com/stinkymonkeyph/gopher-blocks/constants"
)

type WalletTransactionIndex struct {
	BlockHeight      int          `json:"block_height"`
	TransactionIndex int          `json:"transaction_index"`
	Transaction      *Transaction `json:"transactions"`
}

type WalletIndex struct {
	Transactions map[string][]*WalletTransactionIndex `json:"transactions"`
}

func NewWalletIndex() *WalletIndex {
	return &WalletIndex{
		Transactions: make(map[string][]*WalletTransactionIndex),
	}
}

func (w *WalletIndex) AddTransaction(address string, blockHeight int, txIndex int, transaction *Transaction) {
	if w.Transactions == nil {
		w.Transactions = make(map[string][]*WalletTransactionIndex)
	}

	wti := &WalletTransactionIndex{
		BlockHeight:      blockHeight,
		TransactionIndex: txIndex,
		Transaction:      transaction,
	}

	log.Printf("Adding transaction: Address=%s, BlockHeight=%d, TxIndex=%d", address, blockHeight, txIndex)
	w.Transactions[address] = append(w.Transactions[address], wti)
	log.Printf("Transactions for %s: %+v", address, w.Transactions[address])
}

func (w *WalletIndex) CalculateBalance(address string) int {
	bal := 0

	for _, txn := range w.Transactions[address] {
		if txn.Transaction.From == address && txn.Transaction.Status == constants.STATUS_SUCCESS {
			bal -= int(txn.Transaction.Value)
		} else if txn.Transaction.To == address && txn.Transaction.Status == constants.STATUS_SUCCESS {
			bal += int(txn.Transaction.Value)
		}
	}

	return bal
}

func (w *WalletIndex) GetWalletTransactions(address string) []*Transaction {
	t := make([]*Transaction, 0)

	if w.Transactions[address] != nil {
		for _, txi := range w.Transactions[address] {
			t = append(t, txi.Transaction)
		}
	}

	return t
}
