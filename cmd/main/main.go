package main

import (
	"fmt"

	"github.com/rasatmaja/zephyr-one/internal/server"
)

func main() {
	fmt.Println("[ MAIN ] Starting Server ...")
	app := server.New()
	app.Start()
}
