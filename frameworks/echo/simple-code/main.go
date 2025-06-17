package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{1, "João"},
	{2, "Maria"},
}

func main() {

	// Criar instância do Echo
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Rotas
	e.GET("/", home)
	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)

	// Iniciar servidor
	e.Logger.Fatal(e.Start(":3000"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "API funcionando!")
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	user.ID = len(users) + 1
	users = append(users, *user)
	return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Usuário ID: "+id)
}
