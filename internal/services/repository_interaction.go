package services

import "context"

type (
	Checker interface {
		IsTaskExists(ctx context.Context, taskID int) bool
	}

	Getter interface {
		GetTask1Result(ctx context.Context, taskID int) (string, error)
	}

	Modifier interface {
		SetTask1Result(ctx context.Context, taskID int, res string) error
		DeleteTask1ID(ctx context.Context, taskID int) error
	}

	TaskRepo interface {
		Checker
		Getter
		Modifier
	}

	Repository interface {
		TaskRepo
	}
)
