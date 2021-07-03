package models

import "time"

// Auth define column table auth
type Auth struct {
	ID         string    `column:"id"`
	Username   string    `column:"username"`
	Passphrase string    `column:"passphrase"`
	CreatedAt  time.Time `column:"created_at"`
	UpdatedAt  time.Time `column:"updated_at"`

	// additionals fields
	columns   []string
	fieldsptr []interface{}
}

// Columns is a method to get columns name
func (auth *Auth) Columns() []string {
	auth.columns = columns(auth)
	return auth.columns
}

// Fields is a method to get fields pointer
func (auth *Auth) Fields() []interface{} {
	auth.fieldsptr = fields(auth)
	return auth.fieldsptr
}
