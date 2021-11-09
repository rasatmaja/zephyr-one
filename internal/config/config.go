package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/rasatmaja/mura/v2"
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
	ServerHost       string `env:"SERVER_HOST" default:"localhost"`
	ServerPort       int    `env:"SERVER_PORT" default:"2929"`
	ServerReadTO     int    `env:"SERVER_READ_TIMEOUT" default:"10"`
	ServerWriteTO    int    `env:"SERVER_WRITE_TIMEOUT" default:"10"`
	ServerIdleTO     int    `env:"SERVER_IDLE_TIMEOUT" default:"10"`
	ServerProduction bool   `env:"SERVER_PRODUCTION" default:"false"`

	//TLS
	TLS bool `env:"TLS" default:"false"`

	// LOG
	LogLevel  string `env:"LOG_LEVEL" default:"ERROR"` // TRACE, DEBUG, INFO, ERROR
	LogOutput string `env:"LOG_OUTPUT" default:"CMD"`  // CMD, JSON

	// DABATASE
	DatabaseType        string `env:"DB_TYPE" default:"POSTGRESQL"`
	DatabaseMaxIDLE     int    `env:"DB_MAX_IDLE" default:"25"`     // 25
	DatabaseMaxOpen     int    `env:"DB_MAX_OPEN" default:"25"`     // 25
	DatabaseMaxLifetime int    `env:"DB_MAX_LIFETIME" default:"10"` // IN MINUTES

	// DATABASE POSTGRESQL
	DBPostgresHost     string `env:"DB_PG_HOST" default:"localhost"`
	DBPostgresPort     int    `env:"DB_PG_PORT" default:"5432"`
	DBPostgresUsername string `env:"DB_PG_USER" default:"root"`
	DBPostgresPassword string `env:"DB_PG_PASSWORD" default:"root"`
	DBPostgresDatabase string `env:"DB_PG_DATABASE" default:"zepheryone"`
	DBPostgresSSLMode  string `env:"DB_PG_SSLMODE" default:"disable"`

	// TOKEN JWT
	TokenType    string `env:"TOKEN_TYPE" default:"BASIC"`
	TokenSignKey string `env:"TOKEN_SIGN_KEY" default:"secret"`
	TokenSignAlg string `env:"TOKEN_SIGN_ALG" default:"HS256"`
	TokenIssuer  string `env:"TOKEN_ISSUER" default:"zephery-one"`
}

// LoadENV ...
func LoadENV() *ENV {
	singleton.Do(func() {
		fmt.Println("[ CNFG ] Starting ENV config ...")
		config := new(Config)
		instance = config.BuildENV()
	})
	return instance
}

// BuildENV ...
func (cfg *Config) BuildENV() *ENV {
	env := new(ENV)
	mura.SetENVPath(cfg.Path)
	mura.Unmarshal(env)
	return env
}

func init() {
	log.SetPrefix("")
}
