package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	// PostgreSQL Driver
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/database/repository"
)

// Queries is extended type of database/Queries
type Queries repository.Queries

// New is function to initialize progresql
func New() *Queries {

	queries := &Queries{}

	db, err := queries.OpenConn()
	if err != nil {
		panic(fmt.Errorf("Unable to open connection, got: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("Unable to ping connection, got: %v", err))
	}
	queries.DB = db
	return queries

}

// OpenConn is a function to open database connection pool
func (qry *Queries) OpenConn() (*sql.DB, error) {

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

// BeginTX is a function to start transaction
func (qry *Queries) BeginTX(ctx context.Context) (repository.IRepository, *sql.Tx, error) {
	db, err := qry.OpenConn()
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to open connection, got: %v", err)
	}

	// begin transactions
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to open transaction, got: %v", err)
	}
	queries := &Queries{DB: tx}

	return queries, tx, nil
}
