package hello

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

// Hello godoc
// @Summary Hello
// @Description Show greeting text with API stats
// @Tags Hello
// @Produce  json
// @Success 200 {string} string "Greeting message with API stats"
// @Router / [get]
func Hello(e *echo.Echo, store *sessions.CookieStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "stat")

		duration, err := calculateTimePassed(session)
		textHello := "Hello World!"
		if err == nil {
			durationText := formatDurationText(duration)
			textHello += durationText
		}

		textHello += formatNumberOfCallingEachAPIText(e, session)

		return c.String(http.StatusOK, textHello)
	}
}

func calculateTimePassed(session *sessions.Session) (time.Duration, error) {
	startServerTimeStr, ok := session.Values["startServerTime"].(string)
	if !ok {
		return 0, errors.New("startServerTime is not string")
	}

	startServerTime, err := time.Parse(time.RFC3339, startServerTimeStr)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	duration := time.Since(startServerTime)

	return duration, nil
}

func formatDurationText(duration time.Duration) string {
	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	durationText := "\n I've been running for " + strconv.Itoa(days) + " days, " + strconv.Itoa(hours) + " hours, " + strconv.Itoa(minutes) + " minutes, " + strconv.Itoa(seconds) + " seconds."
	return durationText
}

func formatNumberOfCallingEachAPIText(e *echo.Echo, session *sessions.Session) string {
	routes := e.Routes()
	numOfCallingApiText := "\nAPI requested from start"
	for _, route := range routes {
		numberOfCalling, ok := session.Values[route.Path].(int)
		// fmt.Println("formattest")
		// fmt.Println("Route Path:", route.Path)
		// fmt.Println("Number of Calling:", numberOfCalling)
		// fmt.Println("OK:", ok)
		if !ok {
			numberOfCalling = 0
		}

		numOfCallingApiText += "\n" + route.Path + " : " + strconv.Itoa(numberOfCalling) + " times"
	}

	return numOfCallingApiText
}
