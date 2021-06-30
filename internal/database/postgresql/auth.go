package postgresql

import (
	"context"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
)

// CreateAuth is a function to insert data into table auth
func (qry *Queries) CreateAuth(ctx context.Context, username, passphrase string) (*models.Auth, error) {
	return nil, nil
}
