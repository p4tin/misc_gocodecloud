package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))
	if num % 2 == 1 || (num >= 6 && num <= 20) {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}
