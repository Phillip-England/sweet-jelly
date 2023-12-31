package database

import (
	"database/sql"
	"fmt"
	"log"
)

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