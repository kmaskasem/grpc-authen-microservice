package utils

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return emailRegex.MatchString(email)
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	var hasUpper, hasLower, hasNumber bool
	for _, c := range password {
		switch {
		case '0' <= c && c <= '9':
			hasNumber = true
		case 'a' <= c && c <= 'z':
			hasLower = true
		case 'A' <= c && c <= 'Z':
			hasUpper = true
		}
	}
	if !hasUpper || !hasLower || !hasNumber {
		return errors.New("password must contain upper-case, lower-case letters and numbers")
	}
	return nil
}
