package logger

import (
	"os"
	"sync"

	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var instance *Logger
var singleton sync.Once

// Logger is a trust to define logger object
type Logger struct{ *zerolog.Logger }

// New is a function to initialize zap logger
func New() *Logger {
	singleton.Do(func() {
		// setup config env
		_ = config.LoadENV()

		// setup logger
		logger := log.Logger
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		instance = &Logger{&logger}
	})
	return instance
}
