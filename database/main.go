package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Usuario struct {
	ID    int
	Nome  string
	Email string
	Idade int
}

func main() {
	testarMySQL()
	testarPostgreSQL()

	fmt.Println("Testando conexões com os bancos de dados...")
}

func testarMySQL() {
	fmt.Println("\nConectando no MySQL...")
	host := "aluno:aluno123@tcp(localhost:3306)/aula_go"

	db, err := sql.Open("mysql", host)
	if err != nil {
		log.Fatal("Erro ao conectar no MySql:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao fazer ping no MySQL:", err)
	}

	fmt.Println("Conectado no MySQL com sucesso!")

	consultarUsuarios(db, "MySQL")

}

func testarPostgreSQL() {
	host := "host=localhost port=5432 user=aluno password=aluno123 dbname=aula_go sslmode=disable"
	db, err := sql.Open("postgres", host)
	if err != nil {
		log.Fatal("Erro ao conectar no PostgreSQL:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao fazer ping no PostgreSQL:", err)
	}

	fmt.Println("Conectado no PostgreSQL com sucesso!")

	consultarUsuarios(db, "PostgresSQL")
}

// DAO ~ Repository
func consultarUsuarios(db *sql.DB, bancoDados string) {
	fmt.Printf("\n Consultando usuários no %s:\n", bancoDados)

	query := "SELECT id, nome, email, idade FROM usuarios LIMIT 3"

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Erro ao consultar usuários no %s: %v", bancoDados, err)
		return
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		err := rows.Scan(&u.ID, &u.Nome, &u.Email, &u.Idade)

		if err != nil {
			log.Printf("Erro ao fazer scan do usuário: %v", err)
			continue
		}
		usuarios = append(usuarios, u)
	}

	for _, u := range usuarios {
		fmt.Printf("  ID: %d | Nome: %s | Email: %s | Idade: %d\n", u.ID, u.Nome, u.Email, u.Idade)
	}
}
