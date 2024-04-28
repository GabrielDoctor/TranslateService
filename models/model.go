package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// DB holds the connection to the Postgres database
var DB *sql.DB

// DictDB holds the connection to the SQLite database
var EnDictDB *sql.DB

var ViDictDB *sql.DB

func ConnectDb() {
	connectPostgres()
	connectEnDict()
	connectViDict()
}

func connectPostgres() {
	dsn := os.Getenv("DB_CONNECT_STRING")
	if dsn == "" {
		log.Fatal("DB_CONNECT_STRING is not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres DB: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping Postgres DB: %v", err)
	}

	fmt.Println("Connected to Postgres DB")
	DB = db
}

func connectEnDict() {
	db, err := sql.Open("sqlite3", "./cnen.db")
	if err != nil {
		log.Fatalf("Failed to connect to SQLite DB: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping SQLite DB: %v", err)
	}

	fmt.Println("Connected to SQLite DB")
	EnDictDB = db
}

func connectViDict() {
	db, err := sql.Open("sqlite3", "./cnvi.db")
	if err != nil {
		log.Fatalf("Failed to connect to SQLite DB: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping SQLite DB: %v", err)
	}

	fmt.Println("Connected to SQLite DB")
	ViDictDB = db
}
