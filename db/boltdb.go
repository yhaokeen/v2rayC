package db

import (
	"log"
	"path/filepath"
	"sync"

	"go.etcd.io/bbolt"
)

var once sync.Once
var db *bbolt.DB

func initDB() {
	confPath := "$HOME/.config/v2raya"
	dbPath := filepath.Join(confPath, "bolt.db")

	var err error
	db, err = bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal("bbolt.Open: %v", err)
	}
}

func DB() *bbolt.DB {
	once.Do(initDB)
	return db
}

func Transaction(db *bbolt.DB, fn func(*bbolt.Tx) (bool, error)) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	dirty, err := fn(tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if !dirty {
		return tx.Rollback()
	}
	return tx.Commit()
}

// If the bucket does not exist, the dirty flag is setted
func CreateBucketIfNotExists(tx *bbolt.Tx, name []byte, dirty *bool) (*bbolt.Bucket, error) {
	bkt := tx.Bucket(name)
	if bkt != nil {
		return bkt, nil
	}
	bkt, err := tx.CreateBucket(name)
	if err != nil {
		return nil, err
	}
	*dirty = true
	return bkt, nil
}

func GetBucketAll(bucket string) (items map[string][]byte, err error) {
	err = Transaction(DB(), func(tx *bbolt.Tx) (bool, error) {
		dirty := false
		if bkt, err := CreateBucketIfNotExists(tx, []byte(bucket), &dirty); err != nil {
			return dirty, err
		} else {
			c := bkt.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				items[string(k)] = v
			}
		}
		return dirty, nil
	})
	return items, err
}
