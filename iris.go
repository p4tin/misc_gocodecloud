package main

import "github.com/kataras/iris"

func main() {
	options := iris.StationOptions{
		Cache:              true,
		Profile:            true,
		ProfilePath:        "/mypath/debug",
	}

	api := iris.Custom(options)
	api.Listen(":8080")
}