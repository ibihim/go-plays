package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Exec replaces the Exec from the previous
// recipe
func Exec(db *sql.DB) error {
	// uncaught error on cleanup, but we always
	// want to cleanup
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db); err != nil {
		return err
	}

	return nil
}
