package handlers

import "github.com/boltdb/bolt"

type Handler struct {
	DB *bolt.DB
}

func (h *Handler) NewTestConnectionHandlers() *IssueHandlers {
	return &IssueHandlers{
		DB: h.DB,
	}
}
