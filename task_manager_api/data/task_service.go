package data

import (
	"time"

	"Go_Rest_API/task_manager_api/models"
)

var tasks = []models.Task{
    {ID: 1, Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: 2, Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: 3, Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

var nextId = 1

func GetAllTask()[]models.Task{
	return tasks
}

func GetTaskById(id int)(models.Task , bool){
	for _ , task := range tasks{
		if task.ID == id{
			return task , true
		}
	}
	return models.Task{} ,  false
}


func CreateTask(task models.Task)models.Task{
	task.ID = nextId
	nextId ++
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id int , updatedTask models.Task)(models.Task , bool){
	for i,task := range tasks{
		if task.ID == id{
			updatedTask.ID = id
			tasks[i] = updatedTask
			return  updatedTask , true
		}
	}
	return models.Task{} , false 

}

func DeleteTask(id int)bool{
	for i, task := range tasks{
		if task.ID == id{
			tasks = append(tasks[:i],tasks[i+1:]... )
			return true
		}
	}
	return false
}
