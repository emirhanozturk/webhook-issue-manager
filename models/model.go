package models

type Issue struct {
	Id         string   `json:"id"`
	Status     string   `json:"status"`
	Title      string   `json:"title"`
	TemplateMD string   `json:"templatemd"`
	User       User     `json:"user"`
	Labels     []string `json:"labels"`
}

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type Token struct {
	TokenID  string `gorm:"primaryKey"`
	TokenStr string
}
