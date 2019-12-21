package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("\n\nWelcome to the game, %s\n", s.User()))
		io.WriteString(s, "Type 'help' to get a list fo commands or 'quit' to save and exit the game\n\n")
		term := terminal.NewTerminal(s, "> ")
		line := ""
		doneForToday := false
		for {
			line, _ = term.ReadLine()
			switch line {
			case "quit":
				doneForToday = true
			case "help":
				io.WriteString(s, "This is the help\n")
			default:
				io.WriteString(s, "Unrecognized command\n")
			}
			if doneForToday {
				break
			}
		}
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
