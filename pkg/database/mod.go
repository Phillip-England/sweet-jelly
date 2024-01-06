package database

import (
	"cfasuite/pkg/model/locationmod"
	"cfasuite/pkg/model/usermod"
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

func DbCreateTestLocations(db *sql.DB) error {
	// Create Southroads
	l1 := locationmod.NewLocationModel(db)
	l1.Name = "Southroads"
	l1.Number = 3253
	controller1 := DbController(l1)
	if err := DbInsert(controller1); err != nil {
		return err
	}

	// Create Utica
	l2 := locationmod.NewLocationModel(db)
	l2.Name = "Utica"
	l2.Number = 3991
	controller2 := DbController(l2)
	if err := DbInsert(controller2); err != nil {
		return err
	}

	fmt.Println("Test locations created successfully")
	return nil
}


func DbCreateUserTable(db *sql.DB) error {
	createUserTableSQL := `
        CREATE TABLE IF NOT EXISTS "user" (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(50) NOT NULL,
            last_name VARCHAR(50) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
			location_number INT NOT NULL,
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

func DbCreateTestUsers(db *sql.DB) error {
	u := usermod.NewUserModel(db)
	u.FirstName = "bob"
	u.LastName = "jonson"
	u.Email = "test@gmail.com"
	u.LocationNumber = 3253
	u.Password = "aspoaspo"
	controller := DbController(u)
	if err := DbExists(controller); err != nil {
		return err
	}
	if err := DbInsert(controller); err != nil {
		return err
	}
	u2 := usermod.NewUserModel(db)
	u2.FirstName = "phillip"
	u2.LastName = "spenson"
	u2.Email = "phillip@gmail.com"
	u2.Password = "aspoaspo"
	u2.LocationNumber = 3253
	controller2 := DbController(u2)
	if err := DbExists(controller2); err != nil {
		return err
	}
	if err := DbInsert(controller2); err != nil {
		return err
	}
	fmt.Println("Test users created successfully")
	return nil
}