package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Inicializa o banco de dados.
func initDatabase() error {
	var err error

	// 1. Conectar com banco
	db, err = connectDB()
	if err != nil {
		return err
	}
	// 2. Criar tabela se não existir
	err = createTable()
	if err != nil {
		return err
	}
	// 3. Inserir dados iniciais
	err = insertInitialData()
	if err != nil {
		return err
	}
	return nil
}

// Função que encerra / fecha conexão com base de dados.
func closeDatabase() {
	if db != nil {
		db.Close()
		log.Println("Conexão com banco fechada")
	}
}

// Função que conecta com base dados.
func connectDB() (*sql.DB, error) {
	connStr := "host=192.168.99.103 port=5432 user=postgres password=postgres dbname=task_api_db sslmode=disable"

	// Abre a conexão com o banco de dados.
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Testa a conexão com o banco de dados.
	err = database.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Conectado ao PostgreSQL!")
	return database, nil
}

// Função que cria a tabela de tasks no banco
func createTable() error {
	query := ` 
	CREATE TABLE IF NOT EXISTS tasks ( 
		id SERIAL PRIMARY KEY, 
		name VARCHAR(255) NOT NULL, 
		done BOOLEAN DEFAULT FALSE 
	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Tabela 'tasks' verificada!")
	return nil
}

// Seeder - Insere informações iniciais no banco vazio
func insertInitialData() error {
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("Dados já existem na tabela")
		return nil
	}

	initialTasks := []Task{
		{Name: "Estudar Go", Done: false},
		{Name: "Ler sobre REST", Done: true},
		{Name: "Estudar sobre HTML", Done: true},
		{Name: "Aprender PostgreSQL", Done: false},
	}

	for _, task := range initialTasks {
		err := createTaskInDB(&task)
		if err != nil {
			return err
		}
	}

	log.Println("Dados iniciais inseridos!")
	return nil
}

// Consulta todos os dados na base de dados.
func getAllTasksFromDB() ([]Task, error) {
	query := "SELECT id, name, done FROM tasks ORDER BY id"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Cria elementos no banco de dados.
func createTaskInDB(task *Task) error {
	query := "INSERT INTO tasks (name, done) VALUES ($1, $2) RETURNING id"
	err := db.QueryRow(query, task.Name, task.Done).Scan(&task.ID)
	return err
}

func getTaskByIDFromDB(id int) (*Task, error) {
	var task Task
	query := "SELECT id, name, done FROM tasks WHERE id = $1"

	err := db.QueryRow(query, id).Scan(&task.ID, &task.Name, &task.Done)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

// Altera valores no banco de dados.
func updateTaskInDB(task *Task) error {
	query := "UPDATE tasks SET name = $1, done = $2 WHERE id = $3"
	result, err := db.Exec(query, task.Name, task.Done, task.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func deleteTaskFromDB(id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	result, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // Nenhuma linha foi afetada
	}
	return nil
}
