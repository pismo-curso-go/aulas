package main

import (
	"log"
	"net/http"
)

func main() {

	// 1. CONECTAR COM O BANCO DE DADOS
	if err := initDatabase(); err != nil {
		log.Fatal("Erro ao inicializar banco:", err)
	}
	defer closeDatabase()

	// 2. CONFIGURAR ROTAS DO SERVIDOR
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", handleTasks)

	// 3. INICIAR SERVIDOR
	log.Println("Servidor iniciado na porta 3000")
	http.ListenAndServe(":3000", mux)
}
