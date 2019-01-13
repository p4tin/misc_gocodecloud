package main

import (
	"fmt"
	"unicode"
	"bufio"
	"os"
	"strings"
)

func main() {
	words := []string{}
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	runeWord := []rune{}
	for _, elem := range strings.TrimSpace(str) {
		if unicode.IsUpper(elem) {
			words = append(words, string(runeWord))
			runeWord = []rune{}
			runeWord = append(runeWord, elem)
		} else {
			runeWord = append(runeWord, elem)
		}
	}
	words = append(words, string(runeWord))

	//fmt.Println(words)
	fmt.Println(len(words))
}
