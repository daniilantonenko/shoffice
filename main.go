package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
	}
	configMail := readConfig("conf.json")

	tmpl.ExecuteTemplate(w, "index", struct {
		Name string
	}{
		Name: configMail.CompanyName,
	})
}

func send(w http.ResponseWriter, r *http.Request) {
	emailForm := r.FormValue("inputEmail")
	commentForm := r.FormValue("inputComment")
	fileForm := r.FormValue("inputFile")

	if emailForm == "" || commentForm == "" || fileForm == "" {
		fmt.Fprintf(w, "Не все данные заполнены!")
	} else {
		fmt.Println(emailForm, commentForm, fileForm)

		configMail := readConfig("conf.json")
		simpleMail := Mail{
			Name:    "da",
			Message: "Simple first test message.",
			Subject: "Simple",
			ToEmail: "daniil454122922@gmail.com"}
		fmt.Println(simpleMail, configMail)
		sendMail(simpleMail, configMail)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func handleRequest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}

func main() {
	// Initializing the Web Server
	handleRequest()
}
