package controller

import "echo-clean-api/internal/usecase"

type UserController struct {
	UseCase *usecase.UserUseCase
}

// "Construtor"
func NewUserController(uc *usecase.UserUseCase) *UserController {
	return &UserController{UseCase: uc}
}
