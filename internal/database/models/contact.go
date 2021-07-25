package models

import (
	"encoding/json"
	"time"
)

// Contact define column table contact
type Contact struct {
	Models
	AuthID        string    `column:"auth_id"`
	ContactTypeID string    `column:"contact_type_id"`
	Contact       string    `column:"contact"`
	IsPrimary     bool      `column:"is_primary"`
	CreatedAt     time.Time `column:"created_at"`
	UpdatedAt     time.Time `column:"updated_at"`
}

// MarshalJSON is a function to marshaling field time
func (c *Contact) MarshalJSON() ([]byte, error) {
	ctc := struct {
		Contact   string `json:"contact"`
		Types     string `json:"type"`
		IsPrimary bool   `json:"is_primary"`
	}{
		Contact:   c.Contact,
		Types:     c.ParseContactTypeID(),
		IsPrimary: c.IsPrimary,
	}
	return json.Marshal(ctc)
}

// ParseContactTypeID is a helper to define contact type bassed on its ID
func (c *Contact) ParseContactTypeID() string {
	switch c.ContactTypeID {
	case "1":
		return "email"
	case "2":
		return "phone"
	default:
		return ""
	}
}

// ParseContactType is a helper to define contact type ID bassed on its slug
func (c *Contact) ParseContactType(types string) {
	switch types {
	case "email":
		c.ContactTypeID = "1"
	case "phone":
		c.ContactTypeID = "2"
	default:
		c.ContactTypeID = "1"
	}
}
