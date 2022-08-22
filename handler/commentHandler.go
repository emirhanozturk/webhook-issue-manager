package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/service"
)

var (
	commentservice service.CommentService = service.NewCommentService()
)

type CommentsHandler interface {
	CreateComment(c *fiber.Ctx) error
	GetComments(c *fiber.Ctx) error
}

type commentshandler struct{}

func NewCommentHandler() CommentsHandler {
	return &commentshandler{}
}

// CreateComment implements CommentsHandler
func (*commentshandler) CreateComment(c *fiber.Ctx) error {
	var commentReq *model.CommentReq
	issueId := c.Params("id")
	err := json.Unmarshal(c.Body(), &commentReq)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{"message": "cannot unmarshal body"})
	}
	id, _ := uuid.NewRandom()
	assignee := &model.Assignee{Id: id.String(), Email: commentReq.Assignee.Email, UserName: commentReq.Assignee.UserName}
	assigneeId, err := assigneeService.CreateAssignee(assignee)
	if err != nil {
		return err
	}
	commentId, _ := uuid.NewRandom()
	comment := &model.Comment{Id: commentId.String(), IssueId: issueId, CreatedAt: time.Now(), Body: commentReq.Body, AssigneeId: assigneeId}

	commentservice.CreateComment(comment)
	return nil
}

// GetComments implements CommentsHandler
func (*commentshandler) GetComments(c *fiber.Ctx) error {
	issueId := c.Params("id")
	comments, err := commentservice.GetComment(issueId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Cannot get comment"})
	}
	return c.Status(http.StatusOK).JSON(comments)

}
