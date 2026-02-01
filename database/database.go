package database

import (
	"database/sql"
	"log"
	"time"

	//_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}



	// Set connection pool settings (optional tapi recommended)
	db.SetConnMaxLifetime(time.Minute * 3) // Kill connections before the pooler does
  db.SetMaxIdleConns(5)                  // Don't keep too many idle connections
  db.SetMaxOpenConns(20)                 // Limit total connections
  
  	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}


	log.Println("Database connected successfully")
	return db, nil
}
