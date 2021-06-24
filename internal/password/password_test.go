package password

import (
	"strings"
	"testing"

	"github.com/rasatmaja/zephyr-one/internal/password/bcrypt"
)

func TestFactory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pwd := Factory()
		if pwd == nil {
			t.Fail()
		}
	})
}

func TestHash(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		plain := "secret"

		pwd := Factory()
		if pwd == nil {
			t.Fail()
		}

		hash, err := pwd.Hash(plain)
		pwdID := strings.Split(hash, "$")[1]

		if err != nil || len(hash) == 0 || pwdID != bcryptID {
			t.Fail()
		}
	})

	t.Run("success-default", func(t *testing.T) {
		plain := "secret"

		pwd := &Password{
			bcrypt: bcrypt.New(),
		}
		if pwd == nil {
			t.Fail()
		}

		hash, err := pwd.Hash(plain)
		pwdID := strings.Split(hash, "$")[1]

		if err != nil || len(hash) == 0 || pwdID != bcryptID {
			t.Fail()
		}
	})
}

func TestCompare(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		plain := "secret"

		pwd := Factory()
		if pwd == nil {
			t.Fail()
		}

		hash, err := pwd.Hash(plain)
		pwdID := strings.Split(hash, "$")[1]
		if err != nil || len(hash) == 0 || pwdID != bcryptID {
			t.Fail()
		}

		match, err := pwd.Compare(plain, hash)
		if err != nil || !match {
			t.Fail()
		}

	})

	t.Run("password-not-mach-default-with-unknown-identifier", func(t *testing.T) {
		plain := "secret"

		pwd := Factory()
		if pwd == nil {
			t.Fail()
		}

		hash := "$2b$12$ajsflasjkldjsakdjlksajdasjdlk"

		match, err := pwd.Compare(plain, hash)
		if err == nil || match {
			t.Fail()
		}

	})

	t.Run("password-not-match", func(t *testing.T) {
		plain := "secret"
		wrong := "wrong-secret"

		pwd := Factory()
		if pwd == nil {
			t.Fail()
		}

		hash, err := pwd.Hash(plain)
		pwdID := strings.Split(hash, "$")[1]
		if err != nil || len(hash) == 0 || pwdID != bcryptID {
			t.Fail()
		}

		match, err := pwd.Compare(wrong, hash)
		if err == nil || match {
			t.Fail()
		}
	})
}
