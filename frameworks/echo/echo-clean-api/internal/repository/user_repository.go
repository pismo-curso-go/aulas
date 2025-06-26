package repository

import (
	"database/sql"
	"echo-clean-api/internal/model"
	"errors"
)

type UserRepository struct {
	DB *sql.DB
}

// "Construtor"
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User

	row := r.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &user, err
}

// Criar usuário no banco de dados.
func (r *UserRepository) Create(user *model.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	return err
}
