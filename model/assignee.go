package model

type Assignee struct {
	Id       string `json:",omitempty" gorm:"primaryKey"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}
