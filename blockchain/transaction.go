package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math"

	"github.com/stinkymonkeyph/gopher-blocks/constants"
)

type Transaction struct {
	From           string `json:"from"`
	To             string `json:"to"`
	Value          uint64 `json:"value"`
	Data           []byte `json:"data"`
	Status         string `json:"status"`
	TransactioHash string `json:"transaction_hash"`
}

func NewTransaction(from string, to string, value uint64, data []byte) *Transaction {
	t := new(Transaction)
	t.From = from
	t.To = to
	t.Value = value
	t.Data = data
	t.TransactioHash = ""
	return t
}

func (t *Transaction) ToJSON() string {
	tb, err := json.Marshal(t)

	if err != nil {
		panic("Something went wrong while serializing transaction object")
	}

	return string(tb)
}

func (t *Transaction) VerifyTransaction() bool {
	if t.Value == 0 {
		return false
	}

	if t.Value > math.MaxInt64 {
		return false
	}

	// TODO: check signature of transaction here

	return true
}

func (t *Transaction) Hash() string {
	ts, err := json.Marshal(t)

	if err != nil {
		panic("Something went wrong while serializing transaction object")
	}

	sum := sha256.Sum256(ts)
	hex := hex.EncodeToString(sum[:32])
	hex = constants.HEX_PREFIX + hex

	return hex
}
