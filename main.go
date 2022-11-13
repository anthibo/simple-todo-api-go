package main

import (
	controllers "example/todo-with-goLang/controller"
	"example/todo-with-goLang/model"

	"github.com/gin-gonic/gin"
)

var todos = []model.Todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Workout ", Completed: false},
}

func main() {
	router := gin.Default()
	//API Endpoints
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodo)
	router.POST("/todos", controllers.AddTodos)
	router.PATCH("/todos/:id", controllers.ToggleTodoStatus)

	model.ConnectDatabase()
	router.Run("localhost:9090")
}
