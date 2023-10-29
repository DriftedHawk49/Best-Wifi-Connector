package main

import (
	clitools "Best-Wifi-Connector/cli-tools"
	"Best-Wifi-Connector/globals"
	"Best-Wifi-Connector/log"

	"go.uber.org/zap"
)

func main() {

	fl := clitools.FlagParser()

	logLevel := globals.LOGLEVEL_INFO

	if fl.Debug {
		logLevel = globals.LOGLEVEL_DEBUG
	}

	logger := log.InitLogger(logLevel)
	logger.Debug("logger initiated successfully", zap.String("log level", logLevel))
	globals.Logger = logger

}
