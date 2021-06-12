package utils

import (
	"fmt"
	"sync"
)

var instance *Registry
var singleton sync.Once

// Registry struct
type Registry struct {
	Assets *Assets
	Cert   ICertificates
}

// New is a function to init utils struct
func New() *Registry {
	singleton.Do(func() {
		fmt.Println("[ UTLS ] Initialize Utilities")
		instance = &Registry{
			Assets: &Assets{},
			Cert: &Cert{
				CertFilename:    "cert.pem",
				PrivKeyFilename: "priv.pem",
				Permission:      0600,
			},
		}
	})
	return instance
}
