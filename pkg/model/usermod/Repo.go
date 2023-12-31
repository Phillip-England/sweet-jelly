package usermod

import "database/sql"

type Repo struct {
	DB *sql.DB
	Users []Model
}

func NewUserRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

func (ur *Repo) GetAll() error {
	query := `SELECT id, first_name, last_name, email, password, photo FROM "user"`

	rows, err := ur.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []Model
	for rows.Next() {
		var u Model
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Photo); err != nil {
			return err
		}
		users = append(users, u)
	}

	ur.Users = users
	return nil
}

func (ur *Repo) GetAllByLocationNumber(locationNumber int) error {
	query := `
		SELECT id, first_name, last_name, email, password, photo
		FROM "user"
		WHERE location_number = $1
	`

	rows, err := ur.DB.Query(query, locationNumber)
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []Model
	for rows.Next() {
		var u Model
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Photo); err != nil {
			return err
		}
		users = append(users, u)
	}

	ur.Users = users
	return nil
}
