package internal

import (
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database for creating connection and handling transaction to the database
type Database struct {
	logger *log.Logger
}

// NewDatabase returns a new database connection with the given logger
func NewDatabase(logger *log.Logger) *Database {
	return &Database{logger}
}

// TODO: refactor to env variables
var dsn string = "user=postgres password=postgres dbname=insta-clone port=5432 sslmode=disable"

// CreateConnection creates a connection to the postgres database
// It is not closing the connection to the database
func (d *Database) CreateConnection() *sql.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		d.logger.Printf("Error connecting to database: %s\n", err)
		os.Exit(1)
	}

	database, err := db.DB()
	if err != nil {
		d.logger.Printf("Error connecting to database: %s\n", err)
		os.Exit(1)
	}

	d.logger.Println("Successfully connected to postgres database")
	return database
}
