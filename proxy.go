package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func report(r *http.Request) {
	r.Host = "localhost:8888"
	r.URL.Host = r.Host
	r.URL.Scheme = "http"
}

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8888",
	})
	proxy.Director = report

	r := mux.NewRouter()

	r.Handle("/hello", proxy)
	r.Handle("/hello/{user}", proxy)
	r.Handle("/bye", proxy)

	http.ListenAndServe(":8181", r)
}
