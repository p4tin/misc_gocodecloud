package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func findWord(word string, list []string) int {
	for x:=0;x<len(list);x++ {
		if word == list[x] {
			return x
		}
	}
	return -1
}

var text string

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	text2, _ := reader.ReadString('\n')
	text3, _ := reader.ReadString('\n')

	wordsOne := []string(strings.Split(strings.TrimSpace(text2), " "))
	wordsTwo := []string(strings.Split(strings.TrimSpace(text3), " "))
	count := 0

	for _, word := range wordsTwo {
		ib := findWord(word, wordsOne)
		if ib > -1 {
			count++
			wordsOne = append(wordsOne[:ib], wordsOne[ib+1:]...)
		} else {
			break
		}
	}
	if count == len(wordsTwo) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
