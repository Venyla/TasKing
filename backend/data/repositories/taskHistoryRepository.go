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

	err := dbConnection.Model(&datamodels.TaskHistory{TaskId: taskId}).
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
	_, creationError := dbConnection.Model(&datamodels.TaskHistory{
		TaskHistoryId: taskHistoryId,
		CreatedBy:     userName,
		CreatedOn:     time.Now().UTC(),
	}).Insert()

	if creationError != nil {
		transaction.Rollback()
		panic(creationError)
	}

	taskHistory := &datamodels.TaskHistory{}
	findError := dbConnection.Model(taskHistory).
		Where("TaskHistoryId = ?", taskHistoryId).
		Select()

	if findError != nil {
		transaction.Rollback()
		panic(findError)
	}

	transaction.Commit()

	return taskHistory
}
