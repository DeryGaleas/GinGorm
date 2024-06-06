package main

import (
	"GinModule/controllers"
	"GinModule/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r:= gin.Default()

	r.POST("/create", controllers.CreateTask)
	r.GET("/tasks", controllers.GetAllTasks)
	r.GET("/task/:id", controllers.GetTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	r.Run()
	
}