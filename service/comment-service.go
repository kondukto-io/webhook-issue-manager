package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
	commentrepository "github.com/webhook-issue-manager/storage/comment-repository"
)

var (
	commentRepo = commentrepository.NewCommentHandler()
)

type CommentService interface {
	CreateComment(commentReq *model.CommentReq) error
	GetComment(issueID string) (*model.CommentDTOArray, error)
}

type commentService struct{}

func NewCommentService() CommentService {
	return &commentService{}
}

// CreateComment implements CommentService
func (*commentService) CreateComment(commentReq *model.CommentReq) error {
	id, _ := uuid.NewRandom()
	assignee := &model.Assignee{Id: id.String(), Email: commentReq.Assignee.Email, UserName: commentReq.Assignee.UserName}
	assigneeId, err := assigneerepo.AddAssignee(assignee)
	if err != nil {
		return err
	}
	commentId, _ := uuid.NewRandom()
	comment := &model.Comment{Id: commentId.String(), IssueId: commentReq.IssueId, CreatedAt: time.Now(), Body: commentReq.Body, AssigneeId: assigneeId}

	err = commentRepo.AddComments(comment)
	if err != nil {
		return err
	}
	return nil
}

// GetComment implements CommentService
func (*commentService) GetComment(issueID string) (*model.CommentDTOArray, error) {
	comments, err := commentRepo.GetComments(issueID)
	if err != nil {
		return nil, err
	}
	var commentDtoArray model.CommentDTOArray
	for _, comment := range comments {
		assignee, err := assigneerepo.GetAssignee(comment.AssigneeId)
		if err != nil {
			return nil, err
		}
		commentDtoArray.CommentDtos = append(commentDtoArray.CommentDtos, model.CommentDto{
			CreatedAt: comment.CreatedAt,
			Body:      comment.Body,
			Assignee:  model.Assignee{Email: assignee.Email, UserName: assignee.UserName},
		})
	}

	return &commentDtoArray, nil
}
