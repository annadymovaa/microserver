package store

import "github.com/annadymova/microserver/http-rest-api/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) CreateUser(u *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
