package service

import (
	"fmt"
	"sync"
	"task/pkg/customerror"
	"task/pkg/logger"
	"task/pkg/models"
)

type Repository struct {
	mu                sync.Mutex
	tasks             map[int]Task
	lastTaskIDCreated int
}

func NewRepository() *Repository {
	return &Repository{
		mu:                sync.Mutex{},
		tasks:             make(map[int]Task),
		lastTaskIDCreated: 1, // start with value 1 for convenience
	}
}

// AddTask adds a task to the repository and returns the ID of the task
// in this case there are no errors, but the function signature was thought to be able to return an error
// in case of using a database
func (r *Repository) AddTask(status string) (result int, err error) {
	// for concurrent access, it's critical to lock the map when accessing it to read or write
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[r.lastTaskIDCreated] = Task{
		Status: status,
	}
	result = r.lastTaskIDCreated
	r.lastTaskIDCreated++

	return result, nil
}

func (r *Repository) GetTask(ID int) (GetTask, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	result, ok := r.tasks[ID]
	if !ok { // you can also check if the id is greater than the lastTaskIDCreated
		// codified error
		return GetTask{}, customerror.NewI18nErrorWithParams(models.TaskIDNotFoundError, map[string]interface{}{
			"taskId": ID,
		})
	}
	return GetTask{ID: ID, Task: result}, nil
}

func (r *Repository) ChangeStatus(ID int, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	task, ok := r.tasks[ID]
	if !ok {
		logger.Error().Msgf("task with ID %d not found", ID)
		return fmt.Errorf("task with ID %d not found", ID)
	}
	task.Status = status
	r.tasks[ID] = task
	return nil
}

func (r *Repository) AddResponse(response AddResponse) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	task, ok := r.tasks[response.ID]
	if !ok {
		logger.Error().Msgf("task with ID %d not found", response.ID)
		return fmt.Errorf("task with ID %d not found", response.ID)
	}
	task.HTTPStatusCode = response.HTTPStatusCode
	task.Headers = response.Headers
	task.Length = response.Length
	task.Status = response.Status
	r.tasks[response.ID] = task
	return nil
}

func (r *Repository) ChangeStatusInError(ID int, status string, err string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	task, ok := r.tasks[ID]
	if !ok {
		logger.Error().Msgf("task with ID %d not found", ID)
		return fmt.Errorf("task with ID %d not found", ID)
	}
	task.Status = status
	task.Error = err
	r.tasks[ID] = task
	return nil
}
