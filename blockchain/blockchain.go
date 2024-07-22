package blockchain

import "encoding/json"

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
