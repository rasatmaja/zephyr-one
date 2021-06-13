package postgresql

import (
	"database/sql"
	"fmt"
	"net/url"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rasatmaja/zephyr-one/internal/database"
)

// Queries ...
type Queries struct {
	db database.IDatabase
}

// WithTx ..
func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

// New is function to initialize progresql
func New() {

	databaseHost := "localhost"
	databasePort := 5432
	databaseUsername := "root"
	databasePassword := "root"
	databaseName := "test"
	databaseSSLMode := "disable"

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(databaseUsername, databasePassword),
		Host:   fmt.Sprintf("%s:%d", databaseHost, databasePort),
		Path:   databaseName,
	}

	q := dsn.Query()
	q.Add("sslmode", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

}
