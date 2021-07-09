package contract

import (
	"context"
	"time"
)

// IToken define interface
type IToken interface {
	Sign(ctx context.Context, payload *Payload) error
	Verify(ctx context.Context, token string) (*Payload, error)
}

// Token define jwt
type Token struct{}

// Payload define basic JWT field
type Payload struct {
	Issuer         string    `json:"iss,omitempty"`
	Subject        string    `json:"sub,omitempty"`
	Audience       []string  `json:"aud,omitempty"`
	ExpirationTime time.Time `json:"exp,omitempty"`
	NotBefore      time.Time `json:"nbf,omitempty"`
	IssuedAt       time.Time `json:"iat,omitempty"`
	JWTID          string    `json:"jti,omitempty"`
}
