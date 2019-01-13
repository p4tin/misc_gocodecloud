package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/eknkc/amber"
)

var compiler *amber.Compiler

type Person struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Parse the input file
	err := compiler.ParseFile("templates/index.amber")
	if err == nil {
		// Compile input file to Go template
		tpl, err := compiler.Compile()
		if err == nil {
			// Check built in html/template documentation for further details
			person := Person{Name: "Astaxie"}
			tpl.Execute(w, person)
		}
	}
}

func init() {
	compiler = amber.New()
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	router.HandleFunc("/", handler)

	err := http.ListenAndServe(":9999", router)
	if err != nil {
		log.Printf("Web Server cannot start: %v\n", err)
	}
}