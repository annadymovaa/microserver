package store

import "github.com/annadymovaa/avito-test/inetrnal/app/model"

type AccRepository struct {
	store *Store
}

func (r *AccRepository) Create(a *model.Account) (*model.Account, error) {
	if err := r.store.db.QueryRow("INSERT INTO accounts(user_id, amount, type_acc) VALUES ($1, '0', 'balance') RETURNING id_account", a.id_user, a.amount, a.type_acc).Scan(&a.id_account); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *AccRepository) FindByUser(id_user uint64) (*model.Account, error) {
	return nil, nil
}
