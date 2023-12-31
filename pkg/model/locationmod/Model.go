package locationmod

import (
	"database/sql"
	"fmt"
)

type Model struct {
	DB *sql.DB
	ID     int
	Name   string
	Number int
}

func NewLocationModel(db *sql.DB) *Model {
	return &Model{
		DB: db,
	}
}

func (l *Model) Insert() error {
	insertLocationSQL := `
		INSERT INTO location (name, number)
		VALUES ($1, $2)
		RETURNING id
	`

	err := l.DB.QueryRow(insertLocationSQL, l.Name, l.Number).Scan(&l.ID)
	if err != nil {
		return err
	}
	return nil
}

func (l *Model) Exists() error {
	checkExistenceSQL := `
		SELECT EXISTS(SELECT 1 FROM location WHERE id = $1)
	`
	var exists bool
	err := l.DB.QueryRow(checkExistenceSQL, l.ID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("location with ID %d does not exist", l.ID)
	}

	return nil
}

func (l *Model) GetById() error {
	getLocationSQL := `
		SELECT name, number
		FROM location
		WHERE id = $1
	`

	err := l.DB.QueryRow(getLocationSQL, l.ID).Scan(&l.Name, &l.Number)
	if err != nil {
		return err
	}

	return nil
}

func (l *Model) Update() error {
	updateLocationSQL := `
		UPDATE location
		SET name = $1, number = $2
		WHERE id = $3
	`

	_, err := l.DB.Exec(updateLocationSQL, l.Name, l.Number, l.ID)
	if err != nil {
		return err
	}

	return nil
}

func (l *Model) Delete() error {
	deleteLocationSQL := `
		DELETE FROM location
		WHERE id = $1
	`

	_, err := l.DB.Exec(deleteLocationSQL, l.ID)
	if err != nil {
		return err
	}

	return nil
}

//=================================================================
// LOCATION REPOSITORY
//=================================================================

