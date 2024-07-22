package blockchain

import "encoding/json"

type Transaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value uint64 `json:"value"`
	Data  []byte `json:"data"`
}

func NewTransaction(from string, to string, value uint64, data []byte) *Transaction {
	t := new(Transaction)
	t.From = from
	t.To = to
	t.Value = value
	t.Data = data

	return t
}

func (t *Transaction) ToJSON() string {
	tb, err := json.Marshal(t)

	if err != nil {
		panic("Something went wrong while serializing transaction object")
	}

	return string(tb)
}
