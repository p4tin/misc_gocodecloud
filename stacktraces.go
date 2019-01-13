package main

import (
	"log"

	"github.com/pkg/errors"
)

//Log checks and logs a error
func Log(err error) {
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, ""))
	}
}

func CalculateAnswer(question string) (string, error) {
	if question == "" {
		return "", errors.New("Please ask a valid question")
	}
	var answer string
	//... Fancy calculation ...
	return answer, nil
}

func main() {
	_, err := CalculateAnswer("")
	Log(err)
}
