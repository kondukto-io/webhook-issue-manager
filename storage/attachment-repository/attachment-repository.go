package attachmentrepository

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/webhook-issue-manager/config"
	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type AttachmentRepository interface {
	AddAttachment(attachmentReq *model.AttachmentReq) error
}

type attachmentRepository struct{}

func NewAttachmentRepository() AttachmentRepository {
	return &attachmentRepository{}
}

// AddAttachment implements AttachmentRepository
func (*attachmentRepository) AddAttachment(attachmentReq *model.AttachmentReq) error {
	var ctx = context.Background()
	minioClient, err := config.MinioConnection()
	if err != nil {
		return fmt.Errorf("failed to connect Minio: %w", err)
	}

	var base64Text = make([]byte, base64.StdEncoding.DecodedLen(len(attachmentReq.Base64Content)))
	base64.StdEncoding.Decode(base64Text, []byte(attachmentReq.Base64Content))

	var fileP = "/opt/.kondukto/screenshots"
	var objectID = attachmentReq.UUID
	var filePath = fmt.Sprintf("%s/%s", fileP, objectID)
	if err = os.WriteFile(filePath, base64Text, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// Make a new bucket
	var bucketName = "attachments"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	var contentType = "image/jpeg"

	_, err = minioClient.FPutObject(ctx, bucketName, objectID, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	attachmentID, _ := uuid.NewRandom()

	var attachment = model.Attachment{
		ID:       attachmentID.String(),
		IssueID:  attachmentReq.IssueID,
		Title:    attachmentReq.Title,
		FilePath: filePath,
	}

	var db = postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	db.Create(&attachment)

	return nil
}
