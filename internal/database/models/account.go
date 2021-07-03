package models

import "time"

// AccountInfo define column table account_info
type AccountInfo struct {
	ID        string    `column:"id"`
	Photo     string    `column:"photo_profile"`
	Name      string    `column:"name"`
	CreatedAt time.Time `column:"created_at"`
	UpdatedAt time.Time `column:"updated_at"`
	// additionals fields
	columns   []string
	fieldsptr []interface{}
}

// Columns is a method to get columns name
func (acc *AccountInfo) Columns() []string {
	acc.columns = columns(acc)
	return acc.columns
}

// Fields is a method to get fields pointer
func (acc *AccountInfo) Fields() []interface{} {
	acc.fieldsptr = fields(acc)
	return acc.fieldsptr
}
