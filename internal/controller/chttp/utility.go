package chttp

type TaskResponse struct {
	TaskID int `json:"task_id"`
}

type ErrResponse struct {
	Error string `json:"error"`
}
