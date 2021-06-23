package crypto

// IPassword is a interface to define password functions
type IPassword interface {
	Hash(plain string) (string, error)
	Compare(plain, hash string) (bool, error)
}
