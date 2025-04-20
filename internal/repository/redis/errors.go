package redis

import "errors"

var (
	ErrConnToRedis = errors.New("error of connection to the redis-server")
)
