package database

import (
	"database/sql"
	"fmt"
	"log"
)

func DbCreateSessionTable(db *sql.DB) error {
	createSessionTableSQL := `
        CREATE TABLE IF NOT EXISTS session (
            id SERIAL PRIMARY KEY,
            token VARCHAR(255) UNIQUE NOT NULL,
            user_id VARCHAR(255) NOT NULL
        );
    `
	_, err := db.Exec(createSessionTableSQL)
	if err != nil {
		log.Fatal("Error creating session table:", err)
		return err
	}
	fmt.Println("Session table created successfully")
	return nil
}