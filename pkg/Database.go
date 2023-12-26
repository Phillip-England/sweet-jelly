package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
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

func DbDeleteAllTables(db *sql.DB) error {
	query := `
        SELECT table_name
        FROM information_schema.tables
        WHERE table_schema = 'public';
    `

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error querying table names:", err)
		return err
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal("Error scanning table name:", err)
			return err
		}
		tableNames = append(tableNames, tableName)
	}

	for _, tableName := range tableNames {
		dropTableSQL := fmt.Sprintf("DROP TABLE IF EXISTS \"%s\" CASCADE;", tableName)
		_, err := db.Exec(dropTableSQL)
		if err != nil {
			log.Fatal("Error deleting table:", err)
			return err
		}
	}

	fmt.Println("All tables deleted successfully")
	return nil
}

func DbCreateUserTable(db *sql.DB) error {
	createUserTableSQL := `
        CREATE TABLE IF NOT EXISTS "user" (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(50) NOT NULL,
            last_name VARCHAR(50) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL,
			photo BYTEA,
			session_token VARCHAR(64)
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
