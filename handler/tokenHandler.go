package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/webhook-issue-manager/service"
)

var (
	tokenService service.TokenService = service.NewTokenService()
)

type TokenHandler interface {
	CreateToken(c *fiber.Ctx) error
	TokenValidatorMiddleware(c *fiber.Ctx) error
}

type handler struct{}

func NewTokenHandler() TokenHandler {
	return &handler{}
}

func (*handler) CreateToken(c *fiber.Ctx) error {
	token, err := tokenService.CreateToken()
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": token.TokenStr})
}

//Token middleware
func (*handler) TokenValidatorMiddleware(c *fiber.Ctx) error {
	secretHeader := c.Get("X-Kondukto-Secret")

	if secretHeader == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "missing secret key"})
	}

	payload, err := verifyToken(secretHeader)
	fmt.Println(err)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": " " + err.Error()})
	}

	token, err := tokenService.GetToken(payload.ID.String())
	if err != nil {
		c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": err.Error()})
	}

	if secretHeader != token.TokenStr {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	return c.Next()

}
