package dto

type GetTaskResponse struct {
	ID             int               `json:"id"`
	Status         string            `json:"status"`
	HTTPStatusCode int               `json:"httpStatusCode"`
	Headers        map[string]string `json:"headers"` // use same convention as in the request
	Length         int               `json:"length"`
}

type CreateTaskRequest struct {
	Method string `json:"method" validate:"required"`
	URL    string `json:"url" validate:"required"`
	// Not a map[string][]string since in the file provided it wasn't written as an array.
	// I assume multiple headers with the different values are separated by comma
	Headers map[string]string `json:"headers"`
}

type CreateTaskResponse struct {
	ID int `json:"id"`
}
