package task

import (
	"github.com/rranand/backdrop/pkg/text"
)

type NewTaskModel struct {
	ID        text.TrimmedString `json:"id,omitempty"`
	UploadUrl text.TrimmedString `json:"upload_url"`
	TaskType  text.TrimmedString `json:"task_type"`
	Status    text.TrimmedString `json:"status"`
	UserID    text.TrimmedString `json:"user_id"`
}
