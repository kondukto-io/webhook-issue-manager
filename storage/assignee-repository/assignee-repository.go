package assigneerepository

import (
	"errors"

	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type AssigneeRepository interface {
	AddAssignee(assignee *model.Assignee) (string, error)
	GetAssignee(assigneeId string) (*model.Assignee, error)
}

type assigneeRepository struct{}

func NewAssigneeHandler() AssigneeRepository {
	return &assigneeRepository{}
}

// AddAssignee implements AssigneeRepository
func (*assigneeRepository) AddAssignee(assignee *model.Assignee) (string, error) {
	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	db.Create(assignee)
	db.Save(assignee)
	return assignee.Id, nil
}

// GetAssignee implements AssigneeRepository
func (*assigneeRepository) GetAssignee(assigneeId string) (*model.Assignee, error) {
	var assignee *model.Assignee
	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	result := db.Where("id = ?", assigneeId).Find(&assignee)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}
	return assignee, nil
}
