package main

import (
	"github.com/everdev/mack"
	"time"
)

func main() {
	mack.Say("I am asking myself - but like a robot... Laugh out loud!")
	time.Sleep(1000)
	mack.Notify("Complete")
}
