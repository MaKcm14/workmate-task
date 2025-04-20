package chttp

import "errors"

var (
	ErrStartingServer = errors.New("error of starting the http-server")
	ErrClientRequest  = errors.New("error of parsing the request")
	ErrServerHandling = errors.New("error of server's handling the request")
	ErrTaskGetting    = errors.New("error of getting the task's result")
)
