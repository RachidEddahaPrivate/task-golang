package service

import (
	"maps"
	"reflect"
	"sync"
	"task/internal/dto"
	"task/pkg/logger"
	"testing"
)

func TestNewService(t *testing.T) {
	logger.InitializeForTest()
	type args struct {
		repository repository
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateTask(t *testing.T) {
	logger.InitializeForTest()
	type fields struct {
		repository repository
	}
	type args struct {
		request dto.CreateTaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.CreateTaskResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
			}
			got, err := s.CreateTask(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetTask(t *testing.T) {
	logger.InitializeForTest()
	type fields struct {
		repository repository
	}
	type args struct {
		taskID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.GetTaskResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
			}
			got, err := s.GetTask(tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() got = %v, want %v", got, tt.want)
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
					URL:    "www.google.com",
				},
			},
			checkFunc: func(t *testing.T, task GetTask) {
				if task.Status != statusError {
					t.Errorf("task status = %v, want %v", task.Status, statusError)
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
