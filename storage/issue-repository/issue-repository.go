package issuerepository

import (
	"errors"
	"fmt"

	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type IssueRepository interface {
	AddIssue(issue *model.Issue) error
	GetDetails(issueId string) (*model.Issue, error)
	UpdateStatus(issueId string, status string) error
}

type issuerepository struct{}

func NewIssueRepository() IssueRepository {
	return &issuerepository{}
}

func (*issuerepository) AddIssue(issue *model.Issue) error {
	db := postgres.Inıt()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if err := db.Create(issue); err.Error != nil {
		return err.Error
	}

	db.Save(&issue)
	return nil
}

func (*issuerepository) GetDetails(issueId string) (*model.Issue, error) {
	var issue model.Issue
	db := postgres.Inıt()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if issueId == "" {
		fmt.Println("TokenID can not be empty")
	}
	result := db.Where("id = ?", issueId).Find(&issue)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}
	return &issue, nil
}

func (*issuerepository) UpdateStatus(issueId string, status string) error {
	var issue *model.Issue
	db := postgres.Inıt()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	db.Model(&issue).Where("id = ?", issueId).Update("status", status)
	return nil
}
