package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
	issuerepository "github.com/webhook-issue-manager/storage/issue-repository"
)

var (
	issueRepo issuerepository.IssueRepository = issuerepository.NewIssueRepository()
)

type IssueService interface {
	CreateIssue(issueReq *model.IssueReq) (*model.Issue, error)
	GetDetails(issueID string) (*model.IssueDTO, error)
	UpdateStatus(issueID string, status string) error
}

type issueservice struct{}

func NewIssueService() IssueService {
	return &issueservice{}
}

func (*issueservice) CreateIssue(issueReq *model.IssueReq) (*model.Issue, error) {
	newAssigneID, _ := uuid.NewRandom()

	var assignee = &model.Assignee{
		Id:       newAssigneID.String(),
		Email:    issueReq.Assignee.Email,
		UserName: issueReq.Assignee.UserName,
	}

	assigneeID, err := assigneerepo.AddAssignee(assignee)
	if err != nil {
		return nil, err
	}

	var issueID = fmt.Sprintf("%d", time.Now().UnixNano())
	var issue = &model.Issue{
		ID:          issueID,
		Status:      issueReq.Status,
		Title:       issueReq.Title,
		Fp:          issueReq.Fp,
		Link:        issueReq.Link,
		Name:        issueReq.Name,
		Path:        issueReq.Path,
		Severity:    issueReq.Severity,
		ProjectName: issueReq.ProjectName,
		TemplateMD:  issueReq.TemplateMD,
		AssigneeID:  assigneeID,
		Labels:      issueReq.Labels,
		VulnDetail:  model.JSONB{issueReq.VulnDetail},
	}

	if err = issueRepo.AddIssue(issue); err != nil {
		return nil, err
	}

	return issue, err
}

func (*issueservice) GetDetails(issueID string) (*model.IssueDTO, error) {
	issue, err := issueRepo.GetDetails(issueID)
	if err != nil {
		return nil, err
	}

	assignee, err := assigneerepo.GetAssignee(issue.AssigneeID)
	if err != nil {
		return nil, err
	}

	var issueDTO = model.IssueDTO{
		ID:         issue.ID,
		Status:     issue.Status,
		Title:      issue.Title,
		TemplateMD: issue.TemplateMD,
		Assignee:   model.Assignee{Email: assignee.Email, UserName: assignee.UserName},
		Labels:     issue.Labels,
	}

	return &issueDTO, nil
}

func (*issueservice) UpdateStatus(issueID string, status string) error {
	if err := issueRepo.UpdateStatus(issueID, status); err != nil {
		return fmt.Errorf("failed to update issue status")
	}

	return nil
}
