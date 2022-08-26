package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
	issuerepository "github.com/webhook-issue-manager/storage/issue-repository"
)

var (
	issueRepo issuerepository.IssueRepository = issuerepository.NewIssueRepository()
)

type IssueService interface {
	CreateIssue(issueReq *model.IssueReq) (*model.Issue, error)
	GetDetails(issueId string) (*model.IssueDTO, error)
	UpdateStatus(issueId string, status string) error
}

type issueservice struct{}

func NewIssueService() IssueService {
	return &issueservice{}
}

func (*issueservice) CreateIssue(issueReq *model.IssueReq) (*model.Issue, error) {
	assigneeID, _ := uuid.NewRandom()

	assignee := &model.Assignee{Id: assigneeID.String(), Email: issueReq.Assignee.Email, UserName: issueReq.Assignee.UserName}

	assigneeId, err := assigneerepo.AddAssignee(assignee)

	if err != nil {
		return nil, err
	}
	issueId := fmt.Sprintf("%d", time.Now().UnixNano())
	issue := &model.Issue{Id: issueId,
		Status: issueReq.Status, Title: issueReq.Title, Fp: issueReq.Fp, Link: issueReq.Link, Name: issueReq.Name,
		Path: issueReq.Path, Severity: issueReq.Severity, ProjectName: issueReq.ProjectName,
		TemplateMD: issueReq.TemplateMD, AssigneeId: assigneeId, Labels: issueReq.Labels, VulnDetail: model.JSONB{issueReq.VulnDetail}}

	err = issueRepo.AddIssue(issue)
	if err != nil {
		return nil, err
	}
	return issue, err
}

func (*issueservice) GetDetails(issueId string) (*model.IssueDTO, error) {

	issue, err := issueRepo.GetDetails(issueId)
	if err != nil {
		return nil, err
	}

	assignee, err := assigneerepo.GetAssignee(issue.AssigneeId)
	if err != nil {
		return nil, err
	}

	issueDTO := model.IssueDTO{Id: issue.Id, Status: issue.Status, Title: issue.Title,
		TemplateMD: issue.TemplateMD, Assignee: model.Assignee{Email: assignee.Email, UserName: assignee.UserName}, Labels: issue.Labels}

	if err != nil {
		return nil, err
	}
	return &issueDTO, nil
}

func (*issueservice) UpdateStatus(issueId string, status string) error {
	err := issueRepo.UpdateStatus(issueId, status)
	if err != nil {
		return err
	}
	return err
}
