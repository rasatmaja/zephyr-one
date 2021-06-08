package utils

import (
	"fmt"
	"os"
)

// Assets struct
type Assets struct{}

// AssetType is type for asset type
type AssetType int

const (
	// AssetFile type
	AssetFile AssetType = iota
	// AssetDir type
	AssetDir
)

// Asset struct to define assets info
type Asset struct {
	Path string
	Type AssetType
}

var assets []Asset

//Register is a function to register asset
func (a Assets) Register(asset Asset) {
	assets = append(assets, asset)
}

// Cleanup is a function to remove registered assets bassed on their path
func (a Assets) Cleanup() {
	if len(assets) != 0 {
		for _, asset := range assets {
			switch asset.Type {
			case AssetFile:
				os.Remove(asset.Path)
			case AssetDir:
				os.RemoveAll(asset.Path)
			default:
				fmt.Println("[ UTLS ] Cant cleanup asset type unkown")
			}
		}
		fmt.Println("[ UTLS ] Assets successfully cleanup")
	} else {
		fmt.Println("[ UTLS ] Assets already empty")
	}
}
