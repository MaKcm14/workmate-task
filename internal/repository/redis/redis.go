package redis

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	log  *slog.Logger
	conn *redis.Client
}

func New(ctx context.Context, log *slog.Logger, pwd, socket string) (RedisRepo, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     socket,
		Password: pwd,
		DB:       0,
	})

	if res := client.Ping(ctx); res.Err() != nil {
		log.Error(fmt.Sprintf("error of ping the redis: %v", res.Err()))
		return RedisRepo{}, ErrConnToRedis
	}

	return RedisRepo{
		log:  log,
		conn: client,
	}, nil
}

func (r RedisRepo) GetTaskResult(ctx context.Context, taskID int) ([]byte, error) {
	const op = "redis.get-task-result"
	return nil, nil
}

func (r RedisRepo) SetTaskResult(ctx context.Context, taskID int, res []byte) error {
	const op = "redis.set-task-result"
	return nil
}

func (r RedisRepo) DeleteTaskID(ctx context.Context, taskID int) error {
	const op = "redis.delete-task-id"
	return nil
}

func (r RedisRepo) IsTaskExists(ctx context.Context, taskID int) bool {
	return false
}
