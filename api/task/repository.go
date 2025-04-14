package task

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
	CreateTask(ctx context.Context) (NewTaskModel, error)
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

func (r *repo) CreateTask(ctx context.Context) (NewTaskModel, error) {
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
