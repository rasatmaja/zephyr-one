package config

import (
	"os"
	"testing"
)

func TestLoadENV(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer func() {
			os.Remove(".env")
			if r := recover(); r != nil {
				t.Errorf("The code did panic")
				t.Fail()
			}
		}()
		os.WriteFile(".env", []byte("SERVER_PRODUCTION=localhost"), 0600)
		LoadENV()
	})
}
