package contract

import (
	"context"
)

// SignAlgorithm types
type SignAlgorithm string

// HS512 const
const HS512 SignAlgorithm = "HS512"

// HS256 const
const HS256 SignAlgorithm = "HS365"

// IToken define interface
type IToken interface {
	Sign(ctx context.Context, payload *Payload) (string, error)
	Verify(ctx context.Context, token string) (*Payload, error)
}

// Token define jwt
type Token struct {
	Issuer  string
	SignKey string
	SignAlg SignAlgorithm
	PrivKey interface{}
	PublKey interface{}
}

// Payload define basic JWT field
type Payload struct {
	Issuer         string   `json:"iss,omitempty"`
	Subject        string   `json:"sub,omitempty"`
	Audience       []string `json:"aud,omitempty"`
	ExpirationTime int64    `json:"exp,omitempty"`
	NotBefore      int64    `json:"nbf,omitempty"`
	IssuedAt       int64    `json:"iat,omitempty"`
	JWTID          string   `json:"jti,omitempty"`
}
