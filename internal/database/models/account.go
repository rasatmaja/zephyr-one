package models

import "time"

// AccountInfo define column table account_info
type AccountInfo struct {
	Models
	ID        string    `column:"id"`
	Photo     string    `column:"photo_profile"`
	Name      string    `column:"name"`
	CreatedAt time.Time `column:"created_at"`
	UpdatedAt time.Time `column:"updated_at"`
}
