package main

import (
	"bytes"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func splitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i + 1) % n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}
	return subs
}

func compare(str string) int {
	count := 0
	sos := []rune{'S', 'O', 'S'}
	input := []rune(str)
	for x:=0;x<3;x++ {
		if sos[x] != input[x] {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	strs := splitSubN(strings.TrimSpace(str), 3)

	count := 0

	for _, sos := range strs {
		count = count + compare(sos)
	}
	fmt.Println(count)
}
