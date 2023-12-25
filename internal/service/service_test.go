package service

import (
	"github.com/go-test/deep"
	"maps"
	"net/http"
	"reflect"
	"sync"
	"task/internal/dto"
	"task/pkg/customerror"
	"task/pkg/logger"
	"task/pkg/models"
	"testing"
)

func TestService_CreateTask(t *testing.T) {
	logger.InitializeForTest()

	// Important: this test mut be executed in cascade
	// because the repository is a global variable
	s := NewService(NewRepository())
	type args struct {
		request dto.CreateTaskRequest
	}

	tests := []struct {
		name      string
		args      args
		want      dto.CreateTaskResponse
		wantErr   bool
		checkFunc func(t *testing.T, s *Service)
	}{
		{
			name: "create first task",
			args: args{
				request: dto.CreateTaskRequest{
					Method: "GET",
					URL:    "http://www.google.com",
				},
			},
			want: dto.CreateTaskResponse{
				ID: 1,
			},
		},
		{
			name: "create second task",
			args: args{
				request: dto.CreateTaskRequest{
					Method: "GET",
					URL:    "http://www.google.com",
				},
			},
			want: dto.CreateTaskResponse{
				ID: 2,
			},
		},
		{
			name: "create third task",
			args: args{
				request: dto.CreateTaskRequest{
					Method: "GET",
					URL:    "http://www.google.com",
				},
			},
			want: dto.CreateTaskResponse{
				ID: 3,
			},
		},
		{
			name: "stress test",
			args: args{
				request: dto.CreateTaskRequest{
					Method: "GET",
					URL:    "http://www.google.com",
				},
			},
			want: dto.CreateTaskResponse{
				ID: 4,
			},
			checkFunc: func(t *testing.T, s *Service) {
				sg := sync.WaitGroup{}
				for i := 0; i < 100; i++ {
					sg.Add(1)
					go func() {
						defer sg.Done()
						_, _ = s.CreateTask(dto.CreateTaskRequest{
							Method: "GET",
							URL:    "googlassadse",
						})
					}()
				}
				sg.Wait()
				_, err := s.GetTask(104)
				if err != nil {
					t.Errorf("error getting task with id %d", 104)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateTask(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
			if tt.checkFunc != nil {
				tt.checkFunc(t, s)
			}
		})
	}
}

func TestService_GetTask(t *testing.T) {
	logger.InitializeForTest()
	initialMap := map[int]Task{
		1: {
			Status: statusNew,
		},
		2: {
			Status: statusInProcess,
		},
		3: {
			Status:         statusDone,
			HTTPStatusCode: http.StatusOK,
			Headers: map[string][]string{
				"Content-Type": {"application/json,application/xml"},
				"Accept":       {"application/json"},
			},
			Length: 20,
			Error:  "",
		},
		4: {
			Status: statusError,
			Error:  "incorrect url",
		},
	}
	s := NewService(&Repository{
		mu:                sync.Mutex{},
		tasks:             initialMap,
		lastTaskIDCreated: len(initialMap),
	})
	type args struct {
		taskID int
	}
	tests := []struct {
		name    string
		args    args
		want    dto.GetTaskResponse
		wantErr error
	}{
		{
			name: "response with correct status",
			args: args{
				taskID: 1,
			},
			want: dto.GetTaskResponse{
				ID:     1,
				Status: statusNew,
			},
		},
		{
			name: "response with in process status",
			args: args{
				taskID: 2,
			},
			want: dto.GetTaskResponse{
				ID:     2,
				Status: statusInProcess,
			},
		},
		{
			name: "response with multiple headers",
			args: args{
				taskID: 3,
			},
			want: dto.GetTaskResponse{
				ID:             3,
				Status:         statusDone,
				HTTPStatusCode: http.StatusOK,
				Headers: map[string]string{
					"Content-Type": "application/json,application/xml",
					"Accept":       "application/json",
				},
			},
		},
		{
			name: "response with error",
			args: args{
				taskID: 4,
			},
			want: dto.GetTaskResponse{
				ID:     4,
				Status: statusError,
			},
		},
		{
			name: "test with incorrect taskID",
			args: args{
				taskID: 5,
			},
			want: dto.GetTaskResponse{},
			wantErr: customerror.NewI18nErrorWithParams(models.TaskIDNotFoundError, map[string]interface{}{
				"taskId": 5,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetTask(tt.args.taskID)
			if diff := deep.Equal(err, tt.wantErr); diff != nil {
				t.Errorf("GetTask() error = %v, wantErr %v, diff %v", err, tt.wantErr, diff)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("GetTask() got = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}

func TestService_makeRequest(t *testing.T) {
	logger.InitializeForTest()

	initialMap := map[int]Task{
		1: {
			Status: statusNew,
		},
	}
	type fields struct {
		repository repository
	}
	type args struct {
		taskID  int
		request dto.CreateTaskRequest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		checkFunc func(t *testing.T, task GetTask)
	}{
		{
			name: "test with incorrect method",
			fields: fields{repository: &Repository{
				mu:                sync.Mutex{},
				tasks:             maps.Clone(initialMap),
				lastTaskIDCreated: len(initialMap),
			}},
			args: args{
				taskID: 1,
				request: dto.CreateTaskRequest{
					Method: "try",
					URL:    "http://www.google.com",
				},
			},
			checkFunc: func(t *testing.T, task GetTask) {
				if task.Status != statusDone {
					t.Errorf("task status = %v, want %v", task.Status, statusDone)
				}
				if task.HTTPStatusCode != http.StatusMethodNotAllowed {
					t.Errorf("task HTTPStatusCode = %v, want %v", task.HTTPStatusCode, http.StatusMethodNotAllowed)
				}
			},
		},
		{
			name: "test with incorrect url",
			fields: fields{repository: &Repository{
				mu:                sync.Mutex{},
				tasks:             maps.Clone(initialMap),
				lastTaskIDCreated: len(initialMap),
			}},
			args: args{
				taskID: 1,
				request: dto.CreateTaskRequest{
					Method: "GET",
					URL:    "http://wwwww.google.com",
				},
			},
			checkFunc: func(t *testing.T, task GetTask) {
				if task.Status != statusError {
					t.Errorf("task status = %v, want %v", task.Status, statusError)
				}
			},
		},
		{
			name: "test with correct request",
			fields: fields{repository: &Repository{
				mu:                sync.Mutex{},
				tasks:             maps.Clone(initialMap),
				lastTaskIDCreated: len(initialMap),
			}},
			args: args{
				taskID: 1,
				request: dto.CreateTaskRequest{
					Method: "GET",
					URL:    "http://www.google.com",
				},
			},
			checkFunc: func(t *testing.T, task GetTask) {
				if task.Status != statusDone {
					t.Errorf("task status = %v, want %v", task.Status, statusDone)
				}
				if task.HTTPStatusCode != http.StatusOK {
					t.Errorf("task HTTPStatusCode = %v, want %v", task.HTTPStatusCode, http.StatusOK)
				}
				if task.Error != "" {
					t.Errorf("task Error = %v, want %v", task.Error, "")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
			}
			s.makeRequest(tt.args.taskID, tt.args.request)
			if tt.checkFunc != nil {
				task, err := s.repository.GetTask(tt.args.taskID)
				if err != nil {
					t.Errorf("error getting task with id %d", tt.args.taskID)
				}
				tt.checkFunc(t, task)
			}
		})
	}
}

func Test_transformHeaders(t *testing.T) {
	logger.InitializeForTest()
	type args struct {
		headers map[string][]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[string]string
	}{
		{
			name: "test without multiple values",
			args: args{
				headers: map[string][]string{
					"Content-Type": {"application/json"},
					"Accept":       {"application/json"},
				}},
			wantResult: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/json",
			},
		},
		{
			name: "test with multiple values",
			args: args{
				headers: map[string][]string{
					"Content-Type": {"application/json", "application/xml"},
					"Accept":       {"application/json"},
				}},
			wantResult: map[string]string{
				"Content-Type": "application/json,application/xml",
				"Accept":       "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := transformHeaders(tt.args.headers); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("transformHeaders() = %v, want %v", gotResult, tt.wantResult)
			}

		})
	}
}
