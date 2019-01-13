package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time" // or "runtime"
)

func cleanup() {
	fmt.Println("cleanup")
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for {
		fmt.Println("sleeping...")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}
}