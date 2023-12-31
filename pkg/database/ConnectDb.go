package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectDb() (*sql.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
		return nil, err
	}
	fmt.Println("Connected to the database")
	return db, nil
}