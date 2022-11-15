package services

import (
	"example/todo-with-goLang/model"
)

func GetTodos() ([]*model.Todo, error) {
	var todos []*model.Todo
	result := model.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func AddTodo(todoInput *model.TodoInput) (*model.Todo, error) {
	newTodo := model.Todo{Item: todoInput.Item, Completed: false}

	result := model.DB.Create(&newTodo)

	if result.Error != nil {
		return nil, result.Error
	}
	return &newTodo, nil
}

func GetTodoById(id int64) (*model.Todo, error) {
	var todo *model.Todo

	result := model.DB.Where(id).Find(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}

func ToggleTodoStatus(id int64) (*model.Todo, error) {
	var todo *model.Todo
	fetchResult := model.DB.Where(id).Find(&todo)

	if fetchResult.Error != nil {
		return nil, fetchResult.Error
	}
	todo.Completed = !todo.Completed

	updateResult := model.DB.Save(&todo)

	if updateResult.Error != nil {
		return nil, fetchResult.Error
	}
	return todo, nil
}
