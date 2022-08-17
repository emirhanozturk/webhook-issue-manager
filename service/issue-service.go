package service

import (
	"github.com/webhook-issue-manager/model"
	issuerepository "github.com/webhook-issue-manager/storage/issue-repository"
)

var (
	issueRepo issuerepository.IssueRepository = issuerepository.NewIssueRepository()
)

type IssueService interface {
	CreateIssue(issue *model.Issue) error
	GetDetails(issueId string) (*model.Issue, error)
	UpdateStatus(issueId string, status string) error
}

type issueservice struct{}

func NewIssueService() IssueService {
	return &issueservice{}
}

func (*issueservice) CreateIssue(Issue *model.Issue) error {
	err := issueRepo.AddIssue(Issue)
	if err != nil {
		return err
	}
	return nil
}

func (*issueservice) GetDetails(issueId string) (*model.Issue, error) {
	issue, err := issueRepo.GetDetails(issueId)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

func (*issueservice) UpdateStatus(issueId string, status string) error {
	err := issueRepo.UpdateStatus(issueId, status)
	if err != nil {
		return err
	}
	return err
}
