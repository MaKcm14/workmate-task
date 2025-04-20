package services

import (
	"context"

	"github.com/MaKcm14/workmate-task/internal/entities/dto"
)

type (
	TestTaskSolver interface {
		StartTestTask1(ctx context.Context, request dto.TaskRequest)
		CheckTaskStatus(ctx context.Context, request dto.TaskRequest) ([]byte, error)
	}

	Solver interface {
		TestTaskSolver
	}
)
