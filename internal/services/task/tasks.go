package task

import (
	"log/slog"

	"github.com/MaKcm14/workmate-task/internal/services"
)

type TaskStarter struct {
	task1Starter
}

func NewStarter(log *slog.Logger, repo services.Repository) *TaskStarter {
	return &TaskStarter{
		task1Starter: task1Starter{
			log:  log,
			repo: repo,
		},
	}
}
