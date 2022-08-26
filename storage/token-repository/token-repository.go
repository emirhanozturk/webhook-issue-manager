package tokenrepository

import (
	"errors"
	"fmt"

	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type TokenRepository interface {
	AddToken(token *model.Token) (*model.Token, error)
	GetToken(tokenId string) (*model.Token, error)
}

type repo struct{}

func NewTokenRepository() TokenRepository {
	return &repo{}
}

// AddToken implements TokenRepository
func (*repo) AddToken(token *model.Token) (*model.Token, error) {
	db := postgres.Inıt()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	err := db.Create(token).Error
	if err != nil {
		return nil, err
	}
	return token, nil
}

// GetToken implements TokenRepository
func (*repo) GetToken(tokenId string) (*model.Token, error) {
	var token model.Token
	db := postgres.Inıt()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if tokenId == "" {
		fmt.Println("TokenID can not be empty")
	}
	result := db.Where("token_id = ?", tokenId).Find(&token)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}
	return &token, nil
}
