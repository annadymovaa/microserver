package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	//"github.com/jmoiron/sqlx"
)

type Store struct {
	config        *Config
	db            *sql.DB
	accRepository *AccRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {

	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	fmt.Println("Connected!")

	return nil
}

func (s *Store) Close() {
	s.db.Close()

}

func (s *Store) Account() *AccRepository {
	if s.accRepository != nil {
		return s.accRepository
	}

	s.accRepository = &AccRepository{
		store: s,
	}

	return s.accRepository
}
