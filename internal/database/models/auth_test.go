package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	t.Run("Column", func(t *testing.T) {
		auth := &Auth{}
		result := auth.Columns()

		expected := []string{"id", "username", "passphrase", "created_at", "updated_at"}

		assert.Equal(t, expected, result)
	})

	t.Run("Column", func(t *testing.T) {
		auth := &Auth{}
		result := auth.Fields()

		expected := []interface{}{&auth.ID, &auth.Username, &auth.Passphrase, &auth.CreatedAt, &auth.UpdatedAt}

		assert.Equal(t, expected, result)
	})
}
