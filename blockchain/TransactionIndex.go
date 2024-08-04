package blockchain

type BlockHeightTxIndex struct {
	BlockHeight int `json:"block_height"`
	TxIndex     int `json:"tx_index"`
}

type TransactionIndex struct {
	TxBlockHeight map[string]*BlockHeightTxIndex `json:"transactions"`
}

func NewTransactionIndex() *TransactionIndex {
	return &TransactionIndex{
		TxBlockHeight: make(map[string]*BlockHeightTxIndex),
	}
}

func (ti *TransactionIndex) AddTransaction(txHash string, txIndex int, blockHeight int) {

	if ti.TxBlockHeight == nil {
		ti.TxBlockHeight = make(map[string]*BlockHeightTxIndex)
	}

	ti.TxBlockHeight[txHash] = &BlockHeightTxIndex{BlockHeight: blockHeight, TxIndex: txIndex}
}

func (ti *TransactionIndex) GetTransactionMetadata(txHash string) *BlockHeightTxIndex {
	return ti.TxBlockHeight[txHash]
}
