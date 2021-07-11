package token

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	env := config.LoadENV()
	t.Run("Sign", func(t *testing.T) {
		env.TokenType = "BASIC"
		ctx, cancle := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancle()

		tkn := Factory()
		time := contract.TimeNow()

		payload := &contract.Payload{
			JWTID:          uuid.NewString(),
			IssuedAt:       time,
			NotBefore:      time,
			ExpirationTime: time.AddDates(0, 0, 30),
		}

		jwt, err := tkn.Sign(ctx, payload)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(jwt))
	})

	t.Run("Verify", func(t *testing.T) {
		env.TokenType = "UNKNOWN"
		ctx, cancle := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancle()

		tkn := Factory()
		time := contract.TimeNow()

		payload := &contract.Payload{
			JWTID:          uuid.NewString(),
			IssuedAt:       time,
			NotBefore:      time,
			ExpirationTime: time.AddDates(0, 0, 30),
		}

		jwt, err := tkn.Sign(ctx, payload)
		_, err = tkn.Verify(ctx, jwt)

		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(jwt))
	})
}
