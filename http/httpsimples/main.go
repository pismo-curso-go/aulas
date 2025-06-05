package main

import (
	"log"
	"net/http"
)

func main() {

	// localhost:5000/home
	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		println("Olá universo, eu sou um log.")
		rw.Write([]byte("Olá Mundo!"))
	})

	// localhost:5000/usuarios
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de usuários!"))
	})

	// localhost:5000/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página Raiz!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
