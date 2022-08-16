package model

type Assignee struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}
