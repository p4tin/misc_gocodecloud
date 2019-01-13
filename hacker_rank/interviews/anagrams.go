package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func findCommonRune(one []rune, two []rune) (int, int) {
	for x:=0;x<len(one);x++ {
		for y:=0;y<len(two);y++ {
			if one[x] == two[y] {
				return x, y
			}
		}
	}
	return -1, -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text2, _ := reader.ReadString('\n')

	runesOne := []rune(strings.TrimSpace(text))
	runesTwo := []rune(strings.TrimSpace(text2))

	count := len(runesOne) + len(runesTwo)
	fmt.Printf("%d\n", count)

	fmt.Printf("%v\n%v\n", runesOne, runesTwo)
	for {
		ia, ib := findCommonRune(runesOne, runesTwo)
		if ia == -1 {
			break
		}
		fmt.Println(ia, ib)
		runesOne = append(runesOne[:ia], runesOne[ia+1:]...)
		runesTwo = append(runesTwo[:ib], runesTwo[ib+1:]...)
		fmt.Printf("%v\n%v\n", runesOne, runesTwo)
		count -= 2
	}
	fmt.Printf("%d\n", count)
}
