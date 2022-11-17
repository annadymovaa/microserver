package transaction

import "github.com/annadymovaa/avito-test/inetrnal/app/account"

type Transaction struct {
	Id_transaction uint64          `json:"id_transaction"`
	Price          int             `json:"price"`
	Ta_type        string          `json:"ta_type" sql:"type_transaction"`
	Ta_time        string          `json:"ta_type" sql:"timestamp"`
	From_acc       account.Account `json:"from_acc"`
	To_acc         account.Account `json:"to_acc"`
}
