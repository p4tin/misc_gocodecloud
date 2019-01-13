package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"time"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	info := strings.Split(strings.TrimSpace(text), " ")
	text2, _ := reader.ReadString('\n')
	nums := strings.Split(strings.TrimSpace(text2), " ")

	//maxLocation, _ := strconv.Atoi(info[0])
	numRotation, _ := strconv.Atoi(info[1])

	fmt.Println(time.Now())

	result := append(nums[numRotation:], nums[:numRotation]...)
	for x:=0;x<len(result);x++ {
		fmt.Printf("%s ", result[x])
	}
	fmt.Println("")
	fmt.Println(time.Now())
}
