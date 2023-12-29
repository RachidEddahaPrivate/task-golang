package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetEchoParamToInt(contex echo.Context, param string) (int, error) {
	value := contex.Param(param)
	if value == "" {
		return -1, fmt.Errorf("missing param %s", param)
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}
	return result, nil
}
