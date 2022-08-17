package model

import "github.com/lib/pq"

type IssueReq struct {
	Id          string         `json:"id" gorm:"primaryKey"`
	Status      string         `json:"status"`
	Title       string         `json:"title"`
	Fp          bool           `json:"fp"`
	Link        string         `json:"link"`
	Name        string         `json:"name"`
	Path        string         `json:"path"`
	Severity    string         `json:"severity"`
	TemplateMD  string         `json:"template_md"`
	ProjectName string         `json:"project_name"`
	Assignee    Assignee       `json:"assignee"`
	Labels      pq.StringArray `json:"labels" gorm:"type:text[]"`
	//VulnDetail  interface{}    `json:"vulnerability"`
}

type Issue struct {
	Id          string         `json:"id" gorm:"primaryKey"`
	Status      string         `json:"status"`
	Title       string         `json:"title"`
	Fp          bool           `json:"fp"`
	Link        string         `json:"link"`
	Name        string         `json:"name"`
	Path        string         `json:"path"`
	Severity    string         `json:"severity"`
	TemplateMD  string         `json:"template_md"`
	ProjectName string         `json:"project_name"`
	AssigneeId  string         `json:"assignee_id"`
	Labels      pq.StringArray `json:"labels" gorm:"type:text[]"`
	//VulnDetail  interface{}    `json:"vulnerability"`
}

type IssueDTO struct {
	Id         string   `json:"id" gorm:"primaryKey"`
	Status     string   `json:"status"`
	Title      string   `json:"title"`
	TemplateMD string   `json:"template_md"`
	Assignee   Assignee `json:"assignee"  gorm:"foreignKey:AssigneeID;references:Id"`
	Labels     []string `json:"labels"`
}
