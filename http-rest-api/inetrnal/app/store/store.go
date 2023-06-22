package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	config        *Config
	db            *sql.DB
	accRepository *Repository
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

func (s *Store) Account() *Repository {
	if s.accRepository != nil {
		return s.accRepository
	}

	s.accRepository = &Repository{
		store: s,
	}

	return s.accRepository
}
