package bcrypt

import "testing"

func TestNew(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bcrypt := New()
		if bcrypt == nil {
			t.Fail()
		}
	})
}

func TestHash(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bcrypt := New()
		hash, err := bcrypt.Hash("test-hash-password")
		if bcrypt == nil || err != nil || len(hash) == 0 {
			t.Fail()
		}
	})
}
