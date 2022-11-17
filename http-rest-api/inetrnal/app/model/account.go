package model

//import "time"

// func Timestamp() {
// 	presentTime := time.Now().Format("01-02-2006 15:04:05")

// }

type Account struct {
	Id_account uint64 `json: "id_account"`
	Id_user    uint64 `json: "id_user`
	Amount     int    `json: "amount"`
	Type_acc   string `json: "type_acc" sql:"type:type_account"`
}

type Service struct {
	Id_service uint64
	Price      int
}

type Order struct {
	Id_order uint64
}

type Transaction struct {
	Id_transaction uint64
	Price          int
	Ta_type        string `sql:"type_transaction"`
	Ta_time        string `sql:"timestamp"`

	From_acc uint64
	To_acc   uint64
}

// func (a *Account) Validate() error {
// 	return validation.ValidateStruct(
// 		a,
// 		validation.Field(&a.id_user, validation.Required, is.Email),
// 		validation.Field(&a.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
// 	)
// }

// // BeforeCreate ...
// func (a *Account) BeforeCreate() error {
// 	if len(a.id_user) > 0 {
// 		enc, err := encryptString(a.id_user)
// 		if err != nil {
// 			return err
// 		}

// 		a.id_user = enc
// 	}

// 	return nil
// }
