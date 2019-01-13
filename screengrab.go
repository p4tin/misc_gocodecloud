package main

import (
	"github.com/vova616/screenshot"
	"image"
	"io/ioutil"
	"image/jpeg"
	"bytes"
)

func main() {
	img, _ := screenshot.CaptureScreen() // *image.RGBA
	myImg := image.Image(img) // can cast to image.Image, but not necessary

	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, myImg, nil)
	send_s3 := buf.Bytes()


	ioutil.WriteFile("/Users/Paul/Desktop/tmp/screen.jpg", send_s3, 0644)
}
