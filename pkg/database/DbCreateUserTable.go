package database

import (
	"database/sql"
	"fmt"
	"log"
)

func DbCreateUserTable(db *sql.DB) error {
	createUserTableSQL := `
        CREATE TABLE IF NOT EXISTS "user" (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(50) NOT NULL,
            last_name VARCHAR(50) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL,
			photo BYTEA,
			family VARCHAR(255),
			hobbies VARCHAR(255),
			dreams VARCHAR(255)
        );
    `
	_, err := db.Exec(createUserTableSQL)
	if err != nil {
		log.Fatal("Error creating user table:", err)
		return err
	}
	fmt.Println("User table created successfully")
	return nil
}