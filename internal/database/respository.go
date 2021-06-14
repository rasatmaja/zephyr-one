package database

import (
	"context"
	"database/sql"
	"fmt"
)

// New is function to initialize progresql
func New() *Queries {

	db, err := OpenConn()
	if err != nil {
		panic(fmt.Errorf("Unable to open connection, got: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("Unable to ping connection, got: %v", err))
	}

	queries := &Queries{db: db}
	return queries

}

// BeginTX is a function to start transaction
func (q *Queries) BeginTX(ctx context.Context) (*Queries, *sql.Tx, error) {
	db, err := OpenConn()
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to open connection, got: %v", err)
	}

	// begin transactions
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to open transaction, got: %v", err)
	}
	qry := &Queries{db: tx}

	return qry, tx, nil
}
