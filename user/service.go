package user

import (
	"context"
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
	err := error(nil)

	if len(userData.Email) == 0 {
		err = s.repo.LoginUserByUsername(ctx, userData)
	} else {
		err = s.repo.LoginUserByEmail(ctx, userData)
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateUser(ctx context.Context, userData *UserModel) error {
	err := s.repo.CreateUser(ctx, userData)

	if err != nil {
		return err
	}
	return nil
}
