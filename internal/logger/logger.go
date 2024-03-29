package logger

import (
	"fmt"
	"os"
	"strings"
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
		fmt.Println("[ LGGR ] Starting Logger ...")
		// setup config env
		env := config.LoadENV()

		// init logger: zerolog
		logger := log.Logger

		switch env.LogOutput {
		case "CMD":
			logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		}
		fmt.Printf("[ LGGR ] Set logger output to %s \n", env.LogOutput)

		var lvl zerolog.Level = zerolog.Disabled
		var err error

		if strings.ToLower(env.LogLevel) != "silent" && strings.ToLower(env.LogLevel) != "disable" {
			lvl, err = zerolog.ParseLevel(strings.ToLower(env.LogLevel))
			if err != nil {
				fmt.Printf("[ LGGR ] Error parse level, got: %s \n", err)

				lvl = zerolog.TraceLevel
				fmt.Println("[ LGGR ] Switch level to TRACE")
			}
		}

		logger = logger.Level(lvl)
		fmt.Printf("[ LGGR ] Set logger level to %s \n", logger.GetLevel())

		instance = &Logger{&logger}
	})
	return instance
}
