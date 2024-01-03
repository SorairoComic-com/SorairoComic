package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	passwordByte := []byte(password)
	hashPasswordByte, _ := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)

	return string(hashPasswordByte)
}

func ComparePassword(password string, database string) bool {
	passwordByte := []byte(password)
	databaseByte := []byte(database)
	err := bcrypt.CompareHashAndPassword(databaseByte, passwordByte)

	if err != nil {
		return false
	}

	return true
}
