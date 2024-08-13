package calculator

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"errors"
)

type CalculatorRequest struct {
	FirstNumber  float64 `json:"firstNumber"`
	SecondNumber float64 `json:"secondNumber"`
	Operator     string  `json:"operator"`
}

type CalculatorResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Result  float64 `json:"result"`
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

// Calculator godoc
// @Summary Calculate numbers with operators
// @Description Received 2 numbers and calculate them with operator
// @Tags Calculator
// @Accept  json
// @Produce  json
// @Param request body CalculatorRequest true "Request Body"
// @Success 200 {object} CalculatorResponse
// @Failure 400 {object} CalculatorResponse
// @Router /calculator [post]
func Calculator(c echo.Context) error {
	var request CalculatorRequest

	if err := c.Bind(&request); err != nil {
		response := CalculatorResponse{
			Success: false,
			Message: "invalid request",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	result, err := calculateNumber(request.FirstNumber, request.SecondNumber, request.Operator)
	if err != nil {
		response := CalculatorResponse{
			Success: false,
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := CalculatorResponse{
		Success: true,
		Message: "success",
		Result:  result,
	}

	return c.JSON(http.StatusOK, response)
}

func calculateNumber(num1 float64, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("invalid operator")
	}
}
