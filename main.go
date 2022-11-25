package main

import (
	"example/todo-with-goLang/db"

	"github.com/gin-gonic/gin"
)

// TODO: Add authentication middleware with JWT or PASETO
func main() {
	router := gin.Default()
	//API Endpoints
	// router.GET("/todos", controllers.GetTodos)
	// router.GET("/todos/:id", controllers.GetTodo)
	// router.POST("/todos", controllers.AddTodo)
	// router.PATCH("/todos/:id", controllers.ToggleTodoStatus)

	db.ConnectDatabase()
	router.Run("localhost:9090")
}
