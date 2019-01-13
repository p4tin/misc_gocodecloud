package main

import (
	"net/http"
	"log"
	"time"

	jwt2 "github.com/adam-hanna/jwt-auth/src"
)

var restrictedRoute jwt2.Auth

var restrictedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the secret area!"))
})

var regularHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
})

func main() {
	authErr := jwt2.New(&restrictedRoute, jwt2.Options{
		PrivateKeyLocation:     "keys/app.rsa", // `$ openssl genrsa -out app.rsa 2048`
		PublicKeyLocation:      "keys/app.rsa.pub", // `$ openssl rsa -in app.rsa -pubout > app.rsa.pub`
		RefreshTokenValidTime:  72 * time.Hour,
		AuthTokenValidTime:     15 * time.Minute,
		Debug:                  false,
	})
	if authErr != nil {
		log.Println("Error initializing the JWT's!")
		log.Fatal(authErr)
	}

	http.HandleFunc("/", regularHandler)
	// this will never be available because we never issue tokens
	// see login_logout example for how to provide tokens
	http.Handle("/restricted", restrictedRoute.Handler(restrictedHandler))

	log.Println("Listening on localhost:3000")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
