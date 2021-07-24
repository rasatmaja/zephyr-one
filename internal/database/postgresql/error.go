package postgresql

import (
	"database/sql"
	"strings"

	zosql "github.com/rasatmaja/zephyr-one/internal/database/sql"
)

// ParseInsertErr is a function to decide common sql error
// bassed on posible error when sql perform insert query
func ParseInsertErr(err error) error {
	if strings.Contains(err.Error(), "23505") {
		return zosql.ErrDataDuplicate
	}
	return err
}

// ParseReadErr is a function to decide common sql error
// bassed on posible error when sql perform select query
func ParseReadErr(err error) error {
	if err == sql.ErrNoRows {
		return zosql.ErrNotFound
	}
	return err
}
