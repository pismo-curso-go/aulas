# Principais Características do Echo

### 1. Roteamento Simples

```go
// Rotas básicas
e.GET("/", homeHandler)
e.POST("/users", createUserHandler)
e.PUT("/users/:id", updateUserHandler)
e.DELETE("/users/:id", deleteUserHandler)

// Grupos de rotas
api := e.Group("/api/v1")
api.GET("/users", getUsersHandler)
api.POST("/users", createUserHandler)
```

### 2. Middlewares Poderosos

```go
   // Middlewares globais
   e.Use(middleware.Logger())
   e.Use(middleware.Recover())
   e.Use(middleware.CORS())

   // Middleware específico para rota
   e.GET("/admin", adminHandler, middleware.BasicAuth(validator))
   e.GET("/user", userHandler, middleware.BasicAuth(validator))
```

### 3. Binding Automático

```go
type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

func createUser(c echo.Context) error {
    user := new(User)
    if err := c.Bind(user); err != nil {
    return err
}
    // user já está preenchido automaticamente!
    return c.JSON(200, user)
}
```

### 4. Validação Integrada

```go
type User struct {
    Name string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}
```

### 5. Comparação com Desenvolvimento Nativo

```go
// Muitas linhas de código
// Verificações manuais
// Serialização manual
// Tratamento de erro verboso
func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", 405)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading body", 400)
        return
    }

    var user User
    if err := json.Unmarshal(body, &user); err != nil {
        http.Error(w, "Invalid JSON", 400)
        return
    }

    // ... mais código
}

-------------------------------------------
// Código com Echo
// Código limpo e direto
func createUser(c echo.Context) error {
    user := new(User)
    if err := c.Bind(user); err != nil {
        return err
    }

    return c.JSON(201, user)
}

```
