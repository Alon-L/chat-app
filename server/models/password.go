package models

import "golang.org/x/crypto/bcrypt"

type Password string

var salt = "_chatapp_"

func (p *Password) Salt() {
	*p += Password(salt)
}

func (p *Password) Hash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(*p), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	*p = Password(hash)
	return nil
}

func (p Password) Compare(password Password) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(password)); err != nil {
		return false
	}
	return true
}
