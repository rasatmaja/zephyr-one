package bcrypt

import "golang.org/x/crypto/bcrypt"

// Bcrypt ...
type Bcrypt struct {
	cost int
}

// New is a function to initialize bcrypt
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

// Compare is a function to compare plain string with hashed string
func (b *Bcrypt) Compare(plain, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil, err
}
