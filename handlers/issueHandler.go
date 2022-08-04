package handlers

import (
	"github.com/boltdb/bolt"
)

type IssueHandlers struct {
	DB *bolt.DB
}
