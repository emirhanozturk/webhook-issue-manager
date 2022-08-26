package commentrepository

import (
	"errors"
	"fmt"

	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type CommentRepository interface {
	AddComments(comment *model.Comment) error
	GetComments(issueId string) ([]*model.Comment, error)
}

type commentrepository struct{}

func NewCommentHandler() CommentRepository {
	return &commentrepository{}
}

// CreateComments implements CommentRepository
func (*commentrepository) AddComments(comment *model.Comment) error {
	db := postgres.Inıt()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	err := db.Create(comment).Error
	if err != nil {
		return err
	}
	db.Save(comment)
	return nil
}

// GetComments implements CommentRepository
func (*commentrepository) GetComments(issueId string) ([]*model.Comment, error) {
	var comment []*model.Comment
	db := postgres.Inıt()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if issueId == "" {
		fmt.Println("")
	}
	result := db.Where("issue_id = ?", issueId).Find(&comment)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}
	return comment, nil
}
