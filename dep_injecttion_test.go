package main

import (
	"fmt"
	"testing"
)

var ScreenPrint = fmt.Printf

func Greet(name string) {
	ScreenPrint("Hello, %s", name)
}

//func main() {
//	Greet("Playground")
//}

func TestGreet(t *testing.T) {
	var result string
	ScreenPrint = func(format string, a ...interface{}) (n int, err error) {
		result = fmt.Sprintf(format, a[0])
		n = len(result)
		err = nil
		return
	}
	expected := "Hello, Tester"
	Greet("Tester")
	if expected != result {
		t.Errorf("got '%s' want '%s'", result, expected)
	}
}
