package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	//resp, err := http.DefaultClient.Get("https://fp-dev.urbn.com/api/catalog/v0/docs")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//err = ioutil.WriteFile("/Users/p4tin/Dropbox/code/gocode/src/misc/echo-swagger/static/dev_fp-us_catalog.json", body, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//// Echo instance
	e := echo.New()
	e.Use(middleware.Logger())

	e.("/", handler)

	e.Logger.Fatal(e.Start(":1323"))
}

func handler(c *echo.Context) error {
	return nil
}