package models

import (
	"example/todo-with-goLang/dto"

	"gorm.io/gorm"
)

type userOrm struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique:true"`
	Password string `json:"password" gorm:"default:false"`
}

type UserRepository interface {
	CreateUser(userInput dto.UserRegister) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	_ = db.AutoMigrate(&User{}) // builds table when enabled
	return &userOrm{db}
}

func (userOrm *userOrm) CreateUser(userInput dto.UserRegister) (*User, error) {
	newUser := User{Username: userInput.Username, Password: userInput.Password}

	result := userOrm.db.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}

func (userOrm *userOrm) GetUserByUsername(username string) (*User, error) {
	var user *User

	result := userOrm.db.Where(username).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
