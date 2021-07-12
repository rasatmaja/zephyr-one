package basic

import "github.com/stretchr/testify/mock"

// Mocks struct
type Mocks struct{ mock.Mock }

// Name mocks
func (m *Mocks) Name() string {
	args := m.Called()
	return args.String(0)
}

// Sign mocks
func (m *Mocks) Sign(headerPayload []byte) ([]byte, error) {
	args := m.Called(headerPayload)
	return args.Get(0).([]byte), args.Error(1)
}

// Size mocks
func (m *Mocks) Size() int {
	args := m.Called()
	return args.Int(0)
}

// Verify mocks
func (m *Mocks) Verify(headerPayload, sig []byte) error {
	args := m.Called(headerPayload, sig)
	return args.Error(0)
}
