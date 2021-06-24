package password

import (
	"strings"

	"github.com/rasatmaja/zephyr-one/internal/password/bcrypt"
)

// IPassword is a interface to define password functions
type IPassword interface {
	Hash(plain string) (string, error)
	Compare(plain, hash string) (bool, error)
}

const (
	bcryptID = "2a"       // bcrypt identifier
	argonID  = "argon2id" // argon2ID identifier
)

// Password is a struct that hold password config
type Password struct {
	types  string
	bcrypt IPassword
	argon  IPassword
}

// Factory is a function to build password config
func Factory() *Password {
	pwd := &Password{}
	pwd.types = "BCRYPT"
	pwd.bcrypt = bcrypt.New()
	return pwd
}

// Hash is function to hash plain password
func (pwd *Password) Hash(plain string) (string, error) {
	switch pwd.types {
	case "BCRYPT":
		return pwd.bcrypt.Hash(plain)
	default:
		return pwd.bcrypt.Hash(plain)
	}
}

// Compare is a function to compare plain password and hashed password
func (pwd *Password) Compare(plain, hash string) (bool, error) {
	pwdID := strings.Split(hash, "$")[1] // get hash identifier
	switch pwdID {
	case bcryptID:
		return pwd.bcrypt.Compare(plain, hash)
	default:
		return pwd.bcrypt.Compare(plain, hash)
	}
}
