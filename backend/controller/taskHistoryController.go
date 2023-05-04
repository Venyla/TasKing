package controller

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"net/http"
	data "webservice/data/dataModels"
	"webservice/service"
)

//var taskHistories = []data.TaskHistory{
//	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("db9c48a8-f491-4120-a11f-21ec27335c2a"), CreatedBy: "Vina"},
//	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("db9c48a8-f491-4120-a11f-21ec27335c2a"), CreatedBy: "Vina"},
//	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("db9c48a8-f491-4120-a11f-21ec27335c2a"), CreatedBy: "Vanessa"},
//	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("a3958663-e7c4-4b19-a43b-b3a06fec33b5"), CreatedBy: "Vanessa"},
//	{TaskHistoryId: uuid.New(), TaskId: uuid.MustParse("a3958663-e7c4-4b19-a43b-b3a06fec33b5"), CreatedBy: "Lukas"},
//}

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
	//historyOfTask := list.New()
	//
	//for _, h := range taskHistories {
	//	if h.TaskId == id {
	//		historyOfTask.PushFront(h)
	//	}
	//}

	context.IndentedJSON(http.StatusOK, c.taskHistoryService.GetTaskHistoryByTaskId(id))
}

func (c *taskHistoryController) PostTaskHistory(context *gin.Context) {
	var newTaskHistory data.TaskHistory

	if err := context.BindJSON(&newTaskHistory); err != nil {
		return
	}

	//taskHistories = append(taskHistories, newTaskHistory)

	taskHistory := c.taskHistoryService.InsertTaskHistory(newTaskHistory.TaskId, newTaskHistory.CreatedBy)

	context.IndentedJSON(http.StatusCreated, taskHistory)
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

//func getRankings(taskId uuid.UUID) map[string]int {
//	var rankings = make(map[string]int)
//
//	for _, h := range taskHistories {
//
//		if h.TaskId == taskId {
//			amount, exists := rankings[h.CreatedBy]
//			if exists {
//				rankings[h.CreatedBy] = amount + 1
//			} else {
//				rankings[h.CreatedBy] = 1
//			}
//		}
//	}
//
//	return rankings
//}
