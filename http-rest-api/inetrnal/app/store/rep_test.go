package store_test

import (
	"testing"

	"github.com/annadymovaa/microserver/inetrnal/app/model"
	"github.com/annadymovaa/microserver/inetrnal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestAccRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("accounts")

	a, err := s.User().CreateBalance(&model.Account{
		Amount: 1000,
	})
	assert.NoError(t, err)
	assert.NotNil(t, a)
}

func TestAccRepository_Find(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().FindByEmail(email)
	s.User().CreateUser(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
