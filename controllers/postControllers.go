package controllers

import (
	"GinModule/initializers"
	"GinModule/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func CreateTask(c *gin.Context){

	var Task models.Task

	c.Bind(&Task)

	task:= models.Task{ Description: Task.Description}

	result := initializers.DB.Create(&task)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error creating task",
		})
	}

	fmt.Println("Resultado:",*result)


	c.JSON(200, gin.H{
		"Task": task,
	})
}

func GetAllTasks(c *gin.Context){
	var tasks []models.Task

	result := initializers.DB.Find(&tasks)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error getting tasks",
		})
	}

	c.JSON(200, gin.H{
		"Tasks": tasks,
	})
}

func GetTask(c *gin.Context){

	id:= c.Param("id")


	var task models.Task
	result := initializers.DB.First(&task, id )
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error getting task",
		})
	}
	c.JSON(200, gin.H{
		"Task": task,
	})
}

func UpdateTask(c *gin.Context){
	id:= c.Param("id")

	var taskStruct models.Task

	var task models.Task

	c.Bind(&taskStruct)

	initializers.DB.First(&task, id)

	result := initializers.DB.Model(&task).Updates(&taskStruct)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error updating task",
		})
	}

	c.JSON(200, gin.H{"message": "Task updated"}) 
}

func DeleteTask(c *gin.Context){
	id:=c.Param("id")

	var task models.Task

	returns := initializers.DB.Delete(&task, id)

	if returns.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error deleting task",
		})
	}
	c.JSON(200, gin.H{"message": "Task deleted"})
}