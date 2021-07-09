package token

import (
	"github.com/rasatmaja/zephyr-one/internal/token/basic"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
)

// Token is a JWT struct
type Token struct{}

// Factory is a function to build JWT
func Factory() contract.IToken {
	// value of TokenType should be replace with env value
	TokenType := "BASIC"

	switch TokenType {
	case "BASIC":
		return basic.New()
	default:
		return basic.New()
	}
}
