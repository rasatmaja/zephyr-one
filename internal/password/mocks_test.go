package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMock(t *testing.T) {
	pwd := &Mock{}

	t.Run("Hash", func(t *testing.T) {
		pwd.On("Hash", mock.Anything).Return("hash", nil)
		hash, err := pwd.Hash("a")
		assert.NoError(t, err)
		assert.Equal(t, "hash", hash)
	})

	t.Run("Compare", func(t *testing.T) {
		pwd.On("Compare", mock.Anything, mock.Anything).Return(true, nil)
		match, err := pwd.Compare("a", "b")
		assert.NoError(t, err)
		assert.Equal(t, true, match)
	})

}
