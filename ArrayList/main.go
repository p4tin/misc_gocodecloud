package main

import "fmt"

type MapList map[interface{}]interface{}

func (ar MapList) Contains(needle interface{}) bool {
	if _, ok := ar[needle]; ok {
		return true
	}
	return false
}

var emails = MapList{
	"me@aol.com":  nil,
	"you@aol.com": nil,
	"him@aol.com": nil,
	"her@aol.com": nil,
}

var numbers = MapList{
	1:   nil,
	23:  nil,
	60:  nil,
	105: nil,
}

func main() {
	fmt.Println(emails.Contains("you@aol.com"))
	fmt.Println(emails.Contains("you2@aol.com"))

	fmt.Println(numbers.Contains(23))
	fmt.Println(numbers.Contains(8))
}
