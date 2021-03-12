package utils

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func GetLMouseClick() (int, int) {
	//robotgo.AddEvent("mleft")
	time.Sleep(1 * time.Second)
	robotgo.Click("left", false)
	return robotgo.GetMousePos()
}

func Scoot(x, y, dx, dy int) (int, int) {
	return x + dx, y + dy
}

// Add a random delay to behave more like human input
func Delay() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(1-2) + 1
}
