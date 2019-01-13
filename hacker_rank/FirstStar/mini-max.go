package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
	"sort"
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
	str, _ := reader.ReadString('\n')
	strArr := strings.Split(strings.TrimSpace(str), " ")
	intArr := convertStrArrToIntArr(strArr)

	fmt.Println(intArr)

	sort.Ints(intArr)
	fmt.Println(intArr)

	minSum := 0
	for i := 0; i < 4; i++ {
		minSum = minSum + intArr[i]
	}

	maxSum := 0
	for i := 1; i < 5; i++ {
		maxSum = maxSum + intArr[i]
	}

	fmt.Println(minSum, maxSum)
}
