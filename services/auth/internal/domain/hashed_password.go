package domain

import "golang.org/x/crypto/bcrypt"

func CreateHashedPassword(password string) (string, error) {

	hPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", ErrMakingHashedPassword
		// make hashhed password and compare in login
	}
	return string(hPassword), nil
}
