package redis

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

const (
	tasksResKey = "tasks:results"
)

type redisConfig struct {
	conn *redis.Client
	log  *slog.Logger
}

func newRedisConfig(ctx context.Context, log *slog.Logger, pwd, socket string) (*redisConfig, error) {
	const op = "redis.new-redis-config"

	client := redis.NewClient(&redis.Options{
		Addr:     socket,
		Password: pwd,
		DB:       0,
	})

	if res := client.Ping(ctx); res.Err() != nil {
		log.Error(fmt.Sprintf("error of the %s: error of ping the redis: %v", op, res.Err()))
		return nil, ErrConnToRedis
	}

	return &redisConfig{
		conn: client,
		log:  log,
	}, nil
}
