package config

import (
	"fmt"
	"os"
)

// Config ...
type Config struct {
	ServerHost    string `mapstructure:"SERVER_HOST"`
	ServerPort    int    `mapstructure:"SERVER_PORT"`
	ServerReadTO  int    `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTO int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
	ServerIdleTO  int    `mapstructure:"SERVER_IDLE_TIMEOUT"`
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
		panic(err)
	}

	if err := vpr.Unmarshal(&cfg); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	vpr.Debug()

	fmt.Println(cfg)
	return cfg
}
