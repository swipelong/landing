package main

import (
	"errors"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

type EmailForm struct{}

func (ef EmailForm) CreateEmailSubmission(email_address string, ip_address string) (int, error) {
	// Validate the email address.
	err := ValidateEmail(email_address)
	if err != nil {
		return 422, errors.New("INVALID_EMAIL_ADDRESS")
	}
	// Enter the email address into our database.
	err = ef.createEmailSubmission(email_address, ip_address)
	if err != nil {
		// Check if the email address was already entered.
		if strings.Contains(err.Error(), "emails_entered_email_address_key") {
			return 422, errors.New("DUPLICATE_EMAIL_ADDRESS")
		}
		log.Fatal(err.Error())
		return 500, errors.New("Unknown error")
	}
	// Send off welcome email.
	e := new(Emailer)
	err = e.Send(email_address, "welcome")
	if err != nil {
		ef.markEmailStatus(email_address, false)
		return 500, errors.New("ESP_SEND_FAILURE")
	}
	ef.markEmailStatus(email_address, true)
	return 200, nil
}

func (ef EmailForm) createEmailSubmission(email_address string, ip_address string) error {
	// List all transaction stages.
	stage1 := `
		INSERT INTO emails_entered (
			email_address, 
			ip_address) 
		VALUES ($1, $2);`
	// Create a new transaction.
	db, err := DAO()
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer db.Close()
	// Prepare all statements and execute all transactions.
	{ // stage 1:
		stmt, err := tx.Prepare(stage1)
		if err != nil {
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(email_address, ip_address); err != nil {
			tx.Rollback()
			return err
		}
	}
	// Commit the transaction.
	return tx.Commit()
}

func (ef EmailForm) markEmailStatus(email_address string, status bool) error {
	// List all transaction stages.
	stage1 := `
		UPDATE emails_entered 
		SET
			status_sent = $1 
		WHERE
			email_address = $2;`
	// Create a new transaction.
	db, err := DAO()
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer db.Close()
	// Prepare all statements and execute all transactions.
	{ // stage 1:
		stmt, err := tx.Prepare(stage1)
		if err != nil {
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(status, email_address); err != nil {
			tx.Rollback()
			return err
		}
	}
	// Commit the transaction.
	return tx.Commit()
}
