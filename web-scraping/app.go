package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type TeamScore struct {
	Home  bool
	Name  string
	Score int
}

type game [2]TeamScore

var games []game

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	//c.OnHTML("g5-component--nhl-scores__panel--primary", func(e *colly.HTMLElement) {
	//	//g5-component--nhl-scores__team-name 0 = Visitor
	//	e.ForEach("g5-component--nhl-scores__team-name", func(i int, element *colly.HTMLElement) {
	//		fmt.Println(element.Text)
	//	})
	//	//Score = g5-component--nhl-scores__team-score
	//})

	c.OnHTML("div[class=g5-component--nhl-scores__panel--primary]", func(e *colly.HTMLElement) {
		// If attribute class is this long string return from callback
		// As this a is irrelevant
		fmt.Println("Game")

	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.nhl.com/scores")
}
