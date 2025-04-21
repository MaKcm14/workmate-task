package chttp

type TaskInfoResponse struct {
	TaskID int `json:"task_id"`
}

type TaskTest1ResultResponse struct {
	TaskID int    `json:"task_id"`
	Result string `json:"result"`
}

type ErrResponse struct {
	Error string `json:"error"`
}
