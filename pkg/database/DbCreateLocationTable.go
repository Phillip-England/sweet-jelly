package database

import (
	"database/sql"
	"fmt"
	"log"
)

func DbCreateLocationTable(db *sql.DB) error {
	createLocationTableSQL := `
        CREATE TABLE IF NOT EXISTS location (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            number INT NOT NULL
        );
    `
	_, err := db.Exec(createLocationTableSQL)
	if err != nil {
		log.Fatal("Error creating location table:", err)
		return err
	}
	fmt.Println("Location table created successfully")
	return nil
}