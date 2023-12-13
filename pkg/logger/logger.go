package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"task/pkg/configuration"
	"time"
)

const (
	permissionNumber = 0666 // 0666 - readable and writable by everyone
	pathToLogFile    = "logs/%s.log"
	dateTimeFormat   = "2006-01-02"
)

var internalLogger *zerolog.Logger

func Initialize(config configuration.ConfigLogger) {
	logLevel := calculateLevelLogging(config.LogLevel)
	logOutput := calculateOutputLogging(config.LogSaveFile)

	logger := zerolog.New(logOutput).
		Level(logLevel).
		With().Timestamp().
		Caller().
		Logger()

	internalLogger = &logger
}

func calculateOutputLogging(outputIsFile bool) io.Writer {
	if outputIsFile {
		fileName := fmt.Sprintf(pathToLogFile, time.Now().UTC().Format(dateTimeFormat))
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, permissionNumber)
		if err != nil {
			panic(err)
		}
		return file
	}
	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.Stamp,
	}
}

func calculateLevelLogging(level string) zerolog.Level {
	levelZeroLogger, err := zerolog.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	return levelZeroLogger
}

func Trace() *zerolog.Event {
	if internalLogger == nil {
		panic("Logger not initialized")
	}
	return internalLogger.Trace()
}

func Debug() *zerolog.Event {
	if internalLogger == nil {
		panic("Logger not initialized")
	}
	return internalLogger.Debug()
}

func Info() *zerolog.Event {
	if internalLogger == nil {
		panic("Logger not initialized")
	}
	return internalLogger.Info()
}

func Warn() *zerolog.Event {
	if internalLogger == nil {
		panic("Logger not initialized")
	}
	return internalLogger.Warn()
}

func Error() *zerolog.Event {
	if internalLogger == nil {
		panic("Logger not initialized")
	}
	return internalLogger.Error()
}

func Fatal() *zerolog.Event {
	if internalLogger == nil {
		panic("Logger not initialized")
	}
	return internalLogger.Fatal()
}

func Panic() *zerolog.Event {
	if internalLogger == nil {
		panic("Logger not initialized")
	}
	return internalLogger.Panic()
}
