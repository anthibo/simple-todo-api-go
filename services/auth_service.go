package services

import (
	"example/todo-with-goLang/dto"
	"example/todo-with-goLang/models"
)

// TODO: implement auth service
type AuthService struct {
	userRepository models.UserRepository
}

func (*AuthService) Login(in *dto.UserLogin) {

}

func (*AuthService) Register(in *dto.UserRegister) {

}
