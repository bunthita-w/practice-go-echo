package main

import (
	"practice-go-echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	routes.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
