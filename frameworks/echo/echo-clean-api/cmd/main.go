package main

import (
	"echo-clean-api/config"
	"echo-clean-api/internal/router"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado. Usando variáveis do ambiente.")
	}

	// Configura e inicializa o banco de dados.
	db := config.InitDB()

	// cria uma nova instancia do echo
	e := echo.New()

	// Inicializa as rotas da api
	router.InitRoutes(e, db)

	// Verifica a porta padrão de acesso a api
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000" // valor padrão se PORT não estiver definida
	}

	// Inicializa o servidor na porta selecionada.
	e.Logger.Fatal(e.Start(":" + port))
}
