package redis

import (
	"context"
	"fmt"
)

type testTask1Repo struct {
	conf *redisConfig
}

func (t testTask1Repo) GetTask1Result(ctx context.Context, taskID int) (string, error) {
	const op = "redis.get-task1-result"

	queryRes := t.conf.conn.HGet(ctx, tasksResKey, fmt.Sprint(taskID))

	res, err := queryRes.Result()

	if err != nil {
		t.conf.log.Warn(fmt.Sprintf("error of the %s: %s", op, err))
		return "", fmt.Errorf("error of the %s: %w: %s", op, ErrQueryExec, err)
	}

	return res, nil
}

func (t testTask1Repo) SetTask1Result(ctx context.Context, taskID int, res string) error {
	const op = "redis.set-task1-result"

	if err := t.conf.conn.HSet(ctx, tasksResKey, fmt.Sprint(taskID), res).Err(); err != nil {
		t.conf.log.Warn(fmt.Sprintf("error of the %s: %s", op, err))
		return fmt.Errorf("error of the %s: %w: %s", op, ErrQueryExec, err)
	}

	return nil
}

func (t testTask1Repo) DeleteTask1ID(ctx context.Context, taskID int) error {
	const op = "redis.delete-task1-id"

	if err := t.conf.conn.HDel(ctx, tasksResKey, fmt.Sprint(taskID)).Err(); err != nil {
		t.conf.log.Warn(fmt.Sprintf("error of the %s: %s", op, err))
		return fmt.Errorf("error of the %s: %w: %s", op, ErrQueryExec, err)
	}

	return nil
}
