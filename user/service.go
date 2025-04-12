package user

import (
	"context"
	"fmt"
)

type Service interface {
	LoginUser(ctx context.Context, expense UserModel) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) LoginUser(ctx context.Context, userModel UserModel) error {
	return s.repo.Save(ctx, userModel)
}

// example error
var ErrInvalidAmount = fmt.Errorf("amount must be greater than 0")
