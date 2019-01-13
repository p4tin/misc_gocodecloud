package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for x:=0;x<10;x++ {
		fmt.Println(r1.Intn(20))
	}
}
