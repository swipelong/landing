package main

import (
	"bytes"
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"os"
)

// Tempaltes

var email_templates = template.Must(template.ParseFiles(
	"email_templates/welcome.html"))

type WelcomeEmail struct {
	EmailAddress string
	Content      string
}

// Emailer

type Emailer struct{}

func (e Emailer) Send(email_address string, template string) error {
	if template == "welcome" {
		err := e.sendWelcomeEmail(email_address)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Unrecognized email template.")
}

func (e Emailer) sendWelcomeEmail(email_address string) error {
	// Create the welcome email.
	we := new(WelcomeEmail)
	we.EmailAddress = email_address
	var template_buffer bytes.Buffer
	err := email_templates.ExecuteTemplate(&template_buffer, "welcome.html", we)
	if err != nil {
		return err
	}
	we.Content = template_buffer.String()
	// Use the SendGrid API to send the email.
	from := mail.NewEmail("SwipeLong", "welcome@swipelong.com") // @TODO put in app config.
	subject := "Welcome to SwipeLong!"
	to := mail.NewEmail("", we.EmailAddress)
	plainTextContent := "Welcome to SwipeLong!" // backup
	htmlContent := we.Content
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err = client.Send(message) // _ = response. ask client, maybe later we'll log the response.
	if err != nil {
		return err
	}
	return nil
}
