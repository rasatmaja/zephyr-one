package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumns(t *testing.T) {
	t.Run("column", func(t *testing.T) {
		auth := &Auth{}
		columns := Columns(auth)

		expectedCols := []string{"id", "username", "passphrase", "created_at", "updated_at"}
		assert.Equal(t, expectedCols, columns)
	})

	t.Run("no-column", func(t *testing.T) {
		type test struct {
			test string `nocolumn:"test"`
		}
		x := &test{}
		columns := Columns(x)

		assert.Empty(t, columns)
	})
}

func TestFields(t *testing.T) {
	t.Run("fields", func(t *testing.T) {
		auth := &Auth{}
		flds := Fields(auth)
		assert.NotEmpty(t, flds)
	})
}
