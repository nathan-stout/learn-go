package database

import (
	"database/sql"
	"fmt"
	"log"
	"server/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB holds the database connection
type DB struct {
	*sql.DB
}

// NewConnection creates a new database connection
func NewConnection(cfg *config.Config) (*DB, error) {
	dsn := cfg.GetDSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Configure connection pool for production use
	db.SetMaxOpenConns(25)   // Maximum number of open connections
	db.SetMaxIdleConns(10)   // Maximum number of idle connections
	db.SetConnMaxLifetime(0) // Maximum amount of time a connection may be reused

	log.Println("Successfully connected to PostgreSQL database")

	return &DB{db}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}

// Migrate runs database migrations
func (db *DB) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS albums (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		artist VARCHAR(255) NOT NULL,
		price DECIMAL(10,2) NOT NULL CHECK (price > 0),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(title, artist)
	);

	-- Create index for better query performance
	CREATE INDEX IF NOT EXISTS idx_albums_title ON albums(title);
	CREATE INDEX IF NOT EXISTS idx_albums_artist ON albums(artist);
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// SeedData inserts initial data if the table is empty
func (db *DB) SeedData() error {
	// Check if data already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM albums").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check existing data: %v", err)
	}

	if count > 0 {
		log.Println("Data already exists, skipping seed")
		return nil
	}

	// Insert seed data
	query := `
	INSERT INTO albums (title, artist, price) VALUES
		('Blue Train', 'John Coltrane', 56.99),
		('Jeru', 'Gerry Mulligan', 17.99),
		('Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99)
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("failed to seed data: %v", err)
	}

	log.Println("Database seeded successfully")
	return nil
}
