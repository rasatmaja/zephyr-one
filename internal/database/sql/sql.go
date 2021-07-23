package sql

import (
	"context"
	"database/sql"
)

// ISQL is a interface for database/sql
type ISQL interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

//ISQLTX is a interface for sql transaction
type ISQLTX interface {
	Rollback() error
	Commit() error
}

// Queries ...
type Queries struct {
	DB ISQL
}
