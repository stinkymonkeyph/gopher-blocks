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

func ReadFromDb() (Blockchain, error) {
	var bc Blockchain

	db, err := badger.Open(badger.DefaultOptions(constants.DB_PATH))

	if err != nil {
		return bc, err
	}

	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(constants.DB_KEY))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		err = json.Unmarshal(val, &bc)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return bc, err
	}

	return bc, nil
}
