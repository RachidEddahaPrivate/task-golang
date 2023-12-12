package main

import (
	"task/internal/controller"
	"task/internal/service"
	"task/pkg/configuration"
	"task/pkg/logger"
	"task/pkg/webutils"
)

func main() {
	config, err := configuration.LoadConfiguration()
	if err != nil {
		panic(err)
	}
	logger.Initialize(config.ConfigLogger)

	e := webutils.NewEcho(config.ConfigEcho)
	logger.Info().Msg("Finished configuration")

	logger.Info().Msg("Starting services")

	repository := service.NewRepository()
	svc := service.NewService(repository)
	controller.NewController(svc).RegisterRoutes(e)

	webutils.StartEcho(e, config.AddressEcho)
}
