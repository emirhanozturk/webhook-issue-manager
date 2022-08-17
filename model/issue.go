package model

import (
	"encoding/json"
	"errors"

	"database/sql/driver"

	"github.com/lib/pq"
)

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
	VulnDetail  interface{}    `json:"vulnerability"`
}

type JSONB []interface{}

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
	VulnDetail  JSONB          `json:"vulnerability" gorm:"type:jsonb"`
}

type IssueDTO struct {
	Id         string   `json:"id" gorm:"primaryKey"`
	Status     string   `json:"status"`
	Title      string   `json:"title"`
	TemplateMD string   `json:"template_md"`
	Assignee   Assignee `json:"assignee"  gorm:"foreignKey:AssigneeID;references:Id"`
	Labels     []string `json:"labels"`
}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
