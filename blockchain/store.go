package blockchain

import (
	"encoding/json"

	"github.com/dgraph-io/badger"
	"github.com/stinkymonkeyph/gopher-blocks/constants"
)

func openDb() (*badger.DB, error) {
	db, err := badger.Open(badger.DefaultOptions(constants.DB_PATH))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func PutIntoDb(bc *Blockchain) error {

	db, err := openDb()

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

	return err
}

func ReadFromDb() (Blockchain, error) {
	var bc Blockchain

	db, err := openDb()

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

	return bc, err
}
