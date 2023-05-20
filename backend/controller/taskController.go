package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"webservice/service"
)

type TaskController interface {
	GetTasks(context *gin.Context)
	GetTaskById(context *gin.Context)
}

type taskController struct {
	taskService service.TaskService
}

func NewTaskController(taskServ service.TaskService) TaskController {
	return &taskController{
		taskService: taskServ,
	}
}

func (c *taskController) GetTasks(context *gin.Context) {
	log.Println("get tasks")
	context.IndentedJSON(http.StatusOK, c.taskService.GetAllTasks())
}

func (c *taskController) GetTaskById(context *gin.Context) {
	id := uuid.Must(uuid.Parse(context.Param("id")))
	context.IndentedJSON(http.StatusOK, c.taskService.GetTask(id))
}
