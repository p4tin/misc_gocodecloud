package main

import (
	"github.com/melvinmt/firebase"
	"fmt"
)

type PersonName struct {
	First string
	Last  string
}

type Person struct {
	Name PersonName
}

func main() {
	var err error

	url := "https://gostandups.firebaseio.com/users/fred/name"

	// Can also be your Firebase secret:
	authToken := "p4um1u3UUIME6D7CL98AiIOXWrKF41U0SYI6CGg7"

	// Auth is optional:
	ref := firebase.NewReference(url).Auth(authToken)

	// Create the value.
	personName := PersonName{
		First: "Fred",
		Last:  "Swanson",
	}

	// Write the value to Firebase.
	if err = ref.Write(personName); err != nil {
		panic(err)
	}

	// Now, we're going to retrieve the person.
	personUrl := "https://gostandups.firebaseio.com/users/fred"

	personRef := firebase.NewReference(personUrl).Export(false)

	fred := Person{}

	if err = personRef.Value(&fred); err != nil {
		panic(err)
	}

	fmt.Println(fred.Name.First, fred.Name.Last) // prints: Fred Swanson
}