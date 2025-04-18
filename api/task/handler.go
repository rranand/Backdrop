package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/rranand/backdrop/api/user"
	"github.com/rranand/backdrop/internal/util"
	"github.com/rranand/backdrop/pkg/constants"
	"github.com/rranand/backdrop/pkg/validator"
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

func (h *Handler) FetchTask(w http.ResponseWriter, r *http.Request) {
	res := util.JSONResponseWriter{ResponseWriter: w}

	authData, ok := r.Context().Value(constants.AuthDataKey).(user.AuthModel)
	if !ok {
		res.SendJSONError("Internal Server Error", http.StatusBadRequest)
		return
	}

	taskID := chi.URLParam(r, "taskID")

	if !validator.IsTaskIDValid(taskID) {
		res.SendJSONError("Invalid Data Provided", http.StatusBadRequest)
		return
	}

	taskData := TaskResponseModel{UploadURL: taskID}

	if err := h.service.FetchTask(r.Context(), &taskData, authData.UserID); err != nil {
		res.SendJSONError(err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(taskData)
}

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	res := util.JSONResponseWriter{ResponseWriter: w}
	ctx := r.Context()

	w.Header().Set("Connection", "close")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	for i := range 10 {
		if err := ctx.Err(); err != nil {
			fmt.Printf("Context error at step %d: %v\n", i+1, err)
			res.SendJSONError("Request cancelled by user", http.StatusBadRequest)
			return
		}
		select {
		case <-ctx.Done():
			res.SendJSONError("Request cancelled by user", http.StatusBadRequest)
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("Running task... Step %d\n", i+1)
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
	// res := util.JSONResponseWriter{ResponseWriter: w}

	// r.ParseMultipartForm(20 << 30) // 10MB

	// _, ok := r.Context().Value(constants.AuthDataKey).(user.AuthModel)
	// if !ok {
	// 	res.SendJSONError("Internal Server Error", http.StatusBadRequest)
	// 	return
	// }

	// file, handler, err := r.FormFile("file")
	// if err != nil {
	// 	http.Error(w, "Error retrieving the file", http.StatusBadRequest)
	// 	return
	// }
	// defer file.Close()

	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	// // Create destination file
	// dst, err := os.Create("/Users/rohitanand/Projects/Backdrop/media/" + handler.Filename)
	// if err != nil {
	// 	http.Error(w, "Unable to create the file", http.StatusInternalServerError)
	// 	return
	// }
	// defer dst.Close()

	// // Copy the uploaded file to destination
	// _, err = io.Copy(dst, file)
	// if err != nil {
	// 	http.Error(w, "Error while saving file", http.StatusInternalServerError)
	// 	return
	// }

}
