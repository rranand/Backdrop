package user

import (
	"github.com/rranand/backdrop/pkg/text"
	"github.com/rranand/backdrop/pkg/validator"
)

type UserModel struct {
	ID       text.TrimmedString `json:"id,omitempty"`
	Username text.TrimmedString `json:"username"`
	Name     text.TrimmedString `json:"name"`
	Password text.TrimmedString `json:"password"`
	Email    text.TrimmedString `json:"email"`
}

type LoginRequestModel struct {
	Identifier text.TrimmedString `json:"identifier"`
	Password   text.TrimmedString `json:"password"`
}

func GetUserFromLoginRequest(loginRequestModel LoginRequestModel) *UserModel {
	userData := UserModel{Password: loginRequestModel.Password}

	if validator.IsEmailValid(string(loginRequestModel.Identifier)) {
		userData.Email = loginRequestModel.Identifier
	} else {
		userData.Username = loginRequestModel.Identifier
	}

	return &userData
}
