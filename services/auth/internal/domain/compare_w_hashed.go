package domain

import "golang.org/x/crypto/bcrypt"

func CompareWithHashedPassword(tablePassword, usersPassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(tablePassword), []byte(usersPassword))
}
