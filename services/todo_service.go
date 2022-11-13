package services

import (
	"example/todo-with-goLang/model"

	"github.com/gin-gonic/gin"
)

func GetTodos() ([]*model.Todo, error) {
}

func AddTodo(newTodo *model.Todo) (*model.Todo, error) {
	model.DB.Create(newTodo)
}

func GetTodo() (*model.Todo, error) {
}

func ToggleTodoStatus(context *gin.Context) (*model.Todo, error) {

}

func GetTodoById(id string) (*model.Todo, error) {

}
