package internal

import (
	"database/sql"
	"fmt"
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

// CreateConnection creates a connection to the postgres database
// It is not closing the connection to the database
func (d *Database) CreateConnection() *sql.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)

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

	err = createTables(db)
	if err != nil {
		d.logger.Printf("Something went wrong while creating tables: %s\n", err)
		os.Exit(1)
	}
	d.logger.Println("Successfully connected to postgres database")
	return database
}

// CreateTables creates all tables that are constructed in the types
func createTables(db *gorm.DB) error {
	return db.AutoMigrate(&Account{})
}
