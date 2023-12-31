package locationmod

import "database/sql"

type Repo struct {
	DB *sql.DB
	Locations []Model
}

func NewLocationRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

func (lr *Repo) GetAll() error {
	query := `SELECT id, name, number FROM location`

	rows, err := lr.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var locations []Model
	for rows.Next() {
		var loc Model
		if err := rows.Scan(&loc.ID, &loc.Name, &loc.Number); err != nil {
			return err
		}
		locations = append(locations, loc)
	}

	lr.Locations = locations
	return nil
}