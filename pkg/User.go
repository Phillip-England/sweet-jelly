package pkg

import (
	"database/sql"
	"fmt"
)

type User struct {
	id int
	firstName string
	lastName string
	email string
	password string
	photo []byte
    sessionToken string
}

func (u *User) Insert(db *sql.DB) error {
	insertUserSQL := `
		INSERT INTO "user" (first_name, last_name, email, password, session_token)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	// Generate a random session token using the function from the pkg package
	sessionToken := GenerateRandomString(64)

	// Hash the user's password before insertion
	hashedPassword, err := HashPassword(u.password)
	if err != nil {
		return err
	}

	fmt.Println(u.password)
	fmt.Println(hashedPassword)

	err = db.QueryRow(insertUserSQL, u.firstName, u.lastName, u.email, hashedPassword, sessionToken).Scan(&u.id)
	if err != nil {
		return err
	}

	// Assign the generated session token to the user struct
	u.sessionToken = sessionToken

	return nil
}

func (u *User) Exists(db *sql.DB) error {
    query := `SELECT EXISTS(SELECT 1 FROM "user" WHERE email = $1)`
    var exists bool
    err := db.QueryRow(query, u.email).Scan(&exists)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("User with email '%s' already exists", u.email)
    }
    return nil
}

func (u *User) GetById(db *sql.DB, id int) error {
	query := `
		SELECT id, first_name, last_name, email, password, photo, session_token
		FROM "user"
		WHERE id = $1
	`

	err := db.QueryRow(query, id).Scan(&u.id, &u.firstName, &u.lastName, &u.email, &u.password, &u.photo, &u.sessionToken)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetByEmail(db *sql.DB, email string) error {
	query := `
		SELECT id, first_name, last_name, email, password, photo, session_token
		FROM "user"
		WHERE email = $1
	`

	err := db.QueryRow(query, email).Scan(&u.id, &u.firstName, &u.lastName, &u.email, &u.password, &u.photo, &u.sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("User with email '%s' not found", email)
		}
		return err
	}

	return nil
}



func (u *User) Update(db *sql.DB) error {
    updateSQL := `
        UPDATE "user"
        SET first_name = $1, last_name = $2, email = $3, photo = $4
        WHERE id = $5
    `
    _, err := db.Exec(updateSQL, u.firstName, u.lastName, u.email, u.photo, u.id)
    if err != nil {
        return err
    }
    return nil
}

//=================================================================
// USER REPOSITORY
//=================================================================

type UserRepository struct {
	users []User
}

func (ur *UserRepository) GetAll(db *sql.DB) error {
    query := `SELECT id, first_name, last_name, email, password, photo FROM "user"`
    
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.id, &u.firstName, &u.lastName, &u.email, &u.password, &u.photo); err != nil {
            return err
        }
        users = append(users, u)
    }

    ur.users = users
    return nil
}