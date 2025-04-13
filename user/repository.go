package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rranand/backdrop/internal/util"
	"github.com/rranand/backdrop/pkg/constants"
	"github.com/rranand/backdrop/pkg/database"
	"github.com/rranand/backdrop/pkg/text"
)

type Repository interface {
	LoginUserByEmail(ctx context.Context, userModel *UserModel) error
	LoginUserByUsername(ctx context.Context, userModel *UserModel) error
	CreateUser(ctx context.Context, userData *UserModel) error
	GenerateLoginToken(ctx context.Context, userData *UserModel, loginRequestData *LoginRequestModel) error
	FetchUser(ctx context.Context, authData *AuthModel) (ProfileModel, error)
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
		SELECT id, email, name 
		FROM users
		WHERE 
		username = $1 AND password = $2; 
	`

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
		SELECT id, username, name 
		FROM users
		WHERE 
		email = $1 AND password = $2; 
	`

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

func (r *repo) GenerateLoginToken(ctx context.Context, userData *UserModel, loginRequestData *LoginRequestModel) error {
	query := `
		INSERT INTO login_data (
			token,
			user_agent,
			ip_address,
			isp,
			state,
			city,
			country,
			device_type,
			user_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)
		RETURNING id;
	`

	currentTime := time.Now()

	jwtToken, err := util.GenerateJWTToken(jwt.MapClaims{
		"ip_address": loginRequestData.IPAddress,
		"user_id":    userData.Username,
		"hash":       util.GenerateUUID(),
		"TimeStamp":  currentTime,
	})

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeoutDuration)
	defer cancel()

	err = r.db.QueryRowContext(ctx, query, jwtToken, loginRequestData.UserAgent, loginRequestData.IPAddress, loginRequestData.ISP, loginRequestData.State, loginRequestData.City, loginRequestData.Country, loginRequestData.DeviceType, userData.ID).Scan(&loginRequestData.ID)

	if err != nil {
		return err
	}
	userData.Token = text.TrimmedString(jwtToken)
	return nil
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

func (r *repo) FetchUser(ctx context.Context, authData *AuthModel) (ProfileModel, error) {
	var profileData ProfileModel

	query := `
		SELECT id, email, username, name, created_at, updated_at FROM users as u
		WHERE username = $1
		; 
	`

	ctx, cancel := context.WithTimeout(ctx, constants.QueryTimeoutDuration)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, authData.Username)
	if err != nil {
		return profileData, err
	}
	defer rows.Close()

	cnt := 0
	var uid string

	for rows.Next() {
		cnt++
		err := rows.Scan(&uid, &profileData.Email, &profileData.Username, &profileData.Name, &profileData.CreatedOn, &profileData.UpdatedOn)
		if err != nil {
			return profileData, err
		}
	}

	if cnt == 0 {
		return profileData, ErrNoRecordFound
	} else if cnt > 1 {
		return ProfileModel{}, nil
	}

	query = `
		SELECT MAX(last_logged_in) FROM login_data as ld
		WHERE user_id = $1
		GROUP BY user_id
		; 
	`

	ctx, cancel = context.WithTimeout(ctx, constants.QueryTimeoutDuration)
	defer cancel()

	rows, err = r.db.QueryContext(ctx, query, uid)
	if err != nil {
		return profileData, err
	}
	defer rows.Close()

	var last_logged_in time.Time

	for rows.Next() {
		cnt++
		err := rows.Scan(&last_logged_in)
		if err != nil {
			return profileData, err
		}
	}
	profileData.LastLoggedIn = last_logged_in

	return profileData, nil
}
