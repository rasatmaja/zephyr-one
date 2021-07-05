package helper

import (
	"crypto/rand"
	"io"
	"math/big"
)

// Strings helper struct
type Strings struct {
	letters string
	reader  io.Reader
}

// NewStrings return a Strings
func NewStrings() *Strings {
	return &Strings{
		letters: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-",
		reader:  rand.Reader,
	}
}

// GenerateRandomString is a function to generate random string securly using crypto package
func (s *Strings) GenerateRandomString(length int) (string, error) {
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(s.reader, big.NewInt(int64(len(s.letters))))
		if err != nil {
			return "", err
		}
		ret[i] = s.letters[num.Int64()]
	}

	return string(ret), nil
}
