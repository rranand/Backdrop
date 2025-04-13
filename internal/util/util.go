package util

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JSONResponseWriter struct {
	http.ResponseWriter
}

func (w JSONResponseWriter) SendJSONError(err string, httpCode int) {
	w.Header().Set("Content-Type", "application/json")
	jsonStr, _ := json.Marshal(map[string]string{"Error": err})
	stringRep := string(jsonStr)
	http.Error(w, stringRep, httpCode)
}

func getJWTSecret() ([]byte, error) {
	secretStr := os.Getenv("JWT_SECRET")
	if secretStr == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}

	secret, err := base64.URLEncoding.DecodeString(secretStr)
	if err != nil {
		secret = []byte(secretStr)
	}

	if len(secret) < 32 {
		return nil, errors.New("JWT_SECRET should be at least 32 bytes")
	}

	return secret, nil
}

func GenerateJWTToken(claims jwt.Claims) (string, error) {
	key, err := getJWTSecret()

	if err != nil {
		return "", errors.New("failed to parse JWT key")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(key)
}

func GenerateUUID() string {
	return uuid.New().String()
}
