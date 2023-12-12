package dto

type GetTaskResponse struct {
	ID             int                 `json:"id"`
	Status         string              `json:"status"`
	HTTPStatusCode int                 `json:"httpStatusCode"`
	Headers        map[string][]string `json:"headers"`
	Length         int                 `json:"length"`
}

type CreateTaskRequest struct {
	Method  string              `json:"method" validate:"required"`
	URL     string              `json:"url" validate:"required"`
	Headers map[string][]string `json:"headers"` // map since in the file provided it wasn't written as an array
}

type CreateTaskResponse struct {
	ID int `json:"id"`
}
