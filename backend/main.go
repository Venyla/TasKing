package main

import (
	"github.com/gin-gonic/gin"
	"webservice/controller"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default()) // Allow all

	router.GET("/api/tasks", controller.GetTasks)
	router.GET("/api/tasks/:id", controller.GetTaskById)
	router.GET("/api/tasks/:id/rankings", controller.GetRankingsByTaskId)
	router.POST("/api/tasks", controller.PostTask)

	router.GET("/api/history", controller.GetTaskHistories)
	router.GET("/api/history/:id", controller.GetTaskHistories)
	router.POST("/api/history", controller.PostTaskHistory)

	router.Run("localhost:8080")
}
