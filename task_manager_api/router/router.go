package router

import (
	"github.com/gin-gonic/gin"
	"Go_Rest_API/task_manager_api/controllers"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()
	
	router.GET("/tasks" , controllers.GetAllTask)
	router.GET("/tasks/:id" , controllers.GetTaskById)
	router.PUT("/tasks/:id" , controllers.UpdateTask)
    router.DELETE("/tasks/:id" , controllers.DeleteTask)
	router.POST("/tasks" , controllers.CreateTask)

	return router
}