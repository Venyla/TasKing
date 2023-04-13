package controller

import (
	"container/list"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"net/http"
	data "webservice/data/dataModels"
)

var taskHistories = []data.TaskHistory{
	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("db9c48a8-f491-4120-a11f-21ec27335c2a"), CreatedBy: "Vina"},
	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("db9c48a8-f491-4120-a11f-21ec27335c2a"), CreatedBy: "Vina"},
	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("db9c48a8-f491-4120-a11f-21ec27335c2a"), CreatedBy: "Vanessa"},
	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("a3958663-e7c4-4b19-a43b-b3a06fec33b5"), CreatedBy: "Vanessa"},
	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("a3958663-e7c4-4b19-a43b-b3a06fec33b5"), CreatedBy: "Lukas"},
}

func GetTaskHistories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, taskHistories)
}

func GetTaskHistoryByTaskId(c *gin.Context) {
	id := uuid.Must(uuid.Parse(c.Param("id")))
	historyOfTask := list.New()

	for _, h := range taskHistories {
		if h.TaskId == id {
			historyOfTask.PushFront(h)
		}
	}

	c.IndentedJSON(http.StatusOK, historyOfTask)
}

func PostTaskHistory(c *gin.Context) {
	var newTaskHistory data.TaskHistory

	if err := c.BindJSON(&newTaskHistory); err != nil {
		return
	}

	taskHistories = append(taskHistories, newTaskHistory)
	c.IndentedJSON(http.StatusCreated, newTaskHistory)
}
