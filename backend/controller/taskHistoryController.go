package controller

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"net/http"
	data "webservice/data/dataModels"
	"webservice/service"
)

type TaskHistoryController interface {
	GetTaskHistories(context *gin.Context)
	GetTaskHistoryByTaskId(context *gin.Context)
	GetRankingsByTaskId(context *gin.Context)
	PostTaskHistory(context *gin.Context)
}

type taskHistoryController struct {
	taskHistoryService service.TaskHistoryService
}

func NewTaskHistoryController(taskHistoryServ service.TaskHistoryService) TaskHistoryController {
	return &taskHistoryController{
		taskHistoryService: taskHistoryServ,
	}
}

func (c *taskHistoryController) GetTaskHistories(context *gin.Context) {
	//c.IndentedJSON(http.StatusOK, taskHistories)
	context.IndentedJSON(http.StatusOK, c.taskHistoryService.GetAllTaskHistories())
}

func (c *taskHistoryController) GetTaskHistoryByTaskId(context *gin.Context) {
	id := uuid.Must(uuid.Parse(context.Param("id")))

	context.IndentedJSON(http.StatusOK, c.taskHistoryService.GetTaskHistoryByTaskId(id))
}

func (c *taskHistoryController) PostTaskHistory(context *gin.Context) {
	var newTaskHistory data.TaskHistory

	if err := context.BindJSON(&newTaskHistory); err != nil {
		return
	}

	insertedtaskHistory := c.taskHistoryService.InsertTaskHistory(newTaskHistory.TaskId, newTaskHistory.CreatedBy)

	context.IndentedJSON(http.StatusCreated, insertedtaskHistory)
}

func (c *taskHistoryController) GetRankingsByTaskId(context *gin.Context) {
	id := uuid.Must(uuid.Parse(context.Param("id")))

	rankings := make(map[string]int)
	taskHistories := c.taskHistoryService.GetTaskHistoryByTaskId(id)

	for _, h := range taskHistories {
		amount, exists := rankings[h.CreatedBy]
		if exists {
			rankings[h.CreatedBy] = amount + 1
		} else {
			rankings[h.CreatedBy] = 1
		}
	}

	context.IndentedJSON(http.StatusOK, rankings)
}
