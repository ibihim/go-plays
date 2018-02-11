package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Create makes a table called example
// and populates it
func Create(db *sql.DB) error {
	// create the database
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20))"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name) values ("Aaron")`); err != nil {
		return err
	}

	return nil
}
