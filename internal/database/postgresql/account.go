package postgresql

import (
	"context"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
)

// CreateAccountInfo is a function to insert data into table account_nfo
func (qry *Queries) CreateAccountInfo(ctx context.Context, id, name string) (*models.AccountInfo, error) {
	// insert data into accoint_inf
	_, err := qry.DB.ExecContext(ctx, "INSERT INTO account_info(id, name) VALUES($1, $2)", id, name)
	if err != nil {
		return nil, err
	}

	// build model account info
	acc := &models.AccountInfo{
		ID:   id,
		Name: name,
	}
	return acc, nil
}
