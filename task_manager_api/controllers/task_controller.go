package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"Go_Rest_API/task_manager_api/data"
	"Go_Rest_API/task_manager_api/models"
)

func GetAllTask(c *gin.Context){
	tasks:= data.GetAllTask()
	c.IndentedJSON(http.StatusOK , tasks)
}

func GetTaskById(c *gin.Context){
	id ,_ := strconv.Atoi(c.Param("id"))
	task , found := data.GetTaskById(id)
	if !found {
		c.IndentedJSON(http.StatusNotFound , gin.H{"error":"Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK , task)
}

func CreateTask(c *gin.Context){
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":"Invalid input"})
		return 
	}
	createdTask := data.CreateTask(task)
	c.IndentedJSON(http.StatusCreated , createdTask)
}

func UpdateTask(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":"Invalid input"})
		return 
	}
	task , updated := data.UpdateTask(id , updatedTask)

	if !updated{
		c.IndentedJSON(http.StatusNotFound , gin.H{"error":"Task Not Found"})
		return 
	}
	c.IndentedJSON(http.StatusOK , task)
}

func DeleteTask(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	deleted := data.DeleteTask(id)
	if !deleted{
		c.IndentedJSON(http.StatusNotFound , gin.H{"error":"Task not Found"})
		return 
	}
	c.IndentedJSON(http.StatusOK , gin.H{"message":"Task Deleted"})
}