package credentials

import "golang.org/x/crypto/bcrypt"

type Password string

var salt = "_chatapp"

func (p *Password) Salt() Password {
	*p += Password(salt)
	return *p
}

func (p *Password) Hash() (Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(*p), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	*p = Password(hash)
	return *p, nil
}

func (p Password) Compare(password Password) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(password)); err != nil {
		return false
	}
	return true
}
