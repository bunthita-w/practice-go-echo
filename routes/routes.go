package routes

import (
	"github.com/labstack/echo/v4"

	"practice-go-echo/services"
)

func Routes(e *echo.Echo) {
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.GET("/", services.Hello)
	e.POST("/calculator", services.Calculator)
}
