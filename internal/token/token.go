package token

import (
	"context"

	"github.com/gbrlsnchs/jwt/v3"
)

// IToken define interface
type IToken interface {
	Sign(ctx context.Context, payload Payload) error
	Verify(ctx context.Context, token string) (error, Payload)
}

// Payload define JWT field
type Payload struct{ jwt.Payload }

// Token is a JWT struct
type Token struct{}
