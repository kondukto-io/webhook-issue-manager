package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/lib/pq"
)

type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityHigh     Severity = "high"
	SeverityMedium   Severity = "medium"
	SeverityLow      Severity = "low"
)

func (s Severity) String() string {
	return string(s)
}

type CreateIssueRequest struct {
	Status              string      `json:"status"`
	Title               string      `json:"title"`
	Fp                  bool        `json:"fp"`
	Link                string      `json:"link"`
	Name                string      `json:"name"`
	Path                string      `json:"path"`
	Severity            string      `json:"severity"`
	TemplateMD          string      `json:"template_md"`
	ProjectName         string      `json:"project_name"`
	Assignee            Assignee    `json:"assignee"`
	Labels              []string    `json:"labels"`
	VulnerabilityDetail interface{} `json:"vulnerability"`
	DueDate             string      `json:"due_date"`
}

type StatusUpdateRequest struct {
	ID       string   `json:"-"`
	Labels   []string `json:"labels,omitempty"`
	Status   string   `json:"status,omitempty"`
	Severity Severity `json:"severity,omitempty"`
}

type JSONB []interface{}

type IssueLinks struct {
	Self string `json:"self"`
	HTML string `json:"html"`
}

type Issue struct {
	ID                  string         `json:"id" gorm:"primaryKey"`
	Status              string         `json:"status"`
	Title               string         `json:"title"`
	Fp                  bool           `json:"fp"`
	Link                string         `json:"link"`
	Name                string         `json:"name"`
	Path                string         `json:"path"`
	Severity            string         `json:"severity"`
	TemplateMD          string         `json:"template_md"`
	ProjectName         string         `json:"project_name"`
	AssigneeID          string         `json:"assignee_id"`
	Labels              pq.StringArray `json:"labels" gorm:"type:text[]"`
	VulnerabilityDetail JSONB          `json:"vulnerability" gorm:"type:jsonb"`
	DueDate             string         `json:"due_date"`
	Links               *IssueLinks    `json:"links,omitempty" gorm:"-"`
}

type IssueDTO struct {
	ID         string      `json:"id"`
	Status     string      `json:"status"`
	Severity   string      `json:"severity"`
	Title      string      `json:"title"`
	TemplateMD string      `json:"template_md"`
	Assignee   Assignee    `json:"assignee"`
	Labels     []string    `json:"labels"`
	DueDate    string      `json:"due_date"`
	Links      *IssueLinks `json:"links,omitempty"`
}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
