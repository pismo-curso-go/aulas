package usecase

import "echo-clean-api/internal/repository"

type UserUseCase struct {
	Repo *repository.UserRepository
}

// "Construtor"
func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{Repo: repo}
}
