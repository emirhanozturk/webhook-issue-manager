package db

import (
	"log"

	"github.com/boltdb/bolt"
)

func Inıt() *bolt.DB {

	db, err := bolt.Open("issues.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
