package handlers

import (
	"github.com/boltdb/bolt"
	"github.com/gofiber/fiber/v2"
)

type TestHandler struct {
	DB *bolt.DB
}

func New(db *bolt.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

func (i *IssueHandlers) TestIssue(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusOK)
	c.JSON(fiber.Map{"connection": true})
	return nil
}
