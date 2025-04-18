package validator

import (
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rranand/backdrop/internal/util"
)

func IsEmailValid(txt string) bool {
	txt = TrimString(txt)
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, txt)
	return match
}

func TrimString(txt string) string {
	return strings.TrimSpace(txt)
}

func IsTextEmpty(txt string) bool {
	return len(TrimString(txt)) == 0
}

func IsJWTValid(tokenStr *string) bool {
	if tokenStr == nil || *tokenStr == "" {
		return false
	}

	key, err := util.GetJWTSecret()

	if err != nil {
		return false
	}

	token, err := jwt.Parse(*tokenStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	switch {
	case token.Valid:
		return true
	default:
		return false
	}
}

func IsTaskIDValid(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
