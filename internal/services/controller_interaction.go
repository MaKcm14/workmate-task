package services

import (
	"context"

	"github.com/MaKcm14/workmate-task/internal/entities/dto"
)

type (
	TestTaskSolver interface {
		StartTestTask1(ctx context.Context, request dto.TaskRequest)
		CheckTask1Status(ctx context.Context, request dto.TaskRequest) (string, error)
	}

	Solver interface {
		TestTaskSolver
	}
)
