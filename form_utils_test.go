package main

import "testing"

func TestValidateEmail(t *testing.T) {
	var tests = []struct {
		email_address string
		valid         bool
	}{
		// VALID
		{"email@example.com", true},
		{"firstname.lastname@example.com", true},
		{"email@subdomain.example.com", true},
		{"firstname+lastname@example.com", true},
		{"email@123.123.123.123", true},
		{"email@example.com", true},
		{"1234567890@example.com", true},
		{"email@example-one.com", true},
		{"_______@example.com", true},
		{"email@example.name", true},
		{"email@example.museum", true},
		{"email@example.co.jp", true},
		{"firstname-lastname@example.com", true},
		{".email@example.com", true},
		{"email.@example.com", true},
		{"Abc..123@example.com", true},
		{"email..email@example.com", true},
		{"email@example", true},
		{"email@example.web", true},
		{"email@111.222.333.44444", true},
		{"alex@gmail.com", true},
		// INVALID
		{"plainaddress", false},
		{"#@%^%#$@#$@#.com", false},
		{"@example.com", false},
		{"Joe Smith <email@example.com>", false},
		{"email.example.com", false},
		{"email@example@example.com", false},
		{"あいうえお@example.com", false},
		{"email@example.com (Joe Smith)", false},
		{"email@-example.com", false},
		{"email@example..com", false},
	}
	for _, test := range tests {
		if test.valid {
			err := ValidateEmail(test.email_address)
			if err != nil {
				t.Errorf("Valid email caught as invalid: " + test.email_address)
			}
		} else {
			err := ValidateEmail(test.email_address)
			if err == nil {
				t.Errorf("Invalid email passed as valid: " + test.email_address)
			}
		}
	}
}
