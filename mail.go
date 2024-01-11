package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type Mail struct {
	Name    string
	Message string
	Subject string
	ToEmail string
}

type ConfigMail struct {
	FromEmail   string
	FromPass    string
	FileFormats []string
}

func sendMail(mail *Mail, config *ConfigMail) {

	// SMTP server configuration.
	fromSmtpHost := "smtp.gmail.com"
	fromSmtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", config.FromEmail, config.FromPass, fromSmtpHost)

	template, _ := template.ParseFiles("templates/mail.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", mail.Subject, mimeHeaders)))

	template.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    mail.Name,
		Message: mail.Message,
	})

	toEmail := []string{
		mail.ToEmail,
	}

	// Sending email.
	error := smtp.SendMail(fromSmtpHost+":"+fromSmtpPort, auth, config.FromEmail, toEmail, body.Bytes())
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println("Email Sent!")
}

func readConfig(fileName string) (config ConfigMail) {
	// TODO: return a generic structure

	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	return config
}
