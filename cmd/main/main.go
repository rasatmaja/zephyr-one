package main

import (
	"github.com/rasatmaja/zephyr-one/internal/server"
)

func main() {
	app := server.New()
	app.Start()
}
