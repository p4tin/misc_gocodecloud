package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
)

// Environment variables for the poller-receiver
var maxPollers int
var maxReceivers int
var pollingPeriod int  //polling period in milli-seconds

var poller_wg sync.WaitGroup
var receiver_wg sync.WaitGroup

func poller(cs chan string, done chan bool) {
	defer poller_wg.Done()
	timeout := make(chan bool, 1)
	go func() {
		for{
			time.Sleep(time.Duration(pollingPeriod) * time.Millisecond)
			timeout <- true
		}
	}()
	for{
		select {
		case <-done:
			log.Println("Poller Done received.")
			return
		case <-timeout:
			cs <- "Hello"
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))  //Fake random busy time
		}
	}
}

func receiver(cs chan string, done chan bool) {
	defer receiver_wg.Done()
	for{
		select {
		case data := <-cs:
			log.Println(data)
		case <-done:
			log.Println("Receiver Done received.")
			return
		}
		//data := <-cs
		//log.Println(data)
	}
}

func init(){
	var err error
	maxPollers, err = strconv.Atoi(os.Getenv("MAX_POLLERS"))
	if err != nil {
		log.Fatal("Could not set maxPollers, please set the 'MAX_POLLERS' environment variable.")
	}
	maxReceivers, err = strconv.Atoi(os.Getenv("MAX_RECEIVERS"))
	if err != nil {
		log.Fatal("Could not set maxReceivers, please set the 'MAX_RECEIVERS' environment variable.")
	}
	pollingPeriod, err = strconv.Atoi(os.Getenv("POLLING_PERIOD"))
	if err != nil {
		log.Fatal("Could not set the pollingPeriod, please set the 'POLLING_PERIOD' environment variable.")
	}
}

func main() {
	cs := make(chan string, 10000)

	poller_done := make(chan bool)
	receiver_done := make(chan bool)

	poller_wg.Add(maxPollers)
	receiver_wg.Add(maxReceivers)

	//Trap signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		<-c

		log.Println("Cleaning up!!!")
		// Stop Pollers
		for x:=0;x<maxPollers;x++ {
			poller_done <- true
		}
		poller_wg.Wait()

		// Stop Receivers
		for x:=0;x< maxReceivers;x++ {
			receiver_done <- true
		}
		receiver_wg.Wait()
		log.Println("Clean up done, exiting!!!")
		os.Exit(1)
	}()

	for x:=0;x<maxPollers;x++ {
		go poller(cs, poller_done)
	}
	for x:=0;x< maxReceivers;x++ {
		go receiver(cs, receiver_done)
	}

	receiver_wg.Wait()
}

