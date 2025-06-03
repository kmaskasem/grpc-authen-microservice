package utils

import (
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

// ตรวจสอบ format อีเมลอย่างง่าย
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// ตรวจสอบความแข็งแรงของพาสเวิร์ด (ตัวอย่าง: ต้องมี 8 ตัวขึ้นไป, มีตัวพิมพ์ใหญ่, พิมพ์เล็ก, ตัวเลข)
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
