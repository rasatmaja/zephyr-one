package server

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/logger"
	"github.com/rasatmaja/zephyr-one/internal/middleware"
	"github.com/rasatmaja/zephyr-one/internal/utils"
)

func TestNew(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
				t.Fail()
			}
		}()

		New()
	})
}

func TestServer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svr := fiber.New()
		app := &App{
			server:     svr,
			logger:     logger.New(),
			middleware: middleware.New(svr),
			env:        config.LoadENV(),
			utils:      utils.New(),
		}
		go OSInterupt(t)
		app.Start()

	})

	t.Run("success-with-tls", func(t *testing.T) {
		// setup temp directory
		tempDir := t.TempDir()
		utils := &utils.Registry{
			Assets: &utils.Assets{},
			Cert: &utils.Cert{
				CertFilename:    fmt.Sprintf("%s/cert.pem", tempDir),
				PrivKeyFilename: fmt.Sprintf("%s/priv.pem", tempDir),
				Permission:      0600,
			},
		}

		// setup server
		svr := fiber.New()
		app := &App{
			server:     svr,
			logger:     logger.New(),
			middleware: middleware.New(svr),
			utils:      utils,
			env:        config.LoadENV(),
		}
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

		svr := fiber.New()
		app := &App{
			server:     svr,
			logger:     logger.New(),
			middleware: middleware.New(svr),
			env:        config.LoadENV(),
			utils:      utils.New(),
		}
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
