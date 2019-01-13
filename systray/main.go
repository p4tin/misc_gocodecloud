package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/getlantern/systray"
)

var (
	timezone string
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("assets/clock.png"))
	localTime := systray.AddMenuItem("Local time", "Local time")
	hcmcTime := systray.AddMenuItem("Ho Chi Minh time", "Asia/Ho_Chi_Minh")
	sydTime := systray.AddMenuItem("Sydney time", "Australia/Sydney")
	gdlTime := systray.AddMenuItem("Guadalajara time", "America/Mexico_City")
	sfTime := systray.AddMenuItem("San Fransisco time", "America/Los_Angeles")
	mChecked := systray.AddMenuItem("Unchecked", "Check Me")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		for {
			select {
			case <-localTime.ClickedCh:
				timezone = "Local"
			case <-hcmcTime.ClickedCh:
				timezone = "Asia/Ho_Chi_Minh"
			case <-sydTime.ClickedCh:
				timezone = "Australia/Sydney"
			case <-gdlTime.ClickedCh:
				timezone = "America/Mexico_City"
			case <-sfTime.ClickedCh:
				timezone = "America/Los_Angeles"
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					mChecked.SetTitle("Unchecked")
				} else {
					mChecked.Check()
					mChecked.SetTitle("Checked")
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	go func() {
		for {
			systray.SetTitle(getClockTime(timezone))
			systray.SetTooltip("Look at me, I'm a tooltip!")
			time.Sleep(1 * time.Second)
		}
	}()
}

func getClockTime(tz string) string {
	t := time.Now()
	utc, _ := time.LoadLocation(tz)

	hour, min, sec := t.In(utc).Clock()
	return ItoaTwoDigits(hour) + ":" + ItoaTwoDigits(min) + ":" + ItoaTwoDigits(sec)
}

// ItoaTwoDigits time.Clock returns one digit on values, so we make sure to convert to two digits
func ItoaTwoDigits(i int) string {
	b := "0" + strconv.Itoa(i)
	return b[len(b)-2:]
}

func onExit() {
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
