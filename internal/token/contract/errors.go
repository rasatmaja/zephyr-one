package contract

import "errors"

var (
	// ErrExp define error for expiration date
	ErrExp = errors.New("Token expired")

	// ErrNbf define error for not before date
	ErrNbf = errors.New("Token cant be use right now")

	// ErrIat define error for Issued at date
	ErrIat = errors.New("Token create from future")

	// ErrIss define error for mismatch issuer
	ErrIss = errors.New("Issuer not match")
)
