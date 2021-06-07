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
		a.CleanupAssets()

		err := a.server.Shutdown()
		if err != nil {
			panic(err)
		}
	}()
}

// AssetType is type for asset type
type AssetType int

const (
	// AssetFile ...
	AssetFile AssetType = iota
	// AssetDir ...
	AssetDir
)

// Assets ...
type Assets struct {
	Path string
	Type AssetType
}

var assets []Assets

// RegisterAssets ...
func (a *App) RegisterAssets(asst Assets) {
	assets = append(assets, asst)
}

// CleanupAssets ...
func (a *App) CleanupAssets() {
	if len(assets) != 0 {
		for _, asset := range assets {
			switch asset.Type {
			case AssetFile:
				os.Remove(asset.Path)
			case AssetDir:
				os.RemoveAll(asset.Path)
			default:
				fmt.Println("[ SRVR ] Cant cleanup asset type unkown")
			}
		}
		fmt.Println("[ SRVR ] Assets successfully cleanup")
	} else {
		fmt.Println("[ SRVR ] Assets already empty")
	}
}
