package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
	issuerepository "github.com/webhook-issue-manager/storage/issue-repository"
)

var (
	issueRepo = issuerepository.NewIssueRepository()
)

type IssueService interface {
	CreateIssue(*model.CreateIssueRequest) (*model.Issue, error)
	GetDetails(string) (*model.IssueDTO, error)
	UpdateIssue(request *model.UpdateIssueRequest) error
}

type issueService struct{}

func NewIssueService() IssueService {
	return &issueService{}
}

func (*issueService) CreateIssue(request *model.CreateIssueRequest) (*model.Issue, error) {
	newAssigneeID, _ := uuid.NewRandom()

	var assignee = &model.Assignee{
		Id:       newAssigneeID.String(),
		Email:    request.Assignee.Email,
		UserName: request.Assignee.UserName,
	}

	assigneeID, err := assigneerepo.AddAssignee(assignee)
	if err != nil {
		return nil, err
	}

	var issueID = fmt.Sprintf("%d", time.Now().UnixNano())
	var issue = &model.Issue{
		ID:                  issueID,
		Status:              request.Status,
		Title:               request.Title,
		Fp:                  request.Fp,
		Link:                request.Link,
		Name:                request.Name,
		Path:                request.Path,
		Severity:            request.Severity,
		ProjectName:         request.ProjectName,
		TemplateMD:          request.TemplateMD,
		AssigneeID:          assigneeID,
		Labels:              request.Labels,
		VulnerabilityDetail: model.JSONB{request.VulnerabilityDetail},
		DueDate:             request.DueDate,
	}

	if err = issueRepo.AddIssue(issue); err != nil {
		return nil, err
	}

	issue.Links = &model.IssueLinks{HTML: fmt.Sprintf("http://localhost:8080/projects/%s/issues/%s", issue.ProjectName, issue.ID)}

	return issue, err
}

func (*issueService) GetDetails(issueID string) (*model.IssueDTO, error) {
	issue, err := issueRepo.GetDetails(issueID)
	if err != nil {
		return nil, err
	}

	assignee, err := assigneerepo.GetAssignee(issue.AssigneeID)
	if err != nil {
		return nil, err
	}

	issueDTO := model.IssueDTO{
		ID:         issue.ID,
		Status:     issue.Status,
		Severity:   issue.Severity,
		Title:      issue.Title,
		TemplateMD: issue.TemplateMD,
		Assignee:   model.Assignee{Email: assignee.Email, UserName: assignee.UserName},
		Labels:     issue.Labels,
		DueDate:    issue.DueDate,
	}

	issueDTO.Links = &model.IssueLinks{HTML: fmt.Sprintf("http://localhost:8080/projects/%s/issues/%s", issue.ProjectName, issue.ID)}

	return &issueDTO, nil
}

func (*issueService) UpdateIssue(request *model.UpdateIssueRequest) error {
	if request.Status != "" {
		if err := issueRepo.UpdateStatus(request.ID, request.Status); err != nil {
			return fmt.Errorf("failed to update issue status: %w", err)
		}
	}

	if request.Severity != "" {
		if err := issueRepo.UpdateSeverity(request.ID, request.Severity.String()); err != nil {
			return fmt.Errorf("failed to update issue severity: %w", err)
		}
	}

	if request.Body != "" {
		if err := issueRepo.UpdateBody(request.ID, request.Body); err != nil {
			return fmt.Errorf("failed to update issue body: %w", err)
		}
	}

	return nil
}
