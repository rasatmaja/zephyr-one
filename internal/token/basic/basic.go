package basic

import (
	"context"
	// blank import
	_ "github.com/gbrlsnchs/jwt/v3"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
)

// Token define a filed to support building basic
type Token contract.Token

// New is a function to initialize basic jwt
func New() *Token {
	return &Token{}
}

// Sign is a function to signning JWT
func (t *Token) Sign(ctx context.Context, payload *contract.Payload) error { return nil }

// Verify is a function to verify JWT
func (t *Token) Verify(ctx context.Context, token string) (*contract.Payload, error) { return nil, nil }
