package chttp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/MaKcm14/workmate-task/internal/entities/dto"
	"github.com/MaKcm14/workmate-task/internal/services"
	"github.com/labstack/echo"
)

type Controller struct {
	log    *slog.Logger
	contr  *echo.Echo
	slver  services.Solver
	taskID int
}

func NewController(log *slog.Logger, e *echo.Echo, slver services.Solver) *Controller {
	return &Controller{
		log:   log,
		contr: e,
		slver: slver,
	}
}

func (c *Controller) Run(socket string) {
	const op = "chttp.run"

	c.log.Info("configuring and starting the server")

	c.config()

	if err := c.contr.Start(socket); err != nil {
		c.log.Error(fmt.Sprintf("error of the %s: %s: %s", op, err, ErrStartingServer))
		panic(fmt.Sprintf("error of the %s: %s: %s", op, err, ErrStartingServer))
	}
}

func (c *Controller) config() {
	// controller's PATH config here.
	c.contr.POST("/test/task1", c.handlerTestTask1)
	c.contr.POST("/results/test/task1", c.handlerTestTask1Result)
}

// POST-запрос, который может содержать в теле все необходимые параметры для выполнения задачи.
func (c *Controller) handlerTestTask1(ctx echo.Context) error {
	const op = "chttp.handler-test-task-1"

	request := dto.TaskRequest{c.taskID}
	c.taskID++

	go c.slver.StartTestTask1(context.Background(), request)

	return ctx.JSON(http.StatusAccepted, TaskResponse{c.taskID})
}

// POST-запрос, в котором передается taskID клиента.
//
//	{
//	    "task_id" : task_id
//	}
func (c *Controller) handlerTestTask1Result(ctx echo.Context) error {
	const op = "chttp.handler-test-task-1-result"

	request, err := parseRequest(ctx)

	if err != nil {
		requestErr := fmt.Sprintf("error of the %s: %s: %s", op, err, ErrClientRequest)
		c.log.Warn(requestErr)
		return ctx.JSON(http.StatusBadRequest, ErrResponse{requestErr})
	}

	_ = request

	// Add here the code of checking the task's status.

	return nil
}

func (c *Controller) Close() {
	c.contr.Close()
}
