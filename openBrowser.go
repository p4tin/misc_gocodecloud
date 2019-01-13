package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	fmt.Println("Listening...")

	err := open.Run("http://localhost:8080/test.html")
	if err != nil {
		fmt.Println(err)
	}

	http.ListenAndServe(":8080", nil)

}
