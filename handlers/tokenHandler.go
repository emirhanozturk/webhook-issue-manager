package handlers

import (
	"errors"
	"fmt"
	"log"

	model "github.com/webhook-issue-manager/models"
	"gorm.io/gorm"
)

type TokenHandlers struct {
	DB *gorm.DB
}

func (h *Handler) NewTokenHandler() *TokenHandlers {
	return &TokenHandlers{DB: h.DB}
}

func (t TokenHandlers) AddToken(tokenID string, tokenStr string) {
	if tokenID == "" {
		log.Fatal("TokenID can not be empty")
	}
	token := model.Token{TokenID: tokenID, TokenStr: tokenStr}
	result := t.DB.Create(&token)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println("Added the database")
}

func (t TokenHandlers) GetToken(tokenID string) (*model.Token, error) {
	var token model.Token
	if tokenID == "" {
		fmt.Println("TokenID can not be empty")
	}
	result := t.DB.Where("token_id = ?", tokenID).Find(&token)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}

	return &token, nil
}
