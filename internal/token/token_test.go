package token

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	t.Run("Sign", func(t *testing.T) {
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
		t.Log(jwt)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(jwt))
	})
}
