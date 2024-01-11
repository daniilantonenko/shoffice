package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	Name    string
	Message string
	Subject string
	ToEmail string
}

type ConfigMail struct {
	EmailServer string
	EmailPort   int
	FromEmail   string
	FromPass    string
	CompanyName string
	FileFormats []string
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

func sendMail(mail Mail, config ConfigMail) {

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
	m.Attach("test.png")

	d := gomail.NewDialer(config.EmailServer, config.EmailPort, config.FromEmail, config.FromPass)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
