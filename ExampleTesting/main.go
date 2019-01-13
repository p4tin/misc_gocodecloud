package main

import (
	"fmt"
	"sync"
)

var wgApp sync.WaitGroup

func aFuncYouWantToTest(txt string) {
	fmt.Printf("Hello %s!!!\n", txt)
}

func setUp() {
	wgApp = sync.WaitGroup{}
	wgApp.Add(1)

	go func() {
		aFuncYouWantToTest("Paul")
		wgApp.Done()
	}()
	aFuncYouWantToTest("Wowzers")

	wgApp.Wait()
}

func main() {
	setUp()
}
