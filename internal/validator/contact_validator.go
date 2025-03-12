package validator

import "regexp"

func IsValidName(name string) bool {
	return len(name) > 1 && len(name) < 50
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func IsValidPhone(phone string) bool {
	re := regexp.MustCompile(`^(?:\+98\d{10}|0\d{10})$`)
	return re.MatchString(phone)
}
