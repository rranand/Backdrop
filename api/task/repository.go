package task

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rranand/backdrop/pkg/constants"
	"github.com/rranand/backdrop/pkg/database"
)

type Repository interface {
	CreateTask(ctx context.Context, newTask *NewTaskModel) error
}

type repo struct {
	db *sql.DB
}

func NewRepository() Repository {
	return &repo{
		db: database.DB,
	}
}

var ErrConflict = fmt.Errorf("conflict records found")
var ErrNoRecordFound = fmt.Errorf("no record found")
var ErrInvalidCredential = fmt.Errorf("credential not valid")

func (r *repo) CreateTask(ctx context.Context, newTask *NewTaskModel) error {
	query := `
		INSERT INTO tasks (
			task_type,
			user_id
		) VALUES (
			$1, $2
		)
		RETURNING id, status;
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeoutDuration)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, newTask.TaskType, newTask.UserID).Scan(&newTask.ID, &newTask.Status)

	if err != nil {
		return err
	}
	return nil
}
