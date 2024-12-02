package customerror

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
)

const (
	customErrorType      = "ERROR"
	DefaultHttpErrorCode = 400
)

type CustomError struct {
	Code     string                 `json:"code,omitempty"`
	HttpCode int                    `json:"httpCode"`
	Params   map[string]interface{} `json:"params,omitempty"`
	Type     string                 `json:"type"`
}

func NewCustomError(code string) *CustomError {
	return &CustomError{
		Code:     code,
		HttpCode: DefaultHttpErrorCode,
		Type:     customErrorType,
	}
}

func NewI18nErrorWithParams(code string, params map[string]interface{}) *CustomError {
	return &CustomError{
		Code:     code,
		Params:   params,
		HttpCode: DefaultHttpErrorCode,
		Type:     customErrorType,
	}
}

func ErrorHandler(err error, c echo.Context) {
	var customError *CustomError
	if errors.As(err, &customError) {
		_ = c.JSON(customError.HttpCode, customError)
		return
	}
	var errorHTTP *echo.HTTPError
	if errors.As(err, &errorHTTP) {
		_ = c.JSON(errorHTTP.Code, errorHTTP)
		return
	}
	_ = c.JSON(DefaultHttpErrorCode, map[string]interface{}{
		"message": err.Error(),
	})
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("CustomError code: %v, params: %v, httpCode: %v", c.Code, c.Params, c.HttpCode)
}
