package config

import (
	"fmt"
	"sync"
)

var instance *ENV
var singleton sync.Once

// Config ...
type Config struct{}

// ENV is a stuct to hold all environemnt variable for this app
type ENV struct {
	// Server
	ServerHost    string `mapstructure:"SERVER_HOST"`
	ServerPort    int    `mapstructure:"SERVER_PORT"`
	ServerReadTO  int    `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTO int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
	ServerIdleTO  int    `mapstructure:"SERVER_IDLE_TIMEOUT"`

	// LOG
	LogLevel  string `mapstructure:"LOG_LEVEL"`  // TRACE, DEBUG, INFO, ERROR
	LogOutput string `mapstructure:"LOG_OUTPUT"` // CMD, JSON
}

// LoadENV ...
func LoadENV() *ENV {
	singleton.Do(func() {
		config := &Config{}
		instance = config.BuildENV()
	})
	return instance
}

// BuildENV ...
func (cfg *Config) BuildENV() *ENV {
	env := &ENV{
		LogLevel:  "TRACE",
		LogOutput: "CMD",
	}

	vpr := GetViper()
	vpr.AddConfigPath(".")
	vpr.SetConfigName("app")
	vpr.SetConfigType("env")

	vpr.AutomaticEnv()
	vpr.BindEnvs(env)

	if err := vpr.ReadInConfig(); err != nil {
		if vpr.IsFileNotFoundError(err) {
			fmt.Println("[ CNFG ] File app.env not found on root directory, using system variable")
		} else {
			panic(err)
		}
	}

	if err := vpr.Unmarshal(&env); err != nil {
		panic(err)
	}

	return env
}
