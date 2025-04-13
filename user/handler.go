package user

import (
	"encoding/json"
	"net/http"

	"github.com/rranand/backdrop/internal/util"
	"github.com/rranand/backdrop/pkg/validator"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	res := util.JSONResponseWriter{ResponseWriter: w}
	var loginRequestData LoginRequestModel

	if err := json.NewDecoder(r.Body).Decode(&loginRequestData); err != nil {
		res.SendJSONError("Invalid Data Provided", http.StatusBadRequest)
		return
	}

	if len(loginRequestData.Identifier) <= 5 || len(loginRequestData.Password) <= 7 {
		res.SendJSONError("Invalid Data Provided", http.StatusBadRequest)
		return
	}

	userData := GetUserFromLoginRequest(loginRequestData)
	if err := h.service.LoginUser(r.Context(), &userData, &loginRequestData); err != nil {
		res.SendJSONError(err.Error(), http.StatusBadRequest)
		return
	}

	userLoginToken := LoginResponseModel{
		Token: string(userData.Token),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userLoginToken)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	res := util.JSONResponseWriter{ResponseWriter: w}
	var userData UserModel

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		res.SendJSONError("Invalid Data Provided", http.StatusBadRequest)
		return
	}

	if !validator.IsEmailValid(string(userData.Email)) || len(userData.Password) <= 7 || len(userData.Username) <= 5 || len(userData.Name) <= 2 {
		res.SendJSONError("Invalid Data Provided", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(r.Context(), &userData); err != nil {
		res.SendJSONError(err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userData)
}
