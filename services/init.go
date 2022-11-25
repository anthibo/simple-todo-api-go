package services

import (
	"example/todo-with-goLang/db"
	"example/todo-with-goLang/models"
)

// inject repositories/services here

var userRepository = models.NewUserRepository(db.DB)

var todoRepository = models.NewTodoRepository(db.DB)

func NewTodoService() *TodoService {
	return &TodoService{
		todoRepository,
	}
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepository,
	}
}
