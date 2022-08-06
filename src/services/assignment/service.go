package assignment

import (
	"context"
	"ggclass_log_service/src/models"
	"time"
)

type IRepository interface {
	Create(ctx context.Context, assignment models.Assignment) error
	GetAll(ctx context.Context) ([]models.Assignment, error)
	FindByAssignmentId(ctx context.Context, assignmentId int) ([]models.Assignment, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, input createAssignmentLogInput) error {
	now := time.Now()
	assignment := models.Assignment{
		AssignmentId: input.AssignmentId,
		Action:       input.Action,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	return s.repository.Create(ctx, assignment)
}

func (s *service) GetAll(ctx context.Context) ([]models.Assignment, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetByAssignmentId(ctx context.Context, assignmentId int) ([]models.Assignment, error) {
	return s.repository.FindByAssignmentId(ctx, assignmentId)
}
