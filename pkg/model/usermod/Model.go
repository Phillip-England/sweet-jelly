package usermod

import (
	"cfasuite/pkg/util"
	"database/sql"
	"fmt"
)

type Model struct {
	DB 			 *sql.DB
	ID           int
	FirstName    string
	LastName     string
	Email        string
	Password     string
	Photo        []byte
	Family       string
	Hobbies      string
	Dreams       string
}

func NewUserModel(db *sql.DB) *Model {
	return &Model{
		DB: db,
	}
} 

func (u *Model) Insert() error {
	insertUserSQL := `
		INSERT INTO "user" (first_name, last_name, email, password, family, hobbies, dreams)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}
	err = u.DB.QueryRow(insertUserSQL, u.FirstName, u.LastName, u.Email, hashedPassword, u.Family, u.Hobbies, u.Dreams).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}


func (u *Model) Exists() error {
	query := `SELECT EXISTS(SELECT 1 FROM "user" WHERE email = $1)`
	var exists bool
	err := u.DB.QueryRow(query, u.Email).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("user with email '%s' already exists", u.Email)
	}
	return nil
}

func (u *Model) GetById() error {
	query := `
		SELECT id, first_name, last_name, email, password, photo, family, hobbies, dreams
		FROM "user"
		WHERE id = $1
	`

	err := u.DB.QueryRow(query, u.ID).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Photo, &u.Family, &u.Hobbies, &u.Dreams)
	if err != nil {
		return err
	}

	return nil
}

func (u *Model) GetByEmail(email string) error {
	query := `
		SELECT id, first_name, last_name, email, password, photo, family, hobbies, dreams
		FROM "user"
		WHERE email = $1
	`

	err := u.DB.QueryRow(query, email).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Photo, &u.Family, &u.Hobbies, &u.Dreams)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user with email '%s' not found", email)
		}
		return err
	}

	return nil
}

func (u *Model) Update() error {
	updateSQL := `
        UPDATE "user"
        SET first_name = $1, last_name = $2, email = $3, photo = $4, family = $5, hobbies = $6, dreams = $7
        WHERE id = $8
    `
	_, err := u.DB.Exec(updateSQL, u.FirstName, u.LastName, u.Email, u.Photo, u.Family, u.Hobbies, u.Dreams, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *Model) Delete() error {
	return nil
}