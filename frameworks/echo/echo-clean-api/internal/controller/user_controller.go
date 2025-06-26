package controller

import (
	"echo-clean-api/internal/model"
	"echo-clean-api/internal/usecase"
	"echo-clean-api/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UseCase *usecase.UserUseCase
}

// "Construtor"
func NewUserController(uc *usecase.UserUseCase) *UserController {
	return &UserController{UseCase: uc}
}

// Triagem e resposta para o cliente.
func (ctrl *UserController) Register(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil || user.Email == "" || user.Password == "" || user.Name == "" {
		return c.JSON(utils.ErrValidationFailed.Code, utils.ErrValidationFailed)
	}

	err := ctrl.UseCase.Register(&user)
	if err != nil {
		if appErr, ok := err.(utils.AppError); ok {
			return c.JSON(appErr.Code, appErr)
		}
		return c.JSON(utils.ErrInternalError.Code, utils.ErrInternalError)
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Usuário registrado com sucesso!"})
}

func (ctrl *UserController) Login(c echo.Context) error {

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil || credentials.Email == "" || credentials.Password == "" {
		return c.JSON(utils.ErrValidationFailed.Code, utils.ErrValidationFailed)
	}

	token, err := ctrl.UseCase.Login(credentials.Email, credentials.Password)

	if err != nil {
		if appErr, ok := err.(utils.AppError); ok {
			return c.JSON(appErr.Code, appErr)
		}
		return c.JSON(utils.ErrInternalError.Code, utils.ErrInternalError)
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
