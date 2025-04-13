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

func GetJWTSecret() ([]byte, error) {
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
	key, err := GetJWTSecret()

	if err != nil {
		return "", errors.New("failed to parse JWT key")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(key)
}

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {
	key, err := GetJWTSecret()

	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		return claims, nil
	} else {
		return nil, errors.New("failed to parse JWT Token")
	}
}

func GenerateUUID() string {
	return uuid.New().String()
}
