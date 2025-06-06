package config

import (
	"fmt"
	"log/slog"
	"os"
)

type ConfigOpt func(*Settings, *slog.Logger) error

func configEnv(key string) (string, error) {
	const op = "config.config-env"

	val := os.Getenv(key)

	if len(val) == 0 {
		return "", fmt.Errorf(
			"error of the %s: the %s var is empty or didn't set: check it and try again", op, key)
	}

	return val, nil
}

func ConfigSocket(st *Settings, log *slog.Logger) error {
	socket, err := configEnv("SOCKET")

	if err != nil {
		log.Error(err.Error())
		return err
	}
	st.Socket = socket

	return nil
}

func ConfigRedis(st *Settings, log *slog.Logger) error {
	pwd, _ := configEnv("REDIS_PWD")
	st.RedisPWD = pwd

	socket, err := configEnv("REDIS_SOCKET")

	if err != nil {
		log.Error(err.Error())
		return err
	}
	st.RedisSocket = socket

	return nil
}
