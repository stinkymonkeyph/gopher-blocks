package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/stinkymonkeyph/gopher-blocks/constants"
)

type Block struct {
	Number       int            `json:"number"`
	PrevHash     string         `json:"prevHash"`
	Timestamp    int64          `json:"timestamp"`
	Nonce        int            `json:"nonce"`
	Transactions []*Transaction `json:"transactions"`
	MerkleRoot   string         `json:"merkle_root"`
}

func NewBlock(prevHash string, blockHeight int, nonce int, txns []*Transaction) *Block {
	block := new(Block)
	block.PrevHash = prevHash
	block.Timestamp = time.Now().UnixMicro()
	block.Nonce = nonce
	block.Number = blockHeight
	if txns != nil {
		block.Transactions = txns
		leaves := CreateLeafNodes(block.Transactions)
		merkleTree := BuildMerkleTree(leaves)
		merkleRoot := merkleTree.Hash
		block.MerkleRoot = fmt.Sprintf("%x", merkleRoot)
	} else {
		block.Transactions = make([]*Transaction, 0)
		block.MerkleRoot = ""
	}

	return block
}

func (b *Block) ToJson() string {
	bb, err := json.Marshal(b)

	if err != nil {
		panic("Something went wrong while serializing block object")
	}

	return string(bb)
}

func (b *Block) Hash() string {
	bs, err := json.Marshal(b)

	if err != nil {
		panic("Something went wrong while serializing block object")
	}

	sum := sha256.Sum256(bs)
	hex := hex.EncodeToString(sum[:32])
	hex = constants.HEX_PREFIX + hex

	return hex
}

func (b *Block) AddTransactionToTheBlock(txn *Transaction) {
	if txn.Status == constants.STATUS_PENDING {
		isValid := txn.VerifyTransaction()

		if isValid {
			txn.Status = constants.STATUS_SUCCESS
		} else {
			txn.Status = constants.STATUS_FAILED
		}

		b.Transactions = append(b.Transactions, txn)

	}
}
