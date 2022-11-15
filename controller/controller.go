package controllers

import (
	"example/todo-with-goLang/model"
	"example/todo-with-goLang/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodos(context *gin.Context) {

	todos, error := services.GetTodos()
	if error != nil {
		context.IndentedJSON(http.StatusInternalServerError, error)
		return
	}
	context.IndentedJSON(http.StatusOK, todos)
}

func AddTodo(context *gin.Context) {
	var newTodo model.TodoInput

	if error := context.BindJSON(&newTodo); error != nil {
		context.IndentedJSON(http.StatusInternalServerError, error)
		return
	}
	createdTodo, error := services.AddTodo(&newTodo)
	if error != nil {
		context.IndentedJSON(http.StatusInternalServerError, error)
		return
	}
	context.IndentedJSON(http.StatusOK, createdTodo)
}

func GetTodo(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	todo, error := services.GetTodoById(id)
	if error != nil {
		context.IndentedJSON(http.StatusInternalServerError, error)
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func ToggleTodoStatus(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	todo, error := services.ToggleTodoStatus(id)
	if error != nil {
		context.IndentedJSON(http.StatusInternalServerError, error)
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}
