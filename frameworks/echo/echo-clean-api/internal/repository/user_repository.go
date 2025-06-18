package repository

import "database/sql"

type UserRepository struct {
	DB *sql.DB
}

// "Construtor"
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}
