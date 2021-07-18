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

//ISQLTX is a interface for sql transaction
type ISQLTX interface {
	Rollback() error
	Commit() error
}

// Queries ...
type Queries struct {
	DB ISQL
}

// IRepository is a interface to define base function for every repo
type IRepository interface {
	BeginTX(ctx context.Context) (IRepository, ISQLTX, error)

	// Auth repo repository
	Auth(ctx context.Context, username string) (*models.Auth, error)
	CreateAuth(ctx context.Context, username, passphrase string) (*models.Auth, error)

	//Account info repository
	Account(ctx context.Context, id string) (*models.AccountInfo, error)
	CreateAccountInfo(ctx context.Context, id, name string) (*models.AccountInfo, error)

	// Contact reposiotry
	CreateContact(ctx context.Context, contact *models.Contact) error
}
