package token

import (
	"github.com/rasatmaja/zephyr-one/internal/token/basic"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
)

// Token is a JWT struct
type Token struct{}

// Factory is a function to build JWT
func Factory() contract.IToken {
	// TODO: value of TokenType should be replace with env value
	TokenType := "BASIC"

	// build token config
	token := &contract.Token{
		// TODO: value should be replace with env
		Issuer:  "account.rasio.dev",
		SignKey: "secret",
		SignAlg: "HS265",
	}

	switch TokenType {
	case "BASIC":
		return basic.New(token)
	default:
		return basic.New(token)
	}
}
