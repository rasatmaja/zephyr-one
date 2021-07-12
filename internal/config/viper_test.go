package config

import "testing"

func TestViper(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		type Squash struct {
			ServerHost string `mapstructure:"SERVER_HOST"`
		}

		type TestENV struct {
			ServerHost string `mapstructure:"SERVER_HOST"`
			ServerPort int

			Sqsh  Squash `mapstructure:",squash"`
			Sqsh2 Squash `mapstructure:"SQUASH"`
		}

		env := &TestENV{
			Sqsh: Squash{},
		}

		cfg := GetViper()
		cfg.BindEnvs(env)

	})

	t.Run("fill-default", func(t *testing.T) {

		type TestENV struct {
			ServerHost  string   `mapstructure:"SERVER_HOST" default:"localhost"`
			ServerPort  int      `mapstructure:"SERVER_PORT" default:"3090"`
			ServerTLS   bool     `mapstructure:"SERVER_TLS" default:"true"`
			ServerNoDFL bool     `mapstructure:"SERVER_DFL"`
			ServerNoTYP struct{} `mapstructure:"SERVER_TYP" default:""`

			// Test parse error
			BoolErr bool `mapstructure:"SERVER_DFL" default:"123"`
			IntErr  int  `mapstructure:"SERVER_DFL" default:"abc"`
		}

		env := &TestENV{}

		cfg := GetViper()
		cfg.FillDefault(env)
		t.Log(env)

	})
}
