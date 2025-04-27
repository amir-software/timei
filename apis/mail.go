package apis

import (
	"crypto/tls"

	gomail "gopkg.in/mail.v2"
)

func SendMain() {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "amirmahdibahrami7@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "amir.programmer.b@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "amirmahdibahrami7@gmail.com", "ajae uiva iudp ffyu") // Use APP password

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return
}
