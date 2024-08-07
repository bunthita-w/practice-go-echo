package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddingCalculator(t *testing.T) {
	e := echo.New()
	request := CalculatorRequest{
		FirstNumber:  1,
		SecondNumber: 2,
		Operator:     "+",
	}

	jsonstring, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/calculator", strings.NewReader(string(jsonstring)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedResult := CalculatorResponse{
		Success: true,
		Message: "success",
		Result:  3,
	}

	if assert.NoError(t, Calculator(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response CalculatorResponse
		// err := json.Unmarshal(rec.Body.Bytes(), &response)
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, expectedResult.Result, response.Result)
	}
}

func TestSubstractCalculator(t *testing.T) {
	e := echo.New()
	request := CalculatorRequest{
		FirstNumber:  1,
		SecondNumber: 2,
		Operator:     "-",
	}

	jsonstring, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/calculator", strings.NewReader(string(jsonstring)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedResult := CalculatorResponse{
		Success: true,
		Message: "success",
		Result:  -1,
	}

	if assert.NoError(t, Calculator(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response CalculatorResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, expectedResult.Result, response.Result)
	}
}

func TestMultiplyCalculator(t *testing.T) {
	e := echo.New()
	request := CalculatorRequest{
		FirstNumber:  1,
		SecondNumber: 2,
		Operator:     "*",
	}

	jsonstring, _ := json.Marshal(request)
	fmt.Println(string(jsonstring))
	req := httptest.NewRequest(http.MethodPost, "/calculator", strings.NewReader(string(jsonstring)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedResult := CalculatorResponse{
		Success: true,
		Message: "success",
		Result:  2,
	}

	if assert.NoError(t, Calculator(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response CalculatorResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, expectedResult.Result, response.Result)
	}
}

func TestDevideCalculator(t *testing.T) {
	e := echo.New()
	request := CalculatorRequest{
		FirstNumber:  1,
		SecondNumber: 2,
		Operator:     "/",
	}

	jsonstring, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/calculator", strings.NewReader(string(jsonstring)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedResult := CalculatorResponse{
		Success: true,
		Message: "success",
		Result:  0.5,
	}

	if assert.NoError(t, Calculator(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response CalculatorResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, expectedResult.Result, response.Result)
	}
}

func TestDevideByZeroCalculator(t *testing.T) {
	e := echo.New()
	request := CalculatorRequest{
		FirstNumber:  1,
		SecondNumber: 0,
		Operator:     "/",
	}

	jsonstring, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/calculator", strings.NewReader(string(jsonstring)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Calculator(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestInvalidOperatorCalculator(t *testing.T) {
	e := echo.New()
	request := CalculatorRequest{
		FirstNumber:  1,
		SecondNumber: 0,
		Operator:     "z",
	}

	jsonstring, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/calculator", strings.NewReader(string(jsonstring)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Calculator(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
