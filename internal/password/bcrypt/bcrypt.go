package bcrypt

import "golang.org/x/crypto/bcrypt"

// Bcrypt ...
type Bcrypt struct {
	cost int
}

// New ...
func New() *Bcrypt {
	bcrypt := &Bcrypt{
		cost: bcrypt.DefaultCost,
	}
	return bcrypt
}

// Hash is a function to hash string with bcrypt algorithm
func (b *Bcrypt) Hash(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), b.cost)
	return string(hash), err
}
