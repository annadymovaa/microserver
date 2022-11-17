package store

import "github.com/annadymovaa/avito-test/inetrnal/app/model"

type Repository struct {
	store *Store
}

func (r *Repository) CreateBalance(a *model.Account) (*model.Account, error) {
	// 	// if err := a.Validate(); err != nil {
	// 	// 	return err
	// 	// }

	// 	// if err := a.BeforeCreate(); err != nil {
	// 	// 	return err
	// 	// }

	if err := r.store.db.QueryRow("INSERT INTO accounts(amount, type_acc) VALUES ($1, 'balance') RETURNING id_account", a.Amount).Scan(&a.Id_account); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *Repository) FindByUser(id_user int) (*model.Account, error) {
	a := &model.Account{}
	if err := r.store.db.QueryRow("select amount from accounts where id_user = $1 and type_acc = 'balance'", id_user).Scan(&a.Amount); err != nil {
		return nil, err
	}
	return a, nil
}

// func (r *Repository) ReserveMoney(id_user int, id_service uint64, id_order uint64, price int) (*model.Account, error) {
// 	a := &model.Account{}
// 	ser := &model.Service{}
// 	o := &model.Order{}
// 	balance := r.FindByUser(id_user)

// 	if err := r.store.db.QueryRow(""); err != nil{
// 		return nil, err
// 	}
// 	return nil, nil
// }
