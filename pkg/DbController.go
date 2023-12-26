package pkg

import "database/sql"

type DbController interface {
	Insert(db *sql.DB) error
	Exists(db *sql.DB) error
	GetById(db *sql.DB, id int) error
	Update(db *sql.DB) error
}

func DbInsert(controller DbController, db *sql.DB) error {
	return controller.Insert(db)
}

func DbExists(controller DbController, db *sql.DB) error {
	return controller.Exists(db)
}

func DbGetById(controller DbController, db *sql.DB, id int) error {
	return controller.GetById(db, id)
}

func DbUpdate(controller DbController, db *sql.DB, id int) error {
	return controller.Update(db)
}

type DbRepository interface {
	GetAll(db *sql.DB) error
}

func DbGetAll(repository DbRepository, db *sql.DB) error {
	return repository.GetAll(db)
}