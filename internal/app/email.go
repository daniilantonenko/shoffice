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

func (cfg *Configuration) Send(mail Mail) error {
	var err error

	/*fileWithPath := filepath.Join(RootDir(), "./web/templates/mail.html")
	if err != nil {
		log.Fatalln(err)
	}*/

	t, err := template.ParseFiles("./web/templates/mail.html")
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, mail); err != nil {
		return err
	}

	templateBody := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", mail.ToEmail)
	m.SetHeader("To", mail.ToEmail)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", templateBody)
	m.Attach(mail.FileName)

	// TODO: func create email server
	senderServer := cfg.EmailHost
	senderPort := cfg.EmailPort
	senderEmail := mail.ToEmail
	senderPass := cfg.EmailPassword

	d := gomail.NewDialer(senderServer, senderPort, senderEmail, senderPass)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Println("Message sent successfully")

	return nil
}
