package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
	commentrepository "github.com/webhook-issue-manager/storage/comment-repository"
)

var (
	commentrepo commentrepository.CommentRepository = commentrepository.NewCommentHandler()
)

type CommentService interface {
	CreateComment(commentReq *model.CommentReq) error
	GetComment(issueId string) (*model.CommentDTOArray, error)
}

type commentservice struct{}

func NewCommentService() CommentService {
	return &commentservice{}
}

// CreateComment implements CommentService
func (*commentservice) CreateComment(commentReq *model.CommentReq) error {
	id, _ := uuid.NewRandom()
	assignee := &model.Assignee{Id: id.String(), Email: commentReq.Assignee.Email, UserName: commentReq.Assignee.UserName}
	assigneeId, err := assigneerepo.AddAssignee(assignee)
	if err != nil {
		return err
	}
	commentId, _ := uuid.NewRandom()
	comment := &model.Comment{Id: commentId.String(), IssueId: commentReq.IssueId, CreatedAt: time.Now(), Body: commentReq.Body, AssigneeId: assigneeId}

	err = commentrepo.AddComments(comment)
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
