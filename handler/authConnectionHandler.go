package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type TestHandler interface {
	TestIssue(c *fiber.Ctx) error
}

type testhandler struct{}

func NewTestConnectionHandlers() TestHandler {
	return &testhandler{}
}

func (*testhandler) TestIssue(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusOK)
	c.JSON(fiber.Map{"connection": true})
	return nil
}
