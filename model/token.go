package model

type Token struct {
	TokenID  string `gorm:"primaryKey"`
	TokenStr string
}

