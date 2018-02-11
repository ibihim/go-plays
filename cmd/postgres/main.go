package main

import "github.com/ibihim/go-plays/pkg/sql"

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	if err := database.Exec(db); err != nil {
		panic(err)
	}
}
