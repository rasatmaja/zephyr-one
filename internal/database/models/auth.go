package models

import "time"

// Auth define column table auth
type Auth struct {
	Models
	ID         string    `column:"id"`
	Username   string    `column:"username"`
	Passphrase string    `column:"passphrase"`
	CreatedAt  time.Time `column:"created_at"`
	UpdatedAt  time.Time `column:"updated_at"`
}
