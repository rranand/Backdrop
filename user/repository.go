package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rranand/backdrop/pkg/constants"
	"github.com/rranand/backdrop/pkg/database"
)

type Repository interface {
	LoginUserByEmail(ctx context.Context, userModel *UserModel) error
	LoginUserByUsername(ctx context.Context, userModel *UserModel) error
	ValidateLoginToken(ctx context.Context, loginToken string) (*UserModel, error)
	CreateUser(ctx context.Context, userData *UserModel) error
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

func (r *repo) LoginUserByUsername(ctx context.Context, userData *UserModel) error {
	query := `
		SELECT u.id, u.email, u.name FROM users as u
		WHERE u.username = $1 AND u.password = $2; `

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeoutDuration)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, userData.Username, userData.Password)
	if err != nil {
		return err
	}
	defer rows.Close()

	cnt := 0

	for rows.Next() {
		cnt++
		err := rows.Scan(&userData.ID, &userData.Email, &userData.Name)
		if err != nil {
			return err
		}
	}

	if cnt == 0 {
		return ErrInvalidCredential
	} else if cnt > 1 {
		userData.ID = ""
		userData.Email = ""
		userData.Name = ""
	}

	return nil
}

func (r *repo) LoginUserByEmail(ctx context.Context, userData *UserModel) error {
	query := `
		SELECT u.id, u.username, u.name FROM users as u
		WHERE u.email = $1 AND u.password = $2; `

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeoutDuration)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, userData.Email, userData.Password)
	if err != nil {
		return err
	}
	defer rows.Close()

	cnt := 0

	for rows.Next() {
		cnt++
		err := rows.Scan(&userData.ID, &userData.Username, &userData.Name)
		if err != nil {
			return err
		}
	}

	if cnt == 0 {
		return ErrInvalidCredential
	} else if cnt > 1 {
		userData.ID = ""
		userData.Username = ""
		userData.Name = ""
	}

	return nil
}

func (r *repo) ValidateLoginToken(ctx context.Context, loginToken string) (*UserModel, error) {

	currentTime := time.Now()

	query := `
		UPDATE login_data 
		SET last_logged_in = $1 
		FROM users
		WHERE login_data.userID = users.id 
		AND login_data.token = $2 
		RETURNING users.id, users.email, users.username, users.name;
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeoutDuration)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, currentTime, loginToken)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userData UserModel
	cnt := 0

	for rows.Next() {
		cnt++
		err := rows.Scan(&userData.ID, &userData.Email, &userData.Username, &userData.Name)
		if err != nil {
			return nil, err
		}
	}

	if cnt == 1 {
		return &userData, nil
	} else {
		return nil, ErrConflict
	}
}

func (r *repo) CreateUser(ctx context.Context, userData *UserModel) error {

	query := `
		INSERT INTO users (username, email, password, name) 
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeoutDuration)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, userData.Username, userData.Email, userData.Password, userData.Name).Scan(&userData.ID)

	if err != nil {
		return err
	}
	return nil
}
