package main

import (
	"fmt"
	"strings"
	"bytes"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	for i:=1; i<=10;i++ {
		fmt.Printf("%03d - %s%s%d\n", i, MakeFirstUpperCase(randomdata.Adjective()),
			MakeFirstUpperCase(randomdata.Noun()),
			randomdata.s
			randomdata.Number(10,99))
	}
}

func MakeFirstUpperCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToUpper([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
