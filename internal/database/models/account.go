package models

import "time"

// AccountInfo define column table account_info
type AccountInfo struct {
	ID        string
	Photo     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
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
