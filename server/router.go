package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handler "github.com/webhook-issue-manager/handler"
)

var (
	tokenHandler   = handler.NewTokenHandler()
	testHandler    = handler.NewTestConnectionHandlers()
	issueHandler   = handler.NewIssueHandler()
	commentHandler = handler.NewCommentHandler()
)

func Router() *fiber.App {
	var app = fiber.New()
	app.Use(logger.New())
	app.Post("/tokens", tokenHandler.CreateToken)

	var v1 = app.Group("api/v1")
	{
		testGroup := v1.Group("test")
		testGroup.Use(tokenHandler.TokenValidatorMiddleware)
		testGroup.Get("/", testHandler.TestIssue)

		issueGroup := v1.Group("issues")
		issueGroup.Use(tokenHandler.TokenValidatorMiddleware)

		issueGroup.Post("", issueHandler.CreateIssue)
		issueGroup.Get("/:id", issueHandler.GetDetails)
		issueGroup.Patch("/:id", issueHandler.UpdateStatus)
		issueGroup.Post("/:id/attachments", issueHandler.AddAttachment)
		issueGroup.Get("/:id/attachments", issueHandler.ListAttachments)

		commentGroup := issueGroup.Group("/:id/comments")
		commentGroup.Post("", commentHandler.CreateComment)
		commentGroup.Get("", commentHandler.GetComments)
	}

	app.Listen(":3000")

	return app
}
