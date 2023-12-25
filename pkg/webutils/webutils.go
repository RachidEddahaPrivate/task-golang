package webutils

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"task/pkg/configuration"
	"task/pkg/customerror"
	"task/pkg/logger"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewEcho(config configuration.ConfigEcho) *echo.Echo {
	logger.Info().Msg("Initializing echo")
	e := echo.New()
	logger.Debug().Msg("Setting up echo validator")
	e.Validator = &customValidator{validator: validator.New()}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.AllowedOrigins},
		AllowCredentials: config.AllowCredentials,
	}))

	e.HideBanner = true

	e.HTTPErrorHandler = customerror.ErrorHandler
	logger.Info().Msg("Finished echo initialization")
	return e
}

func StartEcho(echo *echo.Echo, address string) {
	err := echo.Start(address)

	if err != nil {
		_ = echo.Shutdown(context.Background())
	}
	logger.Fatal().Msgf("Cannot start Echo: %v", err)
}
