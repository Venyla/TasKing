package service

import (
	"github.com/google/uuid"
	datamodels "webservice/data/dataModels"
	"webservice/data/repositories"
)

type TaskService interface {
	GetAllTasks() []datamodels.Task
	GetTask(taskId uuid.UUID) datamodels.Task
}

type taskService struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) TaskService {
	return &taskService{
		taskRepository: taskRepo,
	}
}

func (service *taskService) GetAllTasks() []datamodels.Task {
	return *service.taskRepository.GetAll()
}

func (service *taskService) GetTask(taskId uuid.UUID) datamodels.Task {
	return *service.taskRepository.GetTask(taskId)
}
