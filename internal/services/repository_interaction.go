package services

import "context"

type (
	TaskRepo interface {
		GetTaskResult(ctx context.Context, taskID int) ([]byte, error)
		SetTaskResult(ctx context.Context, taskID int, res []byte) error
		DeleteTaskID(ctx context.Context, taskID int) error
		IsTaskExists(ctx context.Context, taskID int) bool
	}

	Repository interface {
		TaskRepo
	}
)
