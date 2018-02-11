package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Create makes a table called example
// and populates it
func Create(db *sql.DB) error {
	// create the database
	createTable := "CREATE TABLE example (name VARCHAR(20), created TIMESTAMP)"
	if _, err := db.Exec(createTable); err != nil {
		return errors.Wrapf(err, "error during db creation: %s\n", createTable)
	}

	insertInto := `INSERT INTO example values ('Aaron', current_timestamp)`
	if _, err := db.Exec(insertInto); err != nil {
		return errors.Wrapf(err, "error during insertion: %s\n", insertInto)
	}

	return nil
}
