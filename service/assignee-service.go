package service

import (
	"github.com/webhook-issue-manager/model"
	assigneerepository "github.com/webhook-issue-manager/storage/assignee-repository"
)

var (
	assigneerepo = assigneerepository.NewAssigneeHandler()
)

type AssigneeService interface {
	CreateAssignee(assignee *model.Assignee) (string, error)
	GetAssignee(assigneeId string) (*model.Assignee, error)
}

type assigneeService struct{}

func NewAssigneeService() AssigneeService {
	return &assigneeService{}
}

// CreateAssignee implements AssigneeService
func (*assigneeService) CreateAssignee(assignee *model.Assignee) (string, error) {
	assigneId, err := assigneerepo.AddAssignee(assignee)
	if err != nil {
		return "", err
	}
	return assigneId, nil
}

// GetAssignee implements AssigneeService
func (*assigneeService) GetAssignee(assigneeId string) (*model.Assignee, error) {
	assignee, err := assigneerepo.GetAssignee(assigneeId)
	if err != nil {
		return nil, err
	}
	return assignee, nil
}
