package chttp

import (
	"fmt"

	"github.com/MaKcm14/workmate-task/internal/entities/dto"
	"github.com/labstack/echo"
)

func parseRequest(ctx echo.Context) (dto.TaskRequest, error) {
	const op = "chttp.parse-request"

	request := dto.TaskRequest{}

	if err := ctx.Bind(&request); err != nil {
		return dto.TaskRequest{}, fmt.Errorf("error of the %s: %s", op, err)
	}

	return request, nil
}
