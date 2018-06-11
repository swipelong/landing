package main

import (
	"errors"
	"regexp"
)

func ValidateEmail(email_address string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(email_address) {
		return nil
	}
	return errors.New("INVALID_EMAIL_ADDRESS")
}
