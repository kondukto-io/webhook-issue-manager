package attachmentrepository

import "github.com/webhook-issue-manager/model"

type AttachmentRepository interface {
	AddAttachment(attachmentReq *model.AttachmentReq) error
	ListAttachments(issueID string) ([]model.Attachment, error)
}

func NewAttachmentRepository() AttachmentRepository {
	return &minioRepository{}
}
