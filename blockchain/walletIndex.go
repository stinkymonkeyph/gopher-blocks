package blockchain

import "log"

type WalletTransactionIndex struct {
	BlockHeight      int          `json:"block_height"`
	TransactionIndex int          `json:"transaction_index"`
	Transaction      *Transaction `json:transactions"`
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
		if txn.Transaction.From == address {
			bal -= int(txn.Transaction.Value)
		} else if txn.Transaction.To == address {
			bal += int(txn.Transaction.Value)
		}
	}

	return bal
}
