package service

import (
	"github.com/webhook-issue-manager/model"
	assigneerepository "github.com/webhook-issue-manager/storage/assignee-repository"
)

var (
	assigneerepo assigneerepository.AssigneeRepository = assigneerepository.NewAssigneeHandler()
)

type AssigneeService interface {
	CreateAssignee(assignee *model.Assignee) (string, error)
}

type assigneeservice struct{}

func NewAssigneeService() AssigneeService {
	return &assigneeservice{}
}

// CreateAssignee implements AssigneeService
func (*assigneeservice) CreateAssignee(assignee *model.Assignee) (string, error) {
	assign, err := assigneerepo.AddAssignee(assignee)
	if err != nil {
		return "", err
	}
	return assign.Id, nil
}
