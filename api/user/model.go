package user

import (
	"time"

	"github.com/rranand/backdrop/pkg/text"
	"github.com/rranand/backdrop/pkg/validator"
)

type UserModel struct {
	ID       text.TrimmedString `json:"id,omitempty"`
	Username text.TrimmedString `json:"username"`
	Name     text.TrimmedString `json:"name"`
	Password text.TrimmedString `json:"password"`
	Email    text.TrimmedString `json:"email"`
	Token    text.TrimmedString `json:"token"`
}

type LoginRequestModel struct {
	ID         text.TrimmedString `json:"id,omitempty"`
	Identifier text.TrimmedString `json:"identifier"`
	Password   text.TrimmedString `json:"password"`
	UserAgent  text.TrimmedString `json:"user_agent"`
	IPAddress  text.TrimmedString `json:"ip_address"`
	ISP        text.TrimmedString `json:"isp"`
	State      text.TrimmedString `json:"state"`
	City       text.TrimmedString `json:"city"`
	Country    text.TrimmedString `json:"country"`
	DeviceType text.TrimmedString `json:"device_type"`
}

type LoginResponseModel struct {
	Token string `json:"token"`
}

type AuthModel struct {
	UserID string `json:"uid"`
}

type AuthResponseModel struct {
	Status text.TrimmedString `json:"status"`
}

type ProfileModel struct {
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	CreatedOn    time.Time `json:"created_at"`
	UpdatedOn    time.Time `json:"updated_at"`
	LastLoggedIn time.Time `json:"last_logged_in"`
}

func GetUserFromLoginRequest(loginRequestModel LoginRequestModel) UserModel {
	userData := UserModel{Password: loginRequestModel.Password}

	if validator.IsEmailValid(string(loginRequestModel.Identifier)) {
		userData.Email = loginRequestModel.Identifier
	} else {
		userData.Username = loginRequestModel.Identifier
	}

	return userData
}
