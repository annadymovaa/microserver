package transaction

import "github.com/annadymovaa/microserver/inetrnal/app/account"

type Transaction struct {
	Id_transaction uint64          `json:"id_transaction"`
	Price          int             `json:"price"`
	Ta_time        string          `json:"ta_type" sql:"timestamp"`
	From_acc       account.Account `json:"from_acc"`
	To_acc         account.Account `json:"to_acc"`
}
