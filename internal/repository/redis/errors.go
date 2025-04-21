package redis

import "errors"

var (
	ErrConnToRedis = errors.New("error of connection to the redis-server")
	ErrQueryExec   = errors.New("error of executing the query: from the redis-server")
)
