package logger

import (
	"fmt"
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
		env := config.LoadENV()

		// init logger: zerolog
		logger := log.Logger

		fmt.Printf("[LOGGER] Set logger output to %s \n", env.LogOutput)
		switch env.LogOutput {
		case "CMD":
			logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		}

		instance = &Logger{&logger}
	})
	return instance
}
