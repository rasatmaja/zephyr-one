package postgresql

import (
	"database/sql"
	"fmt"
	"net/url"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rasatmaja/zephyr-one/internal/config"
)

// OpenConn is a function to open database connection pool
func OpenConn() (*sql.DB, error) {

	// build config env
	env := config.LoadENV()

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

	// open connection to PostgreSQL server
	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("cannot open connection, got: %v", err)
	}
	return db, nil
}
