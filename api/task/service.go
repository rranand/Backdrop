package task

import (
	"context"
)

type Service interface {
	CreateTask(ctx context.Context, newTask *NewTaskModel) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateTask(ctx context.Context, newTask *NewTaskModel) error {
	err := s.repo.CreateTask(ctx, newTask)

	if err != nil {
		return err
	}
	return nil
}
