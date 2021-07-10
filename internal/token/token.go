package token

import (
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/token/basic"
	"github.com/rasatmaja/zephyr-one/internal/token/contract"
)

// Token is a JWT struct
type Token struct{}

// Factory is a function to build JWT
func Factory() contract.IToken {

	env := config.LoadENV()

	// build token config
	token := &contract.Token{
		// TODO: value should be replace with env
		Issuer:  env.TokenIssuer,
		SignKey: env.TokenSignKey,
		SignAlg: contract.SignAlgorithm(env.TokenSignAlg),
	}

	switch env.TokenType {
	case "BASIC":
		return basic.New(token)
	default:
		return basic.New(token)
	}
}
