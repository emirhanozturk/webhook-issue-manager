package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handler "github.com/webhook-issue-manager/handler"
)

var (
	tokenHandler   handler.TokenHandler    = handler.NewTokenHandler()
	testHandler    handler.TestHandler     = handler.NewTestConnectionHandlers()
	issueHandler   handler.IssueHandler    = handler.NewIssueHandler()
	commentHandler handler.CommentsHandler = handler.NewCommentHandler()
)

func Router() *fiber.App {

	app := fiber.New()
	app.Use(logger.New())
	app.Post("/tokens", tokenHandler.CreateToken)
	v1 := app.Group("api/v1")
	{
		testGroup := v1.Group("test")
		testGroup.Use(tokenHandler.TokenValidatorMiddleware)
		testGroup.Get("/", testHandler.TestIssue)

		issueGroup := v1.Group("issues")
		issueGroup.Use(tokenHandler.TokenValidatorMiddleware)

		issueGroup.Post("", issueHandler.CreateIssue)
		issueGroup.Get("/:id", issueHandler.GetDetails)
		issueGroup.Patch("/:id", issueHandler.Update)
		issueGroup.Post("/:id/attachments", issueHandler.AddAttachment)

		commentGroup := issueGroup.Group("/:id/comments")
		commentGroup.Post("", commentHandler.CreateComment)
		commentGroup.Get("", commentHandler.GetComments)

	}

	app.Listen(":3000")

	return app
}
