package user

import "context"

type Repository interface {
	Save(ctx context.Context, userModel *UserModel) error
}

type repo struct {
	db map[int64]*UserModel
}

func NewRepository() Repository {
	return &repo{
		db: make(map[int64]*UserModel),
	}
}

func (r *repo) Save(ctx context.Context, userModel *UserModel) error {
	r.db[userModel.ID] = userModel
	return nil
}
