package assigneerepository

import (
	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type AssigneeRepository interface {
	AddAssignee(assignee *model.Assignee) (*model.Assignee, error)
	GetAssignee(assigneeId int) (*model.Assignee, error)
}

type assignerepository struct{}

func NewAssigneeHandler() AssigneeRepository {
	return &assignerepository{}
}

// AddAssignee implements AssigneeRepository
func (*assignerepository) AddAssignee(assignee *model.Assignee) (*model.Assignee, error) {
	db := postgres.Inıt()
	db.Create(assignee)
	db.Save(assignee)
	return assignee, nil
}

// GetAssignee implements AssigneeRepository
func (*assignerepository) GetAssignee(assigneeId int) (*model.Assignee, error) {
	panic("unimplemented")
}
