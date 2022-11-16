package store

import "github.com/annadymovaa/avito-test/inetrnal/app/model"

type AccRepository struct {
	store *Store
}

func (r *AccRepository) CreateBalance(a *model.Account) (*model.Account, error) {
	// 	// if err := a.Validate(); err != nil {
	// 	// 	return err
	// 	// }

	// 	// if err := a.BeforeCreate(); err != nil {
	// 	// 	return err
	// 	// }

	if err := r.store.db.QueryRow("INSERT INTO accounts(type_acc) VALUES ('balance') RETURNING id_account").Scan(); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *AccRepository) FindByUser(id_user uint64) (*model.Account, error) {
	return nil, nil
}
