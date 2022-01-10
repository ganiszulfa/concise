package validator

import "net/mail"

func IsNotValidEmail(email string) bool {
	return !IsValidEmail(email)
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
