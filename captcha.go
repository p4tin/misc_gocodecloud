package main

import (
	"github.com/fogleman/gg"
	"log"
	"math/rand"
	"time"
)

var angles []float64 = []float64{ 330, 340, 350, 0, 10, 20, 30 }
var fonts []string = []string{
	"/Library/Fonts/Arial Bold.ttf",
	"/Library/Fonts/Bradley Hand Bold.ttf",
	"/Library/Fonts/Comic Sans MS.ttf",
	"/Library/Fonts/Trebuchet MS Bold.ttf",
	"/Library/Fonts/DIN Alternate Bold.ttf",
	"/Library/Fonts/Georgia Bold.ttf",
	"/Library/Fonts/Tahoma Bold.ttf",
}
var logos []string = []string{
	"adcaptcha",
	"hlt",
	"ford",
	"intel",
	"vintage",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) []rune {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return b
}

func getRandomInt(size int) int {
	flt := rand.Intn(size)
	return flt
}

func getRandomFloat() float64 {
	flt := rand.Float64()
	return flt
}

func main() {
	runes := RandStringRunes(6)

	im, err := gg.LoadImage("static/logos/" + logos[getRandomInt(5)] + ".png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(259, 102)
	dc.SetRGB(0, 0, 25)

	xpos := float64(36)
	dc.DrawImage(im, 0, 0)


	for x:=0;x<6;x++ {
		if err := dc.LoadFontFace(fonts[getRandomInt(7)], 36); err != nil {
			panic(err)
		}
		a := float64(getRandomInt(7))
		dc.Rotate(gg.Radians(a))
		dc.DrawStringAnchored(string(runes[x]), xpos + float64(x*36), 102/2, getRandomFloat(), getRandomFloat())
		dc.Rotate(gg.Radians(360-a))
	}


	dc.Clip()
	dc.SavePNG("static/out.png")
}