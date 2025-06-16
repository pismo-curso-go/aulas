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
	mux.HandleFunc("/tasks/", handleSpecificTask)

	// 3. INICIAR SERVIDOR
	log.Println("Servidor iniciado na porta 3000")
	log.Println("Endpoints disponíveis:")
	log.Println("   GET    /tasks       - Listar todas as tasks")
	log.Println("   POST   /tasks       - Criar nova task")
	log.Println("   GET    /tasks/{id}  - Buscar task por ID")
	log.Println("   PUT    /tasks/{id}  - Atualizar task")
	log.Println("   DELETE /tasks/{id}  - Deletar task")

	http.ListenAndServe(":3000", mux)
}
