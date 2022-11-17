package store_test

import (
	"testing"

	"github.com/annadymovaa/avito-test/inetrnal/app/account"
	"github.com/annadymovaa/avito-test/inetrnal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestAccRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("accounts")

	a, err := s.Account().CreateBalance(&account.Account{
		Amount: 1000,
	})
	assert.NoError(t, err)
	assert.NotNil(t, a)
}

func TestAccRepository_Find(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("accounts")

	id_user := 15
	_, err := s.Account().FindByUser(id_user)
	assert.Error(t, err)

}
