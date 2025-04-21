package services

import "context"

type (
	Task1Checker interface {
		IsTaskExists(ctx context.Context, taskID int) bool
	}

	Task1Getter interface {
		GetTask1Result(ctx context.Context, taskID int) (string, error)
	}

	Task1Modifier interface {
		SetTask1Result(ctx context.Context, taskID int, res string) error
		DeleteTask1ID(ctx context.Context, taskID int) error
	}

	TaskRepo interface {
		Task1Checker
		Task1Getter
		Task1Modifier
	}
)
