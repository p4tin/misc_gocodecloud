package main

import "fmt"

func main() {
	for i, j := 0, 1; j < 1000; i, j = i+j,i {
		fmt.Println(i)
	}
}
