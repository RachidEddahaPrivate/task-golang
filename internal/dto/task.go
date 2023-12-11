package dto

type GetTaskResponse struct {
}

type CreateTaskRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type CreateTaskResponse struct {
	ID int `json:"id"`
}
