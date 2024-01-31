package issuerepository

import (
	"errors"

	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type IssueRepository interface {
	AddIssue(issue *model.Issue) error
	GetDetails(issueID string) (*model.Issue, error)
	UpdateStatus(issueID string, status string) error
}

type issueRepository struct{}

func NewIssueRepository() IssueRepository {
	return &issueRepository{}
}

func (*issueRepository) AddIssue(issue *model.Issue) error {
	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if err := db.Create(issue); err.Error != nil {
		return err.Error
	}

	db.Save(&issue)
	return nil
}

func (*issueRepository) GetDetails(issueID string) (*model.Issue, error) {
	if issueID == "" {
		return nil, errors.New("issue id can not be empty")
	}

	var issue model.Issue
	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	result := db.Where("id = ?", issueID).Find(&issue)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}
	return &issue, nil
}

func (*issueRepository) UpdateStatus(issueID string, status string) error {
	if issueID == "" {
		return errors.New("issue id can not be empty")
	}

	var issue *model.Issue
	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	db.Model(&issue).Where("id = ?", issueID).Update("status", status)
	return nil
}
