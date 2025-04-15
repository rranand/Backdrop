package task

import (
	"github.com/rranand/backdrop/pkg/text"
)

type NewTaskModel struct {
	ID       string             `json:"id,omitempty"`
	TaskType text.TrimmedString `json:"task_type"`
	Status   string             `json:"status,omitempty"`
	UserID   string             `json:"user_id,omitempty"`
}

type NewTaskResponseModel struct {
	UploadURL string `json:"upload_url"`
	TaskType  string `json:"task_type"`
	Status    string `json:"status,omitempty"`
}

type TaskResponseModel struct {
	UploadURL   string `json:"upload_url"`
	DownloadURL string `json:"download_url"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}
