package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	handler "github.com/webhook-issue-manager/handlers"
	"github.com/webhook-issue-manager/model"
	db "github.com/webhook-issue-manager/storage/postgres"
)

var (
	secretKey       = "secretKey"
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired Token")
)

func Router() *fiber.App {

	app := fiber.New()
	handler := handler.New(db.Inıt())
	app.Use(logger.New())
	app.Post("/tokens", createToken)
	v1 := app.Group("api/v1")
	{
		testGroup := v1.Group("test")
		testGroup.Use(tokenValidatorMiddleware)
		testHandler := handler.NewTestConnectionHandlers()
		testGroup.Get("/", testHandler.TestIssue)

	}

	app.Listen(":3000")

	return app
}

func tokenValidatorMiddleware(c *fiber.Ctx) error {
	secretHeader := c.Get("X-Kondukto-Secret")

	if secretHeader == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "missing secret key"})
	}

	payload, err := verifyToken(secretHeader)
	fmt.Println(err)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": " " + err.Error()})
	}

	handler := handler.New(db.Inıt())
	tokenHandler := handler.NewTokenHandler()
	token, err := tokenHandler.GetToken(payload.ID.String())
	if err != nil {
		fmt.Println(err.Error())
	}

	if token.TokenStr == secretHeader {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Succesfull"})
	}

	return c.Next()

}

func createToken(c *fiber.Ctx) error {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "TokenID couldn't create"})
	}
	claims := jwt.MapClaims{
		"id":         tokenID,
		"issued_at":  time.Now(),
		"expired_at": time.Now().Add(time.Minute * 1),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	var tokenIDStr = tokenID.String()
	handler := handler.New(db.Inıt())
	tokenHandler := handler.NewTokenHandler()
	tokenHandler.AddToken(tokenIDStr, t)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": t})
}

func verifyToken(token string) (*model.Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &model.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*model.Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil

}
