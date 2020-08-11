package main

import (
	"fmt"
	"misc/0-herowars-bot/utils"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		fx := -1
		fy := -1
		for fx == -1 {
			fmt.Println("Reference button not found yet...")
			screenbit := robotgo.CaptureScreen()

			fx, fy = robotgo.FindPic("./return_to_city_button.png", screenbit, 0.1)

			fmt.Println("FindBitmap...", fx, fy)
			time.Sleep(2 * time.Second)
		}
		fx = fx/2 + 30
		fy = fy/2 + 30

		robotgo.Move(fx, fy)

		fmt.Println(utils.GetLMouseClick())
		fmt.Println(utils.GetLMouseClick())

		fx, fy = utils.Scoot(fx, fy, 375, -285)
		robotgo.Move(fx, fy)
		fmt.Println(utils.GetLMouseClick())

		fx, fy = utils.Scoot(fx, fy, 0, 230)
		robotgo.Move(fx, fy)
		fmt.Println(utils.GetLMouseClick())

		fx, fy = utils.Scoot(fx, fy, 145, 70)
		robotgo.Move(fx, fy)
		fmt.Println(utils.GetLMouseClick())

		fx, fy = utils.Scoot(fx, fy, 30, -25)
		robotgo.Move(fx, fy)
		time.Sleep(3 * time.Second)
		fmt.Println(utils.GetLMouseClick())

		time.Sleep(45 * time.Second)
	}
	//robotgo.Click()

}
