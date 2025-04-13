package user

import (
	"encoding/json"
	"fmt"
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
	var loginRequestModel LoginRequestModel

	if err := json.NewDecoder(r.Body).Decode(&loginRequestModel); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if len(loginRequestModel.Identifier) <= 5 || len(loginRequestModel.Password) <= 8 {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	userData := GetUserFromLoginRequest(loginRequestModel)
	if err := h.service.LoginUser(r.Context(), userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := util.GenerateRandomToken(32)

	for err != nil {
		token_temp, err_temp := util.GenerateRandomToken(32)
		token = token_temp
		err = err_temp
	}

	userObj, _ := json.Marshal(userData)
	fmt.Println("User : ", string(userObj), "\n", "Token : ", token)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "login success"})
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData UserModel
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if !validator.IsEmailValid(string(userData.Email)) || len(userData.Password) <= 7 || len(userData.Username) <= 5 || len(userData.Name) <= 2 {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(r.Context(), &userData); err != nil {
		res := util.JSONResponseWriter{ResponseWriter: w}
		res.SendJSONError(err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "signup success"})
}
