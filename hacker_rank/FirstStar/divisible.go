package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func convertStrArrToIntArr(strArr []string) []int {
	var intArr = []int{}
	for _, i := range strArr {
		j, err := strconv.Atoi(strings.TrimSpace(i))
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
	array_str := strings.Split(str, " ")
	socks := convertStrArrToIntArr(array_str)
	divisor := socks[1]
	str, _ = reader.ReadString('\n')
	array_str = strings.Split(str, " ")
	dividends := convertStrArrToIntArr(array_str)

	numEvenDivided := 0
	for x:=0;x<len(dividends)-1;x++ {
		for y:=x+1;y<len(dividends);y++ {
			if ((dividends[x] + dividends[y]) % divisor) == 0 {
				numEvenDivided++
			}
		}
	}
	fmt.Println(numEvenDivided)
}
