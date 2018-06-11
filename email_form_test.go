package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateEmailSubmission(t *testing.T) {
	fmt.Println(os.Getwd())
	ef := new(EmailForm)
	var tests = []struct {
		email_address  string
		ip_address     string
		msg_want       string
		http_code_want int
	}{
		{"alex@gmail.com", "123123123", "", 200},
		{"alex@gmail.com", "123123123", "DUPLICATE_EMAIL_ADDRESS", 422},
		{"new@gmail.com", "123123123", "", 200},
		{"invalid", "123123123", "INVALID_EMAIL_ADDRESS", 422},
		{"", "123123123", "INVALID_EMAIL_ADDRESS", 422},
	}
	for _, test := range tests {
		http_code, err := ef.CreateEmailSubmission(
			test.email_address,
			test.ip_address)
		if err != nil {
			if (err.Error() != test.msg_want) || (http_code != test.http_code_want) {
				t.Errorf("CreateEmailSubmission(%q, %q) == %q, %d. Wanted %q, %d",
					test.email_address,
					test.ip_address,
					err.Error(),
					http_code,
					test.msg_want,
					test.http_code_want)
			}
		}
	}
}
