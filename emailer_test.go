package main

import "testing"

func TestSend(t *testing.T) {
	var tests = []struct {
		email_address string
	}{
		{"email@example.com"},
		{"firstname.lastname@example.com"},
		{"email@subdomain.example.com"},
		{"firstname+lastname@example.com"},
		{"email@123.123.123.123"},
		{"email@example.com"},
		{"1234567890@example.com"},
		{"email@example-one.com"},
	}
	for _, test := range tests {
		e := new(Emailer)
		err := e.Send(test.email_address, "welcome")
		if err != nil {
			t.Errorf("Problem sending email to " + test.email_address + ": " + err.Error())
		}
	}
}
