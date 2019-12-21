package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

func CaseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func main() {
	resp, err := soup.Get("https://news.ycombinator.com/front")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "storylink").Text()
	//for _, link := range links {
	//	if CaseInsensitiveContains(link.Text(), "Go") || CaseInsensitiveContains(link.Text(), "Python") {
	//		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	//	}
	//}
	fmt.Printf("+%v\n", links)
}
