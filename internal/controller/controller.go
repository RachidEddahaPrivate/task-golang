package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"task/internal/dto"
	"task/pkg/logger"
	"task/pkg/utils"
)

type Controller struct {
	service service
}

type service interface {
	GetTask(int) (dto.GetTaskResponse, error)
	CreateTask(dto.CreateTaskRequest) (dto.CreateTaskResponse, error)
}

func NewController(service service) *Controller {
	if service == nil {
		panic(service)
	}
	return &Controller{service: service}
}

func (c *Controller) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/task") // you could version the api here (e.g. api/v1/task) and use a middleware to check the token
	g.GET("/:taskId", c.getTask)
	g.POST("", c.createTask)
}

func (c *Controller) createTask(context echo.Context) error {
	request := dto.CreateTaskRequest{}
	err := context.Bind(&request)
	if err != nil {
		return err
	}
	err = context.Validate(request)
	if err != nil {
		return err
	}
	logger.Debug().Msgf("recived request: %+v", request)
	createdTask, err := c.service.CreateTask(request)
	if err != nil {
		return err
	}
	logger.Debug().Msgf("created task %d", createdTask.ID)
	return context.JSON(http.StatusOK, createdTask)
}

func (c *Controller) getTask(context echo.Context) error {
	// I assume the task id is a number, because in a real application I would use a database,
	// where the id is provided and it is a number

	taskID, err := utils.GetEchoParamToInt(context, "taskId")
	if err != nil {
		return err
	}

	logger.Debug().Msgf("caller want task %d", taskID)

	task, err := c.service.GetTask(taskID)
	if err != nil {
		return err
	}
	logger.Debug().Msgf("the task %d is : %+v", taskID, task)
	return context.JSON(http.StatusOK, task)
}
