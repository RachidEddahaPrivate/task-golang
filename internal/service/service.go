package service

import (
	"context"
	"net/http"
	"task/internal/dto"
	"task/pkg/logger"
	"time"
)

// repository is an interface to the repository layer
// in a real application this repository would be implemented with a database such as postgres using gorm as orm
// since it is a simple application, I will use implement the repository to store information in memory on runtime
// inevitably, this means that the data will be lost when the application is stopped
type repository interface {
	AddTask(status string) (int, error)
	GetTask(ID int) (GetTask, error)
	ChangeStatus(ID int, status string) error
	AddResponse(response AddResponse) error
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	if repository == nil {
		panic(repository)
	}
	return &Service{
		repository: repository,
	}
}

const (
	statusNew       = "new"
	statusInProcess = "in_process"
	statusDone      = "done"
	statusError     = "error"

	timeoutHTTPRequest = 10 * time.Second
)

func (s *Service) GetTask(taskID int) (dto.GetTaskResponse, error) {
	task, err := s.repository.GetTask(taskID)
	if err != nil {
		return dto.GetTaskResponse{}, err
	}
	return dto.GetTaskResponse{
		ID:             task.ID,
		Status:         task.Status,
		HTTPStatusCode: task.HTTPStatusCode,
		Headers:        task.Headers,
		Length:         task.Length,
	}, nil
}

func (s *Service) CreateTask(request dto.CreateTaskRequest) (dto.CreateTaskResponse, error) {
	taskID, err := s.repository.AddTask(statusNew)
	if err != nil {
		return dto.CreateTaskResponse{}, err
	}

	go func() {
		s.makeRequest(taskID, request)
	}()

	return dto.CreateTaskResponse{
		ID: taskID,
	}, nil
}

func (s *Service) makeRequest(taskID int, request dto.CreateTaskRequest) {
	err := s.repository.ChangeStatus(taskID, statusInProcess)
	if err != nil {
		s.changeStatusToError(taskID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutHTTPRequest)
	defer cancel() // cancel whenever the function returns

	req, err := http.NewRequestWithContext(ctx, request.Method, request.URL, nil)
	if err != nil {
		logger.Error().Msgf("failed to create request for task %d, err= %v", taskID, err)

		s.changeStatusToError(taskID)
		return
	}

	for k, v := range request.Headers {
		req.Header[k] = v
	}

	client := http.Client{Timeout: timeoutHTTPRequest}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error().Msgf("failed to make request for task %d, err= %v", taskID, err)

		s.changeStatusToError(taskID)
		return
	}

	err = s.repository.AddResponse(AddResponse{
		ID:             taskID,
		Status:         statusDone,
		HTTPStatusCode: resp.StatusCode,
		Headers:        resp.Header,
		Length:         int(resp.ContentLength),
	})
	if err != nil {
		logger.Error().Msgf("failed to add response for task %d, err= %v", taskID, err)

		s.changeStatusToError(taskID)
	}
	return
}

func (s *Service) changeStatusToError(taskID int) {
	err := s.repository.ChangeStatus(taskID, statusError)
	if err != nil {
		logger.Error().Err(err).Msgf("failed to change status of task %d to error", taskID)
		return
	}
	return
}
