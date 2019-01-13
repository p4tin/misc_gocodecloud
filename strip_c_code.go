package main

import "fmt"

func main() {
	var n int
	var a int
	var b int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var sum int
		fmt.Scan(&a, &b)
		sum = a + b
		fmt.Println(uint(sum))
	}
}
