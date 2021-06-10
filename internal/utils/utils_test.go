package utils

import "testing"

func TestUtils(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		New()
		if instance == nil {
			t.Error("instance shouldnt be nil")
			t.Fail()
		}
	})
}
