package server

import (
	"os"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := New()
		go OSInterupt(t)
		app.Start()

	})

	t.Run("success-with-tls", func(t *testing.T) {
		app := New()
		app.env.TLS = true
		go OSInterupt(t)
		app.Start()
		app.env.TLS = false

	})

	t.Run("error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		app := New()
		app.env.ServerHost = "9009009090"
		app.Start()

	})
}

func OSInterupt(t *testing.T) {
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	time.Sleep(2 * time.Second)
	proc.Signal(os.Interrupt)
}
