package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(password string) string {
	cost := bcrypt.DefaultCost // Or a higher value like bcrypt.MinCost + 2 for increased security

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	return string(hashedPassword)
}

func ComparePassword(rawPassword, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, err // password doesn't match
		}
		return false, err // unexpected error
	}
	return true, nil // password matches
}
