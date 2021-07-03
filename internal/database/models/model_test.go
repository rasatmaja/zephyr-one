package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModels(t *testing.T) {
	type Test struct {
		Models
		A string `column:"a"`
		B string `column:"b"`
		C string
	}
	t.Run("Column", func(t *testing.T) {
		x := &Test{}
		result := x.Columns(x)
		expected := []string{"a", "b"}
		assert.Equal(t, expected, result)
	})

	t.Run("Field", func(t *testing.T) {
		x := &Test{}
		result := x.Fields(x)
		expected := []interface{}{&x.A, &x.B}
		assert.Equal(t, expected, result)
	})
}
