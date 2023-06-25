package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                uint64 `json:"id_user"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Email             string `json:"email"`
	Pass              string `json:"password"`
	EncryptedPassword string `json:"encrypted_pass"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u, validation.Field(
			&u.Email, validation.Required, is.Email),
		validation.Field(
			&u.Pass, validation.By(requiredIf(
				u.EncryptedPassword == "")), validation.Length(6, 100)))
}

func (u *User) BeforeCreate() error {
	if len(u.Pass) > 0 {
		enc, err := encryptString(u.Pass)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
