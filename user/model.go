package user

import (
	"github.com/rranand/backdrop/pkg/validator"
)

type UserModel struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequestModel struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func GetUserFromLoginRequest(loginRequestModel LoginRequestModel) *UserModel {
	userData := UserModel{Password: loginRequestModel.Password}

	if validator.IsEmailValid(loginRequestModel.Identifier) {
		userData.Email = loginRequestModel.Identifier
	} else {
		userData.Username = loginRequestModel.Identifier
	}

	return &userData
}
