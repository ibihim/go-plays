package main

import (
	"github.com/ibihim/go-plays/pkg/sql"
	"github.com/pkg/errors"
	"fmt"
	"os"
)

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%v+\n", err)
	os.Exit(1)
}

func main() {
	db, err := database.Setup()
	if err != nil {
		exit(errors.Wrap(err, "Database setup error"))
	}

	if err := database.Exec(db); err != nil {
		exit(errors.Wrap(err, "Database exec error"))
	}
}
