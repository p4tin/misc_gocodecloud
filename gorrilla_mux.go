package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
)

// hello world, the web server
func IndexServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello index\n")
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(w, "hello, "+vars["name"]+"!\n")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexServer)
	r.HandleFunc("/hello/{name}", HelloServer)
	http.ListenAndServe(":12345", r)
}