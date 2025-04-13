package user

import (
	"context"
	"errors"

	"github.com/rranand/backdrop/internal/util"
)

type Service interface {
	LoginUser(ctx context.Context, userData *UserModel, loginRequestModel *LoginRequestModel) error
	CreateUser(ctx context.Context, userData *UserModel) error
	AuthUser(ctx context.Context, authData *AuthModel) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) LoginUser(ctx context.Context, userData *UserModel, loginRequestModel *LoginRequestModel) error {
	err := error(nil)

	if len(userData.Email) == 0 {
		err = s.repo.LoginUserByUsername(ctx, userData)
	} else {
		err = s.repo.LoginUserByEmail(ctx, userData)
	}

	if err != nil {
		return err
	}

	err = s.repo.GenerateLoginToken(ctx, userData, loginRequestModel)

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

func (s *service) AuthUser(ctx context.Context, authData *AuthModel) error {
	claims, err := util.ParseJWT(string(authData.Token))

	if err != nil {
		return err
	}

	if claims["user_id"] != string(authData.Username) {
		return errors.New("invalid login session")
	}

	err = s.repo.ValidateLoginToken(ctx, authData)

	if err != nil {
		return err
	}
	return nil
}
