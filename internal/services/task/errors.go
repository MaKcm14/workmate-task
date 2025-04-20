package task

import "errors"

var (
	ErrRepoResponse = errors.New("error of executing the repo's request")
	ErrTaskResponse = errors.New("error of task response")
)
