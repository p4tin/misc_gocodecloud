package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin))
	aliceStr, _ := reader.ReadString('\n')
	bobStr, _ := reader.ReadString('\n')
	fmt.Printf("%s\n%s", aliceStr, bobStr)
	aliceStrArr := strings.Split(strings.TrimSpace(aliceStr), " ")
	bobStrArr := strings.Split(strings.TrimSpace(bobStr), " ")

	alice := 0
	bob := 0
	for i:=0;i < len(aliceStrArr);i++ {
		a, _ := strconv.Atoi(string(aliceStrArr[i]))
		b, _ := strconv.Atoi(string(bobStrArr[i]))
		if a > b {
			alice++
		} else if b > a {
			bob++
		}
	}
	fmt.Println(alice, bob)
}
