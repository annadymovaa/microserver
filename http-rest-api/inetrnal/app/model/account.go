package model

import validation "github.com/go-ozzo/ozzo-validation"

type Account struct {
	Id_account uint64  `json:"id_account"`
	Id_user    User    `json:"id_user"`
	Amount     float64 `json:"amount"`
}

func (a *Account) Validate() error {
	return validation.ValidateStruct(a)
}
