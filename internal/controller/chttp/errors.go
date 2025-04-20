package chttp

import "errors"

var (
	ErrStartingServer = errors.New("error of starting the http-server")
	ErrClientRequest  = errors.New("error of parsing the request")
)
