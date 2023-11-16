package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
	var commentReqArray *model.CommentReqArray
	var issueID = c.Params("id")

	if err := json.Unmarshal(c.Body(), &commentReqArray); err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{"message": "cannot unmarshal body"})
	}

	for _, commentReq := range commentReqArray.CommentReq {
		commentReq.IssueId = issueID
		commentservice.CreateComment(&commentReq)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Comment created"})
}

// GetComments implements CommentsHandler
func (*commentshandler) GetComments(c *fiber.Ctx) error {
	issueID := c.Params("id")
	comments, err := commentservice.GetComment(issueID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Cannot get comment"})
	}
	return c.Status(http.StatusOK).JSON(comments)

}
