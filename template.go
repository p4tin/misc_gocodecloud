package main

import (
	"html/template"
	"os"
)


type Page struct {
	Name 	string
	Ec	EntitiesClass
}

type EntitiesClass struct {
	Name string
	Messages []string
}

// In the template, we use rangeStruct to turn our struct values
// into a slice we can iterate over
var htmlTemplate = "{{ range .Ec.Messages }}{{ . }}{{ end }}"

func main() {
	data := EntitiesClass{Name: "test1"}
	data.Messages = make([]string, 3)
	data.Messages[0] = "New game 1\n"
	data.Messages[1] = "New game 2\n"

	page := Page{Name: "booboo", Ec: data}

	t := template.New("t")
	t, err := t.Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, page)
	if err != nil {
		panic(err)
	}

}