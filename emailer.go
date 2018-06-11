package main

import (
	"errors"
	"html/template"
	"os"
)

// Tempaltes

var email_templates = template.Must(template.ParseFiles(
	"email_templates/welcome.html"))

type WelcomeEmail struct {
	EmailAddress string
}

// Emailer

type Emailer struct{}

func (e Emailer) Send(email_address string, template string) error {
	if template == "welcome" {
		we := new(WelcomeEmail)
		we.EmailAddress = email_address
		err := email_templates.ExecuteTemplate(os.Stdout, "welcome.html", we)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Unrecognized email template.")
}
