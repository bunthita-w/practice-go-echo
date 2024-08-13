package timer

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type TimerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// StartTimer godoc
// @Summary Start timer
// @Description Start timer and return unique ID
// @Tags Timer
// @Produce  json
// @Success 200 {object} TimerResponse
// @Failure 500 {object} TimerResponse
// @Router /startTimer [post]
func StartTimer(store *sessions.CookieStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := uuid.New().String()
		session, _ := store.Get(c.Request(), "stat")
		if session.Values[uid] == nil {
			session.Values[uid] = time.Now().Format(time.RFC3339)
			err := session.Save(c.Request(), c.Response())
			if err != nil {
				return err
			}
		}

		response := TimerResponse{
			Success: true,
			Message: "success",
			Result:  uid,
		}

		return c.JSON(http.StatusOK, response)
	}
}

// CheckTimer godoc
// @Summary Check timer
// @Description Return the time passed in seconds by ID.
// @Tags Timer
// @Produce  json
// @Param id path string true "Timer ID"
// @Success 200 {object} TimerResponse
// @Failure 500 {object} TimerResponse
// @Router /checkTimer/{id} [post]
func CheckTimer(store *sessions.CookieStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		session, _ := store.Get(c.Request(), "stat")
		passedTime, err := findTimePassedFromId(id, session, c)
		if err != nil {
			response := TimerResponse{
				Success: false,
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		response := TimerResponse{
			Success: true,
			Message: "success",
			Result:  fmt.Sprintf("%.0f", passedTime) + " seconds",
		}

		return c.JSON(http.StatusOK, response)
	}
}

// StopTimer godoc
// @Summary Stop timer
// @Description Stop timer and return the time passed in seconds by ID.
// @Tags Timer
// @Produce  json
// @Param id path string true "Timer ID"
// @Success 200 {object} TimerResponse
// @Failure 500 {object} TimerResponse
// @Router /stopTimer/{id} [delete]
func StopTimer(store *sessions.CookieStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		session, _ := store.Get(c.Request(), "stat")
		passedTime, err := findTimePassedFromId(id, session, c)
		if err != nil {
			response := TimerResponse{
				Success: false,
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		response := TimerResponse{
			Success: true,
			Message: "success",
			Result:  fmt.Sprintf("%.0f", passedTime) + " seconds",
		}

		delete(session.Values, id)
		saveErr := session.Save(c.Request(), c.Response())
		if saveErr != nil {
			response := TimerResponse{
				Success: false,
				Message: saveErr.Error(),
			}
			return c.JSON(http.StatusInternalServerError, response)
		}

		return c.JSON(http.StatusOK, response)
	}
}

func findTimePassedFromId(id string, session *sessions.Session, c echo.Context) (float64, error) {
	generatedTime, ok := session.Values[id].(string)
	if !ok {
		return 0, errors.New("cannot find this id")
	}

	parsedTime, err := time.Parse(time.RFC3339, generatedTime)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	passedTime := time.Since(parsedTime).Seconds()

	return passedTime, nil
}
