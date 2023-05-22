package repositories

import (
	"github.com/google/uuid"
	"time"
	"webservice/data"
	datamodels "webservice/data/dataModels"
)

type TaskHistoryRepository struct{}

func (r TaskHistoryRepository) GetAllTaskHistories() *[]datamodels.TaskHistory {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	taskHistories := new([]datamodels.TaskHistory)
	err := dbConnection.Model(&datamodels.TaskHistory{}).
		Select(taskHistories)
	if err != nil {
		return nil
	}

	return taskHistories
}

func (r TaskHistoryRepository) FindTaskHistoryByTaskId(taskId uuid.UUID) *[]datamodels.TaskHistory {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	taskHistories := new([]datamodels.TaskHistory)

	err := dbConnection.Model(&datamodels.TaskHistory{}).
		Where("task_id = ?", taskId).
		Select(taskHistories)

	if err != nil {
		return nil
	}

	return taskHistories
}

func (r TaskHistoryRepository) Insert(taskId uuid.UUID, userName string) *datamodels.TaskHistory {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	transaction, transactionError := dbConnection.Begin()
	if transactionError != nil {
		panic(transactionError)
	}

	taskHistoryId := uuid.New()
	taskHistory := datamodels.TaskHistory{
		TaskHistoryId: taskHistoryId,
		TaskId:        taskId,
		CreatedBy:     userName,
		CreatedOn:     time.Now(),
	}

	_, creationError := dbConnection.Model(&taskHistory).Insert()

	if creationError != nil {
		transaction.Rollback()
		panic(creationError)
	}

	insertedTaskHistory := new(datamodels.TaskHistory)
	findError := dbConnection.Model(&datamodels.TaskHistory{}).
		Where("task_history_id = ?", taskHistoryId).
		Select(insertedTaskHistory)

	if findError != nil {
		transaction.Rollback()
		panic(findError)
	}

	transaction.Commit()

	return insertedTaskHistory
}
