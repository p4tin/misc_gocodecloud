package main

import (
	"log"
	"github.com/ajstarks/svgo"
	"net/http"
)

func main() {
	http.Handle("/circle", http.HandlerFunc(circle))
	http.Handle("/circle2", http.HandlerFunc(circle2))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}

func circle2(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	width := 500
	height := 500
	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}
