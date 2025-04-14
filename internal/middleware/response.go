package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/rranand/backdrop/api/user"
	"github.com/rranand/backdrop/internal/util"
	"github.com/rranand/backdrop/pkg/constants"
	"github.com/rranand/backdrop/pkg/database"
	"github.com/rranand/backdrop/pkg/validator"
)

var (
	PublicURL = []string{
		"/auth/v1/login",
		"/auth/v1/signup",
	}
)

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func ValidateAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !slices.Contains(PublicURL, r.URL.Path) {
			res := util.JSONResponseWriter{ResponseWriter: w}
			var authData user.AuthModel

			json.NewDecoder(r.Body).Decode(&authData)

			if len(authData.Username) <= 5 || !validator.IsJWTValid(string(authData.Token)) {
				res.SendJSONError("unauthorized access", http.StatusUnauthorized)
				return
			}

			query := `
				UPDATE login_data
				SET last_logged_in = CURRENT_TIMESTAMP
				FROM users
				WHERE 
					users.username = $1
					AND login_data.token = $2
					AND login_data.user_id = users.id
				RETURNING login_data.id
				;
			`

			ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeoutDuration)
			defer cancel()

			var updatedID string

			err := database.DB.QueryRowContext(ctx, query, authData.Username, authData.Token).Scan(&updatedID)

			if err != nil {
				res.SendJSONError("unauthorized access", http.StatusUnauthorized)
				return
			}
			ctx = context.WithValue(r.Context(), constants.AuthDataKey, authData)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
