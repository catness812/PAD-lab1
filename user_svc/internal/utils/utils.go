package utils

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func CleanUsername(username string) string {
	cleanUsername := html.EscapeString(strings.TrimSpace(username))

	return cleanUsername
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidatePassword(hashedPass, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
}
