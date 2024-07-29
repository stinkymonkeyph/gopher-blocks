package blockchain

import (
	"encoding/json"

	"github.com/dgraph-io/badger"
	"github.com/stinkymonkeyph/gopher-blocks/constants"
)

func PutIntoDb(bc *Blockchain) error {
	db, err := badger.Open(badger.DefaultOptions(constants.DB_PATH))

	if err != nil {
		return err
	}

	defer db.Close()

	value, err := json.Marshal(bc)

	if err != nil {
		return err
	}

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(constants.DB_KEY), value)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}
