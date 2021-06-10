package utils

import (
	"os"
	"testing"
)

func TestAssets(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		asst := Assets{}

		// create test file
		if err := os.WriteFile("test.dump", []byte("test"), 0600); err != nil {
			t.Log(err)
			t.Fail()
		}
		asst.Register(Asset{Path: "test.dump", Type: AssetFile})

		// create directory
		if err := os.Mkdir("test_dir", 0600); err != nil {
			t.Log(err)
			t.Fail()
		}
		asst.Register(Asset{Path: "test_dir", Type: AssetDir})

		const unknownType AssetType = iota + 6
		asst.Register(Asset{Path: "unknown", Type: unknownType})

		asst.Cleanup()
		assets = make([]Asset, 0)
	})

	t.Run("empty-assets/success", func(t *testing.T) {
		assts := Assets{}
		assts.Cleanup()
	})
}
