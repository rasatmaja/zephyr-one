package postgresql

import (
	"context"
	"fmt"
	"strings"

	"github.com/rasatmaja/zephyr-one/internal/database/models"
)

// CreateAccountInfo is a function to insert data into table account_nfo
func (qry *Queries) CreateAccountInfo(ctx context.Context, id, name string) (*models.AccountInfo, error) {
	// insert data into accoint_inf
	_, err := qry.DB.ExecContext(ctx, "INSERT INTO account_info(id, name) VALUES($1, $2)", id, name)
	if err != nil {
		return nil, ParseInsertErr(err)
	}

	// build model account info
	acc := &models.AccountInfo{
		ID:   id,
		Name: name,
	}
	return acc, nil
}

// Account is a repo to pull all account data by its username
func (qry *Queries) Account(ctx context.Context, id string) (*models.AccountInfo, error) {
	acc := &models.AccountInfo{}

	// build table name, column and var pointer
	table := "account_info"
	columns := strings.Join(acc.Columns(acc), ",")
	fields := acc.Fields(acc)

	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1 LIMIT 1", columns, table)
	row := qry.DB.QueryRowContext(ctx, query, id)

	// scan all column and put the value into var
	err := row.Scan(fields...)
	if err != nil {
		return nil, ParseReadErr(err)
	}

	return acc, nil
}
