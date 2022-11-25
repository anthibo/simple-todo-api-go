package models

import (
	"gorm.io/gorm"
)

type todoOrm struct {
	db *gorm.DB
}

// TODO: add a relation between users and todo (one to many)
type Todo struct {
	gorm.Model
	Item      string `json:"item"`
	Completed bool   `json:"completed" gorm:"default:false"`
}

type TodoRepository interface {
	CreateTodo(task string) (*Todo, error)
	GetUserTasks(id int64) ([]*Todo, error)
	ToggleTodoStatus(id int64) (*Todo, error)
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	_ = db.AutoMigrate(&Todo{}) // builds table when enabled
	return &todoOrm{db}
}

func (*Todo) CreateTodo(todo *Todo) {

}

func (todoOrm *todoOrm) CreateTodo(task string) (*Todo, error) {
	newTodo := Todo{Item: task, Completed: false}

	result := todoOrm.db.Create(&newTodo)

	if result.Error != nil {
		return nil, result.Error
	}
	return &newTodo, nil
}

func (todoOrm *todoOrm) GetUserTasks(id int64) ([]*Todo, error) {
	//TODO: implement get user tasks
	var todo *Todo

	result := todoOrm.db.Where(id).Find(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return []*Todo{{}}, nil
}

func (todoOrm *todoOrm) ToggleTodoStatus(id int64) (*Todo, error) {
	var todo *Todo
	fetchResult := todoOrm.db.Where(id).Find(&todo)

	if fetchResult.Error != nil {
		return nil, fetchResult.Error
	}

	todo.Completed = !todo.Completed

	updateResult := todoOrm.db.Save(&todo)

	if updateResult.Error != nil {
		return nil, fetchResult.Error
	}

	return todo, nil
}
