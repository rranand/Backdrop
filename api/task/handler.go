package task

import (
	"encoding/json"
	"net/http"

	"github.com/rranand/backdrop/api/user"
	"github.com/rranand/backdrop/internal/util"
	"github.com/rranand/backdrop/pkg/constants"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	res := util.JSONResponseWriter{ResponseWriter: w}

	authData, ok := r.Context().Value(constants.AuthDataKey).(user.AuthModel)
	if !ok {
		res.SendJSONError("Internal Server Error", http.StatusBadRequest)
		return
	}

	var newTask NewTaskModel
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		res.SendJSONError("Invalid Data Provided", http.StatusBadRequest)
		return
	}

	if len(newTask.TaskType) <= 7 {
		res.SendJSONError("Invalid Data Provided", http.StatusBadRequest)
		return
	}

	newTask.UserID = authData.UserID

	if err := h.service.CreateTask(r.Context(), &newTask); err != nil {
		res.SendJSONError(err.Error(), http.StatusBadRequest)
		return
	}

	newTaskResponseData := NewTaskResponseModel{UploadURL: newTask.ID, TaskType: string(newTask.TaskType), Status: newTask.Status}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTaskResponseData)
}
