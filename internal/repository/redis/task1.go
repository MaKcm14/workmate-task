package redis

import (
	"context"
	"fmt"
)

type task1Repo struct {
	conf *redisConfig
}

func (r RedisRepo) GetTask1Result(ctx context.Context, taskID int) (string, error) {
	const op = "redis.get-task1-result"

	queryRes := r.conf.conn.HGet(ctx, tasksResKey, fmt.Sprint(taskID))

	res, err := queryRes.Result()

	if err != nil {
		r.conf.log.Warn(fmt.Sprintf("error of the %s: %s", op, err))
		return "", fmt.Errorf("error of the %s: %w: %s", op, ErrQueryExec, err)
	}

	return res, nil
}

func (r RedisRepo) SetTask1Result(ctx context.Context, taskID int, res string) error {
	const op = "redis.set-task1-result"

	if err := r.conf.conn.HSet(ctx, tasksResKey, fmt.Sprint(taskID), res).Err(); err != nil {
		r.conf.log.Warn(fmt.Sprintf("error of the %s: %s", op, err))
		return fmt.Errorf("error of the %s: %w: %s", op, ErrQueryExec, err)
	}

	return nil
}

func (r RedisRepo) DeleteTask1ID(ctx context.Context, taskID int) error {
	const op = "redis.delete-task1-id"

	if err := r.conf.conn.HDel(ctx, tasksResKey, fmt.Sprint(taskID)).Err(); err != nil {
		r.conf.log.Warn(fmt.Sprintf("error of the %s: %s", op, err))
		return fmt.Errorf("error of the %s: %w: %s", op, ErrQueryExec, err)
	}

	return nil
}
