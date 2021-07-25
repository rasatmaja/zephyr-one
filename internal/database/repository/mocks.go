package repository

import (
	"context"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
	zosql "github.com/rasatmaja/zephyr-one/internal/database/sql"
	"github.com/stretchr/testify/mock"
)

// Mock define testify/mock type
type Mock struct{ mock.Mock }

// BeginTX mock
func (m *Mock) BeginTX(ctx context.Context) (IRepository, zosql.ISQLTX, error) {
	args := m.Called(ctx)
	return args.Get(0).(IRepository), args.Get(1).(zosql.ISQLTX), args.Error(2)
}

// Auth mock
func (m *Mock) Auth(ctx context.Context, username string) (*models.Auth, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*models.Auth), args.Error(1)
}

// CreateAuth mock
func (m *Mock) CreateAuth(ctx context.Context, username, passphrase string) (*models.Auth, error) {
	args := m.Called(ctx, username, passphrase)
	return args.Get(0).(*models.Auth), args.Error(1)
}

// Account mock
func (m *Mock) Account(ctx context.Context, id string) (*models.AccountInfo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.AccountInfo), args.Error(1)
}

// CreateAccountInfo mock
func (m *Mock) CreateAccountInfo(ctx context.Context, id, name string) (*models.AccountInfo, error) {
	args := m.Called(ctx, id, name)
	return args.Get(0).(*models.AccountInfo), args.Error(1)
}

// CreateContact mock
func (m *Mock) CreateContact(ctx context.Context, contact *models.Contact) error {
	args := m.Called(ctx, contact)
	return args.Error(0)
}

// Contacts mocks
func (m *Mock) Contacts(ctx context.Context, authID string, types ...string) ([]*models.Contact, error) {
	args := m.Called(ctx, authID, types)
	return args.Get(0).([]*models.Contact), args.Error(1)
}
