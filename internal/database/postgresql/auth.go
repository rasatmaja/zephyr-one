package postgresql

import (
	"context"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
)

// CreateAuth is a function to insert data into table auth
func (qry *Queries) CreateAuth(ctx context.Context, username, passphrase string) (*models.Auth, error) {
	// insert username and passphrase into database
	_, err := qry.DB.ExecContext(ctx, "INSERT INTO auth(username, passphrase) VALUES($1,$2)", username, passphrase)

	if err != nil {
		return nil, err
	}

	auth := &models.Auth{}
	// select auth data
	row := qry.DB.QueryRowContext(ctx, "SELECT id, username, passphrase FROM auth WHERE username = $1", username)
	err = row.Scan(&auth.ID, &auth.Username, &auth.Passphrase)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
