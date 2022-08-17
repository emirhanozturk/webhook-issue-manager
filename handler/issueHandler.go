package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/service"
)

var (
	issueService    service.IssueService    = service.NewIssueService()
	assigneeService service.AssigneeService = service.NewAssigneeService()
)

type IssueHandler interface {
	CreateIssue(c *fiber.Ctx) error
	GetDetails(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type issuehandler struct{}

func NewIssueHandler() IssueHandler {
	return &issuehandler{}
}

func (*issuehandler) CreateIssue(c *fiber.Ctx) error {
	var issueReq *model.IssueReq
	err := json.Unmarshal(c.Body(), &issueReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	assigneeID, _ := uuid.NewRandom()

	assignee := model.Assignee{Id: assigneeID.String(), Email: issueReq.Assignee.Email, UserName: issueReq.Assignee.UserName}

	assigneeId, err := assigneeService.CreateAssignee(&assignee)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	issueId := fmt.Sprintf("%d", time.Now().UnixNano())
	issue := model.Issue{Id: issueId,
		Status: issueReq.Status, Title: issueReq.Title, Fp: issueReq.Fp, Link: issueReq.Link, Name: issueReq.Name,
		Path: issueReq.Path, Severity: issueReq.Severity, ProjectName: issueReq.ProjectName,
		TemplateMD: issueReq.TemplateMD, AssigneeId: assigneeId, Labels: issueReq.Labels}

	err = issueService.CreateIssue(&issue)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Issue created"})
}

func (*issuehandler) GetDetails(c *fiber.Ctx) error {
	issueId := c.Params("id")
	issue, err := issueService.GetDetails(issueId)
	if err != nil {
		return err
	}

	issueBytes, err := json.Marshal(issue)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Issue": string(issueBytes)})
}

func (*issuehandler) Update(c *fiber.Ctx) error {
	var issue *model.Issue
	issueId := c.Params("id")

	err := json.Unmarshal(c.Body(), issue)
	if err != nil {
		return err
	}

	err = issueService.UpdateStatus(issueId, issue.Status)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Issue updated"})
}
