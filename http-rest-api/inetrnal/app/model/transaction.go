package model

import validation "github.com/go-ozzo/ozzo-validation"

type Transaction struct {
	Id_transaction uint64  `json:"id_transaction"`
	Price          float64 `json:"price"`
	Ta_time        string  `json:"ta_time" sql:"timestamp"`
	From_acc       *uint64 `json:"from_acc"`
	To_acc         *uint64 `json:"to_acc"`
}

func (t *Transaction) Validate() error {
	return validation.ValidateStruct(t)
}
