package service

import (
	"github.com/google/uuid"
	datamodels "webservice/data/dataModels"
	"webservice/data/repositories"
)

type TaskHistoryService interface {
	GetAllTaskHistories() []datamodels.TaskHistory
	GetTaskHistoryByTaskId(taskHistoryId uuid.UUID) []datamodels.TaskHistory
	InsertTaskHistory(taskId uuid.UUID, userName string) datamodels.TaskHistory
}

type taskHistoryService struct {
	taskHistoryRepository repositories.TaskHistoryRepository
}

func NewTaskHistoryService(taskHistoryRepo repositories.TaskHistoryRepository) TaskHistoryService {
	return &taskHistoryService{
		taskHistoryRepository: taskHistoryRepo,
	}
}

func (service *taskHistoryService) GetAllTaskHistories() []datamodels.TaskHistory {
	return *service.taskHistoryRepository.GetAllTaskHistories()
}

func (service *taskHistoryService) GetTaskHistoryByTaskId(taskId uuid.UUID) []datamodels.TaskHistory {
	return *service.taskHistoryRepository.FindTaskHistoryByTaskId(taskId)
}

func (service *taskHistoryService) InsertTaskHistory(taskId uuid.UUID, userName string) datamodels.TaskHistory {
	return *service.taskHistoryRepository.Insert(taskId, userName)
}
