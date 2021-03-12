package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"

	"github.com/vova616/screenshot"
)

func main() {
	img, _ := screenshot.CaptureScreen() // *image.RGBA
	myImg := image.Image(img)            // can cast to image.Image, but not necessary

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, myImg, nil)
	send_s3 := buf.Bytes()

	ioutil.WriteFile("/Users/Paul/Desktop/tmp/screen.jpg", send_s3, 0644)
}
