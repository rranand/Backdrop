package task

import (
	"context"
)

type Service interface {
	CreateTask(ctx context.Context, newTask *NewTaskModel) error
	FetchTask(ctx context.Context, taskData *TaskResponseModel, uid string) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateTask(ctx context.Context, newTask *NewTaskModel) error {
	err := s.repo.CreateTask(ctx, newTask)

	if err != nil {
		return err
	}
	return nil
}

func (s *service) FetchTask(ctx context.Context, taskData *TaskResponseModel, uid string) error {
	err := s.repo.FetchTask(ctx, taskData, uid)

	if err != nil {
		return err
	}

	taskData.Message = getTaskMessage(taskData.Status)
	if taskData.Status != "COMPLETED" {
		taskData.DownloadURL = ""
	}
	return nil
}

func getTaskMessage(status string) string {
	switch status {
	case "NOT_UPLOADED":
		return "File not uploaded yet. Please upload to start the task."
	case "UPLOADING":
		return "Your file is being uploaded. Please wait..."
	case "FAILED":
		return "There was a problem processing your file. Try again."
	case "PROCESSING":
		return "Your file is currently being processed."
	case "COMPLETED":
		return "Your file is ready to download."
	case "CANCELLED":
		return "Task cancelled. It will not be processed further."
	default:
		return "Unknown task status."
	}
}
