package models

import (
	"encoding/json"
	"time"
)

// AccountInfo define column table account_info
type AccountInfo struct {
	Models
	ID        string    `column:"id"`
	Photo     string    `column:"photo_profile"`
	Name      string    `column:"name"`
	CreatedAt time.Time `column:"created_at"`
	UpdatedAt time.Time `column:"updated_at"`
}

// MarshalJSON is a function to marshaling field account info
func (c *AccountInfo) MarshalJSON() ([]byte, error) {
	ctc := struct {
		Name             string `json:"name"`
		Photo            string `json:"photo"`
		RegistrationDate string `json:"registration_date"`
	}{
		Name:             c.Name,
		Photo:            c.Photo,
		RegistrationDate: c.CreatedAt.Format("Monday 02, January 2006 15:04 MST"),
	}
	return json.Marshal(ctc)
}
