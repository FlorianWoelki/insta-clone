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
	db     *gorm.DB
}

// NewDatabase returns a new database connection with the given logger
func NewDatabase(logger *log.Logger) *Database {
	return &Database{logger, nil}
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

	d.db = db
	d.logger.Println("Successfully connected to postgres database")
	return database
}

// CreateTables creates all tables that are constructed in the types
func (d *Database) CreateTables() {
	if d.db != nil {
		d.db.AutoMigrate(&User{})
	}
}
