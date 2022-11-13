package controllers

import (
	"example/todo-with-goLang/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "")
}

func AddTodos(context *gin.Context) {
	var newTodo model.Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	// todos = append(todos, newTodo)
	// context.IndentedJSON(http.StatusOK, todos)
}

func GetTodo(context *gin.Context) {
	// id := context.Param("id")
	// todo, err := getTodoById(id)

	// if err != nil {
	// 	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	// }
	// context.IndentedJSON(http.StatusOK, todo)
}

func ToggleTodoStatus(context *gin.Context) {
	// id := context.Param("id")
	// todo, err := getTodoById(id)

	// if err != nil {
	// 	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	// }

	// todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, "todo")
}
