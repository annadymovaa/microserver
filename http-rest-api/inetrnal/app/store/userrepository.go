package store

import "github.com/annadymovaa/avito-test/inetrnal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) Find(id_account int) (*model.User, error) {
	return nil, nil
}
