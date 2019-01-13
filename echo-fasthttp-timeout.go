package main

import (
	"log"
	"time"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
)

func main() {
	echo := echo.New()
	echo.GET("/", func(c echo.Context) error {
		//log.Println("Request GET '/'")
		res := []string{"Hello", "World"}
		return c.JSON(http.StatusOK, res)
	})
	fasthttp_engine := fasthttp.New(":8080")
	fasthttp_engine.ReadTimeout = 2 * time.Second
	fasthttp_engine.WriteTimeout = 2 * time.Second
	log.Println("Starting Server on Port :8080")
	log.Fatal(echo.Run(fasthttp_engine))
}
