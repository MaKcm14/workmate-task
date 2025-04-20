package config

import (
	"fmt"
	"log/slog"
)

// Settings defines the app's settings.
type Settings struct {
	Socket string

	RedisPWD    string
	RedisSocket string
}

func NewSettings(log *slog.Logger, opts ...ConfigOpt) (Settings, error) {
	st := Settings{}

	for _, opt := range opts {
		if err := opt(&st, log); err != nil {
			log.Error(fmt.Sprintf("error of configuring the app: %s", err))
			return st, err
		}
	}

	return st, nil
}
