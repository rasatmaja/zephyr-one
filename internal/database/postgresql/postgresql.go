package postgresql

import (
	"database/sql"
	"fmt"
	"net/url"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/database"
)

// Queries ...
type Queries struct {
	db  database.ISQL
	env *config.ENV
}

// WithTx ..
func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

// New is function to initialize progresql
func New() *Queries {

	// build config env
	env := config.LoadENV()
	queries := &Queries{
		env: env,
	}

	databaseHost := env.DBPostgresHost
	databasePort := env.DBPostgresPort
	databaseUsername := env.DBPostgresUsername
	databasePassword := env.DBPostgresPassword
	databaseName := env.DBPostgresDatabase
	databaseSSLMode := env.DBPostgresSSLMode

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(databaseUsername, databasePassword),
		Host:   fmt.Sprintf("%s:%d", databaseHost, databasePort),
		Path:   databaseName,
	}

	q := dsn.Query()
	q.Add("sslmode", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	queries.db = db
	return queries

}
