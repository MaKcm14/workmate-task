package redis

import (
	"context"
	"fmt"
	"log/slog"
)

type RedisRepo struct {
	testTask1Repo
}

func New(ctx context.Context, log *slog.Logger, pwd, socket string) (RedisRepo, error) {
	conf, err := newRedisConfig(ctx, log, pwd, socket)

	if err != nil {
		return RedisRepo{}, err
	}

	return RedisRepo{
		testTask1Repo: testTask1Repo{conf},
	}, nil
}

func (r RedisRepo) IsTaskExists(ctx context.Context, taskID int) bool {
	const op = "redis.is-task-exists"

	if res, err := r.conf.conn.HMGet(ctx, tasksResKey, fmt.Sprint(taskID)).Result(); len(res) != 0 && res[0] == nil {
		return false
	} else if err != nil {
		r.conf.log.Warn(fmt.Sprintf("error of the %s: %s", op, err))
		return false
	}

	return true
}
