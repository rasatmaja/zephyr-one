package repository

import (
	"context"
	"database/sql"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
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
	DB ISQL
}

// IRepository is a interface to define base function for every repo
type IRepository interface {
	BeginTX(ctx context.Context) (IRepository, *sql.Tx, error)
	CreateAuth(ctx context.Context, username, passphrase string) (*models.Auth, error)
}
