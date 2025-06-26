package usecase

import (
	"echo-clean-api/internal/model"
	"echo-clean-api/internal/repository"
	"echo-clean-api/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	Repo *repository.UserRepository
}

// "Construtor"
func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{Repo: repo}
}

func (uc *UserUseCase) Register(user *model.User) error {

	existing, _ := uc.Repo.GetByEmail(user.Email)
	if existing != nil {
		return utils.ErrUserAlreadyExists
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	return uc.Repo.Create(user)
}

func (uc *UserUseCase) Login(email, password string) (string, error) {

	user, _ := uc.Repo.GetByEmail(email)
	if user == nil {
		return "", utils.ErrInvalidCredentials
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", utils.ErrInvalidCredentials
	}

	token, _ := utils.GenerateJWT(user.Email)
	return token, nil

}
