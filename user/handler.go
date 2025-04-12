package user

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	userObj, _ := json.Marshal(userData)
	fmt.Println("User : ", string(userObj))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "login success"})
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData UserModel
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if !validator.IsEmailValid(string(userData.Email)) || len(userData.Password) <= 8 || len(userData.Username) <= 5 || len(userData.Name) <= 2 {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	userObj, _ := json.Marshal(userData)
	fmt.Println("User : ", string(userObj))

	if err := h.service.CreateUser(r.Context(), &userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "signup success"})
}
