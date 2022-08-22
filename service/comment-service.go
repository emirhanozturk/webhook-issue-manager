package service

import (
	"github.com/webhook-issue-manager/model"
	commentrepository "github.com/webhook-issue-manager/storage/comment-repository"
)

var (
	commentrepo commentrepository.CommentRepository = commentrepository.NewCommentHandler()
)

type CommentService interface {
	CreateComment(comment *model.Comment) error
	GetComment(issueId string) (*model.CommentDTOArray, error)
}

type commentservice struct{}

func NewCommentService() CommentService {
	return &commentservice{}
}

// CreateComment implements CommentService
func (*commentservice) CreateComment(comment *model.Comment) error {
	err := commentrepo.AddComments(comment)
	if err != nil {
		return err
	}
	return nil
}

// GetComment implements CommentService
func (*commentservice) GetComment(issueId string) (*model.CommentDTOArray, error) {
	comments, err := commentrepo.GetComments(issueId)
	if err != nil {
		return nil, err
	}
	var commentDtoArray model.CommentDTOArray
	for _, comment := range comments {
		assignee, err := assigneerepo.GetAssignee(issueId)
		if err != nil {
			return nil, err
		}
		commentDtoArray.CommentDtos = append(commentDtoArray.CommentDtos, model.CommentDto{
			CreatedAt: comment.CreatedAt,
			Body:      comment.Body,
			Assignee:  model.Assignee{Email: assignee.Email, UserName: assignee.UserName},
		})
	}

	return &commentDtoArray, nil
}
