package main

import (
	"net/http"
	"time"
	"log"
)

func busyfunc(x int) {
	x++
	time.Sleep(5 * time.Second)
	log.Println("Done busyFunc")
}

func hello(w http.ResponseWriter, r *http.Request) {
	busyfunc(100)
	log.Println("Done the Handler")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
