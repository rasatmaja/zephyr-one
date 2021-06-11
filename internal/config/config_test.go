package config

import (
	"os"
	"testing"
)

func TestLoadENV(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("The code did panic")
				t.Fail()
			}
		}()
		LoadENV()
	})
}

func TestBuildENV(t *testing.T) {
	t.Run("error-read-config", func(t *testing.T) {
		defer func() {
			os.Remove("app")
			if r := recover(); r == nil {
				t.Errorf("The code did panic")
				t.Fail()
			}
		}()
		os.WriteFile("app", []byte(""), 0600)
		config := &Config{
			Filename: "app",
			Type:     "unknown",
			Path:     ".",
		}
		config.BuildENV()
	})

	t.Run("error-unmarshall", func(t *testing.T) {
		defer func() {
			os.Remove("app")
			if r := recover(); r == nil {
				t.Errorf("The code did panic")
				t.Fail()
			}
		}()
		os.WriteFile("app", []byte("SERVER_PRODUCTION=localhost"), 0600)
		config := &Config{
			Filename: "app",
			Type:     "env",
			Path:     ".",
		}
		config.BuildENV()
	})
}
