package redis

import "context"

type task1Repo struct {
	conf *redisConfig
}

func (r RedisRepo) GetTask1Result(ctx context.Context, taskID int) ([]byte, error) {
	const op = "redis.get-task-result"
	return nil, nil
}

func (r RedisRepo) SetTask1Result(ctx context.Context, taskID int, res []byte) error {
	const op = "redis.set-task-result"
	return nil
}

func (r RedisRepo) DeleteTask1ID(ctx context.Context, taskID int) error {
	const op = "redis.delete-task-id"
	return nil
}
