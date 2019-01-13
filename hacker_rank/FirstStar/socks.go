package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
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
	_, _ = reader.ReadString('\n')
	str, _ := reader.ReadString('\n')
	array_str := strings.Split(str, " ")
	socks := convertStrArrToIntArr(array_str)
	sort.Ints(socks)
	numPairs := 0
	for x:=0;x<len(socks)-1;x++ {
		if socks[x] == socks[x+1] {
			numPairs++
			x++
		}
	}
	fmt.Println(numPairs)
}
