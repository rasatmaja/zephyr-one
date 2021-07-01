package postgresql

import (
	"context"

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
