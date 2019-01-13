package main

import (
	"github.com/gocolly/colly"
	"fmt"
)

type Article struct {

}

type Newspaper struct {
	Articles []Article
}
type comment struct {
	Author  string `selector:"a.hnuser"`
	URL     string `selector:".age a[href]" attr:"href"`
	Comment string `selector:".comment"`
	Replies []*comment
	depth   int
}

func main() {
	c := colly.NewCollector()

	// Extract comment
	c.OnHTML("article", func(e *colly.HTMLElement) {
		fmt.Println((*e).DOM.Get(0))
		fmt.Println("\n\n============\n\n")
	})

	c.Visit("https://www.cnn.com/specials/last-50-stories")
}