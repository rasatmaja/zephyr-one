package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumns(t *testing.T) {
	t.Run("column", func(t *testing.T) {
		type test struct {
			a string `column:"a"`
			b string `column:"b"`
		}
		x := &test{}
		columns := columns(x)
		expectedCols := []string{"a", "b"}
		assert.Equal(t, expectedCols, columns)
	})

	t.Run("no-column", func(t *testing.T) {
		type test struct {
			test string `nocolumn:"test"`
		}
		x := &test{}
		columns := columns(x)

		assert.Empty(t, columns)
	})
}

func TestFields(t *testing.T) {
	t.Run("fields", func(t *testing.T) {
		type test struct {
			A string `column:"a"`
			B string `column:"b"`
			C string
		}
		x := &test{}
		flds := fields(x)

		expected := []interface{}{&x.A, &x.B}

		assert.Equal(t, expected, flds)
		assert.NotEmpty(t, flds)
	})
}
