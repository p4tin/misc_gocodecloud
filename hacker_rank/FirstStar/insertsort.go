package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
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

func printArray(arr []int) {
	for x:=0;x<len(arr);x++ {
		fmt.Printf("%d ", arr[x])
	}
	fmt.Println("")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	inputStr, _ := reader.ReadString('\n')
	intsStr := strings.Split(strings.TrimSpace(inputStr), " ")
	intArray := convertStrArrToIntArr(intsStr)
	last := len(intArray) - 1

	element := intArray[last]
	last--
	for last >= 0 {
		if last == 0 && intArray[last] > element {
			intArray[last+1] = intArray[last]
			printArray(intArray)
			intArray[last] = element
			printArray(intArray)
			break
		}
		if intArray[last] > element {
			intArray[last+1] = intArray[last]
			printArray(intArray)
		} else {
			intArray[last+1] = element
			printArray(intArray)
			break
		}
		last--
	}
}
