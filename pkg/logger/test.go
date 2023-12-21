package logger

import (
	"task/pkg/configuration"
)

func InitializeForTest() {
	if internalLogger != nil {
		return
	}
	Initialize(configuration.ConfigLogger{
		LogLevel:    "trace",
		LogSaveFile: false,
	})
}
