package service

import (
	"github.com/webhook-issue-manager/model"
	attachmentrepository "github.com/webhook-issue-manager/storage/attachment-repository"
)

var (
	attachmentRepo = attachmentrepository.NewAttachmentRepository()
)

type AttachmentService interface {
	CreateAttachment(attachmentReq *model.AttachmentReq) error
	ListAttachments(issueID string) ([]model.Attachment, error)
}

type attachmentService struct{}

func NewAttachmentService() AttachmentService {
	return &attachmentService{}
}

func (*attachmentService) CreateAttachment(attachmentReq *model.AttachmentReq) error {
	err := attachmentRepo.AddAttachment(attachmentReq)
	if err != nil {
		return err
	}
	return nil
}

func (*attachmentService) ListAttachments(issueID string) ([]model.Attachment, error) {
	attachments, err := attachmentRepo.ListAttachments(issueID)
	if err != nil {
		return nil, err
	}
	return attachments, nil
}
