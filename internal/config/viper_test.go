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
}
