package contract

import (
	"context"
)

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
	ExpirationTime *Time    `json:"exp,omitempty"`
	NotBefore      *Time    `json:"nbf,omitempty"`
	IssuedAt       *Time    `json:"iat,omitempty"`
	JWTID          string   `json:"jti,omitempty"`
}
