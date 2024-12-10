package attachmentrepository

import (
	"encoding/base64"
	"fmt"
	"github.com/webhook-issue-manager/storage/postgres"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
)

type localRepository struct{}

func (*localRepository) AddAttachment(attachmentReq *model.AttachmentReq) error {
	decodedContent, err := base64.StdEncoding.DecodeString(attachmentReq.Base64Content)
	if err != nil {
		return fmt.Errorf("failed to decode base64 content: %w", err)
	}

	fileDir := "/opt/.kondukto/screenshots"
	if err := os.MkdirAll(fileDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := filepath.Join(fileDir, attachmentReq.Title)
	if err := os.WriteFile(filePath, decodedContent, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	attachmentID, _ := uuid.NewRandom()

	attachment := model.Attachment{
		ID:       attachmentID.String(),
		IssueID:  attachmentReq.IssueID,
		Title:    attachmentReq.Title,
		FilePath: filePath,
	}

	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if err := db.Create(&attachment).Error; err != nil {
		return fmt.Errorf("failed to save attachment to database: %w", err)
	}

	return nil
}

func (*localRepository) ListAttachments(issueID string) ([]model.Attachment, error) {
	var attachments []model.Attachment

	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if err := db.Where("issue_id = ?", issueID).Find(&attachments).Error; err != nil {
		return nil, fmt.Errorf("failed to list attachments: %w", err)
	}

	return attachments, nil
}
