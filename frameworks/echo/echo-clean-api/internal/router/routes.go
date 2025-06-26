package router

import (
	"database/sql"
	"echo-clean-api/internal/controller"
	"echo-clean-api/internal/repository"
	"echo-clean-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *sql.DB) {
	// Padrão de injeção de dependencias por camadas
	repo := repository.NewUserRepository(db)

	// Camada de negócio (também conhecida como Service)
	uc := usecase.NewUserUseCase(repo)

	// Handlers ~ Controllers -
	ctrl := controller.NewUserController(uc)

	e.POST("/register", ctrl.Register)
	e.POST("/login", ctrl.Login)

	// Request -> Controller ou Hendlers (Triagem)
	// UseCase OU Services (Camada de negócio)
	// Repository (Dados ~ DAO) -> Models / Dominios / Entidades

}
