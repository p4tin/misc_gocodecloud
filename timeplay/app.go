package main

import (
	"fmt"
	"time"
)

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func main() {
	deliveryDate := time.Date(2019, time.Month(9), 6, 0, 0, 0, 0, time.Local)
	tomorrow := Bod(time.Now().AddDate(0, 0, 1))

	fmt.Println(deliveryDate)
	fmt.Println(tomorrow)

	if !deliveryDate.Before(tomorrow) {
		fmt.Println("1. send date")
	}

	if deliveryDate.After(tomorrow) {
		fmt.Println("2. send date")
	}

	t := time.Now().UTC()
	fmt.Println(t.Format(time.RFC3339))
}
