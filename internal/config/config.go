package config

import (
	"fmt"
	"os"
	"sync"
)

var instance *ENV
var singleton sync.Once

// Config ...
type Config struct {
	Filename string
	Type     string
	Path     string
}

// ENV is a stuct to hold all environemnt variable for this app
type ENV struct {
	// Server
	ServerHost       string `mapstructure:"SERVER_HOST"`
	ServerPort       int    `mapstructure:"SERVER_PORT"`
	ServerReadTO     int    `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTO    int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
	ServerIdleTO     int    `mapstructure:"SERVER_IDLE_TIMEOUT"`
	ServerProduction bool   `mapstructure:"SERVER_PRODUCTION"`

	//TLS
	TLS bool `mapstructure:"TLS"`

	// LOG
	LogLevel  string `mapstructure:"LOG_LEVEL"`  // TRACE, DEBUG, INFO, ERROR
	LogOutput string `mapstructure:"LOG_OUTPUT"` // CMD, JSON

	// DABATASE
	DatabaseType        string `mapstructure:"DB_TYPE"`
	DatabaseMaxIDLE     int    `mapstructure:"DB_MAX_IDLE"`     // 25
	DatabaseMaxOpen     int    `mapstructure:"DB_MAX_OPEN"`     // 25
	DatabaseMaxLifetime int    `mapstructure:"DB_MAX_LIFETIME"` // IN MINUTES

	// DATABASE POSTGRESQL
	DBPostgresHost     string `mapstructure:"DB_PG_HOST"`
	DBPostgresPort     int    `mapstructure:"DB_PG_PORT"`
	DBPostgresUsername string `mapstructure:"DB_PG_USER"`
	DBPostgresPassword string `mapstructure:"DB_PG_PASSWORD"`
	DBPostgresDatabase string `mapstructure:"DB_PG_DATABASE"`
	DBPostgresSSLMode  string `mapstructure:"DB_PG_SSLMODE"`
}

// LoadENV ...
func LoadENV() *ENV {
	singleton.Do(func() {
		fmt.Println("[ CNFG ] Starting ENV config ...")
		config := &Config{
			Filename: "app",
			Type:     "env",
			Path:     ".",
		}
		instance = config.BuildENV()
	})
	return instance
}

// BuildENV ...
func (cfg *Config) BuildENV() *ENV {
	env := &ENV{
		ServerProduction: false,
		LogLevel:         "TRACE",
		LogOutput:        "CMD",
	}

	vpr := GetViper()
	vpr.AddConfigPath(cfg.Path)
	vpr.SetConfigName(cfg.Filename)
	vpr.SetConfigType(cfg.Type)

	vpr.AutomaticEnv()
	vpr.BindEnvs(env)

	if err := vpr.ReadInConfig(); err != nil {
		if vpr.IsFileNotFoundError(err) {
			pwd, _ := os.Getwd()
			fmt.Printf("[ CNFG ] File app.env not found on '%s',\n", pwd)
			fmt.Println("[ CNFG ] Using system variable")
		} else {
			panic(err)
		}
	}

	if err := vpr.Unmarshal(&env); err != nil {
		panic(err)
	}

	return env
}
