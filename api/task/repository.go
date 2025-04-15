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
	FetchTask(ctx context.Context, taskData *TaskResponseModel, uid string) error
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

func (r *repo) FetchTask(ctx context.Context, taskData *TaskResponseModel, uid string) error {
	query := `
		SELECT
		t.download_url, t.status
		FROM tasks as t
		JOIN users ON 
		t.user_id = users.id
		WHERE
		t.id = $1 AND 
		users.id = $2
	`
	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeoutDuration)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, taskData.UploadURL, uid)
	if err != nil {
		return err
	}

	defer rows.Close()

	cnt := 0

	for rows.Next() {
		cnt++
		err := rows.Scan(&taskData.DownloadURL, &taskData.Status)
		if err != nil {
			return err
		}
	}

	if cnt > 1 {
		taskData.DownloadURL = ""
		taskData.Status = ""
		return ErrConflict
	} else if cnt == 0 {
		return ErrNoRecordFound
	} else {
		return nil
	}
}
