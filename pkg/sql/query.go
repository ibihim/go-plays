package database

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	_ "github.com/lib/pq"
)

// Query grabs a new connection
// creates tables, and later drops them
// and issues some queries
func Query(db *sql.DB) error {
	name := "Aaron"
	rows, err := db.Query("SELECT name FROM example where name = $1", name)
	if err != nil {
		return errors.Wrapf(err, "error during select exec for name %s\n", name)
	}

	defer rows.Close()

	for rows.Next() {
		var e Example
		fmt.Printf("rows: %v+", rows)
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return errors.Wrapf(err, "error during reading row: %v\n", e)
		}

		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}

	return rows.Err()
}
