package config

import (
	"fmt"
	"os"

	"github.com/rasatmaja/zephyr-one/internal/logger"
)

var log *logger.Logger

// Config ...
type Config struct {
	ServerHost    string `mapstructure:"SERVER_HOST"`
	ServerPort    int    `mapstructure:"SERVER_PORT"`
	ServerReadTO  int    `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTO int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
	ServerIdleTO  int    `mapstructure:"SERVER_IDLE_TIMEOUT"`
}

func init() {
	log = logger.New()
}

// BuildConfig ...
func BuildConfig() *Config {
	cfg := &Config{}

	vpr := GetViper()
	vpr.AddConfigPath(".")
	vpr.SetConfigName("app")
	vpr.SetConfigType("env")

	vpr.AutomaticEnv()
	vpr.BindEnvs(cfg)

	if err := vpr.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		log.Panic().Msg(err.Error())
	}

	if err := vpr.Unmarshal(&cfg); err != nil && !os.IsNotExist(err) {
		log.Panic().Msg(err.Error())
	}

	vpr.Debug()

	fmt.Println(cfg)
	return cfg
}
