package repository

import (
	"context"
	"database/sql"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
	"github.com/stretchr/testify/mock"
)

// Mock define testify/mock type
type Mock struct{ mock.Mock }

// BeginTX mock
func (m *Mock) BeginTX(ctx context.Context) (IRepository, *sql.Tx, error) {
	args := m.Called(ctx)
	return args.Get(0).(IRepository), args.Get(1).(*sql.Tx), args.Error(2)
}

// CreateAuth mock
func (m *Mock) CreateAuth(ctx context.Context, username, passphrase string) (*models.Auth, error) {
	args := m.Called(ctx, username, passphrase)
	return args.Get(0).(*models.Auth), args.Error(1)
}

// CreateAccountInfo mock
func (m *Mock) CreateAccountInfo(ctx context.Context, id, name string) (*models.AccountInfo, error) {
	args := m.Called(ctx, id, name)
	return args.Get(0).(*models.AccountInfo), args.Error(1)
}