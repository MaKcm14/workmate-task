package task

import (
	"log/slog"

	"github.com/MaKcm14/workmate-task/internal/services"
)

type TaskStarter struct {
	testTask1Starter
}

func NewStarter(log *slog.Logger, repo services.TaskRepo) *TaskStarter {
	return &TaskStarter{
		testTask1Starter: testTask1Starter{
			log:  log,
			repo: repo,
		},
	}
}
