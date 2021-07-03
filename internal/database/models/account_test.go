package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	t.Run("Column", func(t *testing.T) {
		acc := &AccountInfo{}
		result := acc.Columns()

		assert.Contains(t, result, "id")
	})

	t.Run("Field", func(t *testing.T) {
		acc := &AccountInfo{}
		result := acc.Fields()

		assert.Contains(t, result, &acc.ID)
	})
}
