package routes

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"

	"practice-go-echo/services/calculator"
	"practice-go-echo/services/hello"
	"practice-go-echo/services/timer"
)

func Routes(e *echo.Echo, store *sessions.CookieStore) {
	e.GET("/", hello.Hello(e, store))
	e.POST("/calculator", calculator.Calculator)
	e.POST("/startTimer", timer.StartTimer(store))
	e.POST("/checkTimer/:id", timer.CheckTimer(store))
	e.DELETE("/stopTimer/:id", timer.StopTimer(store))
}
