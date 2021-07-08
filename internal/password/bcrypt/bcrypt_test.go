package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bcrypt := New()
		assert.NotNil(t, bcrypt)
	})
}

func TestHash(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bcrypt := New()
		hash, err := bcrypt.Hash("test-hash-password")
		assert.NotNil(t, bcrypt)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, hash)
	})
}

func TestCompare(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		plain := "secret"
		bcrypt := New()
		if bcrypt == nil {
			t.Fail()
		}
		hash, err := bcrypt.Hash(plain)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(hash))
		match, err := bcrypt.Compare(plain, hash)

		assert.NoError(t, err)
		assert.Equal(t, true, match)
	})

	t.Run("password-not-match", func(t *testing.T) {
		plain := "secret"
		wrong := "wrong-secret"
		bcrypt := New()
		hash, err := bcrypt.Hash(plain)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(hash))
		match, err := bcrypt.Compare(wrong, hash)

		assert.NoError(t, err)
		assert.Equal(t, false, match)
	})

	t.Run("error", func(t *testing.T) {
		plain := "secret"
		wronghash := "wrong-secret"
		bcrypt := New()
		assert.NotNil(t, bcrypt)
		match, err := bcrypt.Compare(plain, wronghash)

		assert.Error(t, err)
		assert.Equal(t, false, match)
	})
}
