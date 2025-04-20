package chttp

type TaskResponse struct {
	TaskID int `json:"task_id"`
}

type TaskTestResultResponse struct {
	TaskID int    `json:"task_id"`
	Result string `json:"result"`
}

type ErrResponse struct {
	Error string `json:"error"`
}
