package password

import "github.com/stretchr/testify/mock"

// Mock define testify/mock
type Mock struct {
	mock.Mock
}

// Hash Mocks
func (m *Mock) Hash(plain string) (string, error) {
	args := m.Called(plain)
	return args.String(0), args.Error(1)
}

// Compare Mocks
func (m *Mock) Compare(plain, hash string) (bool, error) {
	args := m.Called(plain, hash)
	return args.Bool(0), args.Error(1)
}
