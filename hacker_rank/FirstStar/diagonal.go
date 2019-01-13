package main

import (
	"math"
	"fmt"
	"strings"
	"bufio"
	"os"
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

func main() {
	cube := [][]int{}
	size := 0

	reader := bufio.NewReader(os.Stdin)

	//Get Cube size
	str, _ := reader.ReadString('\n')
	size, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		panic(err)
	}

	for x:=0;x<size;x++ {
		str, _ = reader.ReadString('\n')
		strArr := strings.Split(strings.TrimSpace(str), " ")
		intArr := convertStrArrToIntArr(strArr)
		cube = append(cube, intArr)
	}

	num1 := 0
	for x:=0;x<len(cube[0]);x++ {
		num1 = num1 + cube[x][x]
	}
	num2 := 0
	y := 0
	for x:=len(cube[0])-1;x>=0;x-- {
		num2 = num2 + cube[y][x]
		y++
	}
	fmt.Println(math.Abs(float64(num1 - num2)))
}
