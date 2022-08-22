package model

import "time"

type CommentReq struct {
	Id       string   `json:"id" gorm:"primaryKey"`
	IssueId  string   `json:"issue_id"`
	Body     string   `json:"body"`
	Assignee Assignee `json:"author"`
}
type Comment struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	IssueId    string    `json:"issue_id"`
	CreatedAt  time.Time `json:"created_at"`
	Body       string    `json:"body"`
	AssigneeId string    `json:"assignee_id"`
}

type CommentDTOArray struct {
	CommentDtos []CommentDto `json:"comments"`
}

type CommentDto struct {
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	Assignee  Assignee  `json:"author"`
}
