package model

type Issue struct {
	Id         string   `json:"id"`
	Status     string   `json:"status"`
	Title      string   `json:"title"`
	TemplateMD string   `json:"templatemd"`
	Assignee   Assignee `json:"assignee"`
	Labels     []string `json:"labels"`
}
