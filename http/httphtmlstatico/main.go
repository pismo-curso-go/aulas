package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		usuario := usuario{"João", "joao.pedro@gmail.com"}
		templates.ExecuteTemplate(rw, "home.html", usuario)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
