package models

import "time"

// AccountInfo define column table account_info
type AccountInfo struct {
	ID        string
	Photo     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
