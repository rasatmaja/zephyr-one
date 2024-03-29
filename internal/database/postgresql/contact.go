package postgresql

import (
	"context"
	"fmt"
	"strings"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
)

// CreateContact is a repo to insert contact data
func (qry *Queries) CreateContact(ctx context.Context, contact *models.Contact) error {

	_, err := qry.DB.ExecContext(ctx, "INSERT INTO contacts(auth_id, contact_type_id, contact) VALUES($1, $2, $3)", contact.AuthID, contact.ContactTypeID, contact.Contact)

	if err != nil {
		return ParseInsertErr(err)
	}

	return nil
}

// Contacts is a repo to get user contacts
func (qry *Queries) Contacts(ctx context.Context, authID string, types ...string) ([]*models.Contact, error) {

	// build table name, column and var pointer
	ctc := &models.Contact{}
	table := "contacts"
	columns := strings.Join(ctc.Columns(ctc), ",")

	query := fmt.Sprintf("SELECT %s FROM %s WHERE auth_id = $1", columns, table)

	if len(types) != 0 && len(types[0]) != 0 {
		ctc.ParseContactType(types[0])
		query = fmt.Sprintf("%s AND contact_type_id = %s", query, ctc.ContactTypeID)
	}
	rows, err := qry.DB.QueryContext(ctx, query, authID)
	if err != nil {
		return nil, ParseReadErr(err)
	}

	// scan all column and put the value into var
	var contacts []*models.Contact
	defer rows.Close()
	for rows.Next() {
		contact := &models.Contact{}
		fields := contact.Fields(contact)
		err := rows.Scan(fields...)
		if err != nil {
			return nil, ParseReadErr(err)
		}
		contacts = append(contacts, contact)
	}

	if err != nil {
		return nil, ParseReadErr(err)
	}

	return contacts, nil
}

// SetPrimaryContact is a repo to update user primary contact
func (qry *Queries) SetPrimaryContact(ctx context.Context, authID, contact string) error {
	_, err := qry.DB.ExecContext(ctx, "UPDATE contacts SET is_primary = (CASE contact WHEN $1 THEN TRUE ELSE FALSE END) WHERE auth_id = $2", contact, authID)
	if err != nil {
		return ParseInsertErr(err)
	}
	return nil
}

// DeleteContact is repo to delete user contact
func (qry *Queries) DeleteContact(ctx context.Context, authID, contact string) error {
	_, err := qry.DB.ExecContext(ctx, "DELETE FROM contacts WHERE contact = $1 AND auth_id = $2", contact, authID)
	if err != nil {
		return ParseInsertErr(err)
	}
	return nil
}
