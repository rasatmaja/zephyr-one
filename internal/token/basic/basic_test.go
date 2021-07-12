package basic

import (
	"context"
	"fmt"
	"testing"

	"github.com/rasatmaja/zephyr-one/internal/token/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	bsc := New(&contract.Token{
		SignKey: "secret",
	})
	assert.NotNil(t, bsc)
}

func TestBuildSignature(t *testing.T) {
	ctx := context.Background()
	tkn := &Token{
		Token: &contract.Token{
			SignKey: "secret",
		},
		signature: nil,
	}
	t.Run("HS265", func(t *testing.T) {
		tkn.SignAlg = contract.HS256
		tkn.buildSignature(ctx)
		assert.NotNil(t, tkn)
	})

	t.Run("HS512", func(t *testing.T) {
		tkn.SignAlg = contract.HS512
		tkn.buildSignature(ctx)
		assert.NotNil(t, tkn)
	})

	t.Run("default", func(t *testing.T) {
		tkn.SignAlg = "default"
		tkn.buildSignature(ctx)
		assert.NotNil(t, tkn)
	})
}

func TestSetTimestamp(t *testing.T) {
	ctx := context.Background()
	tkn := &Token{
		Token: &contract.Token{
			SignKey: "secret",
		},
		signature: nil,
	}
	payload := &contract.Payload{}
	err := tkn.setTimestamps(ctx, payload)
	assert.NoError(t, err)
	assert.NotNil(t, payload)
}

func TestSetIssuer(t *testing.T) {
	ctx := context.Background()
	tkn := &Token{
		Token: &contract.Token{
			SignKey: "secret",
		},
		signature: nil,
	}
	t.Run("empty-issuer", func(t *testing.T) {
		payload := &contract.Payload{}
		err := tkn.setIssuer(ctx, payload)
		assert.EqualError(t, err, contract.ErrEmptyIss.Error())
	})

	t.Run("success", func(t *testing.T) {
		tkn.Issuer = "test"
		payload := &contract.Payload{}
		err := tkn.setIssuer(ctx, payload)
		assert.NoError(t, err)
		assert.Equal(t, tkn.Issuer, payload.Issuer)
	})
}

func TestSign(t *testing.T) {
	ctx := context.Background()
	tkn := &Token{
		Token: &contract.Token{
			SignKey: "secret",
		},
		signature: nil,
	}
	t.Run("empty-issuer", func(t *testing.T) {
		payload := &contract.Payload{}
		jwt, err := tkn.Sign(ctx, payload)
		assert.EqualError(t, err, contract.ErrEmptyIss.Error())
		assert.Equal(t, 0, len(jwt))
	})

	t.Run("error-sign", func(t *testing.T) {

		// start mock
		alg := &Mocks{}
		alg.On("Name").Return("unknown")
		alg.On("Size").Return(1)
		alg.On("Sign", mock.Anything).Return([]byte(""), fmt.Errorf("error"))
		tkn.signature = alg

		tkn.Issuer = "test"
		payload := &contract.Payload{}

		jwt, err := tkn.Sign(ctx, payload)
		assert.EqualError(t, err, "error")
		assert.Equal(t, 0, len(jwt))
	})

	t.Run("success", func(t *testing.T) {

		// start mock
		alg := &Mocks{}
		alg.On("Name").Return("unknown")
		alg.On("Size").Return(0)
		alg.On("Sign", mock.Anything).Return([]byte(""), nil)
		tkn.signature = alg

		tkn.Issuer = "test"
		payload := &contract.Payload{}

		jwt, err := tkn.Sign(ctx, payload)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(jwt))
	})
}
