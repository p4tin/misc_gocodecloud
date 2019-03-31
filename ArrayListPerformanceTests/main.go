package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

type MapList map[interface{}]interface{}

func (ar MapList) Contains(needle interface{}) bool {
	if _, ok := ar[needle]; ok {
		return true
	}
	return false
}

const (
	listSize = 1000000
)

func main() {
	fmt.Println("List Size:", listSize)
	maplist := make(MapList)
	list := make([]string, 0)

	start := time.Now()
	for x := 0; x < listSize; x++ {
		str := RandStringRunes(10)
		maplist[str] = nil
		list = append(list, str)
	}
	sort.Strings(list)

	elapsed := time.Since(start)
	fmt.Println("Making and Sorting the list", elapsed)

	start = time.Now()

	Contains(list, list[listSize-1])
	elapsed = time.Since(start)
	fmt.Println("Contains took", elapsed)

	start = time.Now()

	maplist[list[listSize-1]] = false
	elapsed = time.Since(start)
	fmt.Println("Map took", elapsed)

	start = time.Now()

	sort.SearchStrings(list, list[listSize-1])
	elapsed = time.Since(start)
	fmt.Println("Find with sort", elapsed)
}
