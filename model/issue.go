package model

type Issue struct {
	Id          string      `json:"id" gorm:"primaryKey"`
	Status      string      `json:"status"`
	Title       string      `json:"title"`
	Fp          bool        `json:"fp"`
	Link        string      `json:"link"`
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Severity    string      `json:"severity"`
	TemplateMD  string      `json:"template_md"`
	ProjectName string      `json:"project_name"`
	AssigneeID  int         `json:"assignee_id"`
	Assignee    Assignee    `json:"assignee" gorm:"foreignKey:AssigneeID;references:Id"`
	Labels      []string    `json:"labels"`
	VulnDetail  interface{} `json:"vulnerability"`
}

type IssueDTO struct {
	Id         string   `json:"id" gorm:"primaryKey"`
	Status     string   `json:"status"`
	Title      string   `json:"title"`
	TemplateMD string   `json:"template_md"`
	AssigneeID int      `json:"assignee_id"`
	Assignee   Assignee `json:"assignee"  gorm:"foreignKey:AssigneeID;references:Id"`
	Labels     []string `json:"labels"`
}
