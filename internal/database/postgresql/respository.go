package postgresql

import (
	"context"
	"database/sql"

	"github.com/rasatmaja/zephyr-one/internal/database"
)

// Queries ...
type Queries struct {
	db database.ISQL
}

// New is function to initialize progresql
func New() *Queries {

	db, err := OpenConn()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	queries := &Queries{db: db}
	return queries

}

// BeginTX is a function to start transaction
func (q *Queries) BeginTX(ctx context.Context) (*Queries, *sql.Tx, error) {
	db, err := OpenConn()
	if err != nil {
		panic(err)
	}

	// begin transactions
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	qry := &Queries{db: tx}

	return qry, tx, nil
}
