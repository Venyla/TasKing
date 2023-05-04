package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webservice/controller"
	"webservice/data/repositories"
	"webservice/service"
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

	server.Group("/api")
	{
		server.GET("/api/tasks", taskController.GetTasks)
		server.GET("/api/tasks/:id", taskController.GetTaskById)
		server.GET("/api/tasks/:id/rankings", taskHistoryController.GetRankingsByTaskId)
		//server.POST("/api/tasks", taskController.PostTask)

		server.GET("/api/history", taskHistoryController.GetTaskHistories)
		server.GET("/api/history/:id", taskHistoryController.GetTaskHistories)
		server.POST("/api/history", taskHistoryController.PostTaskHistory)
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
