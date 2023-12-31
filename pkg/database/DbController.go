package database

type DbController interface {
	Insert() error
	Exists() error
	GetById() error
	Update() error
	Delete() error
}

func DbInsert(controller DbController) error {
	return controller.Insert()
}

func DbExists(controller DbController) error {
	return controller.Exists()
}

func DbGetById(controller DbController) error {
	return controller.GetById()
}

func DbUpdate(controller DbController) error {
	return controller.Update()
}

func DbDelete(controller DbController) error {
	return controller.Delete()
}

type DbRepository interface {
	GetAll() error
}

func DbGetAll(repository DbRepository) error {
	return repository.GetAll()
}