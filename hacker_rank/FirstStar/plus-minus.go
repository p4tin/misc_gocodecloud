package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func convertStrArrToIntArr(strArr []string) []int {
	var intArr = []int{}
	for _, i := range strArr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intArr = append(intArr, j)
	}
	return intArr
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	inputStr, _ := reader.ReadString('\n')
	intsStr := strings.Split(strings.TrimSpace(inputStr), " ")
	intArray := convertStrArrToIntArr(intsStr)
	count := float64(len(intArray))

	pos := 0
	neg := 0
	zer := 0

	for _, j := range intArray {
		if j == 0 {
			zer++
		} else if j > 0 {
			pos++
		} else {
			neg++
		}
	}
	fmt.Printf("%0.6f\n%0.6f\n%0.6f\n", float64(pos) / count, float64(neg) / count, float64(zer) / count)
}
