package timer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestStartTimer(t *testing.T) {
	e := echo.New()

	store := sessions.NewCookieStore([]byte("store-key-for-test"))

	req := httptest.NewRequest(http.MethodPost, "/startTimer", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := StartTimer(store)

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response TimerResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Equal(t, "success", response.Message)
		assert.NotEmpty(t, response.Result)
	}
}

func TestCheckTimer(t *testing.T) {
	e := echo.New()

	store := sessions.NewCookieStore([]byte("store-key-for-test"))

	testUID := "test-uuid"
	req := httptest.NewRequest(http.MethodPost, "/checkTimer/"+testUID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	session, _ := store.Get(req, "stat")
	session.Values[testUID] = time.Now().Format(time.RFC3339)
	session.Save(req, rec)

	c.SetParamNames("id")
	c.SetParamValues(testUID)

	handler := CheckTimer(store)

	time.Sleep(1 * time.Second)

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response TimerResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Equal(t, "success", response.Message)
		assert.Contains(t, response.Result, " seconds")
	}
}

func TestStopTimer(t *testing.T) {
	e := echo.New()

	store := sessions.NewCookieStore([]byte("store-key-for-test"))

	testUID := "test-uid"
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/stopTimer/%s", testUID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(testUID)

	session, _ := store.Get(req, "stat")
	session.Values[testUID] = time.Now().Add(-5 * time.Second).Format(time.RFC3339)
	session.Save(req, rec)

	handler := StopTimer(store)

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response TimerResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Equal(t, "success", response.Message)
		assert.Contains(t, response.Result, " seconds")

		session, _ := store.Get(req, "stat")
		_, exists := session.Values[testUID]
		assert.False(t, exists)
	}
}
