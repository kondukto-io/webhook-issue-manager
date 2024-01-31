package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/service"
)

var (
	issueService      = service.NewIssueService()
	attachmentService = service.NewAttachmentService()
)

type IssueHandler interface {
	CreateIssue(c *fiber.Ctx) error
	GetDetails(c *fiber.Ctx) error
	UpdateStatus(c *fiber.Ctx) error
	AddAttachment(c *fiber.Ctx) error
}

type issueHandler struct{}

func NewIssueHandler() IssueHandler {
	return &issueHandler{}
}

func (*issueHandler) CreateIssue(c *fiber.Ctx) error {
	var issueReq *model.CreateIssueRequest
	if err := json.Unmarshal(c.Body(), &issueReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	issue, err := issueService.CreateIssue(issueReq)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(issue)
}

func (*issueHandler) GetDetails(c *fiber.Ctx) error {
	var issueID = c.Params("id")
	issueDTO, err := issueService.GetDetails(issueID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(issueDTO)
}

func (*issueHandler) UpdateStatus(c *fiber.Ctx) error {
	var request *model.StatusUpdateRequest
	var issueID = c.Params("id")

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return err
	}

	request.ID = issueID
	if err := issueService.UpdateStatus(request); err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"id": issueID, "status": request.Status})
}

// AddAttachment implements AttachmentHandler
func (*issueHandler) AddAttachment(c *fiber.Ctx) error {
	var attachmentReqArray *model.AttachmentReqArray
	var issueID = c.Params("id")

	if err := json.Unmarshal(c.Body(), &attachmentReqArray); err != nil {
		return err
	}
	for _, attachmentReq := range attachmentReqArray.AttachmentReq {
		attachmentReq.IssueID = issueID

		if err := attachmentService.CreateAttachment(&attachmentReq); err != nil {
			return err
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "attachments added successfully"})
}
