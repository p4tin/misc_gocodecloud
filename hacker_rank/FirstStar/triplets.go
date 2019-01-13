package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	aliceStr, _ := reader.ReadString('\n')
	bobStr, _ := reader.ReadString('\n')
	aliceStrArr := strings.Split(aliceStr, " ")
	bobStrArr := strings.Split(bobStr, " ")

	alice := 0
	bob := 0
	for i, _ := range aliceStrArr {
		a, _ := strconv.Atoi(string(aliceStrArr[i]))
		b, _ := strconv.Atoi(string(bobStrArr[i]))
		if a > b {
			alice++
		} else {
			bob++
		}
	}
	fmt.Println(alice, bob)
}
