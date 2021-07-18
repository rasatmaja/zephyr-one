package postgresql

import (
	"context"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
)

// CreateContact is a repo to insert contact data
func (qry *Queries) CreateContact(ctx context.Context, contact *models.Contact) error {

	_, err := qry.DB.ExecContext(ctx, "INSERT INTO contact(auth_id, contact_type_id, contact) VALUES($1, $2, $3)", contact.AuthID, contact.ContactTypeID, contact.Contact)

	if err != nil {
		return err
	}

	return nil
}
