package database

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/rasatmaja/zephyr-one/internal/config"

	// PostgreSQL Driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

// ISQL is a interface for database/sql
type ISQL interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// Queries ...
type Queries struct {
	db ISQL
}

// IRepository is a interface to define base function for every repo
type IRepository interface {
	BeginTX(ctx context.Context) (*Queries, *sql.Tx, error)
}

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

	db.SetMaxOpenConns(env.DatabaseMaxOpen)
	db.SetMaxIdleConns(env.DatabaseMaxIDLE)
	db.SetConnMaxLifetime(time.Duration(env.DatabaseMaxLifetime) * time.Minute)

	return db, nil
}
