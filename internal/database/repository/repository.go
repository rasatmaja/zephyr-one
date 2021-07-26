package repository

import (
	"context"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
	zosql "github.com/rasatmaja/zephyr-one/internal/database/sql"
)

// IRepository is a interface to define base function for every repo
type IRepository interface {
	BeginTX(ctx context.Context) (IRepository, zosql.ISQLTX, error)

	// Auth repo repository
	Auth(ctx context.Context, username string) (*models.Auth, error)
	CreateAuth(ctx context.Context, username, passphrase string) (*models.Auth, error)

	//Account info repository
	Account(ctx context.Context, id string) (*models.AccountInfo, error)
	CreateAccountInfo(ctx context.Context, id, name string) (*models.AccountInfo, error)

	// Contact reposiotry
	CreateContact(ctx context.Context, contact *models.Contact) error
	Contacts(ctx context.Context, authID string, types ...string) ([]*models.Contact, error)
	SetPrimaryContact(ctx context.Context, authID, contact string) error
}
