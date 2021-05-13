package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var instance *Logger
var singeton sync.Once

// Logger is a trust to define logger object
type Logger struct{ *zerolog.Logger }

// New is a function to initialize zap logger
func New() *Logger {
	singeton.Do(func() {
		// setup logger
		logger := log.Logger
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		instance = &Logger{&logger}
	})
	return instance
}
