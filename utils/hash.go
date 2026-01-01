package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Implementation of password hashing
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	// Implementation of password hash comparison
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
