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

func TestCompare(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		plain := "secret"
		bcrypt := New()
		if bcrypt == nil {
			t.Fail()
		}
		hash, err := bcrypt.Hash(plain)
		if err != nil || len(hash) == 0 {
			t.Fail()
		}
		match, err := bcrypt.Compare(plain, hash)

		if err != nil || !match {
			t.Fail()
		}
	})

	t.Run("password-not-match", func(t *testing.T) {
		plain := "secret"
		wrong := "wrong-secret"
		bcrypt := New()
		if bcrypt == nil {
			t.Fail()
		}
		hash, err := bcrypt.Hash(plain)
		if err != nil || len(hash) == 0 {
			t.Fail()
		}
		match, err := bcrypt.Compare(wrong, hash)

		if err != nil || match {
			t.Fail()
		}
	})

	t.Run("error", func(t *testing.T) {
		plain := "secret"
		wronghash := "wrong-secret"
		bcrypt := New()
		if bcrypt == nil {
			t.Fail()
		}
		match, err := bcrypt.Compare(plain, wronghash)

		if err == nil || match {
			t.Fail()
		}
	})
}
