package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Controller - /tasks (get, post)
func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetAllTasks(w, r)
	case http.MethodPost:
		handleCreateTask(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// Controller - /tasks/ (getById, Put, Delete)
func handleSpecificTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGetTaskByID(w, r, id)
	case http.MethodPut:
		handleUpdateTask(w, r, id)
	case http.MethodDelete:
		handleDeleteTask(w, r, id)
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

// USE_CASE ~ SERVICE
func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	err := json.NewDecoder(r.Body).Decode(&newTask)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Regra de negócio - Nome não pode ser vazio
	if err := newTask.validateTask(); err != nil {
		http.Error(w, err.Message, err.Code)
		return
	}

	err = createTaskInDB(&newTask)
	if err != nil {
		log.Printf("Erro ao criar task: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// USE_CASE ~ SERVICE
func handleGetTaskByID(w http.ResponseWriter, _ *http.Request, id int) {
	task, err := getTaskByIDFromDB(id)
	if err != nil {
		log.Printf("Erro ao buscar task %d: %v", id, err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	if task == nil {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// USE_CASE ~ SERVICE
func handleUpdateTask(w http.ResponseWriter, r *http.Request, id int) {
	var updatedTask Task

	err := json.NewDecoder(r.Body).Decode(&updatedTask)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := updatedTask.validateTask(); err != nil {
		http.Error(w, err.Message, err.Code)
		return
	}

	updatedTask.ID = id

	err = updateTaskInDB(&updatedTask)
	if err != nil {
		if err == sql.ErrNoRows {
			// Tarefa não existe
			http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
			return
		}
		log.Printf("Erro ao atualizar task %d: %v", id, err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Retorna tarefa atualizada
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

// USE_CASE ~ SERVICE
func handleDeleteTask(w http.ResponseWriter, _ *http.Request, id int) {
	err := deleteTaskFromDB(id)

	if err != nil {
		if err == sql.ErrNoRows {
			// Tarefa não existe
			http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
			return
		}

		log.Printf("Erro ao deletar task %d: %v", id, err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Sucesso - sem conteúdo na resposta (204 No Content)
	w.WriteHeader(http.StatusNoContent)
}
