//libadd.go
package main

import (
	"C"
	"fmt"
)

func init() {
	fmt.Println("In init")
}

//export add
func add(left, right int) int {
	return left + right
}

func main() {
}
