package user

import "context"

type Repository interface {
	Save(ctx context.Context, userModel *UserModel) error
}

type repo struct {
	db map[string]*UserModel
}

func NewRepository() Repository {
	return &repo{
		db: make(map[string]*UserModel),
	}
}

func (r *repo) Save(ctx context.Context, userModel *UserModel) error {
	r.db[string(userModel.ID)] = userModel
	return nil
}
