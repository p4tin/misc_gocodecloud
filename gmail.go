package main

import (
	"log"
	"net/smtp"
)

func main() {
	sendEmail("hello there")
}

func sendEmail(body string) {
	from := "pfort69@gmail.com"
	pass := "iv17wt0c"
	to := "pafortin@cp-soft.com"

	msg := "From: GoStandUps - Do Not Reply<pfort69@gmail.com>\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
