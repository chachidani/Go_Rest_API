package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct{
	ID          string    `json:"id"`
 Title       string    `json:"title"`
 Description string    `json:"description"`
 DueDate     time.Time `json:"due_date"`
 Status      string    `json:"status"`
}

var tasks = []Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}


func main(){
	router := gin.Default()
	router.GET("/ping" , func(ctx *gin.Context){
		ctx.JSON(200 , gin.H{
			"message":"pong",
		})
	})
	router.GET("/tasks" , getTasks)
	router.GET("/tasks/:id" , getTaskById)
	router.PUT("/tasks/:id" , upDateTask)
    router.DELETE("/tasks/:id" , deletTask)
	router.POST("/tasks" , postTask)
	router.Run("localhost:8080")

}

func getTasks(c *gin.Context){
	c.IndentedJSON(http.StatusOK , tasks)
}
func getTaskById(c *gin.Context){
	id := c.Param("id")
	for _ , task := range tasks{
		if id == task.ID{
			c.IndentedJSON(http.StatusOK , task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound , gin.H{"error":"Task not found"})
}

func upDateTask(c *gin.Context){
	id := c.Param("id")
	var updatedTask Task

	if err := c.BindJSON(&updatedTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":err.Error()})
	}

	for i , task := range tasks{
		if id == task.ID{
			if updatedTask.Title != ""{
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != ""{
				tasks[i].Description = updatedTask.Description
			}

			c.IndentedJSON(http.StatusOK , gin.H{"message":"Task updated"})
			return 
			
		}

	}
 c.IndentedJSON(http.StatusNotFound , gin.H{"message":"Task not found"})

}


func deletTask(c *gin.Context){
	id := c.Param("id")
	for i , task := range tasks{
		if id == task.ID{
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK , gin.H{"message":"Task Removed"})
			return 
		}

	}

	c.IndentedJSON(http.StatusNotFound , gin.H{"message": "Task not found"})
}

func postTask(c *gin.Context){
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":err.Error()})
		return 
	}
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated , gin.H{"message":"Task Created"})


}