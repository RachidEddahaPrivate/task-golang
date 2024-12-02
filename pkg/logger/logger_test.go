package logger

import (
	"github.com/magiconair/properties/assert"
	"github.com/rs/zerolog"
	"task/pkg/configuration"
	"testing"
)

func TestInitialize(t *testing.T) {
	type args struct {
		config      configuration.ConfigLogger
		wantedLevel zerolog.Level
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "initialize logger with trace level",
			args: args{
				config: configuration.ConfigLogger{
					LogLevel:    "trace",
					LogSaveFile: false,
				},
				wantedLevel: zerolog.TraceLevel,
			},
		},
		{
			name: "initialize logger with debug level",
			args: args{
				config: configuration.ConfigLogger{
					LogLevel:    "debug",
					LogSaveFile: false,
				},
				wantedLevel: zerolog.DebugLevel,
			},
		},
		{
			name: "initialize logger with info level",
			args: args{
				config: configuration.ConfigLogger{
					LogLevel:    "info",
					LogSaveFile: false,
				},
				wantedLevel: zerolog.InfoLevel,
			},
		},
		{
			name: "initialize logger with error level",
			args: args{
				config: configuration.ConfigLogger{
					LogLevel:    "error",
					LogSaveFile: false,
				},
				wantedLevel: zerolog.ErrorLevel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Initialize(tt.args.config)
			assert.Equal(t, internalLogger.GetLevel(), tt.args.wantedLevel)
		})
	}
}
