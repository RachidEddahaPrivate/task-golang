package main

import (
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
	
	webutils.StartEcho(e, config.AddressEcho)
}
