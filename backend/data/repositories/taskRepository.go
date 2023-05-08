package repositories

import (
	"github.com/google/uuid"
	"webservice/data"
	datamodels "webservice/data/dataModels"
)

type TaskRepository struct{}

func (r TaskRepository) GetAll() *[]datamodels.Task {
	return getAllTasks()
}

func (r TaskRepository) GetTask(taskId uuid.UUID) *datamodels.Task {
	return findTask(taskId)
}

func findTask(taskId uuid.UUID) *datamodels.Task {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	task := &datamodels.Task{TaskId: taskId}
	err := dbConnection.Model(task).Select()

	if err != nil {
		return nil
	}

	return task
}

func getAllTasks() *[]datamodels.Task {
	dbConnection := data.GetDbConnection()
	defer dbConnection.Close()

	tasks := new([]datamodels.Task)
	err := dbConnection.Model(&datamodels.Task{}).
		Select(tasks)

	if err != nil {
		return nil
	}

	return tasks
}
