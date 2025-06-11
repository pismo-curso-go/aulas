package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Name: "Estudar Go", Done: false},
	{ID: 2, Name: "Ler sobre REST", Done: true},
	{ID: 3, Name: "Estudar sobre HTML", Done: true},
}

func main() {
	mux := http.NewServeMux() //porteiro

	mux.HandleFunc("/tasks", handleTasks)          // /tasks
	mux.HandleFunc("/tasks/", handleSpecificTasks) // /tasks/282

	log.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", mux)
}

// Controller - /tasks (get, post)
func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// http://localhost/tasks/3123/users - ROTA
// http://localhost/tasks/3123?nome=matheus&page=3&offset=20&ordem=ASC - BUSCA
// body { "nome": "matheus" } - CORPO

func handleSpecificTasks(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTaskById(w, r, id)
	case http.MethodDelete:
		deleteTaskById(w, r, id)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// Service ~ UseCase
func getTasks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func getTaskById(w http.ResponseWriter, _ *http.Request, id int) {
	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
}

func deleteTaskById(w http.ResponseWriter, _ *http.Request, id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
}
