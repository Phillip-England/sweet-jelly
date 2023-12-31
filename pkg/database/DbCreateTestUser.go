package database

import (
	"cfasuite/pkg/model/usermod"
	"database/sql"
	"fmt"
)

func DbCreateTestUser(db *sql.DB) error {
	u := usermod.NewUserModel(db)
	u.FirstName = "bob"
	u.LastName = "jonson"
	u.Email = "test@gmail.com"
	u.Password = "aspoaspo"
	controller := DbController(u)
	if err := DbExists(controller); err != nil {
		return err
	}
	if err := DbInsert(controller); err != nil {
		return err
	}
	fmt.Println("Test user created successfully")
	return nil
}