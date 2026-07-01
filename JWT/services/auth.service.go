package services

import (
	"golang.org/x/crypto/bcrypt"
)

// hashes the user password
func HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
  return hashedPassword
}

// checks if the provided password matches
func VerifyPasswod(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}