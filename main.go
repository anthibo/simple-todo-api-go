package main

import (
	controllers "example/todo-with-goLang/controller"
	"example/todo-with-goLang/model"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//API Endpoints
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodo)
	router.POST("/todos", controllers.AddTodo)
	// router.PATCH("/todos/:id", controllers.ToggleTodoStatus)

	model.ConnectDatabase()
	router.Run("localhost:9090")
}
