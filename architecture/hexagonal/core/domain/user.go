package domain

import (
	"errors"
	"net/mail"
)

type User struct {
	ID    string
	Name  string
	Email string
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is mandatory")
	}

	if u.Email == "" {
		return errors.New("email is mandatory")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("email invalid")
	}

	return nil
}
