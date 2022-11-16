package model

type Account struct {
	id_account uint64
	id_user    uint64
	amount     int
	type_acc   string `sql:"type:type_account"`
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
