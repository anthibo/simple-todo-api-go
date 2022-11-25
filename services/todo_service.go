package services

import (
	"example/todo-with-goLang/db"
	"example/todo-with-goLang/dto"
	"example/todo-with-goLang/models"
)

type TodoService struct {
	todoRepository models.TodoRepository
}

// TODO: implement todo service struct and add these function to the struct
func (*TodoService) GetUserTasks(id int64) ([]*models.Todo, error) {
	tasks, err := todoRepository.GetUserTasks(id)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (*TodoService) AddTodo(todoInput *dto.TodoInput) (*models.Todo, error) {
	newTodo := models.Todo{Item: todoInput.Item, Completed: false}

	result := db.DB.Create(&newTodo)

	if result.Error != nil {
		return nil, result.Error
	}
	return &newTodo, nil
}

func (*TodoService) GetTodoById(id int64) (*models.Todo, error) {
	var todo *models.Todo

	result := db.DB.Where(id).Find(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}

func ToggleTodoStatus(id int64) (*models.Todo, error) {
	var todo *models.Todo
	fetchResult := db.DB.Where(id).Find(&todo)

	if fetchResult.Error != nil {
		return nil, fetchResult.Error
	}
	todo.Completed = !todo.Completed

	updateResult := db.DB.Save(&todo)

	if updateResult.Error != nil {
		return nil, fetchResult.Error
	}
	return todo, nil
}
