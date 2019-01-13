//libadd.go
// - go build -buildmode=c-shared -o libadd.so libadd.go
package main

import (
	"C"
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wg sync.WaitGroup

//export add
func add(left, right int) int {
	return left + right
}

//export longRun
func longRun() {
	wg.Add(15)
	rand.Seed(time.Now().UTC().UnixNano())
	for i:=1;i<16;i++ {
		go waitTime(i, rand.Int63n(5))
	}
	wg.Wait()
}

func waitTime(i int, t int64) {
	defer wg.Done()
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("Completed Thread:", i)
}

func main() {
}

