package dto

type GetTaskResponse struct {
	ID             int    `json:"id"`
	Status         string `json:"status"`
	HTTPStatusCode int    `json:"httpStatusCode"`
	// I will use map[string][]string to handle the case where the same header is provided with different values
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
}

type CreateTaskRequest struct {
	Method string `json:"method" validate:"required"`
	URL    string `json:"url" validate:"required"`
	// Not a map[string][]string since in the file provided it wasn't written as an array.
	Headers map[string]string `json:"headers"`
	// I cannot handle the case where the same header is provided multiple times with different values
	// as map[string]string. If the same header is provided multiple times, the last value is the one used
	// it would be a non-valid json
}

type CreateTaskResponse struct {
	ID int `json:"id"`
}
