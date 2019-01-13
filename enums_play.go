package main

import (
	"fmt"
)

// Declare a new type named Weekday which will unify our enum values
// It has an underlying type of unsigned integer (uint).
type Weekday int

// Declare typed constants each with type of Weekday
const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

// String returns the name of the day
func (day Weekday) String() string {
	names := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday",
		"Thursday", "Friday", "Saturday"}

	return names[day]
}

// Weekend return true for a weekend day
func (day Weekday) Weekend() bool {
	switch day {
	case Sunday, Saturday:
		return true
	default:
		return false
	}
}



func main() {
	// will display 0
	fmt.Println(Sunday)
	// will display 6
	fmt.Println(Saturday)

	fmt.Println(Saturday.Weekend())
}

