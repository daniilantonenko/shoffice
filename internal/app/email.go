package app

import (
	"bytes"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	Name     string
	Message  string
	Subject  string
	ToEmail  string
	FileName string
}

func SendMail(mail Mail, config Config) error {

	var err error
	t, _ := template.ParseFiles("templates/mail.html")

	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, mail); err != nil {
		log.Println(err)
	}

	templateBody := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", config.FromEmail)
	m.SetHeader("To", mail.ToEmail)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", templateBody)
	m.Attach(mail.FileName)

	d := gomail.NewDialer(config.EmailServer, config.EmailPort, config.FromEmail, config.FromPass)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
