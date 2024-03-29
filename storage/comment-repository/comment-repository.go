package commentrepository

import (
	"errors"
	"fmt"

	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type CommentRepository interface {
	AddComments(comment *model.Comment) error
	GetComments(issueID string) ([]*model.Comment, error)
}

type commentRepository struct{}

func NewCommentHandler() CommentRepository {
	return &commentRepository{}
}

// CreateComments implements CommentRepository
func (*commentRepository) AddComments(comment *model.Comment) error {
	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	err := db.Create(comment).Error
	if err != nil {
		return err
	}
	db.Save(comment)
	return nil
}

// GetComments implements CommentRepository
func (*commentRepository) GetComments(issueID string) ([]*model.Comment, error) {
	var comment []*model.Comment
	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if issueID == "" {
		fmt.Println("")
	}
	result := db.Where("issue_id = ?", issueID).Find(&comment)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}
	return comment, nil
}
