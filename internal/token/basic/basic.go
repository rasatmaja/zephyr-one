package basic

import (
	"context"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
)

// Token define a filed to support building basic
type Token struct {
	*contract.Token
	signature jwt.Algorithm
}

// New is a function to initialize basic jwt
func New(token *contract.Token) *Token {
	ctx, cancle := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancle()
	basic := &Token{
		Token: token,
	}
	basic.buildSignature(ctx)
	return basic
}

// Sign is a function to signning JWT
func (t *Token) Sign(ctx context.Context, payload *contract.Payload) (string, error) {
	token, err := jwt.Sign(payload, t.signature)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

// Verify is a function to verify JWT
func (t *Token) Verify(ctx context.Context, token string) (*contract.Payload, error) {
	payload := &contract.Payload{}
	_, err := jwt.Verify([]byte(token), t.signature, payload)
	return payload, err
}

func (t *Token) buildSignature(ctx context.Context) {
	switch t.SignAlg {
	case contract.HS256:
		t.signature = jwt.NewHS256([]byte(t.SignKey))
	case contract.HS512:
		t.signature = jwt.NewHS512([]byte(t.SignKey))
	default:
		t.signature = jwt.NewHS256([]byte(t.SignKey))
	}
}
