package main

import (
	"math/rand"
	"time"
	"fmt"
)

// Implementations

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for i:=1; i<=10;i++ {
		fmt.Printf("%03d - %s\n", i, RandStringRunes(10))
	}
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
