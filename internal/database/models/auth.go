package models

import "time"

// Auth define column table auth
type Auth struct {
	ID         string
	Username   string
	Passphrase string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
