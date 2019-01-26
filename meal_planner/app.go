package main

import (
	"bufio"
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"
)

var Days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var date_layout = "01-02-2006"

type Meal struct {
	Name       string
	Difficulty string
	PrepTime   string
	LastUsed   time.Time
}

func pruneMeals(meals *[]Meal) *[]Meal {
	var elibibleMeals []Meal
	lastWeek := time.Now().AddDate(0, 0, -7)
	fmt.Println("LastWeek:", lastWeek)
	for _, meal := range *meals {
		if meal.LastUsed.Before(lastWeek) {
			fmt.Println("\tmeal Time:", meal.LastUsed)
			elibibleMeals = append(elibibleMeals, meal)
		}
	}
	return &elibibleMeals
}

func main() {
	csvFile, _ := os.Open("meals.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var meals []Meal
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		lastUsed, err := time.Parse(date_layout, line[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		meals = append(meals, Meal{
			Name:       line[0],
			Difficulty: line[1],
			PrepTime:   line[2],
			LastUsed:   lastUsed,
		})
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	emailBody := ""
	meals = *pruneMeals(&meals)
	if len(meals) < 7 {
		fmt.Println("I do not have enought meals to get you a new meal plan.  Please add recipies or update the last used date in the CSV")
		os.Exit(0)
	}
	for x := 0; x < 7; x++ {
		ndx := r1.Intn(len(meals))
		emailBody = fmt.Sprintf("%s\n%s\n Supper - %s\n", emailBody, Days[x], meals[ndx].Name)
		meals = append(meals[:ndx], meals[ndx+1:]...)
	}
	//send_mail(emailBody)
	fmt.Println(emailBody)
}

/**** Mail ****/

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func send_mail(body string) {
	mail := Mail{}
	mail.senderId = "pafortin@cp-soft.com"
	mail.toIds = []string{"pafortin@cp-soft.com", "cafortin640@hotmail.com"}
	mail.subject = "Meal Plan for week of starting 14 Jan 2019"
	mail.body = body

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}

	log.Println(smtpServer.host)
	//build an auth
	auth := smtp.PlainAuth("pafortin@cp-soft.com", mail.senderId, "Iv17wt@c", smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")

}
