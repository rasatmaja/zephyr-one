package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger is a trust to define logger object
type Logger struct{ *zerolog.Logger }

// New is a function to initialize zap logger
func New() *Logger {
	// setup logger
	logger := log.Logger
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return &Logger{&logger}
}
