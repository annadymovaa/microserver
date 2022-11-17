package store

import "github.com/annadymovaa/avito-test/inetrnal/app/account"

type Repository struct {
	store *Store
}

func (r *Repository) CreateBalance(a *account.Account) (*account.Account, error) {
	// 	// if err := a.Validate(); err != nil {
	// 	// 	return err
	// 	// }

	// 	// if err := a.BeforeCreate(); err != nil {
	// 	// 	return err
	// 	// }

	if err := r.store.db.QueryRow("INSERT INTO accounts(amount, type_acc) VALUES ($1, 'balance') RETURNING id_account;", a.Amount).Scan(&a.Id_account); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *Repository) FindByUser(id_user int) (*account.Account, error) {
	a := &account.Account{}
	if err := r.store.db.QueryRow("select amount from accounts where id_user = $1 and type_acc = 'balance';", id_user).Scan(&a.Amount); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *Repository) ReserveMoney(id_user int, id_service uint64, id_order uint64, price int) (*account.Account, error) {

	a := &account.Account{}
	// 	ser := &model.Service{}
	// 	o := &model.Order{}
	//balance := r.store.db.QueryRow("select amount from accounts where id_user = $1 and type_acc = 'balance';", id_user).Scan(&a.Amount)

	userid := r.store.db.QueryRow("SELECT amount, id_account, type_acc FROM accounts where id_user=$1;", id_user).Scan()

	if err := userid; err != nil {
		return nil, err
	}

	summa := r.store.db.QueryRow("SELECT price FROM services WHERE id_service=$1;", id_service).Scan()

	if err := summa; err != nil {
		return nil, err
	}

	//unreserved = balance - summa
	if err := r.store.db.QueryRow("UPDATE accounts SET amount=amount-price WHERE id_user=$1 and type_acc = 'balance';", id_user).Scan(); err != nil {
		return nil, err
	}

	return a, nil
}
