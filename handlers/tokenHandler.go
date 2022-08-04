package handlers

import "github.com/boltdb/bolt"

type TokenHandler struct {
	DB *bolt.DB
}
