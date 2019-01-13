package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

//
//func main() {
//	rand.Seed(time.Now().UnixNano())
//	start := time.Now()
//	results := Google("golang")
//	elapsed := time.Since(start)
//	fmt.Println(results)
//	fmt.Println(elapsed)
//}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := First("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"),
		fakeSearch("replica 3"),
		fakeSearch("replica 4"),
		fakeSearch("replica 5"))
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}
