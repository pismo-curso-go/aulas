package utils

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e AppError) Error() string {
	return e.Message
}

var (
	ErrInvalidCredentials = AppError{"Credenciais inválidas", 401}
	ErrUserAlreadyExists  = AppError{"Usuário já existe", 409}
	ErrValidationFailed   = AppError{"Dados inválidos", 400}
	ErrUserNotFound       = AppError{"Usuário não encontrado", 404}
	ErrInternalError      = AppError{"Erro interno do servidor", 500}
)
