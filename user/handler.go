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

	if len(validator.TrimString(loginRequestModel.Identifier)) <= 5 || len(validator.TrimString(loginRequestModel.Password)) <= 5 {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	userData := GetUserFromLoginRequest(loginRequestModel)
	if err := h.service.LoginUser(r.Context(), userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var loginRequestModel LoginRequestModel
	if err := json.NewDecoder(r.Body).Decode(&loginRequestModel); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	userData := GetUserFromLoginRequest(loginRequestModel)
	if err := h.service.LoginUser(r.Context(), userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("User : ", userData)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
