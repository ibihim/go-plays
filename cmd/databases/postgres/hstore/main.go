package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/lib/pq/hstore"
)

type pgConfig struct {
	host           string
	port           int
	user, password string
	dbname         string
}

func main() {
	dbConf := loadDBConf()
	db, err := openDB(dbConf)
	if err != nil {
		log.Fatal(fmt.Errorf("open database: %w", err))
	}
	defer db.Close()

	if err = openCnn(db); err != nil {
		log.Fatal(fmt.Errorf("open connection: %w", err))
	}

	log.Println("Connection to datbase established")

	if err := setupDB(context.Background(), db); err != nil {
		log.Fatal(fmt.Errorf("set up database: %w", err))
	}

	log.Println("Database setup complete")

	if err := selectValues(context.Background(), db); err != nil {
		log.Fatal("select values: %w", err)
	}

	log.Println("Done selecting values")
}

func loadDBConf() *pgConfig {
	var host string
	if host = os.Getenv("DB_HOST"); host == "" {
		host = "db"
	}

	var p string
	if p = os.Getenv("DB_PORT"); p == "" {
		p = "5432"
	}
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatal(fmt.Errorf("load db config: %w"), err)
	}

	var user string
	if user = os.Getenv("DB_USER"); user == "" {
		user = "postgres"
	}

	var password string
	if password = os.Getenv("DB_PASSWORD"); password == "" {
		password = "password"
	}

	var dbname string
	if dbname = os.Getenv("DB_NAME"); dbname == "" {
		dbname = "test"
	}

	return &pgConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

func openDB(conf *pgConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d "+
			"user=%s password=%s "+
			"dbname=%s sslmode=disable",
		conf.host, conf.port,
		conf.user, conf.password,
		conf.dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return db, nil
}

func openCnn(db *sql.DB) error {
	var err error
	retries := constantBackOff(8, 250*time.Millisecond)
	for _, backOff := range retries {
		if err = db.Ping(); err != nil {
			log.Println("db.Ping fail. retry")
			time.Sleep(backOff)
		}
	}

	return err
}

func constantBackOff(n int, amount time.Duration) []time.Duration {
	tries := make([]time.Duration, n)

	for i := range tries {
		tries[i] = amount
	}

	return tries
}

func setupDB(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, "CREATE EXTENSION hstore;"); err != nil {
		return fmt.Errorf("db exec - create extension hstore: %w", err)
	}

	if _, err := db.ExecContext(ctx, `
CREATE TABLE books (
	id		serial	primary key,
	title	VARCHAR (255),
	attr	hstore
);
	`); err != nil {
		return fmt.Errorf("db exec - create table books: %w", err)
	}

	if _, err := db.ExecContext(ctx, `
INSERT INTO books (title, attr)
VALUES (
	'PostgreSQL Tutorial',
	'"paperback" => "243",
	"publisher" => "postgresqltutorial.com",
	"language" => "English",
	"ISBN-13" => "978-1449370000",
	"weight" => "11.2 ounces"'
);
	`); err != nil {
		return fmt.Errorf("db exec - insert into books: %w", err)
	}

	if _, err := db.ExecContext(ctx, `
INSERT INTO books (title, attr)
VALUES (
	'PostgreSQL Cheat Sheet',
	'"paperback" => "5",
	"publisher" => "postgresqltutorial.com",
	"language" => "English",
	"ISBN-13" => "978-1449370001",
	"weight" => "1 ounces",
	"foo" => "bar"'
);
	`); err != nil {
		return fmt.Errorf("db exec - insert into books: %w", err)
	}

	return nil
}

func selectValues(ctx context.Context, db *sql.DB) error {
	rows, err := db.QueryContext(ctx, "SELECT attr FROM books")
	if err != nil {
		return fmt.Errorf("query - SELECT attr FROM books: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		scanHstore(rows)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows err: %w", err)
	}

	return nil
}

func scanHstore(rows *sql.Rows) error {
	var hs hstore.Hstore
	if err := rows.Scan(&hs); err != nil {
		return fmt.Errorf("rows scan : %w", err)
	}

	log.Println("=========================================")

	for _, key := range []string{"ISBN-13", "language", "paperback", "publisher", "weight"} {
		value, ok := hs.Map[key]
		if !ok {
			log.Println(key, "has no value")
			continue
		}

		if !value.Valid {
			log.Println(value, "is not valid")
			continue
		}

		log.Println(key, "\t", value.String)
	}

	return nil
}
