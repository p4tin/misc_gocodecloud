package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

//gfcaaaecbg
//547602
//164280

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	maxCharsStr, _ := reader.ReadString('\n')
	maxChars, err := strconv.Atoi(strings.TrimSpace(maxCharsStr))
	if err != nil {
		panic(err)
	}
	str = strings.TrimSpace(str)
	multi := maxChars / len(str)
	remain := maxChars % len(str)
	substr := str[:remain]

	fmt.Println(str, multi, remain, substr)

	count := (strings.Count(str, "a") * multi)
	if remain > 0 {
		count = count + strings.Count(substr, "a")
	}

	fmt.Println(count)
}
