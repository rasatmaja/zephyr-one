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
	ServerHost       string `mapstructure:"SERVER_HOST" default:"localhost"`
	ServerPort       int    `mapstructure:"SERVER_PORT" default:"2929"`
	ServerReadTO     int    `mapstructure:"SERVER_READ_TIMEOUT" default:"10"`
	ServerWriteTO    int    `mapstructure:"SERVER_WRITE_TIMEOUT" default:"10"`
	ServerIdleTO     int    `mapstructure:"SERVER_IDLE_TIMEOUT" default:"10"`
	ServerProduction bool   `mapstructure:"SERVER_PRODUCTION" default:"false"`

	//TLS
	TLS bool `mapstructure:"TLS" default:"false"`

	// LOG
	LogLevel  string `mapstructure:"LOG_LEVEL" default:"ERROR"` // TRACE, DEBUG, INFO, ERROR
	LogOutput string `mapstructure:"LOG_OUTPUT" default:"CMD"`  // CMD, JSON

	// DABATASE
	DatabaseType        string `mapstructure:"DB_TYPE" default:"POSTGRESQL"`
	DatabaseMaxIDLE     int    `mapstructure:"DB_MAX_IDLE" default:"25"`     // 25
	DatabaseMaxOpen     int    `mapstructure:"DB_MAX_OPEN" default:"25"`     // 25
	DatabaseMaxLifetime int    `mapstructure:"DB_MAX_LIFETIME" default:"10"` // IN MINUTES

	// DATABASE POSTGRESQL
	DBPostgresHost     string `mapstructure:"DB_PG_HOST" default:"localhost"`
	DBPostgresPort     int    `mapstructure:"DB_PG_PORT" default:"5432"`
	DBPostgresUsername string `mapstructure:"DB_PG_USER" default:"root"`
	DBPostgresPassword string `mapstructure:"DB_PG_PASSWORD" default:"root"`
	DBPostgresDatabase string `mapstructure:"DB_PG_DATABASE" default:"zepheryone"`
	DBPostgresSSLMode  string `mapstructure:"DB_PG_SSLMODE" default:"disable"`

	// TOKEN JWT
	TokenType    string `mapstructure:"TOKEN_TYPE" default:"BASIC"`
	TokenSignKey string `mapstructure:"TOKEN_SIGN_KEY" default:"secret"`
	TokenSignAlg string `mapstructure:"TOKEN_SIGN_ALG" default:"HS256"`
	TokenIssuer  string `mapstructure:"TOKEN_ISSUER" default:"zephery-one"`
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
	env := &ENV{}

	vpr := GetViper()
	vpr.AddConfigPath(cfg.Path)
	vpr.SetConfigName(cfg.Filename)
	vpr.SetConfigType(cfg.Type)

	vpr.FillDefault(env)
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
