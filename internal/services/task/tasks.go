package task

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/MaKcm14/workmate-task/internal/entities/dto"
	"github.com/MaKcm14/workmate-task/internal/services"
)

type TaskStarter struct {
	log  *slog.Logger
	repo services.Repository
}

func NewStarter(log *slog.Logger, repo services.Repository) *TaskStarter {
	return &TaskStarter{
		log:  log,
		repo: repo,
	}
}

func (t *TaskStarter) StartTestTask1(ctx context.Context, request dto.TaskRequest) {
	const op = "task.start-test-task1"

	time.Sleep(time.Minute * 3)

	res := "This is a test result for the test-task-1: Hello, World!"

	t.repo.SetTaskResult(context.Background(), request.TaskID, []byte(res))
}

func (t *TaskStarter) CheckTaskStatus(ctx context.Context, request dto.TaskRequest) ([]byte, error) {
	const op = "task.check-task-status"

	if t.repo.IsTaskExists(ctx, request.TaskID) {
		res, err := t.repo.GetTaskResult(ctx, request.TaskID)

		if err != nil {
			t.log.Warn(fmt.Sprintf("error of %s: %s: %s", op, ErrRepoResponse, err))
			return nil, fmt.Errorf("error of %s: %w: %s", op, ErrRepoResponse, err)
		}
		return res, nil
	}
	return nil, ErrTaskResponse
}
