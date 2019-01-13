package main

import (
	"sync"
)

const maxConcurrency = 4 // for example

var throttle = make(chan int, maxConcurrency)

func main() {
	const N = 100 // for example
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		throttle <- 1 // whatever number
		wg.Add(1)
		go f(i, &wg, throttle)
	}
	wg.Wait()
}

func f(i int, wg *sync.WaitGroup, throttle chan int) {
	defer wg.Done()
	// whatever processing
	println(i)
	<-throttle
}