package basic

import (
	"context"
	"fmt"
	"testing"

	"github.com/gbrlsnchs/jwt/v3"
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

func TestVerify(t *testing.T) {
	ctx := context.Background()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6ZXByaHlyLW9uZSIsInN1YiI6InJhc2lvMjkiLCJleHAiOjE2Mjg2NzcwMTIsIm5iZiI6MTYyNjA4NTAxMiwiaWF0IjoxNjI2MDg1MDEyLCJqdGkiOiIwMmY5Zjc5Zi0xMWM2LTRhNmQtYjM4OC01NjRiM2RlMzc5ZDEifQ.pmo0AQVknJ8BhIZBNYFNVyNm8545FTvTVu0x1IT4lfI"

	tkn := &Token{
		Token: &contract.Token{
			SignKey: "secret",
		},
		signature: nil,
	}
	tkn.Issuer = "test"

	t.Run("error-verify", func(t *testing.T) {

		// start mock
		alg := &Mocks{}
		alg.On("Name").Return("unknown")
		alg.On("Size").Return(1)
		alg.On("Verify", mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
		tkn.signature = alg

		payload, err := tkn.Verify(ctx, token)
		assert.Error(t, err)
		assert.EqualError(t, err, "error")
		assert.Nil(t, payload)
	})

	t.Run("token-expired", func(t *testing.T) {

		tkn.buildSignature(ctx)
		tkn.Issuer = "test"

		payload := &contract.Payload{
			ExpirationTime: contract.TimeNow().AddDates(0, 0, -1),
		}
		tokenjwt, err := jwt.Sign(payload, tkn.signature)
		payloadRes, err := tkn.Verify(ctx, string(tokenjwt))
		assert.Error(t, err)
		assert.EqualError(t, err, contract.ErrExp.Error())
		assert.Nil(t, payloadRes)
	})

	t.Run("token-issued-from-future", func(t *testing.T) {

		tkn.buildSignature(ctx)

		payload := &contract.Payload{
			ExpirationTime: contract.TimeNow().AddDates(0, 0, 5),
			IssuedAt:       contract.TimeNow().AddDates(0, 0, 1),
		}
		tokenjwt, err := jwt.Sign(payload, tkn.signature)
		payloadRes, err := tkn.Verify(ctx, string(tokenjwt))
		assert.Error(t, err)
		assert.EqualError(t, err, contract.ErrIat.Error())
		assert.Nil(t, payloadRes)
	})

	t.Run("token-not-before", func(t *testing.T) {

		tkn.buildSignature(ctx)

		payload := &contract.Payload{
			ExpirationTime: contract.TimeNow().AddDates(0, 0, 5),
			IssuedAt:       contract.TimeNow(),
			NotBefore:      contract.TimeNow().AddDates(0, 0, 1),
		}
		tokenjwt, err := jwt.Sign(payload, tkn.signature)
		payloadRes, err := tkn.Verify(ctx, string(tokenjwt))
		assert.Error(t, err)
		assert.EqualError(t, err, contract.ErrNbf.Error())
		assert.Nil(t, payloadRes)
	})

	t.Run("token-mismatch-issuer", func(t *testing.T) {

		tkn.buildSignature(ctx)

		payload := &contract.Payload{
			ExpirationTime: contract.TimeNow().AddDates(0, 0, 5),
			IssuedAt:       contract.TimeNow(),
			NotBefore:      contract.TimeNow(),
			Issuer:         "not-test",
		}
		tokenjwt, err := jwt.Sign(payload, tkn.signature)
		payloadRes, err := tkn.Verify(ctx, string(tokenjwt))
		assert.Error(t, err)
		assert.EqualError(t, err, contract.ErrIss.Error())
		assert.Nil(t, payloadRes)
	})

	t.Run("success", func(t *testing.T) {

		tkn.buildSignature(ctx)

		payload := &contract.Payload{
			ExpirationTime: contract.TimeNow().AddDates(0, 0, 5),
			IssuedAt:       contract.TimeNow(),
			NotBefore:      contract.TimeNow(),
			Issuer:         "test",
		}
		tokenjwt, err := jwt.Sign(payload, tkn.signature)
		payloadRes, err := tkn.Verify(ctx, string(tokenjwt))
		assert.NoError(t, err)
		assert.NotNil(t, payloadRes)
	})

}
