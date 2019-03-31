//libadd.go
package main

import (
	"C"
	"fmt"
)

type IntStruct struct {
	X int
	Y int
}

func init() {
	fmt.Println("In init")
}

//export add
func add(left, right int) int {
	return left + right
}

//export structadd
func structadd(i interface{}) int {
	s := i.(*IntStruct)
	return s.X + s.Y
}

func main() {
}
