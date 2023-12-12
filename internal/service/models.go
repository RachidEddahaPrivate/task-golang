package service

// to separate the entity for the repository from the dto of the request

type AddTask struct {
	Method  string
	URL     string
	Headers map[string][]string
}

type GetTask struct {
	ID             int
	Status         string
	HTTPStatusCode int
	Headers        map[string][]string
	Length         int
}

type AddResponse struct {
	ID             int
	Status         string
	HTTPStatusCode int
	Headers        map[string][]string
	Length         int
}
