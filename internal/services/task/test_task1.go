package task

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/MaKcm14/workmate-task/internal/entities/dto"
	"github.com/MaKcm14/workmate-task/internal/services"
)

type testTask1Starter struct {
	log  *slog.Logger
	repo services.TaskRepo
}

func (t *testTask1Starter) StartTestTask1(ctx context.Context, request dto.TaskRequest) {
	const op = "task.start-test-task1"

	time.Sleep(time.Minute)

	res := "This is a test result for the test-task-1: Hello, World!"

	if err := t.repo.SetTask1Result(context.Background(), request.TaskID, res); err != nil {
		t.log.Warn(fmt.Sprintf("error of the %s: %s", op, err))
	}
}

func (t *testTask1Starter) CheckTask1Status(ctx context.Context, request dto.TaskRequest) (string, error) {
	const op = "task.check-task-status"

	if t.repo.IsTaskExists(ctx, request.TaskID) {
		res, err := t.repo.GetTask1Result(ctx, request.TaskID)

		if err != nil {
			t.log.Warn(fmt.Sprintf("error of %s: %s: %s", op, ErrRepoResponse, err))
			return "", fmt.Errorf("error of %s: %w: %s", op, ErrRepoResponse, err)
		}

		if err := t.repo.DeleteTask1ID(ctx, request.TaskID); err != nil {
			t.log.Warn(fmt.Sprintf("error of %s: %s: %s", op, ErrRepoResponse, err))
			return "", fmt.Errorf("error of %s: %w: %s", op, ErrRepoResponse, err)
		}

		return res, nil
	}
	return "", ErrTaskResponse
}
