package model

type Account struct {
	id_account uint64
	id_user    uint64
	amount     int
	type_acc   string `sql:"type:type_account"`
}
