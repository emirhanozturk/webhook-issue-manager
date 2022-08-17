package model

type Assignee struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}
