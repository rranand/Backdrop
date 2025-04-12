package user

import (
	"context"
	"fmt"
)

type Service interface {
	LoginUser(ctx context.Context, userData *UserModel) error
	CreateUser(ctx context.Context, userData *UserModel) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) LoginUser(ctx context.Context, userData *UserModel) error {
	return s.repo.Save(ctx, userData)
}

func (s *service) CreateUser(ctx context.Context, userData *UserModel) error {
	return s.repo.Save(ctx, userData)
}

// example error
var ErrInvalidAmount = fmt.Errorf("amount must be greater than 0")
