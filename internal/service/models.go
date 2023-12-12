package service

// to separate the entity for the repository from the dto of the request

type AddTask struct {
	Status string
}

type GetTask struct {
	ID int
	Task
}

type AddResponse struct {
	ID             int
	Status         string
	HTTPStatusCode int
	Headers        map[string][]string
	Length         int
}

type Task struct {
	Status         string
	HTTPStatusCode int
	Headers        map[string][]string
	Length         int
}
