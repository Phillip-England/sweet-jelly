package sessionmod

import (
	"cfasuite/pkg/util"
	"database/sql"
)

type Model struct {
	DB *sql.DB
	ID           int
	Token 		 string
	UserID 		 int
}

func NewSessionModel(db *sql.DB) *Model {
	return &Model{
		DB: db,
	}
}

func (s *Model) Insert() error {
	s.Token = util.GenerateRandomString(64)
	query := `
		INSERT INTO session (token, user_id)
		VALUES ($1, $2)
		RETURNING id;
	`
	err := s.DB.QueryRow(query, s.Token, s.UserID).Scan(&s.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Model) GetByToken() error {
	query := `
	SELECT id, user_id
	FROM session
	WHERE token = $1;
	`
	err := s.DB.QueryRow(query, s.Token).Scan(&s.ID, &s.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Model) Exists() error {
	return nil
}

func (s *Model) GetById() error {
	return nil
}

func (s *Model) Update() error {
	return nil
}

func (s *Model) Delete() error {
	query := `
		DELETE FROM session
		WHERE token = $1;
	`
	_, err := s.DB.Exec(query, s.Token)
	if err != nil {
		return err
	}
	return nil
}
