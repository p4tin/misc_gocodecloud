package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	str, _ := reader.ReadString('\n')
	array_str := strings.Split(str, " ")
	total := 0
	for _, i := range array_str {
		j, _ := strconv.Atoi(string(i))
		total = total + j
	}
	fmt.Println(total)
}
