package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TestHandlers struct {
	DB *gorm.DB
}

func (h *Handler) NewTestConnectionHandlers() *IssueHandlers {
	return &IssueHandlers{
		DB: h.DB,
	}
}

func (i *IssueHandlers) TestIssue(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusOK)
	c.JSON(fiber.Map{"connection": true})
	return nil
}
