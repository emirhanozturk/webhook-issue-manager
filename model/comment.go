package model

import "time"

type Comment struct {
	Comments []comment `json:"comments"`
}

type comment struct {
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	Author    author    `json:"author"`
}

type author struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
