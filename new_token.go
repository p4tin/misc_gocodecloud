package main

import (
	"context"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

var MyKey = "token"

type Claims struct {
	Username string `json:"username"`
	// recommended having
	jwt.StandardClaims
}

// create a JWT and put in the clients cookie
func SetToken(res http.ResponseWriter, req *http.Request) {
	// Expires the token and cookie in 1 hour
	fmt.Println("SetToken")
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	expireCookie := time.Now().Add(time.Hour * 1)

	// We'll manually assign the claims but in production you'd insert values from a database
	claims := Claims {
		"myusername",
		jwt.StandardClaims {
			ExpiresAt: expireToken,
			Issuer:    "localhost:9000",
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := token.SignedString([]byte("secret"))
	fmt.Println(signedToken)

	// Place the token in the client's cookie
	cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(res, &cookie)

	// Redirect the user to his profile
	http.Redirect(res, req, "/profile", 307)
}

// middleware to protect private pages
func Validate(page http.HandlerFunc) http.HandlerFunc {
	fmt.Println("Here")
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		fmt.Println(req.Cookies())
		// If no Auth cookie is set then return a 404 not found
		cookie, err := req.Cookie("Auth")
		if err != nil {
			http.NotFound(res, req)
			return
		}

		// Return a Token using the cookie
		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error){
			// Make sure token's signature wasn't changed
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			http.NotFound(res, req)
			return
		}

		// Grab the tokens claims and pass it into the original request
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(req.Context(), MyKey, *claims)
			page(res, req.WithContext(ctx))
		} else {
			http.NotFound(res, req)
			return
		}
	})
}

// only viewable if the client has a valid token
func ProtectedProfile(res http.ResponseWriter, req *http.Request){
	claims, ok := req.Context().Value(MyKey).(Claims)
	if !ok {
		http.NotFound(res, req)
		return
	}

	fmt.Fprintf(res, "Hello %s", claims.Username)
}

// deletes the cookie
func Logout(res http.ResponseWriter, req *http.Request) {
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(res, &deleteCookie)
	return
}

func main(){
	http.HandleFunc("/settoken", SetToken)
	http.HandleFunc("/profile", Validate(ProtectedProfile))
	http.HandleFunc("/logout", Validate(Logout))
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
