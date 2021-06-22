package database

import (
	"fmt"

	"github.com/rasatmaja/zephyr-one/internal/config"
	"github.com/rasatmaja/zephyr-one/internal/database/postgresql"
	"github.com/rasatmaja/zephyr-one/internal/database/repository"
)

// Factory is a function to build database connection pool
// bassed on database type
func Factory() repository.IRepository {
	env := config.LoadENV()
	switch env.DatabaseType {
	case "POSTGRESQL":
		return postgresql.New()
	default:
		fmt.Println("[ DTBS ] Database Type unrecognize, using default database: POSTGRESQL")
		return postgresql.New()
	}
}
