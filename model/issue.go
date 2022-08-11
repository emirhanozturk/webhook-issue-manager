package model

type Issue struct {
	Id         string      `json:"id" gorm:"primaryKey"`
	Status     string      `json:"status"`
	Title      string      `json:"title"`
	TemplateMD string      `json:"template_md"`
	Assignee   Assignee    `json:"assignee"`
	Labels     []string    `json:"labels"`
	VulnDetail interface{} `json:"vulnerability"`
}
