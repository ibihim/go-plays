package database

import (
	"database/sql"
	"fmt"
	"time"
	//	"os"
	_ "github.com/lib/pq"
)

// Example hold the results of our queries
type Example struct {
	Name    string
	Created *time.Time
}

// Setup configures and returns our database
// connection pooled
func Setup() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		"ibihim",      //os.Getenv("PS_USER"),
		"helloGithub", //os.Getenv("PS_PASSWORD"),
		"test",
		"disable",
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
