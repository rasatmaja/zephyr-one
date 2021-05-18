package config

import (
	"os"
	"sync"

	"github.com/rasatmaja/zephyr-one/internal/logger"
)

var instance *ENV
var singleton sync.Once

// Config ...
type Config struct{ log *logger.Logger }

// ENV is a stuct to hold all environemnt variable for this app
type ENV struct {
	ServerHost    string `mapstructure:"SERVER_HOST"`
	ServerPort    int    `mapstructure:"SERVER_PORT"`
	ServerReadTO  int    `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTO int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
	ServerIdleTO  int    `mapstructure:"SERVER_IDLE_TIMEOUT"`
}

// LoadENV ...
func LoadENV() *ENV {
	singleton.Do(func() {
		config := &Config{log: logger.New()}
		instance = config.BuildENV()
	})
	return instance
}

// BuildENV ...
func (cfg *Config) BuildENV() *ENV {
	env := &ENV{}

	vpr := GetViper()
	vpr.AddConfigPath(".")
	vpr.SetConfigName("app")
	vpr.SetConfigType("env")

	vpr.AutomaticEnv()
	vpr.BindEnvs(env)

	if err := vpr.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		cfg.log.Panic().Msg(err.Error())
	}

	if err := vpr.Unmarshal(&env); err != nil && !os.IsNotExist(err) {
		cfg.log.Panic().Msg(err.Error())
	}

	return env
}
