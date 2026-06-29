package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the user password
func HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
  return hashedPassword
}

// CheckPassword checks if the provided password matches
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}