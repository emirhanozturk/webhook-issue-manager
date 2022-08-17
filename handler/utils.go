package handlers

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/webhook-issue-manager/model"
)

var (
	secretKey       = "secretKey"
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired Token")
)

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
