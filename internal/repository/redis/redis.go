package redis

import (
	"context"
	"log/slog"
)

type RedisRepo struct {
	task1Repo
}

func New(ctx context.Context, log *slog.Logger, pwd, socket string) (RedisRepo, error) {
	conf, err := newRedisConfig(ctx, log, pwd, socket)

	if err != nil {
		return RedisRepo{}, err
	}

	return RedisRepo{
		task1Repo: task1Repo{conf},
	}, nil
}

func (r RedisRepo) IsTaskExists(ctx context.Context, taskID int) bool {
	return false
}
