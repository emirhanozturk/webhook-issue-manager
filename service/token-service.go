package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
	tokenrepository "github.com/webhook-issue-manager/storage/token-repository"
)

var (
	secretKey = "secretKey"
)

type TokenService interface {
	CreateToken() (*model.Token, error)
	GetToken(tokenId string) (*model.Token, error)
}

type tokenservice struct{}

var (
	repo tokenrepository.TokenRepository = tokenrepository.NewTokenRepository()
)

func NewTokenService() TokenService {
	return &tokenservice{}
}

func (*tokenservice) CreateToken() (*model.Token, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	claims := jwt.MapClaims{
		"id":         tokenID,
		"issued_at":  time.Now(),
		"expired_at": time.Now().Add(time.Hour * 24 * 7),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	token := model.Token{TokenID: tokenID.String(), TokenStr: tokenStr}

	addedToken, err := repo.AddToken(&token)
	if err != nil {
		return nil, err
	}

	return addedToken, nil
}

func (*tokenservice) GetToken(tokenId string) (*model.Token, error) {
	token, err := repo.GetToken(tokenId)
	if err != nil {
		return nil, err
	}
	return token, nil
}
