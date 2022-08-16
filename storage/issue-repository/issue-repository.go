package issuerepository

import (
	"errors"
	"fmt"

	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type IssueRepository interface {
	AddIssue(issue *model.Issue) error
	GetDetails(issueId string) (*model.IssueDTO, error)
	UpdateStatus(issueId string, status string) error
}

type issuerepository struct{}

func NewIssueRepository() IssueRepository {
	return &issuerepository{}
}

func (*issuerepository) AddIssue(issue *model.Issue) error {
	db := postgres.Inıt()
	if err := db.Model(&model.Issue{}).Create(issue); err.Error != nil {
		return err.Error
	}

	db.Save(&issue)
	return nil
}

func (*issuerepository) GetDetails(issueId string) (*model.IssueDTO, error) {
	var issueDTO model.IssueDTO
	db := postgres.Inıt()
	if issueId == "" {
		fmt.Println("TokenID can not be empty")
	}
	result := db.Where("id = ?", issueId).Find(&issueDTO)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}
	return &issueDTO, nil
}

func (*issuerepository) UpdateStatus(issueId string, status string) error {
	var issue *model.Issue
	db := postgres.Inıt()
	db.Model(&issue).Where("id = ?", issueId).Update("status", status)
	return nil
}
