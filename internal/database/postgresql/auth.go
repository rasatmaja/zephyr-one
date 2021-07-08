package postgresql

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/rasatmaja/zephyr-one/internal/database/models"
)

// CreateAuth is a function to insert data into table auth
func (qry *Queries) CreateAuth(ctx context.Context, username, passphrase string) (*models.Auth, error) {
	// build auth model
	auth := &models.Auth{
		ID:         uuid.NewString(),
		Username:   username,
		Passphrase: passphrase,
	}

	// insert username and passphrase into database
	_, err := qry.DB.ExecContext(ctx, "INSERT INTO auth(id, username, passphrase) VALUES($1, $2, $3)", auth.ID, auth.Username, auth.Passphrase)

	if err != nil {
		return nil, err
	}

	return auth, nil
}

// Auth is a repo to pull all auth data by its username
func (qry *Queries) Auth(ctx context.Context, username string) (*models.Auth, error) {
	auth := &models.Auth{}

	// build table name, column and var pointer
	table := "auth"
	columns := strings.Join(auth.Columns(auth), ",")
	fields := auth.Fields(auth)

	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = $1 LIMIT 1", columns, table)
	row := qry.DB.QueryRowContext(ctx, query, username)

	// scan all column and put the value into var
	err := row.Scan(fields...)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
