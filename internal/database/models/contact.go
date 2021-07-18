package models

import "time"

// Contact define column table contact
type Contact struct {
	Models
	AuthID        string    `column:"auth_id"`
	ContactTypeID string    `column:"contact_type_id"`
	Contact       string    `column:"contact"`
	CreatedAt     time.Time `column:"created_at"`
	UpdatedAt     time.Time `column:"updated_at"`
}
