package main

import (
	"fmt"
	"misc/0-herowars-bot/utils"
	"time"

	"github.com/go-vgo/robotgo"
)

var coords [2]int
var trials = map[string][][2]int{
	"8.2": {{434, 213}, {170, 72}, {20, -37}},
}

var task = [][2]int{{434, 213}, {170, 72}, {20, -37}}

func main() {
	for {
		fx := -1
		fy := -1
		for fx == -1 {
			fmt.Println("Reference button not found yet...")
			screenbit := robotgo.CaptureScreen()

			fx, fy = robotgo.FindPic("./continue_button.png", screenbit, 0.1)

			fmt.Println("FindBitmap...", fx, fy)
			time.Sleep(2 * time.Second)
		}
		fx = fx/2 + 30
		fy = fy/2 + 30

		robotgo.Move(fx, fy)

		fmt.Println(utils.GetLMouseClick())
		fmt.Println(utils.GetLMouseClick())

		fx, fy = utils.Scoot(fx, fy, task[0][0]/2, task[0][1]/2)
		robotgo.Move(fx, fy)
		fmt.Println(utils.GetLMouseClick())

		//fx, fy = utils.Scoot(fx, fy, task[1][0], task[1][1])
		//robotgo.Move(fx, fy)
		//fmt.Println(utils.GetLMouseClick())

		//fx, fy = utils.Scoot(fx, fy, task[2][0], task[2][1])
		//robotgo.Move(fx, fy)
		//fmt.Println(utils.GetLMouseClick())

		//fx, fy = utils.Scoot(fx, fy, task[3][0], task[3][1])
		//robotgo.Move(fx, fy)
		//time.Sleep(3 * time.Second)
		//fmt.Println(utils.GetLMouseClick())

		time.Sleep(60 * time.Second)
	}
}
