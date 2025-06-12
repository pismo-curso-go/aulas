package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Controller - /tasks (get, post)
func handleTasks(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetAllTasks(w, r)
	// case http.MethodPost:
	// 	handleCreateTask(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// USE_CASE ~ SERVICE
func handleGetAllTasks(w http.ResponseWriter, _ *http.Request) {
	tasks, err := getAllTasksFromDB()

	if err != nil {
		log.Printf("Erro ao buscar tasks: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
