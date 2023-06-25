package store

import (
	"errors"

	"github.com/annadymovaa/microserver/inetrnal/app/model"
)

var (
	from *uint64
	to   *uint64
)

type Repository struct {
	store *Store
}

func (r *Repository) CreateUser(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	if err := r.store.db.QueryRow(
		"insert into users (username, surname, email, pass) values ($1, $2, $3, $4) returning id_user",
		u.Name, u.Surname, u.Email, u.Pass).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *Repository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"select id from users where email = $1",
		u.Email).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *Repository) CreateBalance(a *model.Account) (*model.Account, error) {

	if err := r.store.db.QueryRow(
		"INSERT INTO accounts(amount) VALUES ($1) RETURNING id_account;", a.Amount).Scan(&a.Id_account); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *Repository) AddTransaction(t *model.Transaction) (*model.Transaction, error) {

	if err := r.store.db.QueryRow(
		"INSERT INTO transactions (price, ta_date, from_acc, to_acc) VALUES ($1, $2, $3, $4) RETURNING id_transaction;",
		t.Price, t.Ta_time, t.From_acc, t.To_acc).Scan(&t.Id_transaction); err != nil {
		return nil, err
	}

	return t, nil
}

func (r *Repository) Debit(t *model.Transaction, a *model.Account) (*model.Account, error) {
	err := r.store.db.QueryRow(
		"select amount from accounts where id_user = $1", t.From_acc).Scan(&a.Id_user, &a.Amount)
	if err != nil {
		return nil, err
	}

	if t.Price > a.Amount {
		return nil, errors.New("Недостаточно средств")
	}

	_, err = r.store.db.Exec("update accounts set amount = amount - $1 where id_user = $2", t.Price, t.From_acc)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (r *Repository) Credit(t *model.Transaction, a *model.Account) (*model.Account, error) {
	err := r.store.db.QueryRow(
		"select amount from accounts where id_user = $1", t.To_acc).Scan(&a.Id_user, &a.Amount)
	if err != nil {
		return nil, err
	}

	_, err = r.store.db.Exec("update accounts set amount = amount + $1 where id_user = $2", t.Price, t.To_acc)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (r *Repository) ChangeBalance(t *model.Transaction) (*model.Account, error) {
	a := &model.Account{}
	from = t.From_acc
	to = t.To_acc
	if from == nil {
		r.Credit(t, a)
	} else if to == nil {
		r.Debit(t, a)
	} else {
		r.Credit(t, a)
		r.Debit(t, a)

	}

	return a, nil
}

// func (r *Repository) ReserveMoney(id_user int, price int) (*model.Account, error) {

// 	a := &model.Account{}

// 	userid := r.store.db.QueryRow("SELECT amount, id_account FROM accounts where id_user=$1;", id_user).Scan()

// 	if err := userid; err != nil {
// 		return nil, err
// 	}

// 	summa := r.store.db.QueryRow("SELECT price FROM transactions WHERE id_service=$1;", id_service).Scan()

// 	if err := summa; err != nil {
// 		return nil, err
// 	}

// 	if err := r.store.db.QueryRow("UPDATE accounts SET amount=amount-price WHERE id_user=$1 and type_acc = 'balance';", id_user).Scan(); err != nil {
// 		return nil, err
// 	}

// 	return a, nil
// }
