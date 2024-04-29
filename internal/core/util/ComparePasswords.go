package util

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
