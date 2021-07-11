package contract

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// Mock define testify/mock
type Mock struct {
	mock.Mock
}

// Sign Mocks
func (m *Mock) Sign(ctx context.Context, payload *Payload) (string, error) {
	args := m.Called(ctx, payload)
	return args.String(0), args.Error(1)
}

// Verify Mocks
func (m *Mock) Verify(ctx context.Context, token string) (*Payload, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(*Payload), args.Error(1)
}
