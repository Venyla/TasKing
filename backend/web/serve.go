package web

import (
	"fmt"
	"webservice/controller"
	"webservice/data/repositories"
	"webservice/service"

	"github.com/gin-gonic/gin"
)

var (
	taskRepo              = repositories.TaskRepository{}
	taskHistoryRepo       = repositories.TaskHistoryRepository{}
	taskService           = service.NewTaskService(taskRepo)
	taskHistoryService    = service.NewTaskHistoryService(taskHistoryRepo)
	taskController        = controller.NewTaskController(taskService)
	taskHistoryController = controller.NewTaskHistoryController(taskHistoryService)
)

func Serve() {
	server := gin.Default()

	// Used for simpler developing - can be removed later or adjusted only for public api route
	//server.Use(CORS(), RequestCancelRecover())

	endpoints := server.Group("/api")
	{
		endpoints.GET("/tasks", taskController.GetTasks)
		endpoints.GET("/tasks/:id", taskController.GetTaskById)
		endpoints.GET("/tasks/:id/rankings", taskHistoryController.GetRankingsByTaskId)

		endpoints.GET("/history", taskHistoryController.GetTaskHistories)
		endpoints.GET("/history/:id", taskHistoryController.GetTaskHistoryByTaskId)
		endpoints.POST("/history", taskHistoryController.PostTaskHistory)
	}

	server.Run(":8080")
}

func RequestCancelRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("A problem occured")
				c.Request.Context().Done()
			}
		}()
		c.Next()
	}
}
