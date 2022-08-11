package commentrepository

import (
	"errors"
	"fmt"

	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type CommentRepository interface {
	GetComments(issueId string) (*model.Comment, error)
}

type commentrepository struct{}

func NewCommentHandler() CommentRepository {
	return &commentrepository{}
}

// GetComments implements CommentRepository
func (*commentrepository) GetComments(issueId string) (*model.Comment, error) {
	var comment model.Comment
	db := postgres.Inıt()
	if issueId == "" {
		fmt.Println("TokenID can not be empty")
	}
	result := db.Where("id = ?", issueId).Find(&comment)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}
	return &comment, nil
}
