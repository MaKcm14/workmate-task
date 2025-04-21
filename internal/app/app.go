package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/MaKcm14/workmate-task/internal/config"
	"github.com/MaKcm14/workmate-task/internal/controller/chttp"
	"github.com/MaKcm14/workmate-task/internal/repository/redis"
	"github.com/MaKcm14/workmate-task/internal/services/task"
	"github.com/labstack/echo"
)

type Service struct {
	logFile *os.File
	log     *slog.Logger
	contr   *chttp.Controller
	st      config.Settings
}

func NewService() Service {
	date := strings.Split(time.Now().String()[:19], " ")

	mainLogFile, err := os.Create(fmt.Sprintf("../../logs/task-server-main-logs_%s___%s.txt",
		date[0], strings.Join(strings.Split(date[1], ":"), "-")))

	if err != nil {
		panic(fmt.Sprintf("error of creating the main-log-file: %v", err))
	}

	log := slog.New(slog.NewTextHandler(mainLogFile, &slog.HandlerOptions{Level: slog.LevelInfo}))

	log.Info("main application's configuring begun")

	st, err := config.NewSettings(log,
		config.ConfigSocket,
		config.ConfigRedis,
	)

	if err != nil {
		mainLogFile.Close()
		panic(err)
	}

	redisConn, err := redis.New(context.Background(), log, st.RedisPWD, st.RedisSocket)

	if err != nil {
		mainLogFile.Close()
		panic(err)
	}

	return Service{
		logFile: mainLogFile,
		contr:   chttp.NewController(log, echo.New(), task.NewStarter(log, redisConn)),
		st:      st,
	}
}

func (s *Service) Run() {
	defer s.close()
	defer s.log.Info("the app was FULLY STOPPED")

	s.contr.Run(s.st.Socket)
}

func (s *Service) close() {
	s.logFile.Close()
	s.contr.Close()
}
