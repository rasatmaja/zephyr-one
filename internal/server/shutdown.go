package server

import (
	"fmt"
	"os"
	"os/signal"
)

// InitializeShutdownSequence is a function initialize shutdown sequence
func (a *App) InitializeShutdownSequence() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println(" :: OS signal received")
		fmt.Println("[ SRVR ] Gracefully shutting down...")
		fmt.Println("[ SRVR ] Running cleanup tasks...")
		a.utils.Assets.Cleanup()

		a.server.Shutdown()
	}()
}
