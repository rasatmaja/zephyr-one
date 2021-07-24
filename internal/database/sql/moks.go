package sql

import "github.com/stretchr/testify/mock"

// Mock define testify/mock type
type Mock struct{ mock.Mock }

// Rollback Mock
func (m *Mock) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

// Commit Mock
func (m *Mock) Commit() error {
	args := m.Called()
	return args.Error(0)
}
