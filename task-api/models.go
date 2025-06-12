package main

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func (t *Task) validateTask() *AppError {
	if t.Name == "" {
		return &ErrTaskNameRequired
	}
	return nil
}

type AppError struct {
	Message string
	Code    int
}

func (e AppError) Error() string {
	return e.Message
}

var (
	ErrTaskNameRequired = AppError{"Nome da task é obrigatório", 400}
	ErrTaskNotFound     = AppError{"Tarefa não encontrada", 404}
	ErrInvalidID        = AppError{"ID inválido", 400}
	ErrInternalServer   = AppError{"Erro interno do servidor", 500}
)
