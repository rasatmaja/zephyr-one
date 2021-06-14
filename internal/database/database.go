package database

import (
	"context"
	"database/sql"

	"github.com/rasatmaja/zephyr-one/internal/config"
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

// OpenConn ...
func OpenConn() (*sql.DB, error) {
	env := config.LoadENV()
	switch env.DatabaseType {
	case "POSTGRESQL":
		return OpenPGConn()
	default:
		return OpenPGConn()
	}
}
