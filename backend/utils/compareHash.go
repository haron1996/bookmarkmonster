package utils

import "golang.org/x/crypto/bcrypt"

func CompareHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
